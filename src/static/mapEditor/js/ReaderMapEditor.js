function ReaderMapEditor(jsonMessage) {

    if (JSON.parse(jsonMessage).event === "MapList") {
        CreateMapList(jsonMessage)
    }

    if (JSON.parse(jsonMessage).event === "MapSelect") {
        createGame(jsonMessage)
    }

    if (JSON.parse(jsonMessage).event === "getAllTypeCoordinate") {
        createCoordinateMenu(JSON.parse(jsonMessage).type_coordinates)
    }

    if (JSON.parse(jsonMessage).event === "loadNewTypeTerrain") {
        if (JSON.parse(jsonMessage).success) {
            let terrainData = new FormData(document.forms.uploadNewTerrain);
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "http://642e0559eb9c.sn.mynetname.net:8080/upload");
            xhr.send(terrainData);

            mapEditor.send(JSON.stringify({
                event: "getAllTypeCoordinate"
            }));

            if (game) {
                let terrainName = terrainData.get("terrainTexture").name;
                game.load.image(terrainName.substr(0, terrainName.lastIndexOf('.')) || terrainName,
                    'http://' + window.location.host + '/assets/map/terrain/' + terrainName);
                setTimeout(function () {
                    game.load.start();
                }, 1500)
            }
        } else {
            alert("Тип с таким именем уже существует");
        }

        document.forms.uploadNewTerrain.reset();
    }

    if (JSON.parse(jsonMessage).event === "loadNewTypeObject") {
        if (JSON.parse(jsonMessage).success) {
            let objectData = new FormData(document.forms.uploadNewObject);
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "http://642e0559eb9c.sn.mynetname.net:8080/upload");
            xhr.send(objectData);

            mapEditor.send(JSON.stringify({
                event: "getAllTypeCoordinate"
            }));

            if (game) {
                let objectName = objectData.get("objectTexture").name;
                game.load.image(objectName.substr(0, objectName.lastIndexOf('.')) || objectName,
                    'http://' + window.location.host + '/assets/map/objects/' + objectName);
                setTimeout(function () {
                    game.load.start();
                }, 1500)
            }
        } else {
            alert("Тип с таким именем уже существует");
        }

        document.forms.uploadNewObject.reset();
    }
}