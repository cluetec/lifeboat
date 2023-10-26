# Lifeboat

## Configuration

| Config Name | Required | Default Value | Description                                                                                                                 |
|-------------|----------|---------------|-----------------------------------------------------------------------------------------------------------------------------|
| `logLevel`  |          | `info`        | Defines the log level which will be used by the application. Possible values are: `trace`, `debug`, `info`, `warn`, `error` |

### Sources

#### HashiCorp Vault

| Config Name                | Required | Default Value | Description                                      |
|----------------------------|----------|---------------|--------------------------------------------------|
| `source.hashiCorp.address` | X        |               | URL under which the vault REST API is accessible |
| `source.hashiCorp.token`   | X        |               | Vault token                                      |

### Destinations

| Config Name         | Required | Default Value | Description                                                     |
|---------------------|----------|---------------|-----------------------------------------------------------------|
| `destinations.path` |          | `.`           | Path under which the backup should be stored on the destination |

#### AzureBlob

| Config Name                           | Required | Default Value | Description                                     |
|---------------------------------------|----------|---------------|-------------------------------------------------|
| `destination.azureBlob.accountName`   | X        |               | Storage account name                            |
| `destination.azureBlob.accountKey`    | X        |               | Storage account key                             |
| `destination.azureBlob.containerName` | X        |               | Container name inside the Azure Storage Account |

## HashiCorp Vault authentication

The policy which is used by the token need to have the `sudo` capability to create backups.

## Additional Features

- Define multiple destinations
- Azure Blob
  - Support additional auth mechanism
- Creating full-backups from
  - SSH/Rsync
  - MongoDB
  - PostgreSQL
  - Disk
- Delete backups that are older than X
- Create incremental backups from
  - MongoDB
  - PostgreSQL
- Define multiple sources
- K8s Volume replication
- Encryption
- Compression

## Alternatives

- [gobackup/gobackup](https://github.com/gobackup/gobackup)

## TODOs

- Make vault token been mountable as file
  - Block that the vault client automatically pull in env vars like `VAULT_ADDRESS`

## Additional documentation

- <https://github.com/Azure/azure-sdk-for-go/tree/main/sdk/storage/azblob#readme>
