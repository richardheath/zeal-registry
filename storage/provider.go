package storage

import (
	"fmt"
	"mime/multipart"
)

// Provider Interfact for storage provider.
type Provider interface {
	initialize(config map[string]interface{}) error
	DownloadPackage(repo, packageName, version, platform string) (DownloadInfo, error)
	PublishPackage(repo, packageName, version, platform string, data *multipart.File) error
}

// DownloadInfo Meatadata for target file for download.
type DownloadInfo struct {
	RedirectURL string
	FilePath    string
}

// InitializeProvider Returns an instance of storage provider of given type.
func InitializeProvider(providerType string, config map[string]interface{}) (Provider, error) {
	var provider Provider
	if providerType == "local" {
		provider = LocalProvider{}
	} else {
		return nil, fmt.Errorf("Storage provider '%s' not supported", providerType)
	}

	err := provider.initialize(config)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
