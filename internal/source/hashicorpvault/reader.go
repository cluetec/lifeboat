/*
 * Copyright 2023 cluetec GmbH
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package hashicorpvault

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/cluetec/lifeboat/internal/config/validator"
	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/kubernetes"
)

const snapshotPath = "/sys/storage/raft/snapshot"

// Reader implements the `io.ReaderClose` interface in order to read the backup from HashiCorp Vault.
type Reader struct {
	client *vault.Client
	reader io.Reader
}

// NewReader initializes a new `Reader` struct which is implementing the `io.ReaderClose` interface.
func NewReader(c *Config) (*Reader, error) {
	if err := validator.Validator.Struct(c); err != nil {
		return nil, err
	}

	slog.Debug("source config validated", "sourceType", Type, "config", c)

	client, err := vault.NewClient(c.GetHashiCorpVaultConfig())
	if err != nil {
		return nil, err
	}

	switch c.AuthMethod {
	case "token":
		client.SetToken(c.Token)
	case "kubernetes":
		k8sAuth, err := auth.NewKubernetesAuth(c.KubernetesAuth.RoleName)
		if err != nil {
			return nil, err
		}

		authInfo, err := client.Auth().Login(context.TODO(), k8sAuth)
		if err != nil {
			return nil, err
		}
		if authInfo == nil {
			return nil, fmt.Errorf("no auth info was returned after login")
		}
	}

	return &Reader{client: client}, nil
}

func (r *Reader) Read(b []byte) (int, error) {
	slog.Debug("read got called", "sourceType", Type)

	if r.reader == nil {
		resp, err := r.client.Logical().ReadRaw(snapshotPath)
		if err != nil {
			slog.Error("failed to called backup endpoint", "error", err)
			return 0, err
		}

		r.reader = resp.Body
	}

	return r.reader.Read(b)
}

func (r *Reader) Close() error {
	slog.Debug("closing reader", "sourceType", Type)
	if closer, ok := r.reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
