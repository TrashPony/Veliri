function addHeightCoordinate() {
    if (game.input.activePointer.leftButton.isDown) {
        mapEditor.send(JSON.stringify({
            event: "addHeightCoordinate",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(this.q),
            r: Number(this.r)
        }));
    }
}

function subtractHeightCoordinate() {
    if (game.input.activePointer.leftButton.isDown) {
        mapEditor.send(JSON.stringify({
            event: "subtractHeightCoordinate",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(this.q),
            r: Number(this.r)
        }));
    }
}

function PlaceCoordinate(event, type) {

    if (game && game.map && game.map.OneLayerMap) {
        let map = game.map.OneLayerMap;

        destroyAllSelectedSprite(map);
        
        for (let q in map) {
            if (map.hasOwnProperty(q)) {
                for (let r in map[q]) {
                    if (map[q].hasOwnProperty(r)) {

                        if (map[q][r].impact) {
                            continue
                        }

                        let coordinateSprite = map[q][r].sprite;
                        let selectedSprite = game.SelectLayer.create(coordinateSprite.x, coordinateSprite.y, 'selectEmpty');
                        selectedSprite.anchor.setTo(0.5);
                        selectedSprite.inputEnabled = true;

                        map[q][r].selectedSprite = selectedSprite;

                        selectedSprite.events.onInputDown.add(function () {
                            if (game.input.activePointer.leftButton.isDown) {
                                console.log(this);
                                mapEditor.send(JSON.stringify({
                                    event: event,
                                    id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                                    id_type: Number(type.id),
                                    q: Number(q),
                                    r: Number(r)
                                }));

                                while (game.SelectLayer && game.SelectLayer.children.length > 0) {
                                    let sprite = game.SelectLayer.children.shift();
                                    sprite.destroy();
                                }
                            } else {
                                destroyAllSelectedSprite(map);
                            }
                        });

                        selectedSprite.events.onInputOver.add(function () {
                            if (type.impact_radius > 0) {
                                radiusAnimate(map[q][r], type)
                            } else {
                                selectedSprite.animations.add('select');
                                selectedSprite.animations.play('select', 5, true);
                            }
                        });

                        selectedSprite.events.onInputOut.add(function () {
                            if (type.impact_radius > 0) {
                                stopRadiusAnimate(map[q][r], type)
                            } else {
                                selectedSprite.animations.getAnimation('select').stop(false);
                                selectedSprite.animations.frame = 0;
                            }
                        });
                    }
                }
            }
        }
    }
}

function destroyAllSelectedSprite(map) {
    for (let q in map) {
        if (map.hasOwnProperty(q)) {
            for (let r in map[q]) {
                if (map[q].hasOwnProperty(r)) {
                    if (map[q][r].selectedSprite) {
                        map[q][r].selectedSprite.destroy();
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

function stopRadiusAnimate(center, type) {
    let centerCoordinate = game.map.OneLayerMap[center.q][center.r];
    let circleCoordinates = getRadius(centerCoordinate.x, centerCoordinate.y, centerCoordinate.z, type.impact_radius);
    for (let i in circleCoordinates) {
        let q = circleCoordinates[i].Q;
        let r = circleCoordinates[i].R;
        if (game.map.OneLayerMap.hasOwnProperty(q) && game.map.OneLayerMap[q].hasOwnProperty(r)) {
            let animateCoordinate = game.map.OneLayerMap[q][r].selectedSprite;
            if (animateCoordinate) {
                animateCoordinate.animations.getAnimation('select').stop(false);
                animateCoordinate.animations.frame = 0;
            }
        }
    }
}

function radiusAnimate(center, type) {
    let centerCoordinate = game.map.OneLayerMap[center.q][center.r];
    let circleCoordinates = getRadius(centerCoordinate.x, centerCoordinate.y, centerCoordinate.z, type.impact_radius);
    for (let i in circleCoordinates) {
        let q = circleCoordinates[i].Q;
        let r = circleCoordinates[i].R;
        if (game.map.OneLayerMap.hasOwnProperty(q) && game.map.OneLayerMap[q].hasOwnProperty(r)) {
            let animateCoordinate = game.map.OneLayerMap[q][r].selectedSprite;
            if (animateCoordinate) {
                animateCoordinate.animations.add('select');
                animateCoordinate.animations.play('select', 5, true);
            }
        }
    }
}