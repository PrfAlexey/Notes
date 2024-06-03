package services

import "github.com/users_service/pkg/apiclient/v1/client"

type UsersService struct {
	client *client.UsersServiceAPI
}
