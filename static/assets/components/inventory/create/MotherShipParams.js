function CreateMotherShipParamsMenu() {
    let menu = document.getElementById("MotherShipParams");

    let blockAttack = document.createElement("div");
    blockAttack.className = "Value params";
    blockAttack.innerHTML = "▶ Атака";
    blockAttack.onclick = openAttack;

    let blockDef = document.createElement("div");
    blockDef.className = "Value params";
    blockDef.innerHTML = "▶ Защита";
    blockDef.onclick = openDef;

    let blockNav = document.createElement("div");
    blockNav.className = "Value params";
    blockNav.innerHTML = "▶ Навигация";
    blockNav.onclick = openNav;

    menu.appendChild(blockAttack);
    createAttackInfo(menu);

    menu.appendChild(blockDef);
    createDefendInfo(menu);

    menu.appendChild(blockNav);
    createNavInfo(menu);

}

function openAttack() {
    this.innerHTML = "▼ Атака";
    document.getElementById("infoParamsAttack").style.display = "block";
    this.onclick = function () {
        document.getElementById("infoParamsAttack").style.display = "none";
        this.innerHTML = "▶ Атака";
        this.onclick = openAttack;
    }
}

function openDef() {
    this.innerHTML = "▼ Защита";
    document.getElementById("infoParamsDefend").style.display = "block";
    this.onclick = function () {
        document.getElementById("infoParamsDefend").style.display = "none";
        this.innerHTML = "▶ Защита";
        this.onclick = openDef;
    }
}

function openNav() {
    this.innerHTML = "▼ Навигация";
    document.getElementById("infoParamsNav").style.display = "block";
    this.onclick = function () {
        document.getElementById("infoParamsNav").style.display = "none";
        this.innerHTML = "▶ Навигация";
        this.onclick = openNav;
    }
}