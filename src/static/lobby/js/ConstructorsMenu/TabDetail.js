function CreateTabDetailMenu() {
    var tabDetailMenu = document.createElement("div");
    tabDetailMenu.id = "tabDetailMenu";

    var chassisMenu = CreateChassisMenu();
    var weaponMenu = CreateWeaponMenu();
    var radarMenu = CreateRadarMenu();
    var bodyMenu = CreateBodyMenu();

    var tabs = CreateTabs();
    tabDetailMenu.appendChild(tabs);

    chassisMenu.style.display = "block";
    weaponMenu.style.display = "none";
    radarMenu.style.display = "none";
    bodyMenu.style.display = "none";

    tabDetailMenu.appendChild(chassisMenu);
    tabDetailMenu.appendChild(weaponMenu);
    tabDetailMenu.appendChild(radarMenu);
    tabDetailMenu.appendChild(bodyMenu);

    return tabDetailMenu;
}

function CreateTabs() {
    var tabs = document.createElement("ul");
    tabs.id = "Tabs";

    var chassisTab = document.createElement("li");
    chassisTab.id = "chassisTab";
    chassisTab.className = "SelectedTab";
    chassisTab.innerHTML = "Шасси";
    chassisTab.onclick = function () {
        SelectTab(this);
    };

    var weaponTab = document.createElement("li");
    weaponTab.id = "weaponTab";
    weaponTab.className = "Tab";
    weaponTab.innerHTML = "Башни";
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

    var bodyTab = document.createElement("li");
    bodyTab.id = "bodyTab";
    bodyTab.className = "Tab";
    bodyTab.innerHTML = "Корпуса";
    bodyTab.onclick = function () {
        SelectTab(this);
    };

    tabs.appendChild(chassisTab);
    tabs.appendChild(weaponTab);
    tabs.appendChild(radarTab);
    tabs.appendChild(bodyTab);

    return tabs;
}

function SelectTab(tab) {

    var chassisTab = document.getElementById("chassisTab");
    chassisTab.className = "Tab";
    var weaponTab = document.getElementById("weaponTab");
    weaponTab.className = "Tab";
    var radarTab = document.getElementById("radarTab");
    radarTab.className = "Tab";
    var bodyTab = document.getElementById("bodyTab");
    bodyTab.className = "Tab";

    var chassisMenu = document.getElementById("chassisMenu");
    chassisMenu.style.display = "none";
    var weaponMenu = document.getElementById("weaponMenu");
    weaponMenu.style.display = "none";
    var radarMenu = document.getElementById("radarMenu");
    radarMenu.style.display = "none";
    var bodyMenu = document.getElementById("bodyMenu");
    bodyMenu.style.display = "none";

    if(tab.id === "chassisTab") {
        chassisMenu.style.display = "block";
    }

    if(tab.id === "weaponTab") {
        weaponMenu.style.display = "block";
    }

    if(tab.id === "radarTab") {
        radarMenu.style.display = "block";
    }

    if(tab.id === "bodyTab") {
        bodyMenu.style.display = "block";
    }

    tab.className = "SelectedTab";
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

function CreateBodyMenu() {

    var weaponMenu = document.createElement("div");
    weaponMenu.id = "bodyMenu";
    weaponMenu.className = "TabBlocks";

    return weaponMenu;
}