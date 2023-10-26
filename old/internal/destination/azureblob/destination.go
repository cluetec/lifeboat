package azureblob

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/bloberror"
	vault "github.com/hashicorp/vault/api"
	"github.com/rs/zerolog"
)

type Destination struct {
	Logger zerolog.Logger
	Config AzureBlobConfig
}

func (d *Destination) Store(response *vault.Response, filename string) {
	client, err := getClient(d.Config.AccountName, d.Config.AccountKey)
	if err != nil {
		d.Logger.Fatal().
			Stack().
			Err(err).
			Msg("Could not stream backup to azure blob")
		return
	}

	_, err = client.CreateContainer(context.TODO(), d.Config.ContainerName, nil)
	if err != nil && !bloberror.HasCode(err, bloberror.ContainerAlreadyExists) {
		d.Logger.Fatal().
			Stack().
			Err(err).
			Msg("Could not create azure blob container")
	}

	_, err = client.UploadStream(context.TODO(),
		d.Config.ContainerName,
		filename,
		response.Body,
		nil)
	if err != nil {
		d.Logger.Fatal().
			Stack().
			Err(err).
			Msg("Could not stream backup to azure blob")
	}
}

func getClient(accountName string, accountKey string) (*azblob.Client, error) {
	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, err
	}

	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	return azblob.NewClientWithSharedKeyCredential(
		fmt.Sprintf("https://%s.blob.core.windows.net/", accountName),
		cred,
		nil)
}
