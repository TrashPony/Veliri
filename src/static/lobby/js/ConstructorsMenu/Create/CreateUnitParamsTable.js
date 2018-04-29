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

    var moveSpeed = CreateTableRow("Speed", "Speed");
    chassisTableParams.appendChild(moveSpeed);

    var initiative = CreateTableRow("Initiative", "Initiative");
    chassisTableParams.appendChild(initiative);

    return chassisTableParams;
}

function CreateWeaponTable() {
    var weaponTableParams = document.createElement("table");
    weaponTableParams.className = "table params";

    var damage = CreateTableRow("Damage", "Damage");
    weaponTableParams.appendChild(damage);

    var rangeAttack = CreateTableRow("RangeAttack", "RangeAttack");
    weaponTableParams.appendChild(rangeAttack);

    var minAttackRange = CreateTableRow("MinAttackRange", "MinAttackRange");
    weaponTableParams.appendChild(minAttackRange);

    var areaAttack = CreateTableRow("AreaAttack", "AreaAttack");
    weaponTableParams.appendChild(areaAttack);

    var typeAttack = CreateTableRow("TypeAttack", "TypeAttack");
    weaponTableParams.appendChild(typeAttack);

    return weaponTableParams;
}

function CreateSurvivalTable() {
    var survivalTableParams = document.createElement("table");
    survivalTableParams.className = "table params";

    var hp = CreateTableRow("HP", "HP");
    survivalTableParams.appendChild(hp);

    var armor = CreateTableRow("Armor", "Armor");
    survivalTableParams.appendChild(armor);

    var evasionCritical = CreateTableRow("EvasionCritical", "EvasionCritical");
    survivalTableParams.appendChild(evasionCritical);

    var vulKinetics = CreateTableRow("VulKinetics", "VulKinetics");
    survivalTableParams.appendChild(vulKinetics);

    var vulThermal = CreateTableRow("VulThermal", "VulThermal");
    survivalTableParams.appendChild(vulThermal);

    var vulEM = CreateTableRow("VulEM", "VulEM");
    survivalTableParams.appendChild(vulEM);

    var vulExplosive = CreateTableRow("VulExplosive", "VulExplosive");
    survivalTableParams.appendChild(vulExplosive);

    return survivalTableParams;
}

function CreateNavigationTable() {
    var navigationTableParams = document.createElement("table");
    navigationTableParams.className = "table params";

    var rangeView = CreateTableRow("RangeView", "RangeView");
    navigationTableParams.appendChild(rangeView);

    var accuracy = CreateTableRow("Accuracy", "Accuracy");
    navigationTableParams.appendChild(accuracy);

    var wallHack = CreateTableRow("WallHack", "WallHack");
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