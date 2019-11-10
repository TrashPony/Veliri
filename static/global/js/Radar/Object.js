function CreateRadarObject(mark, object) {

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
}