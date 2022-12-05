package access_token

import (
	"strings"

	"github.com/BFDavidGamboa/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(acessTokenId string) (*AccessToken, *errors.RestErr) {
	acessTokenId = strings.TrimSpace(acessTokenId)
	if len(acessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(acessTokenId)

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
