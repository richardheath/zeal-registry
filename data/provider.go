package data

import (
	"fmt"
)

// Provider Interface for data provider.
type Provider interface {
	Initialize(config map[string]interface{}) error
	SearchByKeywords(keywords []string) ([]string, error)
	SearchByPackageVersions(name, platform, filter string) ([]string, error)

	PackageExist(name, version string) (bool, error)
	GetPackageMetadata(name, version string) (PackageMetadata, error)
	CreatePackage(name, version string, platforms, keywords []string) error
	AddPackagePlatform(name, version, platform string) error
}

// PackageMetadata Package metadata.
type PackageMetadata struct {
	Name      string
	Version   string
	Keywords  []string
	Platforms []string
}

// InitializeProvider Initialize data provider.
func InitializeProvider(providerType string, config map[string]interface{}) (Provider, error) {
	var provider Provider
	if providerType == "local" {
		provider = &LocalProvider{}
	} else {
		return nil, fmt.Errorf("Data provider not supported: %s", providerType)
	}

	err := provider.Initialize(config)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
