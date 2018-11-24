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

    for (let i = 0; i < maps.length; i++) {
        if (document.getElementById(maps[i].Id)) {
            continue;
        }
        let option = document.createElement("option");
        option.id = maps[i].Id;
        option.value = maps[i].Id;
        option.innerHTML = maps[i].Name;
        mapSelect.appendChild(option);
    }
}

function selectMap() {
    let mapSelector = document.getElementById("mapSelector");
    let selectedValue = mapSelector.options[mapSelector.selectedIndex].value;

    mapEditor.send(JSON.stringify({
        event: "SelectMap",
        id: Number(selectedValue)
    }));
}

function createGame(jsonMessage) {
    game = CreateGame(JSON.parse(jsonMessage).map);

    // TODO создание эввентов на удаление терейнов game.map
}