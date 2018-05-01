function UnitConfig(unit) {
    UpdateUnitInfo(unit);
    if (unit !== undefined) {
        if (unit.chassis !== null) {
            SelectDetail(unit.chassis, "chassisElement", ChassisMouseOver);
        }
        if (unit.weapon !== null) {
            SelectDetail(unit.weapon, "weaponElement", WeaponMouseOver);
        }
        if (unit.tower !== null) {
            SelectDetail(unit.tower, "towerElement", TowerMouseOver);
        }
        if (unit.body !== null) {
            SelectDetail(unit.body, "bodyElement", BodyMouseOver);
        }
        if (unit.radar !== null) {
            SelectDetail(unit.radar, "radarElement", RadarMouseOver);
        }
    }
}