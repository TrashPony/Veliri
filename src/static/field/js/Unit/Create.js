function InitUnit(jsonMessage) {
    var unitStat = JSON.parse(jsonMessage).unit;
    RemoveSelectMoveCoordinate();
    CreateUnit(unitStat)
}

function CreateUnit(unitStat) {

    var x = unitStat.x;
    var y = unitStat.y;
    var type = unitStat.type;
    var owned = unitStat.owner;
    var action = unitStat.action;
    var hp = unitStat.hp;

    var cell = GameMap.OneLayerMap[x][y].sprite;
    var unit = game.add.sprite(cell.x + tileWidth / 2, cell.y + tileWidth / 2, type);

    game.physics.arcade.enable(unit);
    unit.inputEnabled = true;             // включаем ивенты на спрайт
    unit.anchor.setTo(0.35, 0.5);         // устанавливаем центр спрайта
    unit.scale.set(.32);                  // устанавливаем размер спрайта от оригинала
    unit.body.collideWorldBounds = true;  // границы страницы

    unit.id = x + ":" + y;
    unit.events.onInputOver.add(mouse_over); // обрабатываем наведение мышки
    unit.events.onInputOut.add(mouse_out);   // обрабатываем убирание мышки

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