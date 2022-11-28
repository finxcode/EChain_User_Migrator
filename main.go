package main

import (
	"echain-user-updater/bootstrap"
	"echain-user-updater/cmd/controller"
	"echain-user-updater/global"
	"fmt"
	"log"
)

func main() {
	bootstrap.InitConfig()
	fmt.Println(global.App.Config.Chain.BlockAddress)
	fmt.Println(global.App.Config.Database.Database)

	global.App.DB = bootstrap.InitDb()

	defer global.App.DB.Close()

	num, err := controller.UpdateWallet(global.App.DB)
	if err != nil {
		log.Fatal(fmt.Sprintf("update failed with error: %s", err.Error()))
	}

	log.Println(fmt.Sprintf("%d user records has been updated sucessfully", num))

	//resluts := cmd.Results{
	//	Res: []cmd.Response{
	//		{
	//			ContractAddress: "dasdsdsd",
	//			PrivateKey:      "ttrtertret",
	//			PubKey:          "czxczxczxc",
	//		},
	//		{
	//			ContractAddress: "qqqqqq",
	//			PrivateKey:      "rrrrrrrr",
	//			PubKey:          "eeeeeee",
	//		},
	//	},
	//}
	//
	//err := helper.WriteData(&resluts)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//n, res, _ := helper.ReadInData()
	//fmt.Println(fmt.Sprintf("readin %d records", n))
	//
	//for _, r := range res.Res {
	//	fmt.Println(fmt.Sprintf("c: %s prv: %s pub: %s", r.ContractAddress, r.PrivateKey, r.PubKey))
	//}

}
