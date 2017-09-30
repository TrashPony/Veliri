function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    if (event === "InitPlayer") {
        alert( "твои деньги" + JSON.parse(jsonMessage).player_price);
        alert( "ход # " + JSON.parse(jsonMessage).game_step);
        alert( "фаза " + JSON.parse(jsonMessage).game_phase);
    }

    if (event === "InitMap") {
        Field(JSON.parse(jsonMessage).x_map, JSON.parse(jsonMessage).y_map)
    }
}