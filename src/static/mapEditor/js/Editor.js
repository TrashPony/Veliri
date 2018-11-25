function addHeightCoordinate() {
    mapEditor.send(JSON.stringify({
        event: "addHeightCoordinate",
        id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
        q: Number(this.q),
        r: Number(this.r)
    }));
}

function subtractHeightCoordinate() {
    mapEditor.send(JSON.stringify({
        event: "subtractHeightCoordinate",
        id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
        q: Number(this.q),
        r: Number(this.r)
    }));
}

function PlaceCoordinate(event) {
    if (game && game.map && game.map.OneLayerMap) {
        let map = game.map.OneLayerMap;

        for (let q in map) {
            if (map.hasOwnProperty(q)) {
                for (let r in map[q]) {
                    if (map[q].hasOwnProperty(r)) {

                        let coordinateSprite = map[q][r].sprite;
                        let selectedSprite = game.SelectLayer.create(coordinateSprite.x + 22, coordinateSprite.y + 15, 'selectEmpty');

                        selectedSprite.inputEnabled = true;

                        selectedSprite.events.onInputDown.add(function () {
                            mapEditor.send(JSON.stringify({
                                event: event,
                                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                                id_type:  Number(this.id),
                                q: Number(q),
                                r: Number(r)
                            }));

                            while (game.SelectLayer && game.SelectLayer.children.length > 0) {
                                let sprite = game.SelectLayer.children.shift();
                                sprite.destroy();
                            }
                        });

                        selectedSprite.events.onInputOver.add(function () {
                            selectedSprite.animations.add('select');
                            selectedSprite.animations.play('select', 5, true);
                        });

                        selectedSprite.events.onInputOut.add(function () {
                            selectedSprite.animations.getAnimation('select').stop(false);
                            selectedSprite.animations.frame = 0;
                        });
                    }
                }
            }
        }
    }
}

function SendCommand(command) {
    mapEditor.send(JSON.stringify({
        event: command,
        id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value)
    }));
}