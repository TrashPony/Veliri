function FieldCreate() {
    Crafty.init();

    var hex = Crafty.hexametric.init(64, 64, 10, 10);

    Crafty.sprite(64, "test/images/tile-1.png", {
        grass: [0, 0, 1, 1]
    });

    Crafty.sprite(1200, "test/images/s135.png", {
        tank: [0, 0, 1, 1]
    });

    for (var x = 25; x > 0; x--) {
        for (var y = 25; y > 0; y--) {
            var tile = Crafty.e("2D, DOM, " + "grass" + ", Mouse")
                .attr('z', x + y + 1)
                .bind("Click", function (e) {

                }).bind("MouseOver", function () {
                    this.css({'filter': 'brightness(50%)'});
                }).bind("MouseOut", function () {
                    this.css({'filter': 'brightness(100%)'});
                });

            hex.place(tile, x, y, 1);
        }
    }


}

function ConnectField() {
    field = new WebSocket("ws://" + window.location.host + "/wsField");
    console.log("Websocket field - status: " + field.readyState);

    field.onopen = function() {
        console.log("CONNECTION field opened..." + this.readyState);
        InitGame();
    };

    field.onmessage = function(msg) {
        console.log("message: " + msg.data);
        ReadResponse(msg.data);
    };

    field.onerror = function(msg) {
        console.log("Error field occured sending..." + msg.data);
    };

    field.onclose = function(msg) {
        // 1006 ошибка при выключение сервера или отказа, 1001 - F5
        console.log("Disconnected field - status " + this.readyState);
        if (msg.code !== 1001) {
            location.href = "../../login";
        }
    };
}

function InitGame() {
    idGame = getCookie("idGame");
    field.send(JSON.stringify({
        event: "InitGame",
        id_game: Number(idGame)
    }));
}

function getCookie(name) {
    var matches = document.cookie.match(new RegExp(
        "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
}

function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    if (event === "InitMap") {
        FieldCreate(jsonMessage)
    }
}