package model

import (
	"echain-user-updater/cmd/app"
	"time"
)

type UserWallet struct {
	UserID     int       `db:"user_id"`
	Address    string    `db:"address"`
	Mnemonic   string    `db:"mnemonic"`
	PrivateKey string    `db:"private_key"`
	CreateTime time.Time `db:"create_time"`
}

func (u *UserWallet) SetUserId(user User) {
	u.UserID = user.ID
}
func (u *UserWallet) SetKeys(response app.Response) {
	u.Address = response.ContractAddress
	u.Mnemonic = response.PubKey
	u.PrivateKey = response.PrivateKey
}
