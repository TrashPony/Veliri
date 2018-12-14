function fillSellTable(order) {

    let table = document.getElementById("marketSellTable");
    let tr = document.createElement("tr");
    tr.className = "marketRow";
    tr.order = order;

    if (!(order.IdItem === filterKey.id && order.TypeItem === filterKey.type)) {
        if (filterKey.type !== '') {
            tr.style.display = "none";
        }
    }

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

    let div = createNumberInput(0, order.Count, order.Count, "штук");
    dialogBlock.appendChild(div);

    let resultSpan = document.createElement("span");
    resultSpan.id = "dialogResultSpan";
    resultSpan.innerHTML = "за <span style='color: chartreuse'>" + order.Count * order.Price + "</span> кредитов";
    dialogBlock.appendChild(resultSpan);

    dialogBlock.appendChild(document.createElement("br"));

    div.inputBlock.oninput = function () {
        resultSpan.innerHTML = "за <span style='color: chartreuse'>" + this.value * order.Price + " </span> кредитов";
    };

    let closeButton = createInput("Отменить", dialogBlock);
    closeButton.onclick = function () {
        dialogBlock.remove();
    };

    let sellButton = createInput("Купить", dialogBlock);
    sellButton.onclick = function () {
        if (div.inputBlock.value > 0) {
            marketSocket.send(JSON.stringify({
                event: 'buy',
                order_id: Number(order.Id),
                quantity: Number(div.inputBlock.value)
            }));
            dialogBlock.remove();
        } else {
            alert("нельзя купить 0 предметов")
        }
    };

    document.body.appendChild(dialogBlock);
}