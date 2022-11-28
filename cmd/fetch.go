package cmd

import (
	"bytes"
	"echain-user-updater/cmd/app"
	"echain-user-updater/config"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(eChainConfig config.Chain) (*app.Response, error) {

	reqBody, err := json.Marshal(app.Request{
		BlockAddress: eChainConfig.BlockAddress,
	})
	if err != nil {
		return nil, errors.New("bad request format")
	}

	req, err := http.NewRequest(eChainConfig.Method, eChainConfig.Url, bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, errors.New("request error")
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("response with error: %s", err.Error()))
	}

	res := app.Response{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("parse response failed with error: %s", err.Error()))
	}

	defer response.Body.Close()

	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("response body format error: %s", err.Error()))
	}

	return &res, nil

}
