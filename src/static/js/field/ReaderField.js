function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    var price;
    var step;
    var phase;
    var clicked_id;
    var cell;
    var log;
    var ready;
    var info;
    var x;
    var y;
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
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        var type = JSON.parse(jsonMessage).type_unit;
        var hp = JSON.parse(jsonMessage).hp;
        var action = JSON.parse(jsonMessage).unit_action;
        var users = JSON.parse(jsonMessage).user_owned;


        clicked_id = x + ":" + y;
        cell = document.getElementById(clicked_id);
        if (type === "tank") cell.className = "fieldUnit tank";
        if (type === "scout") cell.className = "fieldUnit scout";
        if (type === "artillery") cell.className = "fieldUnit artillery";
        cell.innerHTML = "hp: " + hp;

        if (JSON.parse(jsonMessage).user_name === users) {
            cell.style.color = "#fbfdff";
            cell.style.borderColor = "#fbfdff";
        } else {
            cell.style.color = "#FF0117";
            cell.style.borderColor = "#FF0117";
        }
    }

    if (event === "InitResp") {
        x = JSON.parse(jsonMessage).respawn_x;
        y = JSON.parse(jsonMessage).respawn_y;
        var coor_id = x + ":" + y;
        cell = document.getElementById(coor_id);
        cell.className = "fieldUnit respawn";
        cell.innerHTML = "Resp: " + JSON.parse(jsonMessage).user_name;
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
    if (event === "MouseOver") {
        info = document.getElementById('unitInfo');
        info.innerHTML =    "Тип Юнита: " + JSON.parse(jsonMessage).type_unit + "<br>" +
                            "Владелец: " + JSON.parse(jsonMessage).user_id + "<br>" +
                            "hp: " + JSON.parse(jsonMessage).hp + "<br>" +
                            "Ходил: " + JSON.parse(jsonMessage).unit_action + "<br>" +
                            "Цель " + JSON.parse(jsonMessage).target + "<br>" +
                            "Урон: " + JSON.parse(jsonMessage).damage + "<br>" +
                            "Скорость: " + JSON.parse(jsonMessage).move_speed + "<br>" +
                            "Инициатива: " + JSON.parse(jsonMessage).init + "<br>" +
                            "Дальность атаки: " + JSON.parse(jsonMessage).range_attack + "<br>" +
                            "Дальность обзора: " + JSON.parse(jsonMessage).range_view + "<br>" +
                            "Площадь атаки: " + JSON.parse(jsonMessage).area_attack + "<br>" +
                            "Тип атаки: " + JSON.parse(jsonMessage).type_attack
    }
    if (event === "Ready") {
        if(JSON.parse(jsonMessage).phase === "") {
            ready = document.getElementById("Ready");
            ready.innerHTML = "Ты готов!";
            ready.style.backgroundColor = "#e1720f"
        }
    }
}