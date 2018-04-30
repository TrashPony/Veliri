
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
        DeleteDetail(this, "picBody")
    };
    unitMenu.appendChild(bodyUnitBox);

    var towerUnitBox = document.createElement("div");
    towerUnitBox.className = "ElementUnitBox right";
    towerUnitBox.id = "towerElement";
    towerUnitBox.onclick = function () {
        DeleteDetail(this, "picTower")
    };
    unitMenu.appendChild(towerUnitBox);

    var chassisUnitBox = document.createElement("div");
    chassisUnitBox.className = "ElementUnitBox left";
    chassisUnitBox.id = "chassisElement";
    chassisUnitBox.onclick = function () {
        DeleteDetail(this, "picChassis")
    };
    unitMenu.appendChild(chassisUnitBox);

    var weaponUnitBox = document.createElement("div");
    weaponUnitBox.className = "ElementUnitBox weapon";
    weaponUnitBox.id = "weaponElement";
    weaponUnitBox.onclick = function () {
        DeleteDetail(this, "picWeapon")
    };
    unitMenu.appendChild(weaponUnitBox);

    var radarUnitBox = document.createElement("div");
    radarUnitBox.className = "ElementUnitBox left";
    radarUnitBox.id = "radarElement";
    radarUnitBox.onclick = function () {
        DeleteDetail(this, "picRadar")
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


function DeleteDetail(box, PicDetail) {
    box.detail = null;
    box.style.backgroundImage = "";
    box.radar = null;
    var pic = document.getElementById(PicDetail);
    if (pic){
        pic.remove();
    }
    SendEventAddOrDelDetail();
}