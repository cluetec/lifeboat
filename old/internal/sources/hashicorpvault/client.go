package hashicorpvault

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/kubernetes"
	"golang.org/x/net/context"
)

const snapshotPath = "/sys/storage/raft/snapshot"

func InitClient(c *Config) (*vault.Client, error) {
	client, err := vault.NewClient(c.GetHashiCorpVaultConfig())
	if err != nil {
		return nil, err
	}

	k8sAuth, err := auth.NewKubernetesAuth(c.KubernetesRole)

	authInfo, err := client.Auth().Login(context.Background(), k8sAuth)
	if err != nil {
		return nil, fmt.Errorf("unable to log in with Kubernetes auth: %w", err)
	}
	if authInfo == nil {
		return nil, fmt.Errorf("no auth info was returned after login")
	}

	return client, nil
}

func Backup(ctx context.Context, client *vault.Client) (*vault.Response, error) {
	return client.Logical().ReadRawWithContext(ctx, snapshotPath)
}
