function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    var price;
    var step;
    var clicked_id;
    var cell;
    var log;
    var ready;
    var info;
    var coor_id;
    var x;
    var y;
    var userOwned;

    if (event === "InitPlayer") {
        price = document.getElementsByClassName('fieldInfo price');
        price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;
        step = document.getElementsByClassName('fieldInfo step');
        step[0].innerHTML = "Ход № " + JSON.parse(jsonMessage).game_step;
        phase = document.getElementsByClassName('fieldInfo phase');
        phase[0].innerHTML = "Фаза: " + JSON.parse(jsonMessage).game_phase;

        if (JSON.parse(jsonMessage).user_ready === "true") {
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
        userOwned = JSON.parse(jsonMessage).user_owned;

        clicked_id = x + ":" + y;
        cell = document.getElementById(clicked_id);
        if (type === "tank") cell.className = "fieldUnit tank";
        if (type === "scout") cell.className = "fieldUnit scout";
        if (type === "artillery") cell.className = "fieldUnit artillery";
        cell.onclick = function () {
            SelectUnit(this.id)
        };
        cell.innerHTML = "hp: " + hp;

        if (JSON.parse(jsonMessage).user_name === userOwned) {
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
        coor_id = x + ":" + y;
        cell = document.getElementById(coor_id);
        cell.className = "fieldUnit respawn";
        cell.innerHTML = "Resp: " + JSON.parse(jsonMessage).user_name;
    }

    if (event === "emptyCoordinate") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        coor_id = x + ":" + y;
        cell = document.getElementById(coor_id);
        if (cell) {
            cell.className = "fieldUnit open";
        }
    }

    if (event === "SelectCoordinateCreate") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        coor_id = x + ":" + y;
        cell = document.getElementById(coor_id);
        if (cell) {
            cell.className = "fieldUnit create";
        }
    }

    if (event === "CreateUnit") {
        if (JSON.parse(jsonMessage).error_type === "") {
            userOwned = JSON.parse(jsonMessage).user_owned;
            x = JSON.parse(jsonMessage).x;
            y = JSON.parse(jsonMessage).y;

            coor_id = x + ":" + y;
            cell = document.getElementById(coor_id);
            cell.onclick = function () {
                SelectUnit(this.id)
            };
            if (typeUnit === "tank") cell.className = "fieldUnit tank";
            if (typeUnit === "scout") cell.className = "fieldUnit scout";
            if (typeUnit === "artillery") cell.className = "fieldUnit artillery";
            price = document.getElementsByClassName('fieldInfo price');
            price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;

            if (JSON.parse(jsonMessage).user_name === userOwned) {
                cell.style.color = "#fbfdff";
                cell.style.borderColor = "#fbfdff";
            } else {
                cell.style.color = "#FF0117";
                cell.style.borderColor = "#FF0117";
            }

        } else {
            var celles = document.getElementsByClassName("fieldUnit create");
            while (0 < celles.length) {
                if (celles[0]) {
                    celles[0].className = "fieldUnit open";
                }
            }
            if (JSON.parse(jsonMessage).error_type === "busy") {
                log = document.getElementById('fieldLog');
                log.innerHTML = "Место занято"
            }

            if (JSON.parse(jsonMessage).error_type === "no many") {
                log = document.getElementById('fieldLog');
                log.innerHTML = "Нет денег"
            }
            if (JSON.parse(jsonMessage).error_type === "not allow") {
                log = document.getElementById('fieldLog');
                log.innerHTML = "Не разрешено"
            }
        }
        typeUnit = null;
    }
    if (event === "MouseOver") {
        info = document.getElementById('unitInfo');
        info.innerHTML = "Тип Юнита: " + JSON.parse(jsonMessage).type_unit + "<br>" +
            "Владелец: " + JSON.parse(jsonMessage).user_owned + "<br>" +
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
        var error = JSON.parse(jsonMessage).error;
        phase = JSON.parse(jsonMessage).phase;
        console.log(phase);
        if (error === "") {
            ready = document.getElementById("Ready");
            var phaseBlock = document.getElementById("phase");

            if (phase === "") {
                ready.innerHTML = "Ты готов!";
                ready.style.backgroundColor = "#e1720f";
            } else {
                ready.innerHTML = "Готов!";
                if (phase === "move") {
                    ready.style.backgroundColor = "#A8ADE1";
                }
                if (phase === "targeting") {
                    ready.style.backgroundColor = "#E1C7A6";
                }
                if (phase === "attack") {
                    ready.style.backgroundColor = "#E12D27";
                }
                phaseBlock.innerHTML = JSON.parse(jsonMessage).phase;
                var cells = document.getElementsByClassName("fieldUnit create");
                while (0 < cells.length) {
                    if (cells[0]) {
                        cells[0].className = "fieldUnit open";
                    }
                }
            }
        } else {
            if (error === "not units"){
                alert("У вас нет юнитов")
            }
        }
    }

    if (event === "SelectUnit") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        coor_id = x + ":" + y;
        cell = document.getElementById(coor_id);
        if (cell) {
            cell.className = "fieldUnit move";
        }
    }

    if (event === "MoveUnit") {
        move = null;
        var moveCells = document.getElementsByClassName("fieldUnit move");
        while (0 < moveCells.length) {
            if (moveCells[0]) {
                moveCells[0].className = "fieldUnit open"; // TODO: ставить реальные статусы ячеек
            }
        }

        error = JSON.parse(jsonMessage).error;
        if (error === "") {
                                // TODO копировать ячейку в новую координату, а старую закрыть
        } else {

        }

    }
}