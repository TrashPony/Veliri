function ChassisMouseOver(chassis) {
    var tipChassis = document.getElementById("tipChassis");

    var tdNameChassis = document.getElementById("nameChassis");
    var tdTypeChassis = document.getElementById("typeChassis");
    var tdCarryingChassis = document.getElementById("carryingChassis");
    var tdManeuverabilityChassis = document.getElementById("maneuverabilityChassis");
    var tdMaxSpeed = document.getElementById("maxSpeed");

    tipChassis.style.display = "block";

    tdNameChassis.innerHTML = chassis.name;
    tdTypeChassis.innerHTML = chassis.type;
    tdCarryingChassis.innerHTML = chassis.carrying;
    tdManeuverabilityChassis.innerHTML = chassis.maneuverability;
    tdMaxSpeed.innerHTML = chassis.max_speed;
}

function WeaponMouseOver(weapon) {
    var tipWeapon = document.getElementById("tipWeapon");

    var tdNameWeapon = document.getElementById("nameWeapon");
    var tdTypeWeapon = document.getElementById("typeWeapon");
    var tdDamage = document.getElementById("damageWeapon");
    var tdMinAttackRange = document.getElementById("minAttackRange");
    var tdAccuracy = document.getElementById("accuracy");
    var tdAreaCovers = document.getElementById("areaCovers");
    var tdWeightWeapon = document.getElementById("weightWeapon");
    var tdRangeWeapon = document.getElementById("RangeWeapon");

    tipWeapon.style.display = "block";

    tdNameWeapon.innerHTML = weapon.name;
    tdTypeWeapon.innerHTML = weapon.type;
    tdDamage.innerHTML = weapon.damage;
    tdMinAttackRange.innerHTML = weapon.min_attack_range;
    tdAccuracy.innerHTML = weapon.accuracy;
    tdAreaCovers.innerHTML = weapon.area_covers;
    tdWeightWeapon.innerHTML = weapon.weight;
    tdRangeWeapon.innerHTML = weapon.range;
}

function TowerMouseOver(tower) {
    var tipTower = document.getElementById("tipTower");

    var tdName = document.getElementById("nameTower");
    var tdType = document.getElementById("typeTower");
    var tdHP = document.getElementById("hpTower");
    var tdPower = document.getElementById("powerTower");
    var tdArmor = document.getElementById("armorTower");

    var tdVulToKinetics = document.getElementById("vulToKineticsTower");
    var tdVulToThermo = document.getElementById("vulToThermoTower");
    var tdVulToEM = document.getElementById("vulToEMTower");
    var tdVulToExplosion = document.getElementById("vulToExplosionTower");

    var tdWeight = document.getElementById("weightTower");

    tipTower.style.display = "block";

    tdName.innerHTML = tower.name;
    tdType.innerHTML = tower.type;
    tdHP.innerHTML = tower.hp;
    tdPower.innerHTML = tower.power_radar;
    tdArmor.innerHTML = tower.armor;

    tdVulToKinetics.innerHTML = tower.vul_to_kinetics;
    tdVulToThermo.innerHTML = tower.vul_to_thermo;
    tdVulToEM.innerHTML = tower.vul_to_em;
    tdVulToExplosion.innerHTML = tower.vul_to_explosion;

    tdWeight.innerHTML = tower.weight;
}

function BodyMouseOver(body) {
    var tipBody = document.getElementById("tipBody");

    var tdName = document.getElementById("nameBody");
    var tdType = document.getElementById("typeBody");
    var tdHP = document.getElementById("hpBody");
    var tdMaxTowerWeight = document.getElementById("maxTowerWeightBody");
    var tdArmor = document.getElementById("armorBody");

    var tdVulToKinetics = document.getElementById("vulToKineticsBody");
    var tdVulToThermo = document.getElementById("vulToThermoBody");
    var tdVulToEM = document.getElementById("vulToEMBody");
    var tdVulToExplosion = document.getElementById("vulToExplosionBody");

    var tdWeight = document.getElementById("weightBody");

    tipBody.style.display = "block";

    tdName.innerHTML = body.name;
    tdType.innerHTML = body.type;
    tdHP.innerHTML = body.hp;
    tdMaxTowerWeight.innerHTML = body.max_tower_weight;
    tdArmor.innerHTML = body.armor;

    tdVulToKinetics.innerHTML = body.vul_to_kinetics;
    tdVulToThermo.innerHTML = body.vul_to_thermo;
    tdVulToEM.innerHTML = body.vul_to_em;
    tdVulToExplosion.innerHTML = body.vul_to_explosion;

    tdWeight.innerHTML = body.weight;
}

function RadarMouseOver(radar) {
    var tipRadar = document.getElementById("tipRadar");

    var tdName = document.getElementById("nameRadar");
    var tdType = document.getElementById("typeRadar");
    var tdPower = document.getElementById("powerRadar");
    var tdThrough = document.getElementById("throughRadar");
    var tdAnalysis = document.getElementById("analysisRadar");
    var tdWeight = document.getElementById("weightRadar");

    tipRadar.style.display = "block";

    tdName.innerHTML = radar.name;
    tdType.innerHTML = radar.type;
    tdPower.innerHTML = radar.power;
    tdThrough.innerHTML = radar.through;
    tdAnalysis.innerHTML = radar.analysis;
    tdWeight.innerHTML = radar.weight;

}