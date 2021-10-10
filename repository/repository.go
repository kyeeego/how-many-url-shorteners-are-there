package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyeeego/how-many-url-shorteners-are-there/domain"
)

type UrlRepository interface {
	Insert(model *domain.UrlModel) (uint, error)
	GetByToken(token string) (domain.UrlModel, error)
	Update(model *domain.UrlModel) error
	GetByOriginalUrl(url string) (domain.UrlModel, error)
}

type Repository struct {
	UrlRepository UrlRepository
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		UrlRepository: newUrlRepository(db),
	}
}
