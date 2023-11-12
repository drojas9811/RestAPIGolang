package models

import (
	"RestAPIGolang/internal/auth"
	"math/rand"
	"time"
)

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *Account) NewAccount(firstName, lastName, password string) error {
	hashedPassword, err := auth.EncryptPassword(password)
	if err != nil {
		return err
	}
	a.FirstName = firstName
	a.LastName = lastName
	a.EncryptedPassword = hashedPassword
	a.Number = int64(rand.Intn(1000000))
	a.CreatedAt = time.Now().UTC()
	return nil
}

func (a *Account) ValidPassword(pw string) bool {
	return auth.ValidatePassword(pw, a.EncryptedPassword) == nil
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}
