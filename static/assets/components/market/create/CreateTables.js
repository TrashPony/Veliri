function CreateBuyTable(BuyOrdersBlock) {
    let buyTable = document.createElement("table");
    buyTable.className = "ordersTable";
    buyTable.id = "marketBuyTable";

    let headRow = document.createElement("tr");

    let td1 = document.createElement("td");
    td1.innerHTML = "Растояние";
    headRow.appendChild(td1);

    let td2 = document.createElement("td");
    td2.innerHTML = "Количество";
    headRow.appendChild(td2);

    let td3 = document.createElement("td");
    td3.innerHTML = "Цена";
    headRow.appendChild(td3);

    let td4 = document.createElement("td");
    td4.innerHTML = "Тип";
    headRow.appendChild(td4);

    let td5 = document.createElement("td");
    td5.innerHTML = "Название";
    headRow.appendChild(td5);

    let td6 = document.createElement("td");
    td6.innerHTML = "Место";
    headRow.appendChild(td6);

    let td7 = document.createElement("td");
    td7.innerHTML = "Мин. выкуп";
    headRow.appendChild(td7);

    let td8 = document.createElement("td");
    td8.innerHTML = "Истекает через";
    headRow.appendChild(td8);

    buyTable.appendChild(headRow);
    BuyOrdersBlock.appendChild(buyTable);
}

function CreateSellTable(SellOrdersBlock) {
    let sellTable = document.createElement("table");
    sellTable.className = "ordersTable";
    sellTable.id = "marketSellTable";

    let headRow = document.createElement("tr");

    let td1 = document.createElement("td");
    td1.innerHTML = "Растояние";
    headRow.appendChild(td1);

    let td2 = document.createElement("td");
    td2.innerHTML = "Количество";
    headRow.appendChild(td2);

    let td3 = document.createElement("td");
    td3.innerHTML = "Цена";
    headRow.appendChild(td3);

    let td4 = document.createElement("td");
    td4.innerHTML = "Тип";
    headRow.appendChild(td4);

    let td5 = document.createElement("td");
    td5.innerHTML = "Название";
    headRow.appendChild(td5);

    let td6 = document.createElement("td");
    td6.innerHTML = "Место";
    headRow.appendChild(td6);

    let td7 = document.createElement("td");
    td7.innerHTML = "Истекает через";
    headRow.appendChild(td7);

    sellTable.appendChild(headRow);
    SellOrdersBlock.appendChild(sellTable);
}
