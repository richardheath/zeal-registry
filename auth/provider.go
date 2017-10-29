package auth

import "fmt"

// Provider Interface for auth provider.
type Provider interface {
	initialize(config map[string]interface{}) error
	Authenticate(username, password, repo string) (success bool, err error)
}

// InitializeProvider Initialize data provider.
func InitializeProvider(providerType string, config map[string]interface{}) (Provider, error) {
	var provider Provider
	if providerType == "local" {
		provider = LocalProvider{}
	} else {
		return nil, fmt.Errorf("Auth provider not supported: %s", providerType)
	}

	err := provider.initialize(config)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
