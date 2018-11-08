function UpdateUnit(unitStat) {

    let unit = GetGameUnitID(unitStat.id);

    if (unitStat.owner === game.user.name) {
        RemoveSelect();
        DeleteMarkTarget(unitStat);
    }

    if (!unit) {
        unit = GetStorageUnit(unitStat.id)
    }

    unit.rotate = unitStat.rotate;
    unit.target = unitStat.target;
    unit.effects = unitStat.effects;
    unit.defend = unitStat.defend;
    unit.body = unitStat.body;
    unit.power = unitStat.power;
    unit.accuracy = unitStat.accuracy;
    unit.action_points = unitStat.action_points;
    unit.armor = unitStat.armor;
    unit.evasion_critical = unitStat.evasion_critical;
    unit.hp = unitStat.hp;
    unit.initiative = unitStat.initiative;
    unit.max_hp = unitStat.max_hp;
    unit.max_power = unitStat.max_power;
    unit.q = unitStat.q;
    unit.r = unitStat.r;
    unit.range_view = unitStat.range_view;
    unit.move = unitStat.move;
    unit.recovery_HP = unitStat.recovery_HP;
    unit.recovery_power = unitStat.recovery_power;
    unit.speed = unitStat.speed;
    unit.wall_hack = unitStat.wall_hack;
    unit.vul_to_em = unitStat.vul_to_em;
    unit.vul_to_explosion = unitStat.vul_to_explosion;
    unit.vul_to_kinetics = unitStat.vul_to_kinetics;
    unit.vul_to_thermo = unitStat.vul_to_thermo;

    if (game.Phase !== "targeting") {
        CreateTargetLine(unitStat)
    }
}

function UpdateMemoryUnit(jsonMessage) {
    game.memoryHostileUnit = JSON.parse(jsonMessage).memory_hostile_unit;
    LoadQueueUnits();
}