function SelectedSprite(event, radius, callBack, onlyObj, onlyTexture, transport, notDestroyOnClick, onlyHandler) {

    if (game && game.map) {
        let map = game.map.OneLayerMap;

        destroyAllSelectedSprite(map);

        for (let x in map) {
            if (map.hasOwnProperty(x)) {
                for (let y in map[x]) {
                    if (map[x].hasOwnProperty(y)) {

                        if (onlyObj && !map[x][y].objectSprite) {
                            continue
                        }

                        if (onlyTexture && map[x][y].texture_over_flore === '') {
                            continue
                        }

                        if (transport && !map[x][y].transport) {
                            continue
                        }

                        if (onlyHandler && map[x][y].handler === "") {
                            continue
                        }

                        let selectedSprite = game.SelectLayer.create(Number(x), Number(y), 'mapEditor');
                        selectedSprite.scale.setTo(0.5);
                        selectedSprite.anchor.setTo(0.5);
                        selectedSprite.inputEnabled = true;

                        if (onlyTexture && map[x][y].texture_over_flore !== '') {
                            let style = {font: "24px Arial", fill: "#ff0000", align: "center"};
                            game.add.text(x - 50, y - 15, map[x][y].texture_over_flore, style, game.redactorButton);
                        }

                        map[x][y].selectedSprite = selectedSprite;

                        selectedSprite.events.onInputDown.add(function () {
                            if (game.input.activePointer.leftButton.isDown) {
                                callBack(x, y);
                            }
                            if (!notDestroyOnClick || game.input.activePointer.rightButton.isDown) {
                                destroyAllSelectedSprite();
                            }
                        });

                        selectedSprite.events.onInputOver.add(function () {

                            if (game.input.activePointer.leftButton.isDown) {
                                callBack(x, y);
                            }

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

function destroyAllSelectedSprite() {
    clear(game.redactorButton);
    clear(game.SelectLayer);
}