function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    if (event === "Join") {
        MyId = JSON.parse(jsonMessage).user_name;

        if (players[JSON.parse(jsonMessage).user_name] == null) {
            addPlayer(JSON.parse(jsonMessage).user_name, JSON.parse(jsonMessage).x, JSON.parse(jsonMessage).y, ((JSON.parse(jsonMessage).rotate * 3.1416) / 180));
        }
    }

    if (event === "Users") {
        if (players[JSON.parse(jsonMessage).user_name] == null) {
            addPlayer(JSON.parse(jsonMessage).user_name, JSON.parse(jsonMessage).x, JSON.parse(jsonMessage).y, ((JSON.parse(jsonMessage).rotate * 3.1416) / 180));
        }
    }

    if (event === "PlayerMove") {
        if (JSON.parse(jsonMessage).user_name !== MyId) {
            players[JSON.parse(jsonMessage).user_name].player.x = JSON.parse(jsonMessage).x;
            players[JSON.parse(jsonMessage).user_name].player.y = JSON.parse(jsonMessage).y;
            players[JSON.parse(jsonMessage).user_name].player.rotation = ((JSON.parse(jsonMessage).rotate * 3.1416) / 180);
        }
    }

    if (event === "Fire") {
        players[JSON.parse(jsonMessage).user_name].weapon.fire();
    }
}