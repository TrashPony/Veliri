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
    for (let i in maps) {

        let option = document.createElement("option");
        option.id = maps[i].id;
        option.value = maps[i].id;

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
        UpdateMap(JSON.parse(jsonMessage).map, game, JSON.parse(jsonMessage).bases);
    } else {

        let loadFunc = function () {
            if (JSON.parse(jsonMessage).bases) {
                CreateLabelBase(JSON.parse(jsonMessage).bases);
            }
            CreateMiniMap();
            CreateGeoData(JSON.parse(jsonMessage).map.geo_data);
            CreateEmittersZone(JSON.parse(jsonMessage).map.emitters);
            CreateAnomalies(JSON.parse(jsonMessage).map.anomalies)
        };

        game = CreateGame(JSON.parse(jsonMessage).map, loadFunc, "mapEditor");
        game.bases = JSON.parse(jsonMessage).bases;
    }
}

function CreateLabelBase(bases) {
    for (let i in bases) {
        if (bases.hasOwnProperty(i) && game.map.OneLayerMap.hasOwnProperty(bases[i].q) && game.map.OneLayerMap.hasOwnProperty(bases[i].r)) {

            let xy = GetXYCenterHex(bases[i].q, bases[i].r);

            let base = game.icon.create(xy.x, xy.y, 'baseIcon');
            base.anchor.setTo(0.5);
            base.scale.setTo(0.1);

            if (game.map.OneLayerMap.hasOwnProperty(bases[i].resp_q) && game.map.OneLayerMap.hasOwnProperty(bases[i].resp_r)) {
                let xy = GetXYCenterHex(bases[i].resp_q, bases[i].resp_r);
                let baseResp = game.icon.create(xy.x, xy.y, 'baseResp');
                baseResp.anchor.setTo(0.5);
                baseResp.scale.setTo(0.05);
            }
        }
    }
}