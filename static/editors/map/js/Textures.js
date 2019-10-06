function PlaceTexture(name) {

    let bmd = game.make.bitmapData(512, 512);
    bmd.alphaMask(name, 'brush');
    let bmdSprite = game.add.sprite(200, 200, bmd);

    setInterval(function () {
        bmdSprite.x = ((game.input.mousePointer.x + game.camera.x) / game.camera.scale.x) - 256;
        bmdSprite.y = ((game.input.mousePointer.y + game.camera.y) / game.camera.scale.y) - 256;
    }, 10);

    game.input.onUp.add(function () {
        game.input.onUp.removeAll();

        let x = (game.input.mousePointer.x + game.camera.x) / game.camera.scale.x;
        let y = (game.input.mousePointer.y + game.camera.y) / game.camera.scale.y;

        game.bmdTerrain.draw(bmd, x - 256, y - 256);
        bmdSprite.destroy();
        bmd.destroy();

        mapEditor.send(JSON.stringify({
            event: "addOverTexture",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            x: Number(x),
            y: Number(y),
            texture_name: name
        }));
    }, this);
}

function RemoveTexture() {
    let callBack = function (x, y) {
        mapEditor.send(JSON.stringify({
            event: "removeOverTexture",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            x: Number(x),
            y: Number(y)
        }));
    };
    SelectedSprite(event, 0, callBack, false, true)
}