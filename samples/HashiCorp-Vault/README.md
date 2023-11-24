# Sample: Backup HashiCorp Vault

In this example we will show you how to back up an HashiCorp Vault instance.

## Requirements

The Vault instance needs to use [raft](https://developer.hashicorp.com/vault/docs/configuration/storage/raft) as the
underlying storage engine.

## Policy

In this sample we are using for simplicity reasons the "root token" to authorize us. In a real world scenario you
would use a separate took or any other
[authentication method supported by vault](https://developer.hashicorp.com/vault/docs/auth).

**What's important here**: Normally your identity shouldn't have root permissions! The only permission you need for
creating the backup/snapshot is the following
([vault policy](https://developer.hashicorp.com/vault/docs/concepts/policies)):

```hcl
path "/sys/storage/raft/snapshot" {
  capabilities = ["read"]
}
```

## Run

### 1. Start the vault instance

The HashiCorp Vault setup is based on a docker compose setup.

```shell
# -d is starting the container in the background
$ docker-compose up -d
```

As we can't simply use the dev mode of Vault, we need to initialize and unseal it first. For this purpose the
`docker-compose.yaml` contains next to the vault container an additional one which is called `vault-init`. This
container contains the bash script `./init-and-fill-vault.sh` which will do all the necessary steps.

On default, we are storing 1000 secrets with a length of 2000 random chars into vault. As this takes some seconds, you
can verify with this command, if the container has successfully executed the script or not. As a hint, it could take
something around 1 minute until the init script finishes.

- Successful: Status of vault-init container == `Exited (0)`
- Not successful: Status of vault-init container == `Exited (1)`

```shell
$ docker-compose ps --all
NAME                           IMAGE                        COMMAND                  SERVICE      CREATED          STATUS                     PORTS
hashicorp-vault-vault-1        hashicorp/vault:1.15         "vault server -confi…"   vault        59 seconds ago   Up 58 seconds              0.0.0.0:8200->8200/tcp
hashicorp-vault-vault-init-1   hashicorp-vault-vault-init   "bash /init.sh"          vault-init   59 seconds ago   Exited (0) 5 seconds ago
```

### 3. Run lifeboat to create the backup

As the root token will be randomly generated everytime you are starting a new vault instance, we are storing it in the
file `./vault-token.txt` so that we can use it in lifeboat to successfully authenticate while doing the backup.
Therefor we need to parse the content of this file into an environment variable which will be then used by
lifeboat.

The following command will trigger a backup and will store it in the `./backup-destination` folder:

```shell
$ SOURCE_TOKEN=$(cat ./vault-token.txt) lb backup --config ./backup-config.yaml
```

## Clean up after run

To clean up everything afterwards, we just need to execute the following commands:

```shell
$ docker-compose down
$ rm -rf .data
$ rm -rf backup-destination/vault-backup.snap
```

## Restore

An official guide how to restore a backup/snapshot can be found here:
<https://developer.hashicorp.com/vault/tutorials/standard-procedures/sop-restore>
