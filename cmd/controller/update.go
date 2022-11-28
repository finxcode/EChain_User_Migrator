package controller

import (
	"echain-user-updater/cmd/service"
	"echain-user-updater/cmd/wrapper"
	"github.com/jmoiron/sqlx"
)

func UpdateWallet(db *sqlx.DB) (int, error) {
	wallets := wrapper.Wrap(db)
	return service.UpdateWallet(wallets, db)
}
