package wrapper

import (
	"echain-user-updater/cmd"
	"echain-user-updater/cmd/model"
	"echain-user-updater/cmd/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func Wrap(db *sqlx.DB) *[]model.UserWallet {
	users, err := service.GetAllUser(db)
	log.Println(fmt.Sprintf("total number of user needed to update: %d", len(*users)))
	var wallet = make([]model.UserWallet, 0)
	if err != nil {
		log.Fatal("error occurred when wrap user wallet: %d", err.Error())
	}
	results := cmd.Build(len(*users))
	for idx, val := range results.Res {
		w := model.UserWallet{}
		w.SetUserId((*users)[idx])
		w.SetKeys(val)
		w.CreateTime = time.Now()
		wallet = append(wallet, w)
	}

	log.Println(fmt.Sprintf("build %d user wallet", len(*users)))

	return &wallet
}
