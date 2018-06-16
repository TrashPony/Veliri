function UpdateUnit(jsonMessage) {

    var unitStat = JSON.parse(jsonMessage).unit;
    var unit = GetGameUnitID(unitStat.id);

    RemoveSelect();
    DeleteMarkTarget(unitStat);

    unit.rotate = unitStat.rotate;
    unit.target = unitStat.target;
    unit.effect = unitStat.effect;

    // todo
}