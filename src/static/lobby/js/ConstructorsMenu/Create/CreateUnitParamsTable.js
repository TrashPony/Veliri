function CreateUnitParams() {
    var unitParams = document.createElement("div");
    unitParams.className = "ConstructorMenu";

    var paramsSpan = document.createElement("div");
    paramsSpan.innerHTML = "Unit params:";
    unitParams.appendChild(paramsSpan);

    var chassisTable = CreateMotionTable();
    unitParams.appendChild(chassisTable);

    var navigationTable = CreateNavigationTable();
    unitParams.appendChild(navigationTable);

    var weaponTable = CreateWeaponTable();
    unitParams.appendChild(weaponTable);

    var survivalTable = CreateSurvivalTable();
    unitParams.appendChild(survivalTable);

    return unitParams
}

function CreateMotionTable() {
    var chassisTableParams = document.createElement("table");
    chassisTableParams.className = "table params";

    var moveSpeed = CreateTableRow("speed", "Speed");
    chassisTableParams.appendChild(moveSpeed);

    var initiative = CreateTableRow("initiative", "Initiative");
    chassisTableParams.appendChild(initiative);

    return chassisTableParams;
}

function CreateWeaponTable() {
    var weaponTableParams = document.createElement("table");
    weaponTableParams.className = "table params";

    var damage = CreateTableRow("damage", "Damage");
    weaponTableParams.appendChild(damage);

    var rangeAttack = CreateTableRow("range_attack", "RangeAttack");
    weaponTableParams.appendChild(rangeAttack);

    var minAttackRange = CreateTableRow("min_attack_range", "MinAttackRange");
    weaponTableParams.appendChild(minAttackRange);

    var areaAttack = CreateTableRow("area_attack", "AreaAttack");
    weaponTableParams.appendChild(areaAttack);

    var typeAttack = CreateTableRow("type_attack", "TypeAttack");
    weaponTableParams.appendChild(typeAttack);

    return weaponTableParams;
}

function CreateSurvivalTable() {
    var survivalTableParams = document.createElement("table");
    survivalTableParams.className = "table params";

    var hp = CreateTableRow("hp", "HP");
    survivalTableParams.appendChild(hp);

    var armor = CreateTableRow("armor", "Armor");
    survivalTableParams.appendChild(armor);

    var evasionCritical = CreateTableRow("evasion_critical", "EvasionCritical");
    survivalTableParams.appendChild(evasionCritical);

    var vulKinetics = CreateTableRow("vul_kinetics", "VulKinetics");
    survivalTableParams.appendChild(vulKinetics);

    var vulThermal = CreateTableRow("vul_thermal", "VulThermal");
    survivalTableParams.appendChild(vulThermal);

    var vulEM = CreateTableRow("vul_em", "VulEM");
    survivalTableParams.appendChild(vulEM);

    var vulExplosive = CreateTableRow("vul_explosive", "VulExplosive");
    survivalTableParams.appendChild(vulExplosive);

    return survivalTableParams;
}

function CreateNavigationTable() {
    var navigationTableParams = document.createElement("table");
    navigationTableParams.className = "table params";

    var rangeView = CreateTableRow("range_view", "RangeView");
    navigationTableParams.appendChild(rangeView);

    var accuracy = CreateTableRow("accuracy", "Accuracy");
    navigationTableParams.appendChild(accuracy);

    var wallHack = CreateTableRow("wall_hack", "WallHack");
    navigationTableParams.appendChild(wallHack);

    return navigationTableParams;
}

function CreateTableRow(id, value) {
    var tr = document.createElement("tr");
    var tdName = document.createElement("td");
    tdName.innerHTML = value;
    var tdValue = document.createElement("td");
    tdValue.innerHTML = "0";
    tdValue.id = id;
    tr.appendChild(tdName);
    tr.appendChild(tdValue);

    return tr;
}