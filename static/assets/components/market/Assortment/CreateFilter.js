let filterKey = {type: '', id: 0};

function selectItem(id, type, name, url) {
    filterKey.type = type;
    filterKey.id = id;

    let marketRows = document.getElementsByClassName("marketRow");
    for (let j = 0; j < marketRows.length; j++) {
        if (marketRows[j].order.IdItem === filterKey.id && marketRows[j].order.TypeItem === filterKey.type) {
            marketRows[j].style.display = "table-row";
        } else {
            marketRows[j].style.display = "none";
        }
    }

    document.getElementById("selectItemIcon").style.background = url +
        name + ".png) center / cover";

    let headEquip = document.getElementById("selectItemName");
    headEquip.innerHTML = "<span>" + name + "</span><br>";

    let placeBuyOrderButton = document.createElement("div");
    placeBuyOrderButton.className = "marketButton";
    placeBuyOrderButton.innerHTML = "Купить";
    placeBuyOrderButton.style.margin = "20px auto";

    placeBuyOrderButton.onclick = function (e) {
        placeBuyDialog(type, id, name, e);
    };

    headEquip.appendChild(placeBuyOrderButton);
}

function placeBuyDialog(type, id, name, e) {

    if (document.getElementById("subMenu")) {
        document.getElementById("subMenu").remove();
    }

    let subMenu = document.createElement("div");
    subMenu.id = "subMenu";

    subMenu.style.top = e.clientY + "px";
    subMenu.style.left = e.clientX + "px";

    let head = document.createElement("h2");
    head.innerHTML = "Покупака " + name;
    subMenu.appendChild(head);

    let divCount = createNumberInput(0, 99999999, 0, "штук");
    subMenu.appendChild(divCount);

    let divPrice = createNumberInput(0, 99999999, 0, "кредитов");
    subMenu.appendChild(divPrice);

    let divMinCount = createNumberInput(0, 99999999, 0, "мин. выкуп");
    subMenu.appendChild(divMinCount);

    let divTime = createNumberInput(0, 99999999, 14, "дней");
    subMenu.appendChild(divTime);

    let divCalculate = document.createElement("div");
    divCalculate.style.textAlign = "center";
    subMenu.appendChild(divCalculate);

    divCount.inputBlock.oninput = function () {
        divCalculate.innerHTML = divCount.inputBlock.value + " за " + divCount.inputBlock.value * divPrice.inputBlock.value + " кредитов";
    };

    divPrice.inputBlock.oninput = function () {
        divCalculate.innerHTML = divCount.inputBlock.value + " за " + divCount.inputBlock.value * divPrice.inputBlock.value + " кредитов";
    };
    divMinCount.inputBlock.oninput = function () {
        while (divCount.inputBlock.value % this.value !== 0) {
            if (divCount.inputBlock.value === "0" || this.value === "0") {
                return
            } else {
                divCount.inputBlock.step = divMinCount.inputBlock.value;
                divCount.inputBlock.value = Number(divCount.inputBlock.value) - 1;
                divCalculate.innerHTML = divCount.inputBlock.value + " за " + divCount.inputBlock.value * divPrice.inputBlock.value + " кредитов";
            }
        }
    };

    let button = document.createElement("input");
    button.type = "button";
    button.className = "lobbyButton inventoryTip";
    button.value = "Разместить заказ";
    button.onclick = function () {
        if (Number(divCount.inputBlock.value) > 0 && Number(divPrice.inputBlock.value) > 0) {
            marketSocket.send(JSON.stringify({
                event: 'placeNewBuyOrder',
                item_id: Number(id),
                item_type: type,
                price: Number(divPrice.inputBlock.value),
                quantity: Number(divCount.inputBlock.value),
                expires: Number(divTime.inputBlock.value),
                min_buy_out: Number(divMinCount.inputBlock.value),
            }));
            subMenu.remove();
        } else {
            alert("ошибка ввода")
        }
    };
    subMenu.appendChild(button);

    let close = document.createElement("input");
    close.type = "button";
    close.className = "lobbyButton inventoryTip";
    close.value = "Отмена";
    close.onclick = function (){
        document.getElementById("subMenu").remove();
    };
    subMenu.appendChild(close);

    document.body.appendChild(subMenu);
}

function createNumberInput(min, max, value, text) {
    let div = document.createElement("div");

    div.inputBlock = document.createElement("input");
    div.inputBlock.type = "number";
    div.inputBlock.min = min;
    div.inputBlock.max = max;
    div.inputBlock.value = value;
    div.appendChild(div.inputBlock);

    div.spanBlock = document.createElement("span");
    div.spanBlock.innerHTML = text;
    div.appendChild(div.spanBlock);

    return div;
}