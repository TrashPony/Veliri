function addHeightCoordinate(q, r) {
    if (game.input.activePointer.leftButton.isDown) {
        mapEditor.send(JSON.stringify({
            event: "addHeightCoordinate",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(q),
            r: Number(r)
        }));
    }
}

function subtractHeightCoordinate(q, r) {
    if (game.input.activePointer.leftButton.isDown) {
        mapEditor.send(JSON.stringify({
            event: "subtractHeightCoordinate",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(q),
            r: Number(r)
        }));
    }
}

function PlaceCoordinate(event, type) {
    let callBack = function (q, r) {
        console.log(this);
        mapEditor.send(JSON.stringify({
            event: event,
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            id_type: Number(type.id),
            q: Number(q),
            r: Number(r)
        }));

        while (game.SelectLayer && game.SelectLayer.children.length > 0) {
            let sprite = game.SelectLayer.children.shift();
            sprite.destroy();
        }
    };

    SelectedSprite(event, type.impact_radius, callBack)
}


function SendCommand(command) {
    mapEditor.send(JSON.stringify({
        event: command,
        id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value)
    }));
}