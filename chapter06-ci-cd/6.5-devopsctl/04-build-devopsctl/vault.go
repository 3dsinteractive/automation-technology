package main

import (
	vault "github.com/hashicorp/vault/api"
)

type IVaultService interface {
	GetSecretS(
		token string,
		path string,
		key string) (string, error)
	GetSecretM(
		token string,
		path string,
		keys []string) (map[string]string, error)
}

type VaultService struct {
	ctx IContext
}

func NewVaultService(ctx IContext) *VaultService {
	return &VaultService{
		ctx: ctx,
	}
}

func (svc *VaultService) getClient(
	vaultAddr string,
	token string) (*vault.Client, error) {

	config := vault.DefaultConfig()
	client, err := vault.NewClient(config)
	if err != nil {
		return nil, err
	}
	client.SetToken(token)
	return client, nil
}

func (svc *VaultService) GetSecretS(
	vaultAddr string,
	token string,
	path string,
	key string) (string, error) {

	retM, err := svc.GetSecretM(vaultAddr, token, path, []string{key})
	if err != nil {
		return "", err
	}
	return retM[key], nil
}

func (svc *VaultService) GetSecretM(
	vaultAddr string,
	token string,
	path string,
	keys []string) (map[string]string, error) {

	retM := map[string]string{}

	client, err := svc.getClient(vaultAddr, token)
	if err != nil {
		return retM, err
	}

	secret, err := client.Logical().Read(path)
	if err != nil {
		return retM, err
	}

	if secret == nil || secret.Data == nil {
		return retM, nil
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return retM, nil
	}

	for _, k := range keys {
		val, _ := data[k].(string)
		retM[k] = val
	}

	return retM, nil
}
