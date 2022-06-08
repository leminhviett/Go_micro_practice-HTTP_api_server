package service

import (
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/dto"
	"github.com/leminhviett/Go_micro_practice-utils_lib/errs"
)

//go:generate mockgen -destination=../test/mocks/service/mockAccountService.go -package=mock_service github.com/leminhviett/Go_micro_practice-HTTP_api_server/service AccountService
type AccountService interface {
	FindById(string) (*dto.AccountResponseDTO, *errs.AppError)
	FindAll()([]*dto.AccountResponseDTO, *errs.AppError)
	InsertNewAccount(*dto.AccountRequestDTO) (*dto.AccountResponseDTO, *errs.AppError)
}
