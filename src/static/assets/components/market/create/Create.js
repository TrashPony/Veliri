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

    let balance = document.createElement("div");
    balance.id = "balance";
    headMarket.appendChild(balance);

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
    ammo.id = "ammoCategoryItem";
    ammo.innerHTML = " ▶ Боеприпасы";
    listItem.appendChild(ammo);

    let weapon = document.createElement("div");
    weapon.className = "categoryItem";
    weapon.id = "weaponCategoryItem";
    weapon.innerHTML = " ▶ Оружие";
    listItem.appendChild(weapon);

    let cabs = document.createElement("div");
    cabs.className = "categoryItem";
    cabs.id = "cabsCategoryItem";
    cabs.innerHTML = " ▶ Корпуса";
    listItem.appendChild(cabs);

    let equip = document.createElement("div");
    equip.className = "categoryItem";
    equip.id = "equipCategoryItem";
    equip.innerHTML = " ▶ Оборудование";
    listItem.appendChild(equip);

    let res = document.createElement("div");
    res.className = "categoryItem";
    res.id = "resCategoryItem";
    res.innerHTML = " ▶ Ресурсы";
    listItem.appendChild(res);
}

function ordersBlockUI(ordersBlock) {
    let menu = document.createElement("div");
    menu.className = "marketTopMenu";

    let allMarket = document.createElement("div");
    allMarket.innerHTML = "Рынок";
    allMarket.className = "activePin";

    let myMarket = document.createElement("div");
    myMarket.innerHTML = "Мои запросы/предложения";
    myMarket.onclick = function(){
        MyOrdersTab(myMarket, allMarket)
    };

    menu.appendChild(allMarket);
    menu.appendChild(myMarket);
    ordersBlock.appendChild(menu);

    let selectItemIcon = document.createElement("div");
    selectItemIcon.id = "selectItemIcon";
    ordersBlock.appendChild(selectItemIcon);

    let selectItemName = document.createElement("div");
    selectItemName.id = "selectItemName";
    ordersBlock.appendChild(selectItemName);

    let ordersBuyBlockHead = document.createElement("div");
    ordersBuyBlockHead.className = "ordersHead";
    ordersBuyBlockHead.innerHTML = "Предложения";
    ordersBuyBlockHead.style.marginTop = "100px";
    ordersBlock.appendChild(ordersBuyBlockHead);

    let sellOrdersBlock = document.createElement("div");
    sellOrdersBlock.id = "sellOrdersBlock";
    ordersBlock.appendChild(sellOrdersBlock);
    CreateSellTable(sellOrdersBlock);

    let ordersSellBlockHead = document.createElement("div");
    ordersSellBlockHead.className = "ordersHead";
    ordersSellBlockHead.innerHTML = "Запросы";
    ordersBlock.appendChild(ordersSellBlockHead);

    let BuyOrdersBlock = document.createElement("div");
    BuyOrdersBlock.id = "BuyOrdersBlock";
    ordersBlock.appendChild(BuyOrdersBlock);
    CreateBuyTable(BuyOrdersBlock);
}

function footUI(foot) {
    let panel = document.createElement("div");
    panel.id = "footPanel";
    foot.appendChild(panel);

    let Close = document.createElement("div");
    Close.className = "marketButton";
    Close.innerHTML = "Закрыть";
    Close.onclick = function () {
        if (document.getElementById("mask")) {
            document.getElementById("mask").remove();
        }

        if (document.getElementById("marketBox")) {
            document.getElementById("marketBox").remove();
        }
        marketSocket.close();
    };
    panel.appendChild(Close);
}