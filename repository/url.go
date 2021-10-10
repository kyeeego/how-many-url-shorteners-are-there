package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyeeego/how-many-url-shorteners-are-there/domain"
)

type urlRepository struct {
	db *gorm.DB
}

func newUrlRepository(db *gorm.DB) *urlRepository {
	return &urlRepository{db}
}

func (r urlRepository) Insert(model *domain.UrlModel) (uint, error) {
	res := r.db.Create(&model)
	return model.ID, res.Error
}

func (r urlRepository) GetByToken(token string) (domain.UrlModel, error) {
	urlModel := domain.UrlModel{}
	res := r.db.First(&urlModel, domain.UrlModel{Token: token})

	return urlModel, res.Error
}

func (r urlRepository) GetByOriginalUrl(url string) (domain.UrlModel, error) {
	urlModel := domain.UrlModel{}
	res := r.db.First(&urlModel, domain.UrlModel{OriginalUrl: url})

	return urlModel, res.Error
}

func (r urlRepository) Update(model *domain.UrlModel) error {
	res := r.db.Save(*model)
	return res.Error
}
