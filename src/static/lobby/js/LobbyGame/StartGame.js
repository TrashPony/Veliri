function StartNewGame(jsonMessage) {
    if (JSON.parse(jsonMessage).error === "") {
        toField = true;
        var idGame = JSON.parse(jsonMessage).id_game;
        document.cookie = "idGame=" + idGame + "; path=/;";
        location.href = "http://" + window.location.host + "/field";
    } else {
        if (JSON.parse(jsonMessage).error === "Players < 2") {
            alert("Ошибка: Мало игроков для старта");
        }
        if (JSON.parse(jsonMessage).error === "error ad to DB") {
            alert("Неизвестная ошибка");
        }
        if (JSON.parse(jsonMessage).error === "PlayerNotReady") {
            alert("Ошибка: не все игроки готовы");
        }
    }
}