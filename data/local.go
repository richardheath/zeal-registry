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

func (handler *LocalProvider) Initialize(config map[string]interface{}) error {
	localDB, err := db.OpenDB(config["path"].(string))
	if err != nil {
		return err
	}

	handler.db = localDB
	return nil
}

func (handler *LocalProvider) SearchByKeywords(keywords []string) ([]string, error) {
	return nil, nil
}

func (handler *LocalProvider) SearchByPackageVersions(name, platform, filter string) ([]string, error) {
	return nil, nil
}

func (handler *LocalProvider) PackageExist(name, version string) (bool, error) {
	return false, nil
}

func (handler *LocalProvider) GetPackageMetadata(name, version string) (PackageMetadata, error) {
	return PackageMetadata{}, nil
}

func (handler *LocalProvider) CreatePackage(name, version string, platforms, keywords []string) error {
	return nil
}

func (handler *LocalProvider) AddPackagePlatform(name, version, platform string) error {
	return nil
}
