package service

import (
	"echain-user-updater/cmd/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func UpdateWallet(wallets *[]model.UserWallet, db *sqlx.DB) (int, error) {
	count := 0
	tx, err := db.Beginx()
	if err != nil {
		fmt.Println("Beginx error:", err)
		panic(err)
	}

	stmt, err := tx.Preparex(db.Rebind("UPDATE integral_wallet_key set address = ?,mnemonic = ?, private_key = ?, create_time = ? where user_id = ?"))
	if err != nil {
		fmt.Println("Prepare error:", err)
		panic(err)
	}
	for _, wallet := range *wallets {
		_, err = stmt.Exec(wallet.Address, wallet.Mnemonic, wallet.PrivateKey, wallet.CreateTime, wallet.UserID)
		if err != nil {
			fmt.Println("Exec error:", err)
			panic(err)
		} else {
			count++
			if count%200 == 0 {
				fmt.Println(fmt.Sprintf("%d of %d wallets has been updated,", count, len(*wallets)))
			}
		}

	}
	err = stmt.Close()
	if err != nil {
		fmt.Println("stmt close error:", err)
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error:", err)
		panic(err)
	}

	return count, err
}
