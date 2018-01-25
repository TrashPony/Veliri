function InitCreateUnit() {
    var mask = document.getElementById("mask");
    mask.style.display = "block";

    var lobbyMenu = document.getElementById("lobby");

    var unitConstructor = document.getElementById("unitConstructor");

    if (!unitConstructor) {
        unitConstructor = document.createElement("div");
        unitConstructor.id = "unitConstructor";
        lobbyMenu.appendChild(unitConstructor);

        var chassisMenu = CreateChassisMenu();
        var weaponMenu = CreateWeaponMenu();
        var unitMenu = CreateUnitMenu();

        unitConstructor.appendChild(chassisMenu);
        unitConstructor.appendChild(unitMenu);
        unitConstructor.appendChild(weaponMenu);

        lobby.send(JSON.stringify({
            event: "GetDetailOfUnits"
        }));

    } else {
        unitConstructor.style.display = "block";
    }
}

function CreateChassisMenu() {

    var chassisMenu = document.createElement("div");
    chassisMenu.id = "chassisMenu";
    chassisMenu.className = "ConstructorMenu";

    var chassisSpan = document.createElement("span");
    chassisSpan.className = "Value";
    chassisSpan.innerHTML = "Шасси:";
    chassisMenu.appendChild(chassisSpan);

    return chassisMenu;
}

function CreateWeaponMenu() {

    var weaponMenu = document.createElement("div");
    weaponMenu.id = "weaponMenu";
    weaponMenu.className = "ConstructorMenu";

    var weaponSpan = document.createElement("span");
    weaponSpan.className = "Value";
    weaponSpan.innerHTML = "Башня:";
    weaponMenu.appendChild(weaponSpan);

    return weaponMenu;
}

function CreateUnitMenu() {
    var unitMenu = document.createElement("div");
    unitMenu.id = "unitMenu";

    var unitSpan = document.createElement("span");
    unitSpan.className = "Value";
    unitSpan.innerHTML = "Юнит";
    unitMenu.appendChild(unitSpan);

    var chassisUnitBox = document.createElement("div");
    chassisUnitBox.className = "ElementUnitBox left";
    unitMenu.appendChild(chassisUnitBox);

    var weaponUnitBox = document.createElement("div");
    weaponUnitBox.className = "ElementUnitBox right";
    unitMenu.appendChild(weaponUnitBox);

    var picUnit = document.createElement("div");
    picUnit.id = "picUnit";
    unitMenu.appendChild(picUnit);

    var chassisTable = CreateChassisTable();
    unitMenu.appendChild(chassisTable);

    var weaponTable = CreateWeaponTable();
    unitMenu.appendChild(weaponTable);

    var acceptButton = document.createElement("input");
    acceptButton.type = "button";
    acceptButton.value = "Accept";
    acceptButton.className = "lobbyButton";
    acceptButton.id = "acceptButton";
    acceptButton.onclick = BackToLobby;
    unitMenu.appendChild(acceptButton);


    return unitMenu;
}

function BackToLobby() {
    var mask = document.getElementById("mask");
    mask.style.display = "none";

    var unitConstructor = document.getElementById("unitConstructor");
    unitConstructor.style.display = "none";
}

function CreateChassisTable() {
    var chassisTableParams = document.createElement("table");
    chassisTableParams.className = "table params chassis";

    var hp = CreateTableRow("HP", "HP");
    chassisTableParams.appendChild(hp);

    var moveSpeed = CreateTableRow("Speed", "Speed");
    chassisTableParams.appendChild(moveSpeed);

    var initiative = CreateTableRow("Initiative", "Initiative");
    chassisTableParams.appendChild(initiative);

    return chassisTableParams;
}

function CreateWeaponTable() {
    var weaponTableParams = document.createElement("table");
    weaponTableParams.className = "table params weapon";

    var damage = CreateTableRow("Damage", "Damage");
    weaponTableParams.appendChild(damage);

    var rangeAttack = CreateTableRow("RangeAttack", "RangeAttack");
    weaponTableParams.appendChild(rangeAttack);

    var rangeView = CreateTableRow("RangeView", "RangeView");
    weaponTableParams.appendChild(rangeView);

    var areaAttack = CreateTableRow("AreaAttack", "AreaAttack");
    weaponTableParams.appendChild(areaAttack);

    var typeAttack = CreateTableRow("TypeAttack", "TypeAttack");
    weaponTableParams.appendChild(typeAttack);

    return weaponTableParams;
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

function moveTip(e) {

    var tipWeapon = document.getElementById("tipWeapon").style;
    var tipChassis = document.getElementById("tipChassis").style;

    var w = 250; // Ширина слоя
    var x = e.pageX; // Координата X курсора
    var y = e.pageY; // Координата Y курсора

    if ((x + w + 10) < document.body.clientWidth) {
        // Показывать слой справа от курсора
        tipChassis.left = x + 'px';
        tipWeapon.left = x + 'px';
    } else {
        // Показывать слой слева от курсора
        tipChassis.left = x - w + 'px';
        tipWeapon.left = x - w + 'px';
    }
    // Положение от верхнего края окна браузера
    tipChassis.top = y + 20 + 'px';
    tipWeapon.top = y + 20 + 'px';
}