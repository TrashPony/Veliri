function PlaceUnit(jsonMessage) {
    if (JSON.parse(jsonMessage).error === null || JSON.parse(jsonMessage).error === undefined) {

        CreateUnit(JSON.parse(jsonMessage).unit);

        var boxUnit = document.getElementById(JSON.parse(jsonMessage).unit.id);
        boxUnit.remove();

    } else {
        var log = document.getElementById('fieldLog');

        if (JSON.parse(jsonMessage).error_type === "busy") {
            log.innerHTML = "Место занято"
        }
        if (JSON.parse(jsonMessage).error_type === "no many") {
            log.innerHTML = "Нет денег"
        }
        if (JSON.parse(jsonMessage).error_type === "not allow") {
            log.innerHTML = "Не разрешено"
        }
    }
}