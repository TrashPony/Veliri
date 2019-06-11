function fillBuyTable(order) {
    let table = document.getElementById("marketBuyTable");
    let tr = document.createElement("tr");
    tr.id = order.Type + order.Id;
    tr.className = "marketRow";
    tr.order = order;

    let td1 = document.createElement("td");
    td1.innerHTML = order.path_jump;
    if (order.path_jump < 0) {
        td1.innerHTML = "<span>База</span>"
    }
    tr.appendChild(td1);

    let td2 = document.createElement("td");
    td2.innerHTML = order.Count;
    tr.appendChild(td2);

    let td3 = document.createElement("td");
    td3.className = "creditsTD";
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

    if (document.getElementById("subMenu")) {
        document.getElementById("subMenu").remove();
    }

    let subMenu = document.createElement("div");
    subMenu.id = "subMenu";
    subMenu.style.top = e.clientY + "px";
    subMenu.style.left = e.clientX + "px";

    let head = document.createElement("h2");
    head.innerHTML = "Продажа    " + order.Item.name;
    subMenu.appendChild(head);

    // todo имеется на складе

    let div = createNumberInput(0, order.Count, order.Count, "штук");
    subMenu.appendChild(div);

    let resultSpan = document.createElement("div");
    resultSpan.innerHTML = "за <span style='color: chartreuse'>" + order.Count * order.Price + "</span> кредитов";
    subMenu.appendChild(resultSpan);

    div.inputBlock.oninput = function () {
        resultSpan.innerHTML = "за <span style='color: chartreuse'>" + this.value * order.Price + " </span> кредитов";
    };

    let closeButton = createInput("Отменить", subMenu);
    closeButton.onclick = function () {
        subMenu.remove();
    };

    let sellButton = createInput("Продать", subMenu);
    sellButton.onclick = function () {
        if (div.inputBlock.value > 0) {
            marketSocket.send(JSON.stringify({
                event: 'sell',
                order_id: Number(order.Id),
                quantity: Number(div.inputBlock.value)
            }));
            subMenu.remove();
        } else {
            alert("нельзя продать 0 предметов")
        }
    };

    document.body.appendChild(subMenu);
}
