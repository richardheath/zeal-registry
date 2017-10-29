package data

import (
	"github.com/HouzuoGuo/tiedot/db"
)

type LocalProvider struct {
	db *db.DB
	// Package file location. Use package/version/platform as key
	// then file path as value.
	filesCollection []string // files collection
	// Package versions, info, search keywords
	packagesCollection []string
	// Pure package metadata
	metadataCollection []string
}

type LocalProviderConfig struct {
	dataFolder string
}

func (handler LocalProvider) initialize(config map[string]interface{}) error {
	localDB, err := db.OpenDB(config["path"].(string))
	if err != nil {
		return err
	}

	handler.db = localDB
	return nil
}

func (handler LocalProvider) SearchPackages(keywords []string) ([]string, error) {
	return nil, nil
}

func (handler LocalProvider) PackageExist(packageName string, version string) (bool, error) {
	return false, nil
}

func (handler LocalProvider) GetPackageVersions(packageName string) ([]PackageVersions, error) {
	return nil, nil
}

func (handler LocalProvider) GetPackageMetadata(packageName string, version string) (PackageMetadata, error) {
	return PackageMetadata{}, nil
}

func (handler LocalProvider) SetPackageConfiguration(packageName string, version string, PackageMetadata map[string]interface{}) error {
	return nil
}

func (handler LocalProvider) AddPackagePlatform(packageName string, version string, platform string) error {
	return nil
}
