# Agile notes

## Planned features

- [ ] Step: 1
  - [x] Init cobra & viper setup
  - [ ] Dynamic config loading with viper (`type`)
- [ ] Step: 2
  - [ ] Init slog setup
- [ ] Implement Filesystem to Filesystem backup
  - [ ] Think about the interfaces
- [ ] Implement Vault to Filesystem backup

For later:

- [ ] Define multiple destinations
- [ ] Azure Blob
  - [ ] Support additional auth mechanism
- [ ] Creating full-backups from
  - [ ] SSH/Rsync
  - [ ] MongoDB
  - [ ] PostgreSQL
  - [ ] Disk
- [ ] Delete backups that are older than X
- [ ] Create incremental backups from
  - [ ] MongoDB
  - [ ] PostgreSQL
- [ ] Define multiple sources
- [ ] K8s Volume replication
- [ ] Encryption
- [ ] Compression
- [ ] Implement command `check-config`
- [ ] Implement command `generate-config`

## Config ideas

```yaml
source:
  type: "hashicorpvault"
  address: http://localhost:8200
  token: xxx

destination:
  type: azureBlob
  containerName: backup
  accountName: xxx
  accountKey: xxx
```

```shell
lb backup
```

---

```yaml
sources:
  production-env:
    type: hashicorpvault
    address: http://localhost:8200
    token: xxx

destination:
  my-backup:
    type: azureBlob
    containerName: xxx
    accountName: xxx
    accountKey: xxxx
```

```shell
lb backup --source=production-env --destination=my-backup
```

---

```yaml
source:
  type: "hashicorpvault"
  name: "production-env"
  address: http://localhost:8200
  token: xxx

sources:
  - type: "hashicorpvault"
    name: "production-env"
    address: http://localhost:8200
    token: xxx

destination:
  type: azureBlob
  name: backup
  containerName: xxx
  accountName: xxx
  accountKey: xxx

destinations:
  - type: azureBlob
    name: backup
    containerName: xxx
    accountName: xxx
    accountKey: xxx
```

```shell
lb backup
lb backup --source=production-env --destination=backup
```

## Excalidraw

<https://excalidraw.com/#room=5962c55eb139a273c2c0,40kFzCW8vuuOOpjpaNdCLg>
