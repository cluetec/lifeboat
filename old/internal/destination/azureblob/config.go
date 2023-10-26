package azureblob

type AzureBlobConfig struct {
	AccountName   string `validate:"required"`
	AccountKey    string `validate:"required"`
	ContainerName string `validate:"required"`
}
