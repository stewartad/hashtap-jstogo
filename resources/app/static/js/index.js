let time = 0.0;
let start, end;
let index = {
    about: function(html) {
        let c = document.createElement("div");
        c.innerHTML = html;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    getPrice: function(price) {
        let c = document.createElement("div");
        c.innerHTML = "Current price for soda: " + price;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    addFolder(name, path) {
        let div = document.createElement("div");
        div.className = "dir";
        div.onclick = function() { index.explore(path) };
        div.innerHTML = `<i class="fa fa-folder"></i><span>` + name + `</span>`;
        document.getElementById("dirs").appendChild(div)
    },
    holdTap() {
        start = new Date();
    },
    releaseTap() {
        end = new Date();
        let time = end - start;
        let message = {"name": "tap", "payload": time};
        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            asticode.loader.hide();
            document.getElementById("cost").innerHTML = `<h1>Cost: ` + `   ` + message.payload.cost + `</h1><br /><h1>Time: ` + time + `</h1><br /><h1>USD: ` + message.payload.usd + `</h1>`
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
            index.explore();
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

            // Process path
            document.getElementById("path").innerHTML = message.payload.path;

            // Process dirs
            document.getElementById("dirs").innerHTML = ""
            for (let i = 0; i < message.payload.dirs.length; i++) {
                index.addFolder(message.payload.dirs[i].name, message.payload.dirs[i].path);
            }

            // Process files
            document.getElementById("files_count").innerHTML = message.payload.files_count;
            //document.getElementById("files_size").innerHTML = message.payload.files_size;
            document.getElementById("cost").innerHTML = `<h1>WTF</h1>` + message.payload;
            document.getElementById("files").innerHTML = ""; 
            if (typeof message.payload.files !== "undefined") {
                document.getElementById("files_panel").style.display = "block";
                let canvas = document.createElement("canvas");
                document.getElementById("files").append(canvas);
                new Chart(canvas, message.payload.files);
            } else {
                document.getElementById("files_panel").style.display = "none";
            }
        })
    },
    listen: function() {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    index.about(message.payload);
                    return {payload: "payload"};
                    break;
                case "getPrice":
                    index.getPrice(message.payload);
                    return {payload: "payload"};
                    break;
                case "check.out.menu":
                    asticode.notifier.info(message.payload);
                    break;
            }
        });
    }
};