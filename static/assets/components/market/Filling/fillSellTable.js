function fillSellTable(order, baseName) {

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
    td1.innerHTML = "0"; // todo захардкожаное растояние
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
    td7.innerHTML = "0";
    tr.appendChild(td7);

    tr.onclick = function (e) {
        buyDialog(order, e)
    };

    table.appendChild(tr)
}

function buyDialog(order, e) {
    let subMenu = document.createElement("div");
    subMenu.id = "subMenu";

    subMenu.style.top = e.clientY + "px";
    subMenu.style.left = e.clientX + "px";

    let head = document.createElement("h2");
    head.innerHTML = "Покупака " + order.Item.name;
    subMenu.appendChild(head);

    let div = createNumberInput(0, order.Count, order.Count, "штук");
    subMenu.appendChild(div);

    let resultSpan = document.createElement("span");
    resultSpan.id = "dialogResultSpan";
    resultSpan.innerHTML = "за <span style='color: chartreuse'>" + order.Count * order.Price + "</span> кредитов";
    subMenu.appendChild(resultSpan);

    subMenu.appendChild(document.createElement("br"));

    div.inputBlock.oninput = function () {
        resultSpan.innerHTML = "за <span style='color: chartreuse'>" + this.value * order.Price + " </span> кредитов";
    };

    let closeButton = createInput("Отменить", subMenu);
    closeButton.onclick = function () {
        subMenu.remove();
    };

    let sellButton = createInput("Купить", subMenu);
    sellButton.onclick = function () {
        if (div.inputBlock.value > 0) {
            marketSocket.send(JSON.stringify({
                event: 'buy',
                order_id: Number(order.Id),
                quantity: Number(div.inputBlock.value)
            }));
            subMenu.remove();
        } else {
            alert("нельзя купить 0 предметов")
        }
    };

    document.body.appendChild(subMenu);
}