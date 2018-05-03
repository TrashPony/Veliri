function UpdateUnitInfo(unit) {
    for (var parameter in unit) {
        if (unit.hasOwnProperty(parameter)) {
            var row = document.getElementById(parameter);
            if (row) {
                row.innerHTML = unit[parameter];
            }
        }
    }

    UpdateUnitPicture(unit);
}

function UpdateUnitPicture(unit) {
    var picUnit = document.getElementById("picUnit");

    var gameWindows = document.getElementById("gameUnit");

    if (gameWindows) {
        gameWindows.remove();
    }

    gameWindows = document.createElement("div");
    gameWindows.id = "gameUnit";

    picUnit.appendChild(gameWindows);

    if (picUnit) {

        /*console.log(unit.chassis);
        console.log(unit.weapon);
        console.log(unit.tower);
        console.log(unit.body);
        console.log(unit.radar);*/

        var game = new Phaser.Game(180, 180, Phaser.AUTO, gameWindows.id, {
            preload: preload,
            create: create,
            update: update
        });

        game.time.desiredFps = 1;
        
        function preload() {
            if (unit.chassis !== null) {
                game.load.image('chassis', '/assets/' + unit.chassis.name + '.png');
            }
            if (unit.weapon !== null) {
                game.load.image('weapon', '/assets/' + unit.weapon.name + '.png');
            }
            if (unit.tower !== null) {
                game.load.image('tower', '/assets/' + unit.tower.name + '.png');
            }
            if (unit.body !== null) {
                game.load.image('body', '/assets/' + unit.body.name + '.png');
            }
            if (unit.radar !== null) {
                game.load.image('radar', '/assets/' + unit.radar.name + '.png');
            }
        }

        function create() {
            game.stage.backgroundColor = '#182d3b';
            var chassis = game.add.sprite(0, 0, 'chassis');
            if (chassis.key === "__missing") {
                chassis.visible = false;
            }
            var weapon = game.add.sprite(0, 0, 'weapon');
            if (weapon.key === "__missing") {
                weapon.visible = false;
            }
            var tower = game.add.sprite(0, 0, 'tower');
            if (tower.key === "__missing") {
                tower.visible = false;
            }
            var body = game.add.sprite(0, 0, 'body');
            if (body.key === "__missing") {
                body.visible = false;
            }
            var radar = game.add.sprite(0, 0, 'radar');
            if (radar.key === "__missing") {
                radar.visible = false;
            }

        }

        function update() {

        }
    }
}