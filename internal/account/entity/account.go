package entity

import (
	"time"

	common_entity "isling-be/internal/common/entity"
)

type Account struct {
	ID                common_entity.AccountID `json:"id" example:"1"`
	Email             string                  `json:"email" example:"isling@isling.me"`
	EncryptedPassword string
	CreatedAt         time.Time `json:"createdAt" example:"2022-12-12T12:12:12"`
	UpdatedAt         time.Time `json:"updatedAt" example:"2022-12-12T12:12:12"`
}

func NewAccount(
	id common_entity.AccountID,
	email,
	encryptedPassword string,
	createdAt, updatedAt time.Time,
) Account {
	return Account{
		ID:                id,
		Email:             email,
		EncryptedPassword: encryptedPassword,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
	}
}

type AccountWithoutPass struct {
	ID        common_entity.AccountID `json:"id" example:"1"`
	Email     string                  `json:"email" example:"isling@isling.me"`
	CreatedAt time.Time               `json:"createdAt" example:"2022-12-12T12:12:12"`
	UpdatedAt time.Time               `json:"updatedAt" example:"2022-12-12T12:12:12"`
}

func (a *Account) ToAccountWithoutPass() *AccountWithoutPass {
	return &AccountWithoutPass{
		ID:        a.ID,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
