function InitStructure(jsonMessage) {
    //TODO {"event":"InitStructure","user_name":"user","structure":{"type":"respawn","name_user":"user","watch_zone":2,"x":0,"y":0}}

    var structureStat = JSON.parse(jsonMessage).structure;
    CreateStructure(structureStat)
}

function CreateStructure(structureStat) {
    var x = structureStat.x;
    var y = structureStat.y;
    var type = structureStat.type;
    var owner = structureStat.owner;

    var structure = game.add.sprite(0, 0, type);
    game.physics.arcade.enable(structure);
    structure.inputEnabled = true;             // включаем ивенты на спрайт

    structure.events.onInputOver.add(mouse_over); // обрабатываем наведение мышки
    structure.events.onInputOut.add(mouse_out);   // обрабатываем убирание мышки

    structure.scale.set(.17);                  // устанавливаем размер спрайта от оригинала
    structure.id = x + ":" + y;

    var style;

    if (MY_ID === owner) {
        style = {font: "128px Arial", fill: "#00ffff"};
    } else {
        style = {font: "128px Arial", fill: "#ff0000"};
    }

    console.log(owner);
    var label_score = game.add.text(170, 40, owner, style);
    structure.addChild(label_score);

    if (cells[x + ":" + y]) {
        cells[x + ":" + y].addChild(structure);
    }
}