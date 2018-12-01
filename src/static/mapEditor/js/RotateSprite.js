function RotateSprite() {
    let coordinate = this;

    let block = document.getElementById("coordinates");

    let rotate = document.createElement("div");
    rotate.id = "rotateBlock";

    let range = document.createElement("input");
    range.id = "rotateRange";
    range.name = "rotateRange";
    range.type = "range";
    range.min = 0;
    range.max = 360;
    range.step = 1;
    range.value = this.objectSprite.angle;
    range.oninput = function () {
        document.getElementById("rotateOutput").innerHTML = range.value;
        coordinate.objectSprite.angle = range.value;
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.angle = range.value;
        }
    };

    let output = document.createElement("div");
    output.innerHTML = "<span> Градусы: </span> <span id='rotateOutput'> " + coordinate.objectSprite.angle + " </span>";

    let apply = document.createElement("input");
    apply.value = "Применить";
    apply.type = "submit";
    apply.onclick = function () {
        mapEditor.send(JSON.stringify({
            event: "rotateObject",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(coordinate.q),
            r: Number(coordinate.r),
            rotate: Number(document.getElementById("rotateRange").value)
        }));

        mapEditor.send(JSON.stringify({
            event: "getAllTypeCoordinate"
        }));
    };

    let cancel = document.createElement("input");
    cancel.value = "Отменить";
    cancel.type = "submit";
    cancel.onclick = function () {
        mapEditor.send(JSON.stringify({
            event: "SelectMap",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value)
        }));
    };

    rotate.appendChild(range);
    rotate.appendChild(output);

    rotate.appendChild(apply);
    rotate.appendChild(cancel);

    block.appendChild(rotate);
}