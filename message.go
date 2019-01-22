/*
**	This handels all the messages that come from or are sent too the javascipt in the application.
**
**  This file was created originally to demo astilectron a library that connected go to electron js.
**	It was rewritten for the purpose of making an app that simulated a self serve soda dispenser that 
**  uses the Hedera Hashgraph cryptocurrency for making transactions
*/

package main

import (
	"encoding/json"
	"time"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/hashgraph/hedera-sdk-go"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {

	//this is when they hold the button
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
	Cost	float64	`json:"cost"`
}

// Drink represents a liquid to be dispensed
type Drink struct {
	name	string	`json:"name"`
	price	float64	`json:"price"`
	flow	float64	`json:"flow"`
}


// Exploration represents a transaction showing both costumer and business balance
type Exploration struct {
	Cost		float64				`json:"cost"`
	Usd			float64				`json:"usd"`
	C_Balance	float64			`json:"cbal"`
	B_Balance	float64				`json:"bbal"`
}

// explore does the actual transaction and returns a stuct with the ending balances of the costumer and business
func explore(timeD int64) (e Exploration, err error) {

	//init drink2
	var soda = Drink{
		name: "MtnDew",
		price: 60000000,
		flow: 1,
	}
	
	

	// cost is calculated by time multiplied by price since the flow is one floz/sec
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

