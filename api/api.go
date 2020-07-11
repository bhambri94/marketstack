package api

import (
	"ShiftAlt/marketstack/configs"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type MarketstackResponse struct {
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Data struct {
		Name        string `json:"name"`
		Acronym     string `json:"acronym"`
		Mic         string `json:"mic"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		City        string `json:"city"`
		Website     string `json:"website"`
		Timezone    struct {
			Timezone string `json:"timezone"`
			Abbr     string `json:"abbr"`
			AbbrDst  string `json:"abbr_dst"`
		} `json:"timezone"`
		Currency struct {
			Code   string `json:"code"`
			Symbol string `json:"symbol"`
			Name   string `json:"name"`
		} `json:"currency"`
		Eod []struct {
			Open      float64 `json:"open"`
			High      float64 `json:"high"`
			Low       float64 `json:"low"`
			Close     float64 `json:"close"`
			Volume    float64 `json:"volume"`
			AdjHigh   float64 `json:"adj_high"`
			AdjLow    float64 `json:"adj_low"`
			AdjClose  float64 `json:"adj_close"`
			AdjOpen   float64 `json:"adj_open"`
			AdjVolume float64 `json:"adj_volume"`
			Symbol    string  `json:"symbol"`
			Exchange  string  `json:"exchange"`
			Date      string  `json:"date"`
		} `json:"eod"`
	} `json:"data"`
}

func GetMarketStackData() MarketstackResponse {
	req, err := http.NewRequest("GET", configs.MarketStackApiPath+configs.MarketStackApiAccessToken, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Cookie", "__cfduid=d43da601d5c117d53139732ee78bb60ad1594395724")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var marketstackResp MarketstackResponse
	json.Unmarshal(body, &marketstackResp)
	return marketstackResp
}
