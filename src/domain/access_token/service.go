package access_token

import (
	"strings"

	"github.com/BFDavidGamboa/bookstore_utils-go/rest_errors"
)

type Repository interface {
	GetById(string) (*AccessToken, rest_errors.RestErr)
	Create(AccessToken) rest_errors.RestErr
	UpdateExpirationTime(AccessToken) rest_errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, rest_errors.RestErr)
	Create(AccessToken) rest_errors.RestErr
	UpdateExpirationTime(AccessToken) rest_errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(acessTokenId string) (*AccessToken, rest_errors.RestErr) {
	acessTokenId = strings.TrimSpace(acessTokenId)
	if len(acessTokenId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(acessTokenId)

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) rest_errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) rest_errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
