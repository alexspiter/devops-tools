package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	drone "github.com/drone/drone-go/drone"
	drone_go "github.com/honestbee/devops-tools/quay/pkg/drone"
	"github.com/honestbee/devops-tools/quay/pkg/github"
	"github.com/honestbee/devops-tools/quay/pkg/quay"
	"github.com/tuannvm/tools/pkg/utils"
)

func getGithubRepos() ([]github.Repo, error) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	githubClient := github.NewClient(ctx, githubToken)

	githubConfig := github.Config{
		Organization: "honestbee",
		Client:       githubClient,
		Repo: github.RepoConfig{
			Type: "private",
		},
	}

	repos, err := githubConfig.ListRepos()
	return repos, err

}

func createQuayRepos(droneRepos []*drone.Repo) ([]quay.RepositoryOutput, error) {
	var quayRepoOuputs []quay.RepositoryOutput

	for _, droneRepo := range droneRepos {
		quayRepoInput := quay.RepositoryInput{
			Namespace:   "honestbee",
			Visibility:  droneRepo.Visibility,
			Repository:  strings.ToLower(droneRepo.Name),
			Description: "",
		}

		quayRepoOutput, err := quayRepoInput.CreateRepository()
		if err != nil {
			fmt.Printf("Error creating %v : %v", quayRepoInput, err)
		}
		quayRepoOuputs = append(quayRepoOuputs, quayRepoOutput)
	}
	return quayRepoOuputs, nil

}

func droneRegistryCreate(c drone.Client, hostname, username, password string, repo *drone.Repo) error {
	registry := &drone.Registry{
		Address:  hostname,
		Username: username,
		Password: password,
	}
	_, err := c.RegistryCreate(repo.Owner, repo.Name, registry)
	if err != nil {
		return err
	}
	return nil
}

func saveToCsv(droneRepos []*drone.Repo) error {
	file, err := os.Create("repos.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data := [][]string{}

	for _, droneRepo := range droneRepos {
		data = append(data, []string{droneRepo.Name})
	}

	utils.SaveToCsv(file, data)
	return nil
}

func droneSecretCreate(c drone.Client, name, value string, repo *drone.Repo) error {
	var defaultSecretEvents = []string{
		drone.EventPush,
		drone.EventTag,
		drone.EventDeploy,
	}
	secret := &drone.Secret{
		Name:   name,
		Value:  value,
		Events: defaultSecretEvents,
	}
	_, err := c.SecretCreate(repo.Owner, repo.Name, secret)
	if err != nil {
		return err
	}
	return nil
}

func droneSecretDelete(c drone.Client, name, value string, repo *drone.Repo) error {
	var defaultSecretEvents = []string{
		drone.EventPush,
		drone.EventTag,
		drone.EventDeploy,
	}
	secret := &drone.Secret{
		Name:   name,
		Value:  value,
		Events: defaultSecretEvents,
	}
	_, err := c.SecretCreate(repo.Owner, repo.Name, secret)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	droneClient, _ := drone_go.NewClient()
	droneRepos, _ := droneClient.RepoList()

	saveToCsv(droneRepos)

	// Create secret

	//secrets := []struct {
	//	Name  string
	//	Value string
	//}{
	//	{
	//		Name:  "quay_username",
	//		Value: os.Getenv("QUAY_USERNAME"),
	//	},
	//	{
	//		Name:  "quay_password",
	//		Value: os.Getenv("QUAY_PASSWORD"),
	//	},
	//	{
	//		Name:  "quay_registry",
	//		Value: "quay.io",
	//	},
	//}

	//for _, droneRepo := range droneRepos {
	//	for _, secret := range secrets {
	//		err := droneSecretCreate(droneClient, secret.Name, secret.Value, droneRepo)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//	}
	//}

	// Create Registry
	//hostname := "quay.io"
	//username := os.Getenv("DRONE_USERNAME")
	//password := os.Getenv("DRONE_PASSWORD")
	//for _, droneRepo := range droneRepos {
	//	err := droneRegistryCreate(droneClient, hostname, username, password, droneRepo)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	// Create quay repo
	quayRepos, err := createQuayRepos(droneRepos)
	if err != nil {
		panic(err)
	}
	for _, quayRepo := range quayRepos {
		fmt.Printf("%v\n", quayRepo.Name)
	}
}
