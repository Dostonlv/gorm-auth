package main

import (
	"github.com/Dostonlv/gorm-auth/api"
	"github.com/Dostonlv/gorm-auth/config"
	"github.com/Dostonlv/gorm-auth/storage"
	"github.com/Dostonlv/gorm-auth/storage/postgres"
	"github.com/gin-gonic/gin"
)

func init() {
	postgres.NewPostgres(config.Load())
	storage.SyncDatabase()

}

func main() {
	cfg := config.Load()
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	api.SetUpAPI(r, cfg)
	r.Run(config.Load().ServicePort)
}
