package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kyeeego/how-many-url-shorteners-are-there/config"
	"github.com/kyeeego/how-many-url-shorteners-are-there/delivery/http"
	"github.com/kyeeego/how-many-url-shorteners-are-there/repository"
	"github.com/kyeeego/how-many-url-shorteners-are-there/server"
	"github.com/kyeeego/how-many-url-shorteners-are-there/service"
	"log"
	"net/url"
)

func main() {

	conf, err := config.Init()
	if err != nil {
		log.Fatalf("Could not load config: %e\n", err)
		return
	}

	dsn := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(conf.DBUsername, conf.DBPassword),
		Host:     fmt.Sprintf("%s:%s", conf.DBHost, conf.DBPort),
		Path:     conf.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		log.Fatalf("Could not connect to database: %e\n", err)
	}

	repos := repository.New(db)
	services := service.New(repos)
	handler := http.New(services)

	s := &server.Server{}

	if err := s.Run(8081, handler.Init()); err != nil {
		log.Fatalf("Got an error while trying to start the server: %e\n", err)
	} else {
		log.Printf("Server is active on port :%d\n", 8081)
	}
}
