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
            CreateAnomalies(JSON.parse(jsonMessage).map.anomalies);
            CreateLabelEntry(JSON.parse(jsonMessage).entry_to_sector);
        };

        game = CreateGame(JSON.parse(jsonMessage).map, loadFunc, "mapEditor");
        game.bases = JSON.parse(jsonMessage).bases;
    }
}

function CreateLabelEntry(entryPoints) {
    for (let i of entryPoints) {
        for (let position of i.positions) {
            let xy = GetXYCenterHex(position.q, position.r);
            let baseResp = game.icon.create(xy.x, xy.y, 'baseResp');
            baseResp.angle = position.resp_rotate;
            baseResp.anchor.setTo(0.5);
            baseResp.scale.setTo(0.05);
        }
    }
}

function CreateLabelBase(bases) {
    for (let i in bases) {
        if (bases.hasOwnProperty(i) && game.map.OneLayerMap.hasOwnProperty(bases[i].q) && game.map.OneLayerMap.hasOwnProperty(bases[i].r)) {

            let xy = GetXYCenterHex(bases[i].q, bases[i].r);

            let base = game.icon.create(xy.x, xy.y, 'baseIcon');
            base.anchor.setTo(0.5);
            base.scale.setTo(0.1);

            for (let j in bases[i].respawns) {
                let respPount = bases[i].respawns[j];
                let xy = GetXYCenterHex(respPount.q, respPount.r);
                let baseResp = game.icon.create(xy.x, xy.y, 'baseResp');

                baseResp.angle = respPount.resp_rotate;
                baseResp.anchor.setTo(0.5);
                baseResp.scale.setTo(0.05);
            }
        }
    }
}