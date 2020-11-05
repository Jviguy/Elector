package Elector
import (
	"encoding/json"
	"github.com/gojektech/heimdall/v6/httpclient"
	"io/ioutil"
	"time"
)

type Client struct {
	http *httpclient.Client
}

type Result struct {
	President struct {
		CalledWinner bool `json:"called-winner"`
		CalledDemocrats int `json:"called-d"`
		CalledRepublican int `json:"called-r"`
		LeanDemocrat int `json:"lean-d"`
		LeanRepublican int `json:"lean-r"`
		LikelyDemocrat int `json:"likely-d"`
		LikelyRepublican int `json:"likely-r"`
		SolidDemocrat int `json:"solid-d"`
		SolidRepublican int `json:"sold-r"`
		Tossup int `json:"toss-up"`
	} `json:"president"`
	Senate struct {
		CalledWinner bool `json:"called-winner"`
		SeatedDemocrats int `json:"seated-d"`
		SeatedRepublicans int `json:"seated-r"`
		CalledDemocrats int `json:"called-d"`
		CalledRepublican int `json:"called-r"`
		LeanDemocrat int `json:"lean-d"`
		LeanRepublican int `json:"lean-r"`
		LikelyDemocrat int `json:"likely-d"`
		LikelyRepublican int `json:"likely-r"`
		SolidDemocrat int `json:"solid-d"`
		SolidRepublican int `json:"sold-r"`
		Tossup int `json:"toss-up"`
	} `json:"senate"`
	House struct {
		CalledWinner string `json:"called-winner"`
		CalledDemocrats int `json:"called-d"`
		CalledRepublican int `json:"called-r"`
		LeanDemocrat int `json:"lean-d"`
		LeanRepublican int `json:"lean-r"`
		LikelyDemocrat int `json:"likely-d"`
		LikelyRepublican int `json:"likely-r"`
		SolidDemocrat int `json:"solid-d"`
		SolidRepublican int `json:"sold-r"`
		Tossup int `json:"toss-up"`
	} `json:"house"`
	LastUpdated string `json:"lastUpdated"`
}

func New(timeout time.Duration) *Client {
	return &Client{http: httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))}
}

func (c Client) GetData(year int) (*Result,error) {
	res,err := c.http.Get("https://www.politico.com/2020-national-results/balance-of-power.json",nil)
	if err != nil{
		return nil,err
	}
	body, err := ioutil.ReadAll(res.Body)
	data := &Result{}
	err = json.Unmarshal(body,data)
	return data , err
}