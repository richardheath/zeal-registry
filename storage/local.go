package storage

import (
	"mime/multipart"
)

type LocalProvider struct {
	packagesDir string
}

func InitializeLocalProvider(packagesDir string) (LocalProvider, error) {
	handler := LocalProvider{}
	handler.packagesDir = packagesDir
	return handler, nil
}

func (handler *LocalProvider) Initialize(config map[string]interface{}) error {
	return nil
}

func (handler *LocalProvider) DownloadPackage(repo, name, version, platform string) (DownloadInfo, error) {
	return DownloadInfo{}, nil
}

func (handler *LocalProvider) PublishPackage(repo, name, version, platform string, data *multipart.File) error {
	return nil
}

func (handler *LocalProvider) DownloadDefinition(repo, name, version, platform string) (DownloadInfo, error) {
	return DownloadInfo{}, nil
}

func (handler *LocalProvider) PublishDefinition(repo, name, version, platform string, data *multipart.File) error {
	return nil
}
