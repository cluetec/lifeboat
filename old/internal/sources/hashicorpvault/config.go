package hashicorpvault

import (
	vault "github.com/hashicorp/vault/api"
)

type Config struct {
	Address        string `validate:"http_url,required"`
	KubernetesRole string `validate:"required"`
}

// GetHashiCorpVaultConfig was implement in regard to the
// `vault.DefaultConfig()` method. While the implementation the
// `config.ReadEnvironment` was left of, to avoid the usage of additional
// environment variables like `VAULT_ADDRESS`.
func (c *Config) GetHashiCorpVaultConfig() *vault.Config {
	config := vault.Config{
		Address: c.Address,
	}

	return &config
}
