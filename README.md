# Lifeboat

Lifeboat is an backup tool provided by [cluetec GmbH](https://cluetec.de).
Lifeboat enables the user to create a backup for a range of different systems (e.g. PostgeSQL, MongoDB) and storing the backup also in different storage systems (e.g. S3 Buckets, Azure Blob, S/FTP, local disk).

## ✅ Supported Systems

Source systems:

- Local filesystem
- PostgeSQL (Not implemented yet)
- MongoDB (Not implemented yet)
- HashiCorp Vault (Not implemented yet)

Destination storage systems:

- Local filesystem
- S3 Bucket (Not implemented yet)
- Azure Storage Account (Not implemented yet)
- S/FTP (Not implemented yet)

## 🔥 Motivation

## ⚙️ Usage

Lifeboat is designed as a CLI tool. That allows the user to use it in a large varient of environments like on a local machine, in a unix cronjob, in kubernetes, in a VM, where ever the user wants.
As the tool needs a quite complex configuration it's possible to provide a config file. In addition it's also possible to put all configuration also in environment variables.

TODO: put here cli help output

### Configurations

The configurations is divided in three different parts:
1. General configs
2. Source system configs
3. Destination storage configs

#### General configs

TODO PUT HERE TABLE WITH THE CONFIGS

#### Source system configs

TODO PUT HERE SUBSECTIONS WITH TABLES OF THE CONFIGS

#### Destination storage configs

TODO PUT HERE SUBSECTIONS WITH TABLES OF THE CONFIGS

## 🤝 Contribution

Everyone is more than welcome to contribute to this project! That's what open source is all about!

To have some contribution guidance, please have a look at [CONTRIBUTING.md](CONTRIBUTING.md).

## License

The project is licensed under the ["Apache-2.0"](./LICENSE) license.
