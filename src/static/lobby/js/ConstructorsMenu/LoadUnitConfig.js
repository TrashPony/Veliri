function UnitConfig(unit) {
    UpdateUnitInfo(unit);

    if (unit.chassis !== null){
        SelectDetail(unit.chassis, "chassisElement", "picChassis", "picDetail chassis");
    }
    if (unit.weapon !== null){
        SelectDetail(unit.weapon, "weaponElement","picWeapon", "picDetail weapon");
    }
    if (unit.tower !== null){
        SelectDetail(unit.tower, "towerElement", "picTower", "picDetail tower");
    }
    if (unit.body !== null){
        SelectDetail(unit.body, "bodyElement", "picBody", "picDetail body");
    }
    if (unit.radar !== null){
        SelectDetail(unit.radar, "radarElement", "picRadar", "picDetail radar");
    }
}