package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var input = `
{
    "address": "823YkiiaTwit1hBEfnfVUfgaQy7fqm28GezYgtYt3e1s",
    "status": "inactive",
    "commission": 6,
    "activationDate": "2022-10-29"
}{
    "address": "4EPmCGjHeTVaeK1cbXdGekRhUvUGGBRH7usgQKFL2fuV",
    "status": "active",
    "commission": 6,
    "activationDate": "2022-11-24"
}
`

type DelegatorStatus struct {
	Address        string  `json:"address"`
	Status         string  `json:"status"` // status enum: active, inactive
	Commission     float64 `json:"commission,omitempty"`
	ActivationDate string  `json:"activationDate,omitempty"`
}

func main() {
	client := &http.Client{}

	url := "https://svc.blockdaemon.com/reporting/staking/v1/solana/mainnet/delegator/status/6AzhJQqax85X43PPPrk6WdLSVFmT4zjmofQr4StusY6A"

	req, _ := http.NewRequest("GET", url, nil)

	// set header
	req.Header.Set("Authorization", "Bearer YOUR_TOKEN")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Body)
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)

	delegatorStatus := make([]DelegatorStatus, 0)

	for {
		var degStats DelegatorStatus

		err := dec.Decode(&degStats)
		if err == io.EOF {
			// all done
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", degStats)
		// append degstats to slice
		delegatorStatus = append(delegatorStatus, degStats)
	}

	fmt.Println("ini degstatus", delegatorStatus)
}
