package cmd

import (
	"echain-user-updater/cmd/app"
	"echain-user-updater/global"
	"echain-user-updater/helper"
	"fmt"
	"log"
)

func Build(amount int) *app.Results {
	results := app.NewResults()
	writeIn := app.NewResults()
	total := 0
	rolling := 0

	count, response, err := helper.ReadInData()
	if err != nil {
		log.Println(fmt.Sprintf("error occured when read in csv file: %s", err.Error()))
	}

	results.Res = *response

	for {
		total++
		res, err := Fetch(global.App.Config.Chain)

		if err != nil {
			log.Print(err)
		}

		if res.Result == 1 {
			count++
			rolling++
			results.Res = append(results.Res, *res)
			writeIn.Res = append(writeIn.Res, *res)
		}
		if count == amount {
			break
		}

		if count%100 == 0 {
			log.Println(fmt.Sprintf("%d wallet records has been built", count))
		}

		if rolling%200 == 0 {
			err = helper.WriteData(writeIn)
			if err != nil {
				log.Println(fmt.Sprintf("error occured when write in csv file: %s", err.Error()))
			} else {
				writeIn = app.NewResults()
			}
		}
	}

	fmt.Println(fmt.Sprintf("successful in %d request: %d", total, count))

	return results
}
