
function fillBuyTable(order) {
    let table = document.getElementById("marketBuyTable");
    let tr = document.createElement("tr");
    tr.className = "marketRow";
    tr.order = order;

    if (!(order.IdItem === filterKey.id && order.TypeItem === filterKey.type)) {
        if (filterKey.type !== '') {
            tr.style.display = "none";
        }
    }

    let td1 = document.createElement("td");
    td1.innerHTML = "База"; // todo захардкожаная база
    tr.appendChild(td1);

    let td2 = document.createElement("td");
    td2.innerHTML = order.Count;
    tr.appendChild(td2);

    let td3 = document.createElement("td");
    td3.innerHTML = order.Price + " cr.";
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
    td7.innerHTML = order.MinBuyOut;
    tr.appendChild(td7);

    let td8 = document.createElement("td");
    td8.innerHTML = "0";
    tr.appendChild(td8);

    tr.onclick = function (e) {
        sellDialog(order, e)
    };

    table.appendChild(tr)
}

function sellDialog(order, e) {
    let dialogBlock = document.createElement("div");
    dialogBlock.id = "dialogBlock";

    dialogBlock.style.top = e.clientY + "px";
    dialogBlock.style.left = e.clientX + "px";

    let head = document.createElement("h2");
    head.innerHTML = "Продажа    " + order.Item.name;
    dialogBlock.appendChild(head);

    // todo имеется на складе

    let div = createNumberInput(0, order.Count, order.Count, "штук");
    dialogBlock.appendChild(div);

    let resultSpan = document.createElement("div");
    resultSpan.innerHTML = "за <span style='color: chartreuse'>" + order.Count * order.Price + "</span> кредитов";
    dialogBlock.appendChild(resultSpan);

    div.inputBlock.oninput = function () {
        resultSpan.innerHTML = "за <span style='color: chartreuse'>" + this.value * order.Price + " </span> кредитов";
    };

    let closeButton = createInput("Отменить", dialogBlock);
    closeButton.onclick = function () {
        dialogBlock.remove();
    };

    let sellButton = createInput("Продать", dialogBlock);
    sellButton.onclick = function () {
        if (div.inputBlock.value > 0) {
            marketSocket.send(JSON.stringify({
                event: 'sell',
                order_id: Number(order.Id),
                quantity: Number(div.inputBlock.value)
            }));
            dialogBlock.remove();
        } else {
            alert("нельзя продать 0 предметов")
        }
    };

    document.body.appendChild(dialogBlock);
}
