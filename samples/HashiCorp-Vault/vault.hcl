ui = true

api_addr = "http://127.0.0.1:8200"
cluster_addr = "https://127.0.0.1:8201"

listener "tcp" {
  tls_disable = 1
  address = "[::]:8200"
  cluster_address = "[::]:8201"
}

disable_mlock = true
storage "raft" {
  path = "/vault/data"
}
