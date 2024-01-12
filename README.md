# Lifeboat

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cluetec/lifeboat)](https://github.com/cluetec/lifeboat/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cluetec/lifeboat)](go.sum)
[![Go Report Card](https://goreportcard.com/badge/github.com/cluetec/lifeboat)](https://goreportcard.com/report/github.com/cluetec/lifeboat)

Lifeboat is a backup tool provided by [cluetec GmbH](https://cluetec.de). Lifeboat enables the user to create backups
for a range of different source systems (e.g. PostgreSQL, MongoDB, HashiCorp Vault) and storing the backup also in
different destination storage systems (e.g. S3 Buckets, Azure Blob, S/FTP, local filesystem).

## ✅ Supported Systems

Source systems:

- [x] Local filesystem
- [ ] PostgreSQL (Not implemented yet)
- [ ] MongoDB (Not implemented yet)
- [ ] HashiCorp Vault (Not implemented yet)

Destination storage systems:

- [x] Local filesystem
- [ ] S3 Bucket (Not implemented yet)
- [ ] Azure Storage Account (Not implemented yet)
- [ ] S/FTP (Not implemented yet)

## 🔥 Motivation

cluetec has been offering [software development services](https://www.cluetec.de/development/digitale-transformation/)
for several years now. We have been contracted with the implementation and initial operation of the software for a large
number of projects. Here we repeatedly encountered the backup issue for various database systems. To avoid having to
copy and adapt shell scripts back and forth every time, we thought about turning these shell scripts into an application
that could be used to back up various database systems with as little adaptation effort as possible.

## 💻 Installation

At the moment we don't provide any installation methods. As we just started the project, we will start with providing
the compiled binaries within the GitHub Releases. Later container images as also helm charts will follow.

## ⚙️ Usage

Lifeboat is a CLI tool which makes it possible to use it in various kind of environments like on a local machine, in a
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

If you want to set a config via an environment variable, just concatenate the yaml structure (in uppercase characters)
with underscores (`_`). For example like this: `SOURCE_FILESYSTEM_PATH`

To get an idea how the configuration can look like, have a look at the [`config.yaml`](./config.yaml).

#### General configuration

| Yaml Config | Default | Required | Description                                                                                    |
|-------------|---------|----------|------------------------------------------------------------------------------------------------|
| `logLevel`  | `info`  | 👎       | Defines the log level of the application. Possible value are: `debug`, `info`, `warn`, `error` |

#### Source system configuration

All configurations for the source system needs to be placed under the `source` object in the configuration.
Furthermore, we need to define, which source system we want to use. This will be done by setting the `type` field, like
it's done in the following example. The possible values can be found in the respective subsections for each source
system.

```yaml
source:
  type: filesystem
```

##### Filesystem

The following configs need to be place under the following yaml structure:

```yaml
source:
  type: filesystem
  filesystem:
    ...
```

| Yaml Config | Default | Required | Description                                                                                                  |
|-------------|---------|----------|--------------------------------------------------------------------------------------------------------------|
| `path`      |         | 👍       | Defines the path in the local filesystem (relative or absolute) to a file or folder that should be backed up. |

#### Destination storage configuration

All configurations for the destination storage systems needs to be placed under the `destination` object in the
configuration. Furthermore, we need to define, which destination storage system we want to use. This will be done by
setting the `type` field, like it's done in the following example. The possible values can be found in the respective
subsections for each destination storage system.

```yaml
destination:
  type: filesystem
```

##### Filesystem

The following configs need to be place under the following yaml structure:

```yaml
destination:
  type: filesystem
  filesystem:
    ...
```

| Yaml Config | Default | Required | Description                                                                              |
|-------------|---------|----------|------------------------------------------------------------------------------------------|
| `path`      |         | 👍       | Defines the path in the local filesystem (relative or absolute) where to store the file. |

## 🤝 Contribution

Everyone is more than welcome to contribute to this project! That's what open source is all about!

To have some contribution guidance, please have a look at [CONTRIBUTING.md](CONTRIBUTING.md).

### 👥 Contributors

<a title="Contributors" href="https://github.com/cluetec/lifeboat/graphs/contributors">
  <img alt="Contributors" src="https://contrib.rocks/image?repo=cluetec/lifeboat" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

## ⚖️ License

The project is licensed under the ["Apache-2.0"](./LICENSE) license.
