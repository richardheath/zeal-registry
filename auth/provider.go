package auth

import "fmt"

// Provider Interface for auth provider.
type Provider interface {
	Initialize(config map[string]interface{}) error
	Authenticate(username, apikey string) (success bool, err error)
	GetUserRepoPermission(username, repo string) string
	GetRepoSettings(repo string) (RepoSettings, error)
}

type RepoSettings struct {
	AllowAnonymousRead  bool `json:"allowAnonymousRead"`
	AllowAnonymousWrite bool `json:"allowAnonymousWrite"`
}

// InitializeProvider Initialize data provider.
func InitializeProvider(providerType string, config map[string]interface{}) (Provider, error) {
	var provider Provider
	if providerType == "local" {
		provider = &LocalProvider{}
	} else {
		return nil, fmt.Errorf("Auth provider not supported: %s", providerType)
	}

	err := provider.Initialize(config)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
