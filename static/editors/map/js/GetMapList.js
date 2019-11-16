function GetMapList() {
    mapEditor.send(JSON.stringify({
        event: "getMapList"
    }));

    mapEditor.send(JSON.stringify({
        event: "getAllTypeCoordinate"
    }));
}

let game;

function CreateMapList(jsonMessage) {

    let maps = JSON.parse(jsonMessage).maps;

    let mapSelect = document.getElementById("mapSelector");
    mapSelect.innerHTML = '';

    mapSelect.innerHTML = `<option value="0">-</option>`;
    for (let i in maps) {

        let option = document.createElement("option");
        option.id = maps[i].id;
        option.value = maps[i].id;
        option.innerHTML = maps[i].Name + `<span style='color: red'> ID: ${maps[i].id}</span>`;

        mapSelect.appendChild(option);
    }
}

function selectMap() {
    let selectedValue = document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value;

    mapEditor.send(JSON.stringify({
        event: "SelectMap",
        id: Number(selectedValue)
    }));
}

let bases = null;

function createGame(jsonMessage) {

    if (game) {
        bases = JSON.parse(jsonMessage).bases;
        UpdateMap(JSON.parse(jsonMessage).map, game, JSON.parse(jsonMessage).bases);
        CreateMeta(JSON.parse(jsonMessage));
    } else {

        let loadFunc = function () {
            CreateMeta(JSON.parse(jsonMessage));
        };

        game = CreateGame(JSON.parse(jsonMessage).map, loadFunc, "mapEditor");
        game.bases = JSON.parse(jsonMessage).bases;
    }
}

function CreateMeta(data) {
    if (data.bases) {
        CreateLabelBase(data.bases);
    }
    CreateMiniMap();
    CreateEmittersZone(data.map.emitters);
    CreateAnomalies(data.map.anomalies);
    CreateLabelEntry(data.entry_to_sector);

    let position = document.getElementById("MousePosition");
    if (!position) {

        position = document.createElement("div");
        position.id = "MousePosition";
        document.body.appendChild(position);

        document.body.onmousemove = function (e) {
            position.style.left = (e.pageX + 3) + "px";
            position.style.top = (e.pageY + 3) + "px";
        };
    }

    game.input.addMoveCallback(function () {
        let x = Math.round((game.input.mousePointer.x + game.camera.x) / game.camera.scale.x);
        let y = Math.round((game.input.mousePointer.y + game.camera.y) / game.camera.scale.y);

        position.innerHTML = `y:${x} / y:${y}`
    }, null);
}