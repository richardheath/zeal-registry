package data

import (
	"fmt"
)

// Provider Interface for data provider.
type Provider interface {
	initialize(config map[string]interface{}) error
	SearchPackages(keywords []string) ([]string, error)
	PackageExist(packageName string, version string) (bool, error)
	GetPackageVersions(packageName string) ([]PackageVersions, error)
	GetPackageMetadata(packageName string, version string) (PackageMetadata, error)
	SetPackageConfiguration(packageName string, version string, PackageMetadata map[string]interface{}) error
	AddPackagePlatform(packageName string, version string, platform string) error
}

// PackageVersions Package versions.
type PackageVersions struct {
	PackageName string
	Versions    []string
}

// PackageMetadata Package metadata.
type PackageMetadata struct {
	PackageName        string
	Version            string
	Configuration      map[string]interface{}
	SupportedPlatforms []string
}

// InitializeProvider Initialize data provider.
func InitializeProvider(providerType string, config map[string]interface{}) (Provider, error) {
	var provider Provider
	if providerType == "local" {
		provider = LocalProvider{}
	} else {
		return nil, fmt.Errorf("Data provider not supported: %s", providerType)
	}

	err := provider.initialize(config)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
