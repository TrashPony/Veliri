function RemoveHandler() {
    let callBack = function (x, y) {
        if (game.input.activePointer.leftButton.isDown) {
            mapEditor.send(JSON.stringify({
                event: "removeHandler",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                x: Number(x),
                y: Number(y)
            }));
        }
    };
    SelectedSprite(event, 0, callBack, false, false, false, false, true)
}

function AddHandler() {

    let transportIcon = game.add.sprite(0, 0, 'baseInIcon');
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

            ChangeHandlerOption(x, y)
        }
    })
}

function ChangeHandlerOption(x, y) {
    if (document.getElementById("handlerBlockOption")) document.getElementById("handlerBlockOption").remove();

    let handlerBlockOption = document.createElement("div");
    handlerBlockOption.id = "handlerBlockOption";

    let typeSelect = document.createElement("select");
    typeSelect.id = "typeSelect";
    typeSelect.innerHTML = "" +
        "<option value disabled selected> Выберите тип перехода </option>" +
        "<option value='base'>Base</option>" +
        "<option value='sector'>Sector</option>";

    let baseSelect = document.createElement("select");
    baseSelect.id = "baseSelect";
    baseSelect.style.opacity = 0;
    baseSelect.innerHTML = "<option value disabled selected> Выберите базу </option>";
    for (let i in game.bases) {
        baseSelect.innerHTML += "<option value='" + game.bases[i].id + "'> " + game.bases[i].name + " </option>"
    }

    let mapOptions = document.getElementById("mapSelector").innerHTML;
    let mapSelect = document.createElement("select");
    mapSelect.style.opacity = 0;
    mapSelect.id = "mapSelect";
    mapSelect.innerHTML = mapOptions;

    let position = document.createElement("div");
    position.style.opacity = 0;
    position.appendChild(PositionsReader(null));

    let apply = document.createElement("input");
    apply.value = "Отменить";
    apply.type = "submit";
    apply.onclick = function () {
        if (document.getElementById('typeSelect').value === 'base') {
            mapEditor.send(JSON.stringify({
                event: "addHandler",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                x: Number(x),
                y: Number(y),
                to_base_id: Number(document.getElementById('baseSelect').value),
                type_handler: document.getElementById('typeSelect').value,
            }));
        } else if (document.getElementById('typeSelect').value === 'sector') {
            mapEditor.send(JSON.stringify({
                event: "addHandler",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                x: Number(x),
                y: Number(y),
                to_pos: JSON.stringify(handlersPos),
                to_map_id: Number(document.getElementById('mapSelect').value),
                type_handler: document.getElementById('typeSelect').value,
            }));
        }
        handlersPos = [];
        handlerBlockOption.remove();
    };

    typeSelect.onchange = function () {
        if (this.value === "base") {
            apply.value = "Применить";
            position.style.opacity = 0;
            mapSelect.style.opacity = 0;
            baseSelect.style.opacity = 1;
        } else if (this.value === "sector") {
            apply.value = "Применить";
            position.style.opacity = 1;
            mapSelect.style.opacity = 1;
            baseSelect.style.opacity = 0;
        } else {
            apply.value = "отменить";
        }
    };

    handlerBlockOption.appendChild(typeSelect);
    handlerBlockOption.appendChild(baseSelect);
    handlerBlockOption.appendChild(mapSelect);
    handlerBlockOption.appendChild(position);
    handlerBlockOption.appendChild(apply);
    document.getElementById("coordinates").appendChild(handlerBlockOption);
}

//{"x":0,"y":0,"resp_rotate":0}
let handlersPos = [];

function EditPos() {
    let callBack = function (x, y) {
        let handlerBlockOption = document.createElement("div");
        handlerBlockOption.id = "handlerBlockOption";

        let mapOptions = document.getElementById("mapSelector").innerHTML;
        let mapSelect = document.createElement("select");
        mapSelect.id = "mapSelect";
        mapSelect.innerHTML = mapOptions;
        mapSelect.value = game.map.OneLayerMap[x][y].to_map_id;

        let position = document.createElement("div");
        position.appendChild(PositionsReader(game.map.OneLayerMap[x][y].positions));

        let apply = document.createElement("input");
        apply.value = "Применить";
        apply.type = "submit";
        apply.onclick = function () {
            mapEditor.send(JSON.stringify({
                event: "addHandler",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                x: Number(x),
                y: Number(y),
                to_pos: JSON.stringify(handlersPos),
                to_map_id: Number(document.getElementById('mapSelect').value),
                type_handler: "sector",
            }));
            handlersPos = [];
            handlerBlockOption.remove();
        };

        handlerBlockOption.appendChild(mapSelect);
        handlerBlockOption.appendChild(position);
        handlerBlockOption.appendChild(apply);
        document.getElementById("coordinates").appendChild(handlerBlockOption);
    };
    SelectedSprite(event, 0, callBack, false, false, false, false, true)
}

function PositionsReader(old) {

    let pos = document.getElementById("posTable");
    if (!pos) {
        pos = document.createElement("table");
        pos.id = "posTable";
    }

    pos.innerHTML = `
        <tr>
            <td>X</td>
            <td>Y</td>
            <td>Angle</td>
            <td></td>
        </tr>
    `;
    if (old) {
        handlersPos = old;

        for (let i in old) {
            pos.innerHTML += `
            <tr>
                <td>${old[i].x}</td>
                <td>${old[i].y}</td>
                <td>${old[i].resp_rotate}</td>
                <td><input onclick="removePosByID(${i})" type="button" value="X"></td>
            </tr>
            `;
        }
    }

    pos.innerHTML += `
        <tr>
            <td><input id="posX" type="number" value="0" placeholder="X"></td>
            <td><input id="posY" type="number" value="0" placeholder="Y"></td>
            <td><input id="posA" type="number" value="0" placeholder="A"></td>
            <td></td>
        </tr>
        <tr>
            <td></td>
            <td></td>
            <td><input id="AddPosButton" type="button" value="Добавить"></td>
            <td></td>
        </tr>
    `;

    setTimeout(function () {
        $('#AddPosButton').click(function () {
            handlersPos.push({
                x: Number($('#posX').val()),
                y: Number($('#posY').val()),
                resp_rotate: Number($('#posA').val()),
            });
            PositionsReader(handlersPos);
        });
    }, 300);

    return pos;
}

function removePosByID(i) {
    handlersPos.splice(Number(i), 1);
    PositionsReader(handlersPos);
}