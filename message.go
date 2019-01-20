package main

import (
	"encoding/json"
	// "io/ioutil"
	// "os"
	// "os/user"
	// "path/filepath"
	// "sort"
	//"strconv"

	"github.com/asticode/go-astichartjs"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
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
	case "tap":
		// Unmarshall
		//var time int64
		var time int64
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &time); err != nil {
				payload = err.Error()
				return
			}
		}

		if payload, err = explore(time); err != nil {
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
	Dirs       []Dir              `json:"dirs"`
	Files      *astichartjs.Chart `json:"files,omitempty"`
	FilesCount int                `json:"files_count"`
	FilesSize  string             `json:"files_size"`
	Path       string             `json:"path"`
}

// PayloadDir represents a dir payload
type Dir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func calCost(time int64) (c Cost) {
	// init drink1
	var MtnDew = Drink{
		name: "Mtn Dew",
		price: 0.04,
		flow: 0.1,
	}

	c = Cost{
		cost: (float64(time) / 1000.0) * MtnDew.price,
	}
	return
}

// explore explores a path.
// If path is empty, it explores the user's home directory
func explore(time int64) (e Exploration, err error) {

	//init drink2
	var soda = Drink{
		name: "MtnDew",
		price: 600000000,
		flow: 1,
	}


	cost := (float64(time) / 1000) * soda.price
	usd := cost * 0.1 / 1000000000
	// Init exploration
	e = Exploration{
		Cost: cost,
		Usd: usd,
		//Dirs: []Dir{},
		//Path: path,
	}





	return
}
