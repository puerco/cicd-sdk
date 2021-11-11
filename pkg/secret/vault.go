// Copyright (c) 2021-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package secret

import (
	"crypto/tls"
	"net/http"
	"os"

	vault "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type VaultManager struct {
	client *vault.Client
}

func NewVaultManager() (*VaultManager, error) {
	v := &VaultManager{}
	config := vault.DefaultConfig() // modify for more granular configuration
	config.HttpClient.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client, err := vault.NewClient(config)
	if err != nil {
		return nil, errors.Wrap(err, "initializing vault client")
	}
	v.client = client

	// Check if auth is set
	if t := os.Getenv("VAULT_TOKEN"); t != "" {
		v.client.SetToken(os.Getenv("VAULT_TOKEN"))
	}
	return v, nil
}

func (v *VaultManager) Get(key string) (*Secret, error) {
	secret, err := v.client.Logical().Read(key)
	if err != nil {
		return nil, errors.Wrap(err, "fetching secret")
	}
	logrus.Infof("%+v", secret)

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("unable to find data entry in response")
	}

	return &Secret{
		Key:   key,
		Value: data,
	}, nil
}

func (v *VaultManager) Write(key string, val Value) error {
	_, err := v.client.Logical().Write(key, val)
	return errors.Wrap(err, "writing secret to Vault")
}

func (v *VaultManager) List() ([]string, error) {
	return []string{}, nil
}
