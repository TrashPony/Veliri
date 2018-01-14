function InitUnit(jsonMessage) {
    var unitStat = JSON.parse(jsonMessage).unit;
    DelMoveCoordinate();
    CreateUnit(unitStat)
}

function CreateUnit(unitStat) {
    var x = unitStat.x;
    var y = unitStat.y;
    var type = unitStat.type;
    var owned = unitStat.owner;
    var action = unitStat.action;
    var hp = unitStat.hp;

    var cell = cells[x + ":" + y];
    var unit = game.add.sprite(cell.x + tileWidth / 2, cell.y + tileWidth / 2, type);

    game.physics.arcade.enable(unit);
    unit.inputEnabled = true;             // включаем ивенты на спрайт
    unit.anchor.setTo(0.35, 0.5);         // устанавливаем центр спрайта
    unit.scale.set(.32);                  // устанавливаем размер спрайта от оригинала
    unit.body.collideWorldBounds = true;  // границы страницы

    unit.id = x + ":" + y;
    unit.events.onInputOver.add(mouse_over); // обрабатываем нажатие мышки
    unit.events.onInputOut.add(mouse_out);   // обрабатываем нажатие мышки

    var style;

    if (MY_ID === owned) {
        style = {font: "52px Arial", fill: "#00ffff"};
        unit.events.onInputDown.add(SelectUnit);
        if (action === false) {
            unit.tint = 0x757575; // накладывает фильтр со светом
        }
    } else {
        style = {font: "52px Arial", fill: "#ff0000"};
        unit.events.onInputDown.add(SelectTarget);
    }

    var label_score = game.add.text(x, y, "hp " + hp, style);
    unit.addChild(label_score);

    units[unit.id] = unit;
}

function BuildUnit(jsonMessage) {
    if (JSON.parse(jsonMessage).error === null) {
        var price = document.getElementsByClassName('fieldInfo price');
        price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;
    } else {
        var log = document.getElementById('fieldLog');

        if (JSON.parse(jsonMessage).error_type === "busy") {
            log.innerHTML = "Место занято"
        }
        if (JSON.parse(jsonMessage).error_type === "no many") {
            log.innerHTML = "Нет денег"
        }
        if (JSON.parse(jsonMessage).error_type === "not allow") {
            log.innerHTML = "Не разрешено"
        }
    }

    var cells = document.getElementsByClassName("fieldUnit create");
    while (0 < cells.length) {
        if (cells[0]) {
            cells[0].className = "fieldUnit open";
        }
    }
    typeUnit = null;
}