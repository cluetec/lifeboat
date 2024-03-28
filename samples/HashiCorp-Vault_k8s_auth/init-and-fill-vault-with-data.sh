#
# Copyright 2023 cluetec GmbH
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

#!/bin/bash

set -o errexit
set -o errtrace
set -o pipefail
set -o nounset

export VAULT_ADDR="http://localhost:8200"

# Initialize vault
initOutput=$(vault operator init)
#initOutput="Unseal Key 1: xxxx
#Unseal Key 2: xxxx
#Unseal Key 3: xxxx
#Unseal Key 4: xxxx
#Unseal Key 5: xxxx
#
#Initial Root Token: xxxxx
#
#Vault initialized with 5 key shares and a key threshold of 3. Please securely
#distribute the key shares printed above. When the Vault is re-sealed,
#restarted, or stopped, you must supply at least 3 of these keys to unseal it
#before it can start servicing requests.
#
#Vault does not store the generated root key. Without at least 3 keys to
#reconstruct the root key, Vault will remain permanently sealed!
#
#It is possible to generate new unseal keys, provided you have a quorum of
#existing unseal keys shares. See \"vault operator rekey\" for more information.: No such file or directory"

# Get unseal keys
unsealKeys=$(echo "$initOutput" | grep "^Unseal Key ")

# Get root token
while IFS= read -r SINGLELINE
do
    re="Initial Root Token: "
    if [[ "${SINGLELINE}" =~ $re ]]; then
        rootToken=$(echo "${SINGLELINE}" | sed "s/${re}//")
    fi
done << EOF
$initOutput
EOF

# Write unseal keys und token into separate files
echo "${unsealKeys}" > vault-unseal-keys.txt
echo "${rootToken}" > vault-token.txt

# Unseal vault
printf "\nFirst unseal command:\n"
vault operator unseal $(echo "${unsealKeys}" | sed -n 1p | sed 's/Unseal Key 1: //')
printf "\nSecond unseal command:\n"
vault operator unseal $(echo "${unsealKeys}" | sed -n 2p | sed 's/Unseal Key 2: //')
printf "\nThird unseal command:\n"
vault operator unseal $(echo "${unsealKeys}" | sed -n 3p | sed 's/Unseal Key 3: //')

# Login to vault
echo "$rootToken" | vault login -

kubectl exec -ti vault-1 -- vault operator raft join http://vault-0.vault-internal:8200
kubectl exec -ti vault-2 -- vault operator raft join http://vault-0.vault-internal:8200

# Unseal other nodes
kubectl exec -ti vault-1 -- vault operator unseal $(echo "${unsealKeys}" | sed -n 1p | sed 's/Unseal Key 1: //')
kubectl exec -ti vault-1 -- vault operator unseal $(echo "${unsealKeys}" | sed -n 2p | sed 's/Unseal Key 2: //')
kubectl exec -ti vault-1 -- vault operator unseal $(echo "${unsealKeys}" | sed -n 3p | sed 's/Unseal Key 3: //')

kubectl exec -ti vault-2 -- vault operator unseal $(echo "${unsealKeys}" | sed -n 1p | sed 's/Unseal Key 1: //')
kubectl exec -ti vault-2 -- vault operator unseal $(echo "${unsealKeys}" | sed -n 2p | sed 's/Unseal Key 2: //')
kubectl exec -ti vault-2 -- vault operator unseal $(echo "${unsealKeys}" | sed -n 3p | sed 's/Unseal Key 3: //')

# Enable & configure k8s auth
vault auth enable kubernetes || true
vault write auth/kubernetes/config kubernetes_host=https://kubernetes.default:443
vault write auth/kubernetes/role/default \
    policies="root" \
    bound_service_account_names="*" \
    bound_service_account_namespaces="*"

# Enable secret engine
vault secrets enable -version=2 -path="secret" kv

# Put secrets into vault
amountOfSecrets=1000
secretLength=2000
for i in `seq 2 $amountOfSecrets`; do
    printf "\nPut secret number ${i} into vault:\n"
    superSecureSecret=$(sed "s/[^a-zA-Z0-9]//g" <<< $(cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9!@#$%*()-+' | fold -w ${secretLength} | head -n 1))
    echo "${superSecureSecret}" | vault kv put secret/${i} content=-
done

printf "\nSuccessful initialized vault and put ${amountOfSecrets} secrets with a length of ${secretLength} random chars into vault\n"
