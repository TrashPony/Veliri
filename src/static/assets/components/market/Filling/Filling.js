function Filling(data) {

    if (data.assortment) {
        FillAssortment(data.assortment)
    }

    deleteOldRows();

    document.getElementById("balance").innerHTML = "Мой баланс: <span>" + data.credits + "</span> cr.";

    for (let i in data.orders) {
        if (data.orders.hasOwnProperty(i)) {
            let order = data.orders[i];

            if (order.Type === "sell") {
                fillSellTable(order);
            } else {
                fillBuyTable(order);
            }
        }
    }
}

function deleteOldRows() {
    let oldRows = document.getElementsByClassName("marketRow");
    while (oldRows.length > 0) {
        oldRows[0].remove();
    }
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