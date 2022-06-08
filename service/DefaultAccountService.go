package service

import (
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/domain"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/dto"
	"github.com/leminhviett/Go_micro_practice-utils_lib/errs"
)

type DefaultAccountService struct {
	repo domain.AccountRepo
}

func NewDefAccountService(repo domain.AccountRepo) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}

func (service DefaultAccountService) FindById(id string) (*dto.AccountResponseDTO, *errs.AppError) {
	domainAcc, err := service.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return domainAcc.ToDTO(), nil
}

func (service DefaultAccountService) FindAll() ([]*dto.AccountResponseDTO, *errs.AppError) {
	domainAccs, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	dtoAccResponse := make([]*dto.AccountResponseDTO, 0)

	for _, domainAcc := range domainAccs {
		dtoAccResponse = append(dtoAccResponse, domainAcc.ToDTO())
	}
	return dtoAccResponse, nil
}

func (service DefaultAccountService) InsertNewAccount(acc *dto.AccountRequestDTO) (*dto.AccountResponseDTO, *errs.AppError) {
	domainAcc := &domain.Account{AccountName: acc.AccountName, Ammount: acc.Ammount}
	domainAcc, err := service.repo.InsertNewAccount(domainAcc)

	if err != nil {
		return nil, err
	}
	return domainAcc.ToDTO(), nil
}
