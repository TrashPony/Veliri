function PlaceTexture(name) {
    let callBack = function (q, r) {
        mapEditor.send(JSON.stringify({
            event: "addOverTexture",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(q),
            r: Number(r),
            texture_name: name
        }));
    };
    SelectedSprite(event, 0, callBack)
}

function RemoveTexture() {
    let callBack = function (q, r) {
        mapEditor.send(JSON.stringify({
            event: "removeOverTexture",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(q),
            r: Number(r)
        }));
    };
    SelectedSprite(event, 0, callBack, false, true)
}