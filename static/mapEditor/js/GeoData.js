let geoPoint;
let radius = 24;
let response = [];

function RemoveGeoData() {

    if (document.getElementById("rotateBlock")) document.getElementById("rotateBlock").remove();
    let rotate = document.createElement("div");
    rotate.style.height = "50px";
    rotate.id = "rotateBlock";

    let apply = document.createElement("input");
    apply.value = "Применить";
    apply.type = "submit";
    apply.onclick = function () {
        for (let i = 0; i < response.length; i++) {
            mapEditor.send(response[i]);
        }
        mapEditor.send(JSON.stringify({
            event: "SelectMap",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
        }));
        CancelGeoData();
    };

    let cancel = document.createElement("input");
    cancel.value = "Отмена";
    cancel.type = "submit";
    cancel.onclick = function () {
        CancelGeoData();
        mapEditor.send(JSON.stringify({
            event: "SelectMap",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
        }));
    };

    rotate.appendChild(cancel);
    rotate.appendChild(apply);
    document.getElementById("coordinates").appendChild(rotate);

    for (let i = 0; i < game.map.geo_data.length; i++) {
        let deleteButton = game.icon.create(game.map.geo_data[i].x, game.map.geo_data[i].y, 'cancelIcon');
        deleteButton.anchor.setTo(0.5);
        deleteButton.scale.set(0.05)

        deleteButton.inputEnabled = true; // включаем ивенты на спрайт
        deleteButton.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
        deleteButton.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

        let overSprite = game.icon.create(game.map.geo_data[i].x, game.map.geo_data[i].y, 'cancelIcon');
        overSprite.anchor.setTo(0.5);
        overSprite.scale.set(0.07);
        overSprite.alpha = 0;

        deleteButton.events.onInputDown.add(function () {
            response.push(JSON.stringify({
                event: "removeGeoData",
                id: game.map.geo_data[i].id,
            }));
            game.map.geo_data[i] = null;
            CreateGeoData(game.map.geo_data);
            deleteButton.destroy();
            overSprite.destroy();
        });

        deleteButton.events.onInputOver.add(function () {
            overSprite.alpha = 1;
        });

        deleteButton.events.onInputOut.add(function () {
            overSprite.alpha = 0;
        });
    }
}

function CancelGeoData() {
    if (game) {
        game.input.onDown.remove(addGeoData, game);
        game.input.moveCallbacks = [];
        if (document.getElementById("rotateBlock")) document.getElementById("rotateBlock").remove();
        if (geoPoint) geoPoint.destroy();
        response = [];
        game.icon.forEach(function (c) {
            c.kill();
        });
    }
}

function AddGeoData() {
    if (game) {

        //создание окна с настройками
        if (document.getElementById("rotateBlock")) document.getElementById("rotateBlock").remove();
        let rotate = document.createElement("div");
        rotate.style.height = "50px";
        rotate.id = "rotateBlock";

        let geoRadius = createRange("geoRadius", 1, 300, 1, 24);
        geoRadius.oninput = function () {
            document.getElementById("geoInput").innerHTML = geoRadius.value;
            let x = geoPoint.x;
            let y = geoPoint.y;
            geoPoint.destroy();
            geoPoint = game.add.graphics(0, 0);
            geoPoint.beginFill(0xFF0000, 0.3);
            geoPoint.drawCircle(geoRadius.value, geoRadius.value, geoRadius.value * 2);
            radius = geoRadius.value;
            geoPoint.x = x;
            geoPoint.y = y;
        };
        let div = document.createElement("div");
        div.innerHTML = "<span> Радиус: </span> <span id='geoInput'> 12 </span>";

        let apply = document.createElement("input");
        apply.value = "Применить";
        apply.type = "submit";
        apply.onclick = function () {
            for (let i = 0; i < response.length; i++) {
                mapEditor.send(response[i]);
            }
            mapEditor.send(JSON.stringify({
                event: "SelectMap",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            }));
            CancelGeoData();
        };

        rotate.appendChild(geoRadius);
        rotate.appendChild(div);
        rotate.appendChild(apply);
        document.getElementById("coordinates").appendChild(rotate);

        geoPoint = game.add.graphics(0, 0);
        geoPoint.beginFill(0xFF0000, 0.3);
        geoPoint.drawCircle(24, 24, 24 * 2);

        function move(e, x, y) {
            geoPoint.x = e.worldX / game.camera.scale.x - radius;
            geoPoint.y = e.worldY / game.camera.scale.y - radius;
        }

        game.input.onDown.add(addGeoData, game);
        game.input.addMoveCallback(move, this);
    }
}

function addGeoData(e) {
    if (game.input.activePointer.leftButton.isDown) {

        response.push(JSON.stringify({
            event: "addGeoData",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            x: Math.round(Number(e.worldX / game.camera.scale.x)),
            y: Math.round(Number(e.worldY / game.camera.scale.y)),
            radius: Number(document.getElementById("geoRadius").value),
        }));

        game.geoData.beginFill(0xFF0000, 0.5);
        game.geoData.drawCircle(
            e.worldX / game.camera.scale.x,
            e.worldY / game.camera.scale.y,
            Number(document.getElementById("geoRadius").value) * 2,
        );
    }
}