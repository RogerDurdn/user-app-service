package main

import (
	"github.com/RogerDurdn/users/domain"
	"github.com/RogerDurdn/users/pkg"
	"github.com/RogerDurdn/users/rest"
)

func main() {
	storage := pkg.NewPostgresStorage()
	service := domain.NewNormalSrv(storage)
	server := rest.NewRest(service, "9092")
	server.Start()
}
