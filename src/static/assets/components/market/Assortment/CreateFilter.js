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

    if (document.getElementById("dialogBlock")) {
        document.getElementById("dialogBlock").remove();
    }

    let dialogBlock = document.createElement("div");
    dialogBlock.id = "dialogBlock";

    dialogBlock.style.top = e.clientY + "px";
    dialogBlock.style.left = e.clientX + "px";

    let head = document.createElement("h2");
    head.innerHTML = "Покупака " + name;
    dialogBlock.appendChild(head);

    let divCount = createNumberInput(0, 99999999, 0, "штук");
    dialogBlock.appendChild(divCount);

    let divPrice = createNumberInput(0, 99999999, 0, "кредитов");
    dialogBlock.appendChild(divPrice);

    let divMinCount = createNumberInput(0, 99999999, 0, "мин. выкуп");
    dialogBlock.appendChild(divMinCount);

    let divTime = createNumberInput(0, 99999999, 14, "дней");
    dialogBlock.appendChild(divTime);

    let divCalculate = document.createElement("div");
    divCalculate.style.textAlign = "center";
    dialogBlock.appendChild(divCalculate);

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
        if (divCount.inputBlock.value > 0 && divPrice.inputBlock.value > 0) {
            marketSocket.send(JSON.stringify({
                event: 'placeNewBuyOrder',
                item_id: Number(id),
                item_type: type,
                price: Number(divPrice.inputBlock.value),
                quantity: Number(divCount.inputBlock.value),
                expires: Number(divTime.inputBlock.value),
                min_buy_out: Number(divMinCount.inputBlock.value),
            }));
            dialogBlock.remove();
        } else {
            alert("ошибка ввода")
        }
    };

    dialogBlock.appendChild(button);
    document.body.appendChild(dialogBlock);
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