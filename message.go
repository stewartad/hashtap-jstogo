package main

import (
	"encoding/json"
	"time"

	"github.com/asticode/go-astichartjs"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/hashgraph/hedera-sdk-go"
)




// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	// case "explore":
	// 	// Unmarshal payload
	// 	var path string
	// 	if len(m.Payload) > 0 {
	// 		// Unmarshal payload
	// 		if err = json.Unmarshal(m.Payload, &path); err != nil {
	// 			payload = err.Error()
	// 			return
	// 		}
	// 	}

	// 	// Explore
	// 	if payload, err = explore(path); err != nil {
	// 		payload = err.Error()
	// 		return
	// 	}




	// case "onload":
	// 	if payload, err = explore(); err != nil {
	// 		payload = err.Error()
	// 		return
	// 	}
	case "tap":
		// Unmarshall
		//var time int64
		var timeD int64
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &timeD); err != nil {
				payload = err.Error()
				return
			}
		}

		if payload, err = explore(timeD); err != nil {
			payload = err.Error()
			return
		}
		
	}
	return
}

// Cost gives all hedera values
type Cost struct {
	cost	float64	`json:"cost"`
}

// Drink represents a liquid to be dispensed
type Drink struct {
	name	string	`json:"name"`
	price	float64	`json:"price"`
	flow	float64	`json:"flow"`
}

// Transaction will store info to send to the Hedera network
type Transaction struct {
	drink		string
	amount		float64
}

// Exploration represents the results of an exploration
type Exploration struct {
	Cost		float64				`json:"cost"`
	Usd			float64				`json:"usd"`
	C_Balance	float64			`json:"cbal"`
	B_Balance	float64				`json:"bbal"`
	Dirs       []Dir              `json:"dirs"`
	Files      *astichartjs.Chart `json:"files,omitempty"`
	FilesCount int                `json:"files_count"`
	FilesSize  string             `json:"files_size"`
	Path       string             `json:"path"`
}

// Balance does stuff
type Balance struct {
	C_Balance	float64			`json:"cbal"`
	B_Balance	float64				`json:"bbal"`
}

// PayloadDir represents a dir payload
type Dir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// func calCost(time int64) (c Cost) {
// 	// init drink1
// 	var MtnDew = Drink{
// 		name: "Mtn Dew",
// 		price: 0.04,
// 		flow: 0.1,
// 	}

// 	c = Cost{
// 		cost: (float64(time) / 1000.0) * MtnDew.price,
// 	}
// 	return
// }

// explore explores a path.
// If path is empty, it explores the user's home directory
func explore(timeD int64) (e Exploration, err error) {

	//init drink2
	var soda = Drink{
		name: "MtnDew",
		price: 60000000,
		flow: 1,
	}
	
	


	cost := (float64(timeD) / 1000) * soda.price

	transferAmount(hedera.AccountID{Account: 1001}, hedera.AccountID{Account: 1002}, int64(cost))

	time.Sleep(2*time.Second)

	customerBal := getAccountBal(hedera.AccountID{Account: 1001})
	bizBal := getAccountBal(hedera.AccountID{Account: 1002})
	usd := cost * 0.1 / 100000000
	// Init exploration
	e = Exploration{
		Cost: cost,
		Usd: usd,
		C_Balance: customerBal,
		B_Balance: bizBal,
		//Dirs: []Dir{},
		//Path: path,
	}
	if err := bootstrap.SendMessage(w, "check.out.menu", "Transaction Success"); err != nil {
		
	}
	return
}

// func getBalance() (b Exploration, err error) {
// 	//Customer
	
// 	// Init exploration
// 	b = Exploration{
		
// 		//Dirs: []Dir{},
// 		//Path: path,
// 	}

// 	return
// }