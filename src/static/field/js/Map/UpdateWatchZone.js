function UpdateWatchZone(watch) {
    if (watch) {
        var closeCoordinate = watch.close_coordinate;
        var openCoordinate = watch.open_coordinate;
        var openUnits = watch.open_unit;
        var openMatherShip = watch.open_mather_ship;

        if (closeCoordinate) {
            CloseCoordinates(closeCoordinate);
        }

        if (openCoordinate) {
            OpenCoordinates(openCoordinate);
        }

        if (openUnits) {
            while (openUnits.length > 0) {
                var openUnit = openUnits.shift();
                CreateUnit(openUnit)
            }
        }

        if (openMatherShip) {
            // TODO добавить структуры
        }
    }
}