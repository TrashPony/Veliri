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
}

function RemoveRadarObject(mark) {
    if (mark.type_object === "transport") {
        for (let idBase in game.bases) {
            for (let idTransport in game.bases[idBase].transports) {
                if (Number(idTransport) === Number(mark.id_object)) {
                    game.bases[idBase].transports[idTransport].sprite.shadow.destroy();
                    game.bases[idBase].transports[idTransport].sprite.destroy();
                    game.bases[idBase].transports[idTransport].sprite = null;
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
}