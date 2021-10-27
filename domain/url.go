package domain

type UrlModel struct {
	ID          uint   `json:"id" gorm:"column:id;type:serial;not null;autoIncrement;primaryKey"`
	OriginalUrl string `json:"original_url" gorm:"column:original_url;type:text;not null"`
	Token       string `json:"token" gorm:"column:token;type:text;not null"`
	Views       int    `json:"views" gorm:"column:views;type:int(16);not null"`
}

type ShortenDto struct {
	UrlToShorten string `json:"url_to_shorten"`
}
