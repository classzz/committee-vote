package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GasTracker struct {
	Code int `json:"code"`
	Data struct {
		Rapid     int64 `json:"rapid"`
		Fast      int64 `json:"fast"`
		Standard  int64 `json:"standard"`
		Slow      int64 `json:"slow"`
		Timestamp int64 `json:"timestamp"`
	} `json:"data"`
}

func GetGasTracker() (*GasTracker, error) {
	resp, err := http.Get("https://www.gasnow.org/api/v3/gas/price?utm_source=Classzz")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}

	var res *GasTracker
	_ = json.Unmarshal(body, &res)

	return res, nil
}
