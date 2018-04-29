function InitCreateUnit() {
    var mask = document.getElementById("mask");
    mask.style.display = "block";
    // TODO передалать конструктор в полную таблицу, а то дивы нахер расползаються
    var lobbyMenu = document.getElementById("lobby");

    var unitConstructor = document.getElementById("unitConstructor");

    if (!unitConstructor) {
        unitConstructor = document.createElement("div");
        unitConstructor.id = "unitConstructor";
        lobbyMenu.appendChild(unitConstructor);

        var constructorTable = document.createElement("table");
        constructorTable.id = "constructorTable";
        var constructorTr = document.createElement("tr");
        var tdUnitParams = document.createElement("td");
        tdUnitParams.className = "ConstructorTD";
        var tdUnitMenu = document.createElement("td");
        tdUnitMenu.className = "ConstructorTD";
        var tdTabDetailMenu = document.createElement("td");
        tdTabDetailMenu.className = "ConstructorTD";

        constructorTable.appendChild(constructorTr);
        constructorTr.appendChild(tdUnitParams);
        constructorTr.appendChild(tdUnitMenu);
        constructorTr.appendChild(tdTabDetailMenu);

        var unitMenu = CreateUnitMenu();
        var unitParams =  CreateUnitParams();
        var tabDetailMenu = CreateTabDetailMenu();

        tdUnitParams.appendChild(unitParams);
        tdUnitMenu.appendChild(unitMenu);
        tdTabDetailMenu.appendChild(tabDetailMenu);

        unitConstructor.appendChild(constructorTable);

        lobby.send(JSON.stringify({
            event: "GetDetailOfUnits"
        }));

    } else {
        unitConstructor.style.display = "block";
    }
}


function CreateUnitParams() {
    var unitParams = document.createElement("div");
    unitParams.className = "ConstructorMenu";

    var paramsSpan = document.createElement("div");
    paramsSpan.innerHTML = "Unit params:";
    unitParams.appendChild(paramsSpan);

    var chassisTable = CreateChassisTable();
    unitParams.appendChild(chassisTable);

    var weaponTable = CreateWeaponTable();
    unitParams.appendChild(weaponTable);

    return unitParams
}

function CreateUnitMenu() {
    var unitMenu = document.createElement("div");
    unitMenu.id = "unitMenu";

    var unitSpan = document.createElement("span");
    unitSpan.innerHTML = "Юнит";
    unitMenu.appendChild(unitSpan);

    var bodyUnitBox = document.createElement("div");
    bodyUnitBox.className = "ElementUnitBox right";
    bodyUnitBox.id = "bodyElement";
    bodyUnitBox.onclick = function () {
        this.style.backgroundImage = "";
        this.body = null;
        var picBody = document.getElementById("picBody");
        if (picBody){
            picBody.remove();
        }
    };
    unitMenu.appendChild(bodyUnitBox);

    var towerUnitBox = document.createElement("div");
    towerUnitBox.className = "ElementUnitBox right";
    towerUnitBox.id = "towerElement";
    towerUnitBox.onclick = function () {
        this.style.backgroundImage = "";
        this.body = null;
        var picTower = document.getElementById("picTower");
        if (picTower){
            picTower.remove();
        }
    };
    unitMenu.appendChild(towerUnitBox);

    var chassisUnitBox = document.createElement("div");
    chassisUnitBox.className = "ElementUnitBox left";
    chassisUnitBox.id = "chassisElement";
    chassisUnitBox.onclick = function () {
        this.style.backgroundImage = "";
        this.chassis = null;
        var picChassis = document.getElementById("picChassis");
        if (picChassis){
            picChassis.remove();
        }
    };
    unitMenu.appendChild(chassisUnitBox);

    var weaponUnitBox = document.createElement("div");
    weaponUnitBox.className = "ElementUnitBox weapon";
    weaponUnitBox.id = "weaponElement";
    weaponUnitBox.onclick = function () {
        this.style.backgroundImage = "";
        this.weapon = null;
        var picWeapon = document.getElementById("picWeapon");
        if (picWeapon){
            picWeapon.remove();
        }
    };
    unitMenu.appendChild(weaponUnitBox);

    var radarUnitBox = document.createElement("div");
    radarUnitBox.className = "ElementUnitBox left";
    radarUnitBox.id = "radarElement";
    radarUnitBox.onclick = function () {
        this.style.backgroundImage = "";
        this.radar = null;
        var picRadar = document.getElementById("picRadar");
        if (picRadar){
            picRadar.remove();
        }
    };
    unitMenu.appendChild(radarUnitBox);

    var picUnit = document.createElement("div");
    picUnit.id = "picUnit";
    unitMenu.appendChild(picUnit);

    var acceptButton = document.createElement("input");
    acceptButton.type = "button";
    acceptButton.value = "Accept";
    acceptButton.className = "lobbyButton";
    acceptButton.id = "acceptButton";
    acceptButton.onclick = BackToLobby;
    unitMenu.appendChild(acceptButton);

    return unitMenu;
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
    var tipTower = document.getElementById("tipTower").style;
    var tipBody = document.getElementById("tipBody").style;
    var tipRadar = document.getElementById("tipRadar").style;


    var w = 250; // Ширина слоя
    var x = e.pageX; // Координата X курсора
    var y = e.pageY; // Координата Y курсора

    if ((x + w + 10) < document.body.clientWidth) {
        // Показывать слой справа от курсора
        tipChassis.left = x + 'px';
        tipWeapon.left = x + 'px';
        tipTower.left = x + 'px';
        tipRadar.left = x + 'px';
        tipBody.left = x + 'px';
    } else {
        // Показывать слой слева от курсора
        tipChassis.left = x - w + 'px';
        tipWeapon.left = x - w + 'px';
        tipTower.left = x - w + 'px';
        tipRadar.left = x - w + 'px';
        tipBody.left = x - w + 'px';
    }
    // Положение от верхнего края окна браузера
    tipChassis.top = y + 20 + 'px';
    tipWeapon.top = y + 20 + 'px';
    tipTower.top = y + 20 + 'px';
    tipRadar.top = y + 20 + 'px';
    tipBody.top = y + 20 + 'px';

}

function BackToLobby() {
    var mask = document.getElementById("mask");
    mask.style.display = "none";

    var unitConstructor = document.getElementById("unitConstructor");
    unitConstructor.style.display = "none";
}
