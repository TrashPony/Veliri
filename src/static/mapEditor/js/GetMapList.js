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

    for (let i in maps) {
        if (maps.hasOwnProperty(i)) {
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

        setTimeout(function () {
            CreateLabelBase(JSON.parse(jsonMessage).bases)
        }, 1500);
    }
}

function CreateLabelBase(bases) {
    for (let i in bases){
        if (bases.hasOwnProperty(i) && game.map.OneLayerMap.hasOwnProperty(bases[i].q) && game.map.OneLayerMap.hasOwnProperty(bases[i].r)) {
            let coordinateBase = game.map.OneLayerMap[bases[i].q][bases[i].r].sprite;
            let base = game.icon.create(coordinateBase.x, coordinateBase.y, 'baseIcon');
            base.anchor.setTo(0.5);
            base.scale.setTo(0.1);

            if (game.map.OneLayerMap.hasOwnProperty(bases[i].resp_q) && game.map.OneLayerMap.hasOwnProperty(bases[i].resp_r)) {
                let coordinateResp = game.map.OneLayerMap[bases[i].resp_q][bases[i].resp_r].sprite;
                let baseResp = game.icon.create(coordinateResp.x, coordinateResp.y, 'baseResp');
                baseResp.anchor.setTo(0.5);
                baseResp.scale.setTo(0.05);
            }
        }
    }
}