function SelectCoordinateUnitCreate(jsonMessage) {
    var place_coordinate = JSON.parse(jsonMessage).place_coordinate;

    for (var i = 0; i < place_coordinate.length; i++) {
        game.map.OneLayerMap[place_coordinate[i].x][place_coordinate[i].y].sprite.tint = 0x0000ff * 2;
    }
}

function RemoveSelectCoordinateUnitCreate() {

}