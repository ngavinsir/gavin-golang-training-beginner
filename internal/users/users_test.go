package users_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/ngavinsir/golangtraining"
	"github.com/ngavinsir/golangtraining/internal/users"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}

	testCases := []struct {
		desc           string
		client         *users.Client
		expectedResult golangtraining.User
	}{
		{
			desc: "success",
			client: func() *users.Client {
				return users.NewUsersClient(httpClient)
			}(),
			expectedResult: golangtraining.User{
				ID:   1,
				Name: "Leanne Graham",
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res, err := tC.client.GetUsers(context.TODO())

			require.Equal(t, nil, err)
			require.Equal(t, tC.expectedResult, res)
		})
	}
}
