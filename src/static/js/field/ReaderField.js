function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    if (event === "InitPlayer") {
        var price = document.getElementsByClassName('fieldInfo price');
        price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;
        var step = document.getElementsByClassName('fieldInfo step');
        step[0].innerHTML = "Ход № " + JSON.parse(jsonMessage).game_step;
        var phase = document.getElementsByClassName('fieldInfo phase');
        phase[0].innerHTML = "Фаза: " + JSON.parse(jsonMessage).game_phase;

        this.phase = JSON.parse(jsonMessage).game_phase;
    }

    if (event === "InitMap") {
        Field(JSON.parse(jsonMessage).x_map, JSON.parse(jsonMessage).y_map)
    }

    if (event === "InitUnit") {
//////////////////////////////////
    }

    if (event === "CreateUnit"){
        var clicked_id = JSON.parse(jsonMessage).x + ":" + JSON.parse(jsonMessage).y;
        var cell = document.getElementById(clicked_id);
        if (typeUnit === "tank") cell.className = "fieldUnit tank";
        if (typeUnit === "scout") cell.className = "fieldUnit scout";
        if (typeUnit === "arta") cell.className = "fieldUnit arta";
        typeUnit = null;
    }
}