
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

}
