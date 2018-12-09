function Filling(data) {

    deleteOldRows();

    for (let i in data.orders) {
        if (data.orders.hasOwnProperty(i)) {
            let order = data.orders[i];
            console.log(order);

            if (order.Type === "sell") {
                fillSellTable(order);
            } else {

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

    let head = document.createElement("h2");
    head.innerHTML = "Покупака " + order.Item.name;
    dialogBlock.appendChild(head);

    let input = document.createElement("input");
    input.type = "number";
    input.min = 0;
    input.max = order.Count;
    input.value = order.Count;
    dialogBlock.appendChild(input);

    let span = document.createElement("span");
    span.innerHTML = "штук";
    dialogBlock.appendChild(span);

    dialogBlock.appendChild(document.createElement("br"));

    let resultSpan = document.createElement("span");
    resultSpan.id = "dialogResultSpan";
    resultSpan.innerHTML = "за <span style='color: chartreuse'>" + order.Count * order.Price + "</span> кредитов";
    dialogBlock.appendChild(resultSpan);

    dialogBlock.appendChild(document.createElement("br"));

    input.oninput = function () {
        resultSpan.innerHTML = "за <span style='color: chartreuse'>" + this.value * order.Price + " </span> кредитов";
    };

    let closeButton = createInput("Отменить", dialogBlock);
    closeButton.onclick = function () {
        dialogBlock.remove();
    };

    let sellButton = createInput("Купить", dialogBlock);
    sellButton.onclick = function () {
        if (input.value > 0) {
            marketSocket.send(JSON.stringify({
                event: 'buy',
                order_id: Number(order.Id),
                quantity: Number(input.value)
            }));
            dialogBlock.remove();
        } else {
            alert("нельзя купить 0 предметов")
        }
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