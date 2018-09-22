function UpdateUnit(jsonMessage) {

    let unitStat = JSON.parse(jsonMessage).unit;
    let unit = GetGameUnitID(unitStat.id);

    RemoveSelect();
    DeleteMarkTarget(unitStat);

    unit.rotate = unitStat.rotate;
    unit.target = unitStat.target;
    unit.effects = unitStat.effects;
    unit.defend = unitStat.defend;

    // todo
}