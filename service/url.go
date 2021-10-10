package service

import (
	"github.com/kyeeego/how-many-url-shorteners-are-there/domain"
	"github.com/kyeeego/how-many-url-shorteners-are-there/repository"
	"github.com/kyeeego/how-many-url-shorteners-are-there/utils"
)

type urlService struct {
	repository *repository.Repository
}

func newUrlService(repo *repository.Repository) *urlService {
	return &urlService{repo}
}

func (s urlService) GetUrlByToken(token string) (domain.UrlModel, error) {
	return s.repository.UrlRepository.GetByToken(token)
}

func (s urlService) Shorten(url string) (domain.UrlModel, error) {

	m, err := s.repository.UrlRepository.GetByOriginalUrl(url)
	if err == nil {
		return m, nil
	}

	model := domain.UrlModel{
		OriginalUrl: url,
		Token:       utils.RandomString(8),
		Views:       0,
	}

	id, err := s.repository.UrlRepository.Insert(&model)
	model.ID = id

	return model, err
}

func (s urlService) Update(model domain.UrlModel) error {
	return s.repository.UrlRepository.Update(&model)
}
