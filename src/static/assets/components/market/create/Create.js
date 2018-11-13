function CreateMarketMenu() {
    if (document.getElementById("mask")) {
        document.getElementById("mask").remove();
    }

    if (document.getElementById("marketBox")) {
        document.getElementById("marketBox").remove();
    }

    let mask = document.createElement("div");
    mask.id = "mask";
    mask.style.display = "block";
    document.body.appendChild(mask);

    let marketBox = document.createElement("div");
    marketBox.id = "marketBox";
    document.body.appendChild(marketBox);

    let leftBar = document.createElement("div");
    leftBar.id = "leftBar";
    marketBox.appendChild(leftBar);

    let headMarket = document.createElement("div");
    headMarket.id = "headMarket";
    leftBar.appendChild(headMarket);

    let listItem = document.createElement("div");
    listItem.id = "listItem";
    leftBar.appendChild(listItem);

    let foot = document.createElement("div");
    foot.id = "footMarket";
    leftBar.appendChild(foot);

    let ordersBlock = document.createElement("div");
    ordersBlock.id = "ordersBlock";
    marketBox.appendChild(ordersBlock);

    headUI(headMarket);
    createListItemUI(listItem);
    ordersBlockUI(ordersBlock);
    footUI(foot);

}

function headUI(headMarket) {
    let headMarketHeading = document.createElement("div");
    headMarketHeading.innerHTML = "Рынок Базы 1";
    headMarketHeading.className = "headMarketHeading";
    headMarket.appendChild(headMarketHeading);

    let searchInput = document.createElement("input");
    searchInput.innerHTML = "Поиск";
    searchInput.className = "searchInput";
    searchInput.type = "text";
    searchInput.placeholder = "поиск";
    headMarket.appendChild(searchInput);
}

function createListItemUI(listItem) {
    let ammo = document.createElement("div");
    ammo.className = "categoryItem";
    ammo.innerHTML = " ▶ Боеприпасы";
    listItem.appendChild(ammo);

    let weapon = document.createElement("div");
    weapon.className = "categoryItem";
    weapon.innerHTML = " ▶ Оружие";
    listItem.appendChild(weapon);

    let cabs = document.createElement("div");
    cabs.className = "categoryItem";
    cabs.innerHTML = " ▶ Корпуса";
    listItem.appendChild(cabs);

    let equip = document.createElement("div");
    equip.className = "categoryItem";
    equip.innerHTML = " ▶ Оборудование";
    listItem.appendChild(equip);

    let res = document.createElement("div");
    res.className = "categoryItem";
    res.innerHTML = " ▶ Ресурсы";
    listItem.appendChild(res);
}

function ordersBlockUI(ordersBlock) {
    let selectItemIcon = document.createElement("div");
    selectItemIcon.id = "selectItemIcon";
    ordersBlock.appendChild(selectItemIcon);

    let sellOrdersBlock = document.createElement("div");
    sellOrdersBlock.id = "sellOrdersBlock";
    ordersBlock.appendChild(sellOrdersBlock);
    CreateSellTable(sellOrdersBlock);

    let BuyOrdersBlock = document.createElement("div");
    BuyOrdersBlock.id = "BuyOrdersBlock";
    ordersBlock.appendChild(BuyOrdersBlock);
    CreateBuyTable(BuyOrdersBlock);
}

function footUI(foot) {

    let panel = document.createElement("div");
    panel.id = "footPanel";
    foot.appendChild(panel);

    let placeBuyOrderButton = document.createElement("div");
    placeBuyOrderButton.className = "marketButton";
    placeBuyOrderButton.innerHTML = "Купить";
    panel.appendChild(placeBuyOrderButton);

    let SellOrderButton = document.createElement("div");
    SellOrderButton.className = "marketButton";
    SellOrderButton.innerHTML = "Продать";
    panel.appendChild(SellOrderButton);
}