package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// LocalProvider Use local JSON file for auth.
type LocalProvider struct {
	config authConfig
}

type authConfig struct {
	Users map[string]userConfig   `json:"users"`
	Repos map[string]RepoSettings `json:"repos"`
}

type userConfig struct {
	Apikey          string            `json:"apikey"`
	RepoPermissions map[string]string `json:"repoPermissions"`
}

// Initialize Load auth config JSON file.
func (provider *LocalProvider) Initialize(config map[string]interface{}) error {
	rawPath, pathGiven := config["path"]
	if !pathGiven {
		return fmt.Errorf("Config path not given")
	}

	configPath := fmt.Sprintf("%v", rawPath)
	provider.config = authConfig{}
	rawConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawConfig, &provider.config)
	if err != nil {
		return err
	}

	return nil
}

// Authenticate Authenticate user.
func (provider *LocalProvider) Authenticate(username, apikey string) (success bool, err error) {
	user, userFound := provider.config.Users[username]
	if !userFound {
		return false, nil
	}

	if user.Apikey != apikey {
		return false, nil
	}

	return true, nil
}

func (provider *LocalProvider) GetUserRepoPermission(username, repo string) string {
	user, userFound := provider.config.Users[username]
	if !userFound {
		return ""
	}

	repoPermission, repoFound := user.RepoPermissions[repo]
	if !repoFound {
		return ""
	}

	return repoPermission
}

func (provider *LocalProvider) GetRepoSettings(repoName string) (RepoSettings, error) {
	repo, repoFound := provider.config.Repos[repoName]
	if !repoFound {
		return RepoSettings{}, fmt.Errorf("repo not found: " + repoName)
	}

	return repo, nil
}
