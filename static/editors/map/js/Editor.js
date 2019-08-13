function addHeightCoordinate(q, r) {
    if (game.input.activePointer.leftButton.isDown) {
        mapEditor.send(JSON.stringify({
            event: "addHeightCoordinate",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(q),
            r: Number(r)
        }));
    }

    let xy = GetXYCenterHex(q, r);
    game.map.OneLayerMap[q][r].level++;
    CreateTerrain(game.map.OneLayerMap[q][r], xy.x, xy.y, q, r)
}

function subtractHeightCoordinate(q, r) {
    if (game.input.activePointer.leftButton.isDown) {
        mapEditor.send(JSON.stringify({
            event: "subtractHeightCoordinate",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(q),
            r: Number(r)
        }));
    }

    let xy = GetXYCenterHex(q, r);
    game.map.OneLayerMap[q][r].level--;
    CreateTerrain(game.map.OneLayerMap[q][r], xy.x, xy.y, q, r)
}

function PlaceCoordinate(event, type) {

    let newType = Object.assign({}, type);

    let callBack = function (q, r) {
        mapEditor.send(JSON.stringify({
            event: event,
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            id_type: Number(type.id),
            q: Number(q),
            r: Number(r)
        }));

        newType.q = q;
        newType.r = r;

        newType.scale = 100;
        newType.shadow = false;
        newType.x_shadow_offset = 10;
        newType.y_shadow_offset = 10;
        newType.level = game.map.OneLayerMap[q][r].level;
        newType.coordinateText = game.map.OneLayerMap[q][r].coordinateText;

        for (let i in game.mapPoints) {

            if (game.mapPoints[i].q === Number(q) && game.mapPoints[i].r === Number(r)) {

                if (game.mapPoints[i].coordinate.objectSprite) {
                    if (game.mapPoints[i].coordinate.objectSprite.shadow) {
                        game.mapPoints[i].coordinate.objectSprite.shadow.destroy();
                    }
                    game.mapPoints[i].coordinate.objectSprite.destroy();
                }

                game.mapPoints[i].coordinate = newType;
                game.map.OneLayerMap[q][r] = newType;

                ReloadCoordinate(game.mapPoints[i]);
            }
        }
    };

    SelectedSprite(event, newType.impact_radius, callBack, null, null, null, true)
}


function SendCommand(command) {
    mapEditor.send(JSON.stringify({
        event: command,
        id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value)
    }));
}