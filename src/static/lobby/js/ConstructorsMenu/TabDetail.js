function CreateTabDetailMenu() {
    var tabDetailMenu = document.createElement("div");
    tabDetailMenu.id = "tabDetailMenu";

    var bodyMenu = CreateBodyMenu();
    var towerMenu = CreateTowerMenu();
    var chassisMenu = CreateChassisMenu();
    var weaponMenu = CreateWeaponMenu();
    var radarMenu = CreateRadarMenu();

    var tabs = CreateTabs();
    tabDetailMenu.appendChild(tabs);

    bodyMenu.style.display = "block";
    towerMenu.style.display = "none";
    chassisMenu.style.display = "none";
    weaponMenu.style.display = "none";
    radarMenu.style.display = "none";

    tabDetailMenu.appendChild(bodyMenu);
    tabDetailMenu.appendChild(towerMenu);
    tabDetailMenu.appendChild(chassisMenu);
    tabDetailMenu.appendChild(weaponMenu);
    tabDetailMenu.appendChild(radarMenu);

    return tabDetailMenu;
}

function CreateTabs() {
    var tabs = document.createElement("ul");
    tabs.id = "Tabs";

    var bodyTab = document.createElement("li");
    bodyTab.id = "bodyTab";
    bodyTab.className = "SelectedTab";
    bodyTab.innerHTML = "Корпуса";
    bodyTab.onclick = function () {
        SelectTab(this);
    };

    var towerTab = document.createElement("li");
    towerTab.id = "towerTab";
    towerTab.className = "Tab";
    towerTab.innerHTML = "Башни";
    towerTab.onclick = function () {
        SelectTab(this);
    };

    var chassisTab = document.createElement("li");
    chassisTab.id = "chassisTab";
    chassisTab.className = "Tab";
    chassisTab.innerHTML = "Шасси";
    chassisTab.onclick = function () {
        SelectTab(this);
    };

    var weaponTab = document.createElement("li");
    weaponTab.id = "weaponTab";
    weaponTab.className = "Tab";
    weaponTab.innerHTML = "Оружие";
    weaponTab.onclick = function () {
        SelectTab(this);
    };

    var radarTab = document.createElement("li");
    radarTab.id = "radarTab";
    radarTab.className = "Tab";
    radarTab.innerHTML = "Радары";
    radarTab.onclick = function () {
        SelectTab(this);
    };

    tabs.appendChild(bodyTab);
    tabs.appendChild(towerTab);
    tabs.appendChild(chassisTab);
    tabs.appendChild(weaponTab);
    tabs.appendChild(radarTab);

    return tabs;
}

function SelectTab(tab) {

    var bodyTab = document.getElementById("bodyTab");
    bodyTab.className = "Tab";
    var towerTab = document.getElementById("towerTab");
    towerTab.className = "Tab";
    var chassisTab = document.getElementById("chassisTab");
    chassisTab.className = "Tab";
    var weaponTab = document.getElementById("weaponTab");
    weaponTab.className = "Tab";
    var radarTab = document.getElementById("radarTab");
    radarTab.className = "Tab";

    var bodyMenu = document.getElementById("bodyMenu");
    bodyMenu.style.display = "none";
    var towerMenu = document.getElementById("towerMenu");
    towerMenu.style.display = "none";
    var chassisMenu = document.getElementById("chassisMenu");
    chassisMenu.style.display = "none";
    var weaponMenu = document.getElementById("weaponMenu");
    weaponMenu.style.display = "none";
    var radarMenu = document.getElementById("radarMenu");
    radarMenu.style.display = "none";

    if(tab.id === "bodyTab") {
        bodyMenu.style.display = "block";
    }

    if(tab.id === "chassisTab") {
        chassisMenu.style.display = "block";
    }

    if(tab.id === "weaponTab") {
        weaponMenu.style.display = "block";
    }

    if(tab.id === "radarTab") {
        radarMenu.style.display = "block";
    }

    if(tab.id === "towerTab") {
        towerMenu.style.display = "block";
    }

    tab.className = "SelectedTab";
}

function CreateBodyMenu() {

    var weaponMenu = document.createElement("div");
    weaponMenu.id = "bodyMenu";
    weaponMenu.className = "TabBlocks";

    return weaponMenu;
}

function CreateTowerMenu() {

    var towerMenu = document.createElement("div");
    towerMenu.id = "towerMenu";
    towerMenu.className = "TabBlocks";

    return towerMenu;
}

function CreateChassisMenu() {

    var chassisMenu = document.createElement("div");
    chassisMenu.id = "chassisMenu";
    chassisMenu.className = "TabBlocks";

    return chassisMenu;
}

function CreateWeaponMenu() {

    var weaponMenu = document.createElement("div");
    weaponMenu.id = "weaponMenu";
    weaponMenu.className = "TabBlocks";

    return weaponMenu;
}

function CreateRadarMenu() {

    var chassisMenu = document.createElement("div");
    chassisMenu.id = "radarMenu";
    chassisMenu.className = "TabBlocks";

    return chassisMenu;
}