package domain

import "github.com/leminhviett/Go_micro_practice-utils_lib/errs"

type AccountRepoStub struct {

}

func NewAccountRepoStub() AccountRepoStub{
	return AccountRepoStub{}
}

func (repo AccountRepoStub) FindAll() ([]Account, *errs.AppError) {
	return []Account{
		{"1", "VietLe", 1000},
		{"2", "VietLe2", 1000},
		{"3", "VietLe3", 1000},

	}, nil
}

func (repo AccountRepoStub) FindById(id string) (*Account, *errs.AppError) {
	return &Account{"1", "VietLe", 1000}, nil
}

func (repo AccountRepoStub) InsertNewAccount(acc *Account) (*Account, *errs.AppError) {
	return acc, nil
}