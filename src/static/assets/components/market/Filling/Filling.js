function Filling(data) {

    deleteOldRows();

    for (let i = 0; i < data.orders.length; i++) {

        let order = data.orders[i];
        console.log(order);

        if (order.Type === "sell") {
            fillSellTable(order);
        } else {

        }
    }
}

function deleteOldRows() {
    let oldRows = document.getElementsByClassName("marketRow");
    while (oldRows.length > 0) {
        oldRows[0].remove();
    }
}

function fillSellTable(order) {
    let table = document.getElementById("marketSellTable");
    let tr = document.createElement("tr");
    tr.className = "marketRow";

    let td1 = document.createElement("td");
    td1.innerHTML = "0";
    tr.appendChild(td1);

    let td2 = document.createElement("td");
    td2.innerHTML = order.Count;
    tr.appendChild(td2);

    let td3 = document.createElement("td");
    td3.innerHTML = order.Price;
    tr.appendChild(td3);

    let td4 = document.createElement("td");
    td4.innerHTML = order.TypeItem;
    tr.appendChild(td4);

    let td5 = document.createElement("td");
    td5.innerHTML = order.Item.name;
    tr.appendChild(td5);

    let td6 = document.createElement("td");
    td6.innerHTML = order.PlaceName;
    tr.appendChild(td6);

    let td7 = document.createElement("td");
    td7.innerHTML = "0";
    tr.appendChild(td7);

    tr.onclick = function (e) {
        buyDialog(order, e)
    };

    table.appendChild(tr)
}

function buyDialog(order, e) {
    let dialogBlock = document.createElement("div");
    dialogBlock.id = "dialogBlock";

    dialogBlock.style.top = e.clientY + "px";
    dialogBlock.style.left = e.clientX + "px";

    dialogBlock.innerHTML = "<h2>Покупака " + order.Item.name + "</h2>" +
        "<div><input type='number' min='0' value='" + order.Count + "' max='" + order.Count + "'> <span>штук</span><br>" +
        "<span> за " + order.Count * order.Price + " кредитов </span></div>";

    let closeButton = createInput("Отменить", dialogBlock);
    closeButton.onclick = function () {
        dialogBlock.remove();
    };

    let sellButton = createInput("Продать", dialogBlock);
    sellButton.onclick = function () {

    };

    document.body.appendChild(dialogBlock);
}

function createInput(value, parrent) {
    let button = document.createElement("input");
    button.type = "button";
    button.className = "lobbyButton inventoryTip";
    button.value = value;
    button.style.pointerEvents = "auto";
    parrent.appendChild(button);
    return button
}