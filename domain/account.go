package domain

import (
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/dto"
	"github.com/leminhviett/Go_micro_practice-utils_lib/errs"
)

type Account struct{
	AccountId string
	AccountName string
	Ammount float64
}

func (acc *Account) ToDTO() *dto.AccountResponseDTO {
	return &dto.AccountResponseDTO{
		AccountId: acc.AccountId,
		AccountName: acc.AccountName,
		Ammount: acc.Ammount,
	}
}

//go:generate mockgen -destination=../test/mocks/domain/mockAccountRepo.go -package=mock_repo github.com/leminhviett/Go_micro_practice-HTTP_api_server/domain AccountRepo
type AccountRepo interface {
	FindById(string) (*Account, *errs.AppError)
	FindAll() ([]Account, *errs.AppError)
	InsertNewAccount(*Account) (*Account, *errs.AppError)
}
