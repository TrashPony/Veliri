function UnitConfig(unit) {
    UpdateUnitInfo(unit);
    if (unit !== undefined) {
        if (unit.chassis !== null) {
            SelectDetail(unit.chassis, "chassisElement", "picChassis", "picDetail chassis", ChassisMouseOver);
        }
        if (unit.weapon !== null) {
            SelectDetail(unit.weapon, "weaponElement", "picWeapon", "picDetail weapon", WeaponMouseOver);
        }
        if (unit.tower !== null) {
            SelectDetail(unit.tower, "towerElement", "picTower", "picDetail tower", TowerMouseOver);
        }
        if (unit.body !== null) {
            SelectDetail(unit.body, "bodyElement", "picBody", "picDetail body", BodyMouseOver);
        }
        if (unit.radar !== null) {
            SelectDetail(unit.radar, "radarElement", "picRadar", "picDetail radar", RadarMouseOver);
        }
    }
}