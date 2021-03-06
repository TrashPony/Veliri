function CreateRadarObject(mark, object) {
    if (mark.type_object === "transport") {
        let sprite = CreateEvacuation(object.x, object.y, object.base_id, object.id);
        QuickUpEvacuation(sprite, object.rotate)
    }

    if (mark.type_object === "box") {
        CreateBox(object)
    }

    if (mark.type_object === "unit") {
        CreateNewUnit(object)
    }

    if (mark.type_object === "reservoir") {
        CreateReservoir(object, Number(object.x), Number(object.y));
    }

    if (mark.type_object === "dynamic_objects") {

        for (let i in game.objects) {
            if (game.objects[i] && game.objects[i].id === object.id) return;
        }

        if (object.texture !== '') {
            CreateObject(object, object.x, object.y);
        }

        if (object.animate_sprite_sheets !== '') {
            CreateAnimate(object, object.x, object.y);
        }
        game.objects.push(object);
    }
}

function RemoveRadarObject(mark) {
    if (mark.type_object === "transport") {
        for (let idBase in game.bases) {
            for (let idTransport in game.bases[idBase].transports) {
                if (Number(idTransport) === Number(mark.id_object)) {
                    if (game.bases[idBase].transports[idTransport].sprite) {
                        if (game.bases[idBase].transports[idTransport].sprite.shadow) {
                            game.bases[idBase].transports[idTransport].sprite.shadow.destroy();
                        }
                        game.bases[idBase].transports[idTransport].sprite.destroy();
                        game.bases[idBase].transports[idTransport].sprite = null;
                    }
                }
            }
        }
    }

    if (mark.type_object === "box") {
        DestroyBox(Number(mark.id_object), false)
    }

    if (mark.type_object === "unit") {
        let unit = game.units[mark.id_object];
        if (unit) removeUnit(unit);
    }

    if (mark.type_object === "reservoir") {
        for (let x in game.map.reservoir) {
            for (let y in game.map.reservoir[x]) {
                let reservoir = game.map.reservoir[x][y];
                if (reservoir && reservoir.sprite && Number(reservoir.id) === Number(mark.id_object)) {
                    removeReservoir(reservoir);
                    game.map.reservoir[x][y] = null;
                }
            }
        }
    }

    if (mark.type_object === "dynamic_objects") {
        for (let i in game.objects) {
            if (game.objects[i] && Number(game.objects[i].id) === Number(mark.id_object)) {
                if (game.objects[i].objectSprite.shadow) {
                    game.objects[i].objectSprite.shadow.destroy();
                }
                game.objects[i].objectSprite.destroy();

                game.objects[i] = null;
            }
        }
    }
}

function removeReservoir(reservoir) {
    if (reservoir.sprite.shadow) {
        reservoir.sprite.shadow.destroy();
    }
    if (reservoir.border) {
        reservoir.border.destroy();
    }
    if (document.getElementById("reservoirTip" + reservoir.x + "" + reservoir.y)) {
        document.getElementById("reservoirTip" + reservoir.x + "" + reservoir.y).remove();
    }
    reservoir.sprite.destroy();
}

function removeAllObj() {
    for (let i = 0; game.boxes && i < game.boxes.length; i++) {
        DestroyBox(Number(game.boxes[i].id), false)
    }

    for (let i in game.units) {
        removeUnit(game.units[i]);
    }

    for (let x in game.map.reservoir) {
        for (let y in game.map.reservoir[x]) {
            let reservoir = game.map.reservoir[x][y];
            if (reservoir && reservoir.sprite) {
                removeReservoir(reservoir);
                game.map.reservoir[x][y] = null;
            }
        }
    }

    for (let idBase in game.bases) {
        for (let idTransport in game.bases[idBase].transports) {
            if (game.bases[idBase].transports[idTransport].sprite) {
                if (game.bases[idBase].transports[idTransport].sprite.shadow) {
                    game.bases[idBase].transports[idTransport].sprite.shadow.destroy();
                }
                game.bases[idBase].transports[idTransport].sprite.destroy();
                game.bases[idBase].transports[idTransport].sprite = null;
            }
        }
    }

    // for (let i in game.objects) {
    //     if (game.objects[i] && game.objects[i].objectSprite) {
    //         if (game.objects[i].objectSprite.shadow) {
    //             game.objects[i].objectSprite.shadow.destroy();
    //         }
    //         game.objects[i].objectSprite.destroy();
    //         game.objects[i] = null;
    //         game.objects.splice(i, 1);
    //     }
    // }
}

function CreateDynamicObjects(dynamicObjects) {
    for (let x in dynamicObjects) {
        for (let y in dynamicObjects[x]) {
            let object = dynamicObjects[x][y];

            let find = false;
            for (let i in game.objects) {
                if (game.objects[i] && game.objects[i].id === object.id) find = true;
            }

            if (find) continue;

            if (object.texture !== '') {
                CreateObject(object, object.x, object.y);
            }
            if (object.animate_sprite_sheets !== '') {
                CreateAnimate(object, object.x, object.y);
            }
            game.objects.push(object)
        }
    }
}