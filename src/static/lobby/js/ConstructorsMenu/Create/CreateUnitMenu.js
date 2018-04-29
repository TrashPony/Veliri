
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