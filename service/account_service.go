package service

import (
	"github.com/Pan9Hu/api-server_v2/model"
)

type AccountService struct {
}

var (
	accountModel model.Account
)

func (ac *AccountService) GetAccountByUID(uid string) {

}

func (ac *AccountService) ListAccount(name, group, phone, email, sex, archGroup string) ([]struct{}, error) {
	if name == "" && group == "" && phone == "" && email == "" && sex == "" && archGroup == "" {
		return accountModel.AllAccount(), nil
	}

	return []struct{}{}, nil
}

func (ac *AccountService) CreateAccount(uid string) {

}

func (ac *AccountService) DeleteAccountByUID(uid string) {

}

func (ac *AccountService) DeleteAccount(uid string) {

}

func (ac *AccountService) UpdateAccount(uid string) {

}
