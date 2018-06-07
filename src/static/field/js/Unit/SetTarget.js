function SetTarget(jsonMessage) {
    var unitStat = JSON.parse(jsonMessage).unit;
    var unit = GetGameUnitID(unitStat.id);

    unit.rotate = unitStat.rotate;
    unit.target = unitStat.target;
    // todo
    console.log();
}