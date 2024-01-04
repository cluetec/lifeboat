# Lifeboat

Lifeboat is a backup tool provided by [cluetec GmbH](https://cluetec.de). Lifeboat enables the user to create a backup
for a range of different systems (e.g. PostgreSQL, MongoDB) and storing the backup also in different storage systems
(e.g. S3 Buckets, Azure Blob, S/FTP, local disk).

## ✅ Supported Systems

Source systems:

- [x] Local filesystem
- [ ] PostgeSQL (Not implemented yet)
- [ ] MongoDB (Not implemented yet)
- [ ] HashiCorp Vault (Not implemented yet)

Destination storage systems:

- [x] Local filesystem
- [ ] S3 Bucket (Not implemented yet)
- [ ] Azure Storage Account (Not implemented yet)
- [ ] S/FTP (Not implemented yet)

## 🔥 Motivation

## 💻 Installation

At the moment we don't provide any installation methods.
As we just started the project, we will start with providing the compiled binaries within the GitHub Releases.

## ⚙️ Usage

Lifeboat is a CLI tool which makes it possible to use it in a various kind of environments like on a local machine, in a
Unix cronjob, in Kubernetes, in a VM, wherever the user wants. As the tool needs a quite complex configuration it's
possible to provide a config file. In addition, it's also possible to provide all configuration via environment
variables.

```shell
$ lb
Lifeboat is a general purpose backup tool which supports backups for arbitrary sources and destinations.

Usage:
  lb [command]

Available Commands:
  backup      Execute the backup.
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for lb

Use "lb [command] --help" for more information about a command.
```

### Configuration

The configuration is divided in three different parts:

1. General configs
2. Source system configs
3. Destination storage configs

#### General configuration

| Yaml Config | Env        | Default | Required | Description                                                                                 |
|-------------|------------|---------|----------|---------------------------------------------------------------------------------------------|
| `logLevel`  | `LOGLEVEL` | `info`  | 👎       | Defines the log level of the application. Valid value are: `debug`, `info`, `warn`, `error` |

#### Source system configuration

##### Filesystem

| Yaml Config | Env        | Default | Required | Description                                                                                 |
|-------------|------------|--------|----------|---------------------------------------------------------------------------------------------|
| `logLevel`  | `LOGLEVEL` |   | 👎       | Defines the log level of the application. Valid value are: `debug`, `info`, `warn`, `error` |

#### Destination storage configuration

TODO PUT HERE SUBSECTIONS WITH TABLES OF THE CONFIGS

## 🤝 Contribution

Everyone is more than welcome to contribute to this project! That's what open source is all about!

To have some contribution guidance, please have a look at [CONTRIBUTING.md](CONTRIBUTING.md).

## License

The project is licensed under the ["Apache-2.0"](./LICENSE) license.
