function PlaceUnit(jsonMessage) {
    if (JSON.parse(jsonMessage).error === null) {
        var price = document.getElementsByClassName('fieldInfo price');
        price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;
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

    var cells = document.getElementsByClassName("fieldUnit create");
    while (0 < cells.length) {
        if (cells[0]) {
            cells[0].className = "fieldUnit open";
        }
    }
    typeUnit = null;
}