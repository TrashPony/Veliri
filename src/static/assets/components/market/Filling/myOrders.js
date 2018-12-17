function FillMyOrders(orders) {
    
    let table = document.getElementById("marketMyTable");
    let ordersBlock = document.getElementById("MyOrdersBlock");

    if (!table) {
        table = createMyTable()
    }
    if (!ordersBlock) {
        return
    }
    ordersBlock.appendChild(table);

    for (let i in orders) {
        if (orders.hasOwnProperty(i)) {
            addMyOrder(orders[i])
        }
    }
}

function addMyOrder(order) {
    let table = document.getElementById("marketMyTable");

    let tr = document.createElement("tr");
    tr.className = "marketRow myOrders";
    tr.order = order;

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

    let td9 = document.createElement("td");
    if (order.Type === "sell") {
        td9.innerHTML = "<span style='color: #ffb300'>" + order.Type + "</span>";
    } else {
        td9.innerHTML = "<span style='color: #25ff00'>" + order.Type + "</span>";
    }
    tr.appendChild(td9);

    let td10 = document.createElement("td");
    td10.innerHTML = order.Price * order.Count + " cr.";
    tr.appendChild(td10);

    let td11 = document.createElement("td");
    tr.appendChild(td11);

    let cancelButton = document.createElement("input");
    cancelButton.type = "button";
    cancelButton.value = "Отменить";
    cancelButton.className = "button cancel order";
    cancelButton.onclick = function(){
        marketSocket.send(JSON.stringify({
            event: "cancelOrder",
            order_id: order.Id
        }));
    };
    td11.appendChild(cancelButton);

    table.appendChild(tr)
}

function createMyTable() {
    let sellTable = document.createElement("table");
    sellTable.className = "ordersTable";
    sellTable.id = "marketMyTable";

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

    let td9 = document.createElement("td");
    td9.innerHTML = "Тип";
    headRow.appendChild(td9);

    let td10 = document.createElement("td");
    td10.innerHTML = "Общая стоимость";
    headRow.appendChild(td10);

    let td11 = document.createElement("td");
    td11.innerHTML = "";
    headRow.appendChild(td11);

    sellTable.appendChild(headRow);

    return sellTable
}