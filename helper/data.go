package helper

import (
	"bufio"
	"echain-user-updater/cmd/app"
	"echain-user-updater/global"
	"encoding/csv"
	"os"
	"strings"
)

func ReadInData() (int, *[]app.Response, error) {
	file, err := os.Open(getFilePath())
	if err != nil {
		return -1, nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	res := app.Results{}

	for fileScanner.Scan() {
		content := strings.Split(fileScanner.Text(), ",")
		if len(content) < 3 {
			continue
		}
		record := app.Response{
			ContractAddress: content[0],
			PrivateKey:      content[1],
			PubKey:          content[2],
		}
		res.Res = append(res.Res, record)
	}

	return len(res.Res), &res.Res, nil
}

func WriteData(results *app.Results) error {
	file, err := os.OpenFile(getFilePath(), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	records := buildRecords(results)

	w := csv.NewWriter(file)

	err = w.WriteAll(*records)

	if err != nil {
		return err
	}
	return nil
}

func IsFileExist(path, filename string) bool {
	if _, err := os.Stat(path + filename); err == nil {
		return true
	}
	return false
}

func getFilePath() string {
	return global.App.Config.Persist.Loc + global.App.Config.Persist.Filename
}

func buildRecords(results *app.Results) *[][]string {
	records := make([][]string, 0)
	for _, res := range results.Res {
		record := make([]string, 0)
		record = append(record, res.ContractAddress)
		record = append(record, res.PrivateKey)
		record = append(record, res.PubKey)
		records = append(records, record)
	}
	return &records
}
