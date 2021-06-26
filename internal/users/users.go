package users

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ngavinsir/golangtraining"
)

type Client struct {
	URL        string
	HTTPClient *http.Client
}

func NewUsersClient(httpClient *http.Client) *Client {
	return &Client{
		URL:        "https://jsonplaceholder.typicode.com/users",
		HTTPClient: httpClient,
	}
}

func (c *Client) GetUsers(ctx context.Context) (golangtraining.User, error) {
	var res golangtraining.User

	u, err := url.Parse(c.URL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("Accept", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	jbyt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	var arrRes []golangtraining.User
	err = json.Unmarshal(jbyt, &arrRes)
	if err != nil {
		return res, err
	}

	return arrRes[0], nil
}
