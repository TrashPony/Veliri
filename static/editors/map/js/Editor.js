function PlaceCoordinate(event, type) {

    let object = game.add.sprite(0, 0, type.texture_object);
    object.anchor.setTo(0.5);
    object.scale.set(0.5);

    setInterval(function () {
        object.x = ((game.input.mousePointer.x + game.camera.x) / game.camera.scale.x);
        object.y = ((game.input.mousePointer.y + game.camera.y) / game.camera.scale.y);
    }, 10);

    game.input.onUp.add(function () {

        object.destroy();
        game.input.onUp.removeAll();

        let x = (game.input.mousePointer.x + game.camera.x) / game.camera.scale.x;
        let y = (game.input.mousePointer.y + game.camera.y) / game.camera.scale.y;

        mapEditor.send(JSON.stringify({
            event: event,
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            id_type: Number(type.id),
            x: Number(x),
            y: Number(y)
        }));

    }, this);
}

function RemoveCoordinate() {
    SelectedSprite("", 0, function (x, y) {
        mapEditor.send(JSON.stringify({
            event: "placeCoordinate",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            id_type: 1,
            x: Number(x),
            y: Number(y)
        }));
    }, true, false, false, false)
}

function SendCommand(command) {
    mapEditor.send(JSON.stringify({
        event: command,
        id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value)
    }));
}