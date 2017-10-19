var idGame;

function InitGame() {
    idGame = getCookie("idGame");
    sendInitGame(idGame);
}

function getCookie(name) {
    var matches = document.cookie.match(new RegExp(
        "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
}

function Field(xSize,ySize) {
    var main = document.getElementById("main");
    main.style.boxShadow = "25px 25px 20px rgba(0,0,0,0.5)";

    for (var y = 0; y < ySize; y++) {
        for (var x = 0; x < xSize; x++) {
            var div = document.createElement('div');
            div.className = "fieldUnit";
            div.id = x + ":" + y;
            div.innerHTML = x + ":" + y;
            div.onclick = function () {
                reply_click(this.id);
            };
            div.onmouseover = function () {
                mouse_over(this.id);
            };
            div.onmouseout = function () {
                mouse_out(this.id)
            };
            main.appendChild(div);
        }
        var nline = document.createElement('div');
        nline.className = "nline";
        nline.innerHTML = "";
        main.appendChild(nline);
    }
}

function sendInitGame(idGame) {
    sock.send(JSON.stringify({
        event: "InitGame",
        id_game: idGame
    }));
}