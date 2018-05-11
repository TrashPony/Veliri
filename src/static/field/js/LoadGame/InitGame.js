var idGame;

function InitGame() {
    idGame = getCookie("idGame");
    field.send(JSON.stringify({
        event: "InitGame",
        id_game: Number(idGame)
    }));
}

function getCookie(name) {
    var matches = document.cookie.match(new RegExp(
        "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
}

function LoadGame(jsonMessage) {
    var gameInfo = JSON.parse(jsonMessage).game_info;
    GameInfo(gameInfo);

    var map = JSON.parse(jsonMessage).map;
    FieldCreate(map);

    var userName = JSON.parse(jsonMessage).user_name;
    var ready = JSON.parse(jsonMessage).ready;
    var equip = JSON.parse(jsonMessage).equip;
    var units = JSON.parse(jsonMessage).units;
    var hostileUnits = JSON.parse(jsonMessage).hostile_units;
    var unitStorage = JSON.parse(jsonMessage).unit_storage;
    var matherShip = JSON.parse(jsonMessage).mather_ship;
    var hostileMatherShips = JSON.parse(jsonMessage).hostile_mather_ships;

    var watch = JSON.parse(jsonMessage).watch;
    LoadOpenCoordinate(watch); // todo карта не успевает загружаться
}
