package service

import (
	"github.com/kyeeego/how-many-url-shorteners-are-there/domain"
	"github.com/kyeeego/how-many-url-shorteners-are-there/repository"
)

type UrlService interface {
	Shorten(url string) (domain.UrlModel, error)
	GetUrlByToken(token string) (domain.UrlModel, error)
	Update(model domain.UrlModel) error
}

type Service struct {
	UrlService UrlService
}

func New(repository *repository.Repository) *Service {
	return &Service{
		UrlService: newUrlService(repository),
	}
}