function UpdateWatchZone(watch) {
    if (watch) {
        let closeCoordinate = watch.close_coordinate;
        let openCoordinate = watch.open_coordinate;
        let openUnits = watch.open_unit;

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
    }
}