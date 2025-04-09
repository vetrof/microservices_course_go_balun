package main

import (
	"bytes"
	"encoding/json"
	"github.com/fatih/color"
	"log"
	"net/http"
)

type Ip struct {
	IP string `json:"ip"`
}

type IPInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func ipInfoClient() (IPInfo, error) {

	ip := Ip{IP: "195.189.71.234"}

	data, err := json.Marshal(ip)
	if err != nil {
		return IPInfo{}, err
	}

	resp, err := http.Post("http://ipinfo.io", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return IPInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return IPInfo{}, err
	}

	var createInfo IPInfo
	if err = json.NewDecoder(resp.Body).Decode(&createInfo); err != nil {
		return IPInfo{}, err
	}

	return createInfo, nil

}

func main() {

	info, err := ipInfoClient()
	if err != nil {
		log.Fatal("failed to create note:", err)
	}

	log.Printf(color.RedString("Info = "), color.GreenString("%+v", info))

}
