package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/BFDavidGamboa/bookstore_oauth-api/src/domain/users"
	"github.com/BFDavidGamboa/bookstore_utils-go/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {

	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	bytes, _ := json.Marshal(request)
	fmt.Println(string(bytes))
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		err := errors.New("Could not receive valid response from client")
		return nil, rest_errors.NewInternalServerError("invalid restClient response when trying to login user ", err)
	}
	if response.StatusCode > 299 {
		fmt.Println(response.String())
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshall users response", err)
	}
	return &user, nil
}
