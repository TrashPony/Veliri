function RemoveTransport() {
    let callBack = function (q, r) {
        if (game.input.activePointer.leftButton.isDown) {
            mapEditor.send(JSON.stringify({
                event: "removeTransport",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                q: Number(q),
                r: Number(r)
            }));
        }
    };
    SelectedSprite(event, 0, callBack, false, false, true)
}

function AddTransport() {
    let callBack = function (q, r) {
        if (game.input.activePointer.leftButton.isDown) {
            mapEditor.send(JSON.stringify({
                event: "addTransport",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                q: Number(q),
                r: Number(r)
            }));
        }
    };
    SelectedSprite(event, 0, callBack)
}