package domain

type UrlModel struct {
	ID          uint
	OriginalUrl string
	Token       string
	Views       int
}

type ShortenDto struct {
	UrlToShorten string
}
