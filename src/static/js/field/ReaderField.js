function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    var price;
    var step;
    var phase;
    var clicked_id;
    var cell;
    var log;
    var ready;

    if (event === "InitPlayer") {
        price = document.getElementsByClassName('fieldInfo price');
        price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;
        step = document.getElementsByClassName('fieldInfo step');
        step[0].innerHTML = "Ход № " + JSON.parse(jsonMessage).game_step;
        phase = document.getElementsByClassName('fieldInfo phase');
        phase[0].innerHTML = "Фаза: " + JSON.parse(jsonMessage).game_phase;

        if(JSON.parse(jsonMessage).user_ready === "true") {
            ready = document.getElementById("Ready");
            ready.innerHTML = "Ты готов!";
            ready.style.backgroundColor = "#e1720f"
        }

        this.phase = JSON.parse(jsonMessage).game_phase;
    }

    if (event === "InitMap") {
        Field(JSON.parse(jsonMessage).x_map, JSON.parse(jsonMessage).y_map)
    }

    if (event === "InitUnit") {
        var x = (JSON.parse(jsonMessage).x).split(':');
        var y = (JSON.parse(jsonMessage).y).split(':');
        var type = (JSON.parse(jsonMessage).type_unit).split(':');

        for (var i = 0; i < x.length; i++) {
            clicked_id = x[i] + ":" + y[i];
            cell = document.getElementById(clicked_id);
            if (type[i] === "tank") cell.className = "fieldUnit tank";
            if (type[i] === "scout") cell.className = "fieldUnit scout";
            if (type[i] === "artillery") cell.className = "fieldUnit artillery";
        }
    }

    if (event === "CreateUnit"){
        if(JSON.parse(jsonMessage).error_type === "") {
            clicked_id = JSON.parse(jsonMessage).x + ":" + JSON.parse(jsonMessage).y;
            cell = document.getElementById(clicked_id);
            if (typeUnit === "tank") cell.className = "fieldUnit tank";
            if (typeUnit === "scout") cell.className = "fieldUnit scout";
            if (typeUnit === "artillery") cell.className = "fieldUnit artillery";
            price = document.getElementsByClassName('fieldInfo price');
            price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;
        } else {
            if(JSON.parse(jsonMessage).error_type === "busy") {
                log = document.getElementById('fieldLog');
                log.innerHTML = "Место занято"
            }

            if(JSON.parse(jsonMessage).error_type === "noMany") {
                log = document.getElementById('fieldLog');
                log.innerHTML = "Нет денег"
            }
        }
        typeUnit = null;
    }
    if (event === "Ready") {
        if(JSON.parse(jsonMessage).phase === "") {
            ready = document.getElementById("Ready");
            ready.innerHTML = "Ты готов!";
            ready.style.backgroundColor = "#e1720f"
        }
    }
}