package service

import (
	"context"
	"fmt"
	"log"

	"github.com/justym/gq-playground/github-client/config"
	"github.com/justym/gq-playground/github-client/model"
	"github.com/machinebox/graphql"
)

//LicenseClient is interface for API Client
type LicenseClient interface {
	Do() error
}

//implement LicenseClient
type licenseClient struct {
	Client   *graphql.Client
	Query    string
	Response model.Response
}

//Do asks GitHub API to get license information
func (l *licenseClient) Do() error {
	req := graphql.NewRequest(l.Query)
	token, err := config.GetGithubToken()
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "bearer "+token)
	ctx := context.Background()
	if err := l.Client.Run(ctx, req, &l.Response); err != nil {
		return err
	}

	msg := fmt.Sprintf("Licencse %s:\n Description: %s", l.Response.License.Name, l.Response.License.Description)
	log.Println(msg)

	return nil
}

func NewLisenseClient(endpoint, q string, options ...graphql.ClientOption) LicenseClient {
	client := graphql.NewClient(endpoint, options...)
	fmt.Println(q)
	return &licenseClient{Client: client, Query: q}
}
