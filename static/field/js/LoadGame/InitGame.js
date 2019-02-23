function InitGame() {
    let idGame = getCookie("idGame");
    field.send(JSON.stringify({
        event: "InitGame",
        id_game: Number(idGame)
    }));
}

function getCookie(name) {
    let matches = document.cookie.match(new RegExp(
        "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
}