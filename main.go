package main

import (
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/app"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/domain"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/service"
)

func main() {
	accHandler := app.NewAccountHandler(service.NewDefAccountService(domain.NewAccountRepoStub()))

	app.Start(accHandler)
}