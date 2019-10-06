function RemoveTransport() {
    let callBack = function (x, y) {
        if (game.input.activePointer.leftButton.isDown) {
            mapEditor.send(JSON.stringify({
                event: "removeTransport",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                x: Number(x),
                y: Number(y)
            }));
        }
    };
    SelectedSprite(event, 0, callBack, false, false, true)
}

function AddTransport() {

    let transportIcon = game.add.sprite(0, 0, 'transportIcon');
    transportIcon.anchor.setTo(0.5);
    transportIcon.scale.set(0.5);

    setInterval(function () {
        transportIcon.x = ((game.input.mousePointer.x + game.camera.x) / game.camera.scale.x);
        transportIcon.y = ((game.input.mousePointer.y + game.camera.y) / game.camera.scale.y);
    }, 10);

    game.input.onUp.add(function () {

        game.input.onUp.removeAll();
        transportIcon.destroy();

        if (game.input.activePointer.leftButton.isDown) {

            let x = (game.input.mousePointer.x + game.camera.x) / game.camera.scale.x;
            let y = (game.input.mousePointer.y + game.camera.y) / game.camera.scale.y;

            mapEditor.send(JSON.stringify({
                event: "addTransport",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                x: Number(x),
                y: Number(y)
            }));
        }
    });
}