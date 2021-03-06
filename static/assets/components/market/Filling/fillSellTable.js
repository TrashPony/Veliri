function fillSellTable(order) {

    let table = document.getElementById("marketSellTable");
    let tr = document.createElement("tr");
    tr.id = order.Type + order.Id;
    tr.className = "marketRow";
    tr.order = order;

    let td1 = document.createElement("td");
    td1.innerHTML = order.path_jump;
    if (order.path_jump <= 0) {
        td1.style.color = "transparent";
        td1.style.textShadow = "none";
        if (order.path_jump === -1) {
            td1.innerHTML += "<span class='basePath'>База</span>"
        } else if (order.path_jump === 0) {
            td1.innerHTML += "<span class='basePath'>Сектор</span>"
        }
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
    td7.innerHTML = "0";
    tr.appendChild(td7);

    tr.onclick = function (e) {
        buyDialog(order, e)
    };

    table.appendChild(tr)
}

function buyDialog(order, e) {

    if (document.getElementById("subMenu")) {
        document.getElementById("subMenu").remove();
    }

    let subMenu = document.createElement("div");
    subMenu.id = "subMenu";
    subMenu.style.top = e.clientY + "px";
    subMenu.style.left = e.clientX + "px";
    subMenu.style.width = "335px";
    subMenu.style.minWidth = "unset";

    subMenu.innerHTML = `
        <h2>Покупака ${order.Item.name}</h2>
        <div class="marketDialogItemIcon">
            ${getBackgroundUrlByItem({type: order.TypeItem, item: {name: order.Item.name, icon: "blueprint"}})}
        </div>
        <form oninput="result.value = count.value * ${order.Price}" style="float: right;">
        
            <div>
                <span style="float: left"> Цена за шт.:</span>  
                <span class="holdInput cr" style="float: right">${order.Price}</span>
            </div>
            
            <div>
                <span style="float: left"> Купить: </span> 
                <input id="buyCount" style="float: right" name="count" type="number" min="1" max="${order.Count}" value="${order.Count}">
            </div>
            
            <div>
                <span style="float: left"> Всего кредитов:  </span> 
                <output style="float: right" name="result" style='color: chartreuse'>${order.Count * order.Price}</output>
            </div>
        </form>
    `;

    let closeButton = createInput("Отменить", subMenu);
    closeButton.style.margin = "5px 7% 0";
    closeButton.onclick = function () {
        subMenu.remove();
    };

    let sellButton = createInput("Купить", subMenu);
    sellButton.style.margin = "5px 7% 0";
    sellButton.onclick = function () {
        marketSocket.send(JSON.stringify({
            event: 'buy',
            order_id: Number(order.Id),
            quantity: Number(document.getElementById('buyCount').value)
        }));
        subMenu.remove();
    };

    document.body.appendChild(subMenu);
}