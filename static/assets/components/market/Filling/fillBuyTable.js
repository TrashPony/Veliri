function fillBuyTable(order) {
    let table = document.getElementById("marketBuyTable");
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
    subMenu.style.width = "335px";
    subMenu.style.minWidth = "unset";

    subMenu.innerHTML = `
     <h2>Продажа ${order.Item.name}</h2>
        <div class="marketDialogItemIcon">
            ${getBackgroundUrlByItem({type: order.TypeItem, item: {name: order.Item.name, icon: "blueprint"}})}
        </div>
        
        <form oninput="result.value = count.value * ${order.Price}" style="float: right;">
            <div>
                <span style="float: left"> Имеется на складе: </span>
                <span class="holdInput count" id="sellCountInStorage" style="float: right">0</span>
            </div>
            
            <div>
                <span style="float: left"> Цена за шт.:</span>  
                <span class="holdInput cr" style="float: right">${order.Price}</span>
            </div>
            
            <div> 
                <span style="float: left"> Продать: </span>
                <input id="buyCount" style="float: right" name="count" type="number" min="1" max="${order.Count}" value="1">
            </div>
            
            <div>
                <span style="float: left"> Всего кредитов:  </span>
                <output style="float: right" name="result" style='color: chartreuse'>${order.Price}</output>
            </div>
        </form>
    `;

    let closeButton = createInput("Отменить", subMenu);
    closeButton.onclick = function () {
        subMenu.remove();
    };

    let sellButton = createInput("Продать", subMenu);
    sellButton.onclick = function () {
        marketSocket.send(JSON.stringify({
            event: 'sell',
            order_id: Number(order.Id),
            quantity: Number(div.inputBlock.value)
        }));
        subMenu.remove();
    };

    document.body.appendChild(subMenu);

    // запрос сколько имеется на складе айтемов соотвествующие этому ордеру, на базе где размещен ордер
    marketSocket.send(JSON.stringify({
        event: 'getItemsInStorage',
        item_id: order.Item.id,
        item_type: order.TypeItem,
        order_id: order.Id,
    }));
}
