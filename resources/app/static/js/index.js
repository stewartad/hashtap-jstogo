/* 
**  All the Javascript functions are in this file.
**  The most important ones are the ones that communicate with the go backend.
*/

let time = 0.0;
let start, end;
let index = {
    //Gets the price of soda that is coded in main.go
    getPrice: function(price) {
        let c = document.createElement("div");
        c.innerHTML = "Current price for soda: " + price + " Tinybars per second";
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },

    //This function was supposed to get the account balances at load time but it doesn't work
    getBal: function() {
        let message = {"name": "onload"};
        astilectron.sendMessage(message, function(message) {
            document.getElementById("bal").innerHTML = `<h1>Customer Balance: ` + `   ` + message.payload.cbal + `</h1><br /><h1>Business Balance: ` + message.payload.bbal + `</h1>`;
        })
    },

    //Makes a new date when the button was just pressed
    holdTap: function() {
        start = new Date();
    },

    //calculates the time the button was held down and calls go functions to complete the transaction.
    releaseTap: function() {
        end = new Date();
        let time = end - start;
        let message = {"name": "tap", "payload": time};
        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            asticode.loader.hide();
            document.getElementById("cost").innerHTML = `<h1>Cost: ` + `   ` + message.payload.cost + ` Tinybars</h1>`;
            document.getElementById("time").innerHTML   = `<h1>Time: ` + time + ` ms</h1>`;
            document.getElementById("usd").innerHTML = `<h1>USD: $` + message.payload.usd + `</h1>`;
            document.getElementById("cBal").innerHTML =  `<h1>Customer Balance: ` + `   ` + message.payload.cbal + ` Hbar</h1>`;
            document.getElementById("bBal").innerHTML = `<h1>Business Balance: ` + message.payload.bbal + ` Hbar</h1>`;
        })
        index.explore();
        
    },


    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();

            

            // Explore default path
            index.explore(0);
        })
        document.getElementById("tap-button").addEventListener('mousedown', function() {
            index.holdTap();
        })
        document.getElementById("tap-button").addEventListener('mouseup', function() {
            index.releaseTap();
        })
    },
    explore: function(path) {
        // Create message
        let message = {"name": "explore"};
        if (typeof path !== "undefined") {
            message.payload = path
        }

        // Send message
        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            // Init
            asticode.loader.hide();

            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }

            // Process costs
            document.getElementById("cost").innerHTML = `<h1>WTF</h1>` + message.payload;
        })
    },
    listen: function() {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "getPrice":
                    index.getPrice(message.payload);
                    return {payload: "payload"};
                    break;
            }
        });
    }
};