function SelectedSprite(event, radius, callBack, onlyObj) {
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

                        if (!map[q][r].sprite) {
                            continue
                        }

                        if (onlyObj && !map[q][r].objectSprite) {
                            continue
                        }

                        let coordinateSprite = map[q][r].sprite;
                        let selectedSprite = game.SelectLayer.create(coordinateSprite.x, coordinateSprite.y, 'selectEmpty');
                        selectedSprite.anchor.setTo(0.5);
                        selectedSprite.inputEnabled = true;

                        map[q][r].selectedSprite = selectedSprite;

                        selectedSprite.events.onInputDown.add(function () {
                            if (game.input.activePointer.leftButton.isDown) {
                                callBack(q, r);
                                destroyAllSelectedSprite(map);
                            } else {
                                destroyAllSelectedSprite(map);
                            }
                        });

                        selectedSprite.events.onInputOver.add(function () {
                            if (radius > 0) {
                                radiusAnimate(map[q][r], radius)
                            } else {
                                selectedSprite.animations.add('select');
                                selectedSprite.animations.play('select', 5, true);
                            }
                        });

                        selectedSprite.events.onInputOut.add(function () {
                            if (radius > 0) {
                                stopRadiusAnimate(map[q][r], radius)
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

function stopRadiusAnimate(center, radius) {
    let centerCoordinate = game.map.OneLayerMap[center.q][center.r];
    let circleCoordinates = getRadius(centerCoordinate.x, centerCoordinate.y, centerCoordinate.z, radius);
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

function radiusAnimate(center, radius) {
    let centerCoordinate = game.map.OneLayerMap[center.q][center.r];
    let circleCoordinates = getRadius(centerCoordinate.x, centerCoordinate.y, centerCoordinate.z, radius);
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