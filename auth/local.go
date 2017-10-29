package auth

// LocalProvider Use local JSON file for auth.
type LocalProvider struct {
}

func (provider LocalProvider) initialize(config map[string]interface{}) error {
	return nil
}

func (provider LocalProvider) Authenticate(username, password, repo string) (success bool, err error) {
	return true, nil
}
