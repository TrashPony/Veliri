var keybord;
var players = {};    // карта со всеми игроками на карте
var game;
var tileWidth = 100; // ширина и высота спрайта в сетке грида
var countWidthGrid = 10;    //
var countHeightGrid = 10;   //
var MyId;            // ид полученый от сервера текущего клиента

function Game() {

    var width;//получаем ширину монитора
    var height; //получаем высоту монитора

    if (window.innerWidth < tileWidth * countWidthGrid) {
        width = window.innerWidth;
    } else {
        width = tileWidth * countWidthGrid
    }

    if (window.innerHeight < tileWidth * countHeightGrid) {
        height = window.innerHeight;
    } else {
        height = tileWidth * countHeightGrid;
    }

    game = new Phaser.Game(width, height, Phaser.CANVAS, 'main', {
        preload: preload,
        create: create,
        update: update,
        render: render
    }); //создаем игровое поле с высотой и шир
}

function preload() {
    game.stage.disableVisibilityChange = true; // не дает оставиться выполнения скрипта если окно скрыто
    game.load.image('player', 'test/images/tank.png');
    game.load.image('bullet', 'test/images/bullet.png');
    game.load.image('floor', 'test/images/terrain.jpg');
    game.load.bitmapFont('carrier_command', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.png', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.xml');
}

function create() {

    webSocket.send(JSON.stringify({
        event: "Join"
    }));

    webSocket.onmessage = function (msg) {
        ReadResponse(msg.data)
    };

    game.physics.startSystem(Phaser.Physics.ARCADE);

    game.time.advancedTiming = true; // настройка fts
    game.time.desiredFps = 60;       // макс фпс 60
    game.time.slowMotion = 0;        // лавный переход в мин фпс

    game.world.setBounds(0, 0, tileWidth*countWidthGrid, tileWidth*countHeightGrid); //размеры карты

    game.stage.backgroundColor = "#242424"; //цвет фона
    keybord = game.input.keyboard.createCursorKeys();  //инициализируем клавиатуру

    for (var x = 0; x < countWidthGrid; x++) {
        for (var y = 0; y < countHeightGrid; y++) {
            console.log("123");
            var floorSprite = game.add.tileSprite(x*tileWidth, y*tileWidth, tileWidth,tileWidth, 'floor');
            game.add.bitmapText(x*tileWidth + tileWidth/2, y*tileWidth + tileWidth/2, 'carrier_command',x + ":" + y,12);
            floorSprite.id = x + ":" + y;
            floorSprite.inputEnabled = true; // включаем ивенты на спрайт
            floorSprite.events.onInputDown.add(test, this); // обрабатываем нажатие мышки
        }
    }

    game.input.onDown.add(function (pointer) {
        if (pointer.leftButton.isDown) {
            sendFire();
            players[MyId].weapon.fire(); //исполняем выстрелы
        }
    });
}

function test(sprite, pointer) {
    if (pointer.rightButton.isDown) {
        players[MyId].player.rotation = game.physics.arcade.angleToXY(players[MyId].player, sprite.x + tileWidth / 2, sprite.y + tileWidth / 2);
        players[MyId].player.body.velocity = game.physics.arcade.velocityFromAngle(players[MyId].player.angle, 40);
        players[MyId].player.movePoint = new Phaser.Point(sprite.x + tileWidth / 2, sprite.y + tileWidth / 2);
        console.log(players[MyId].player.movePoint.x)
    }
}

function update() {
    if (players[MyId]) {
        characterController(); //управление игроком
    }
}

function render() {
    //game.debug.cameraInfo(game.camera, 32, 32);
    if (players[MyId]) {
        game.debug.spriteInfo(players[MyId].player, 32, 32);
        game.debug.spriteCoords(players[MyId].player, 32, 500);
    }
}

function characterController() {

    if (players[MyId].player.movePoint == null) {
        players[MyId].player.body.angularVelocity = 0;
        players[MyId].player.body.velocity.x = 0;
        players[MyId].player.body.velocity.y = 0;
    } else {
        sendPosition();

        var xTarget = Math.round(players[MyId].player.movePoint.x);
        var xPlayer = Math.round(players[MyId].player.x);
        var yTarget = Math.round(players[MyId].player.movePoint.y);
        var yPlayer = Math.round(players[MyId].player.y);

        if (xTarget === xPlayer && yTarget === yPlayer) {
            players[MyId].player.movePoint = null;
        }
    }

    if (game.input.keyboard.isDown(Phaser.Keyboard.A) || keybord.left.isDown) {
        players[MyId].player.body.angularVelocity = -100;
        sendPosition("A");
    }
    if (game.input.keyboard.isDown(Phaser.Keyboard.D) || keybord.right.isDown) {
        players[MyId].player.body.angularVelocity = 100;
        sendPosition("D");
    }
    if (game.input.keyboard.isDown(Phaser.Keyboard.W) || keybord.up.isDown) {
        players[MyId].player.body.velocity = game.physics.arcade.velocityFromAngle(players[MyId].player.angle, 200);
        sendPosition("W");
    }
    if (game.input.keyboard.isDown(Phaser.Keyboard.S) || keybord.down.isDown) {
        players[MyId].player.body.velocity = game.physics.arcade.velocityFromAngle(players[MyId].player.angle - 180, 100);
        sendPosition("S");
    }
} //управление

function sendPosition(character) {
    webSocket.send(JSON.stringify({
        event: "PlayerMove",
        user_name: MyId,
        rotate: Math.round(Number((players[MyId].player.rotation * 180) / 3.1416)),
        x: Math.round(Number(players[MyId].player.x)),
        y: Math.round(Number(players[MyId].player.y))
    }));
} //отправляем инфу о том, куда игрок двинулся на сервер

function sendFire() {
    webSocket.send(JSON.stringify({
        event: "Fire",
        user_name: MyId
    }));
}

function addPlayer(playerId, x, y, rotate) {
    var player = game.add.sprite(x, y, "player");
    game.physics.arcade.enable(player);
    player.smoothed = false;
    player.anchor.setTo(0.35, 0.5);        // устанавливаем центр спрайта
    player.scale.set(.4);                  // устанавливаем размер спрайта от оригинала
    player.body.collideWorldBounds = true; // границы страницы
    player.rotation = rotate;              // задаем угол поворота спрайта
    player.id = playerId;

    player.inputEnabled = true; // включаем ивенты на спрайт
    player.events.onInputOver.add(mouse_over, this); // обрабатываем навод мышки на спрайт
    player.events.onInputOut.add(mouse_out, this); // обрабатываем уход мышки со спрайта

    var weapon = game.add.weapon(30, 'bullet'); //подключаем возможность выстрелов
    weapon.bulletKillType = Phaser.Weapon.KILL_WORLD_BOUNDS;
    weapon.bulletSpeed = 1200; // скорость полета снаряда
    weapon.fireRate = 500;     // скорострельность оружия
    weapon.trackSprite(player, 0, 0, true);
    weapon.trackOffset.x = 80;
    weapon.trackOffset.y = 0;

    //players[playerId] = {player, weapon};

    game.camera.follow(players[MyId].player);
}