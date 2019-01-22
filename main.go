/*
**	Project Name:			HashTap
**	Project Authors:		Ryan McCommon, Aaron McCommon, Austen Stewart
**	Project Description:	HashTap was a proof of concept project made during Hack Arizona 2019.
**							The challenge was to create a good use case for Hedera Hashgraph's cryptocurrency.
**							Our idea was to create an application that simulated the use of a self serve
**							soda machine that charged the costumer per Fl.Oz. using the cryptocurrency.
**							This project was the winner of the challenge.
 */

package main

import (
	"flag"

	"encoding/json"

	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	astilog "github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "enables the debug mode")
	w       *astilectron.Window
	price   int
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug: *debug,

		//Make the file menu displayed on the top bar
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("File"),
			//In the file menu put an option called price
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astilectron.PtrStr("Price"),
					//Price displays the price of soda per fl oz
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if err := bootstrap.SendMessage(w, "getPrice", price, func(m *bootstrap.MessageIn) {
							var s string
							if err := json.Unmarshal(m.Payload, &s); err != nil {
								astilog.Error(errors.Wrap(err, "unmarshalling payload failed"))
								return
							}
							astilog.Infof("Price modal has been displayed and payload is %s!", s)
						}); err != nil {
							astilog.Error(errors.Wrap(err, "sending price event failed"))
						}
						return
					},
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		//This had an old function in it but now its used to store the price constant
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			price = 60000000
			return nil
		},
		RestoreAssets: RestoreAssets,
		//This sets up the window that everything is displayed in
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#333"),
				Center:          astilectron.PtrBool(true),
				Height:          astilectron.PtrInt(700),
				Width:           astilectron.PtrInt(700),
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
