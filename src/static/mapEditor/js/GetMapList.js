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
        if (maps[i].global) {
            option.innerHTML = maps[i].Name + "<span style='color: red'> Глоб.</span>";
        } else {
            option.innerHTML = maps[i].Name;
        }
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

function createGame(jsonMessage) {
    if (game) {
        UpdateMap(JSON.parse(jsonMessage).map, game);
    } else {
        game = CreateGame(JSON.parse(jsonMessage).map);
        appendRedactorEventsToFloor(game)
    }
}

function appendRedactorEventsToFloor(game) {
    let map = game.map.OneLayerMap;
}

function addButtons(map) {
    for (let q in map) {
        if (map.hasOwnProperty(q)) {
            for (let r in map[q]) {
                if (map[q].hasOwnProperty(r)) {

                    if (map[q][r].impact) {
                        continue
                    }


                }
            }
        }
    }
}