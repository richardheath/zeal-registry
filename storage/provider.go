package storage

import (
	"fmt"
	"mime/multipart"
)

// Provider Interfact for storage provider.
type Provider interface {
	Initialize(config map[string]interface{}) error
	DownloadPackage(repo, name, version, platform string) (DownloadInfo, error)
	PublishPackage(repo, name, version, platform string, data *multipart.File) error
	DownloadDefinition(repo, name, version, platform string) (DownloadInfo, error)
	PublishDefinition(repo, name, version, platform string, data *multipart.File) error
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
		provider = &LocalProvider{}
	} else {
		return nil, fmt.Errorf("Storage provider '%s' not supported", providerType)
	}

	err := provider.Initialize(config)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
