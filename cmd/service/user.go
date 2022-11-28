package service

import (
	"echain-user-updater/cmd/model"
	"github.com/jmoiron/sqlx"
)

func GetAllUser(db *sqlx.DB) (*[]model.User, error) {
	var users = make([]model.User, 0)
	db.Select(&users, "SELECT id FROM user_info where id != 1 limit 2")
	return &users, nil
}

func GetAllUserWithoutWallet(db *sqlx.DB) (*[]model.User, error) {
	var users = make([]model.User, 0)
	db.Select(&users, "select id from user_info where "+
		"(select count(1) from integral_wallet_key where user_info.ID = integral_wallet_key.user_id) = 0 and id != 1")
	return &users, nil
}
