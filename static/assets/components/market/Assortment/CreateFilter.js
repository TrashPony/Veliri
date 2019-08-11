let filterKey = {type: '', id: 0};
let searchFilter = '';
let radiusFilter = 2;

let buySortingRules = {columnNumber: "", sorting: "", type: ""}; // ASC - по возрастанию, DESC - по убыванию
let sellSortingRules = {columnNumber: "", sorting: "", type: ""}; // ASC - по возрастанию, DESC - по убыванию
let userSortingTable = {columnNumber: "", sorting: "", type: ""}; // ASC - по возрастанию, DESC - по убыванию

function clearFilter() {
    document.getElementById("selectItemName").innerHTML = '';
    document.getElementById("selectItemIcon").innerHTML = '';
    filterKey = {type: '', id: 0};
    filterOrders();
}

function filterOrders() {
    let marketRows = document.getElementsByClassName("marketRow");

    sortingTable('#marketSellTable', sellSortingRules);
    sortingTable('#marketBuyTable', buySortingRules);
    sortingTable('#marketMyTable', userSortingTable);

    for (let j = 0; j < marketRows.length; j++) {

        let row = marketRows[j];

        // перваначально фильтруется дальность рынка
        if (radiusFilter === 1 && row.order.path_jump > 0) {
            row.style.display = "none";
            continue
        }

        if (radiusFilter === 0 && row.order.path_jump !== -1) {
            row.style.display = "none";
            continue
        }

        // если айтем выбран то фокусим его и игнорируем именной поиск
        // если айтем не выбран то показываем весь рынок
        if (filterKey.id === 0 && filterKey.type === '') {
            // если именой фильтр не пуст то фильтруем по нему
            if (row.order.Item.name.indexOf(searchFilter) + 1 || searchFilter === '') {
                row.style.display = "table-row";
            } else {
                row.style.display = "none";
            }
        } else {
            if (row.order.IdItem === filterKey.id && row.order.TypeItem === filterKey.type) {
                row.style.display = "table-row";
            } else {
                row.style.display = "none";
            }
        }
    }
}

function sortingTable(tableID, filter) {

    let rows = $(tableID + " .marketRow");

    rows.sort(function (a, b) {

        a = $(a).find('td:eq(' + filter.columnNumber + ')').contents().get(0).nodeValue;
        b = $(b).find('td:eq(' + filter.columnNumber + ')').contents().get(0).nodeValue;

        switch (filter.type) {
            case 'text':
                return filter.sorting === 'ASC' ? a.localeCompare(b) : b.localeCompare(a);
            case 'number':
                return filter.sorting === 'ASC' ? a - b : b - a;
            case 'date':
                let dateFormat = function (dt) {
                    [m, d, y] = dt.split('/');
                    return [y, m - 1, d];
                };

                //convert the date string to an object using `new Date`
                a = new Date(...dateFormat(a));
                b = new Date(...dateFormat(b));

                //You can use getTime() to convert the date object into numbers.
                //getTime() method returns the number of milliseconds between midnight of January 1, 1970
                //So since a and b are numbers now, you can use the same process if the type is number. Just deduct the values.
                return filter.sorting === 'ASC' ? a.getTime() - b.getTime() : b.getTime() - a.getTime();
        }
    });

    rows.appendTo(tableID);
}

function selectItem(id, type, name) {
    filterKey.type = type;
    filterKey.id = id;

    filterOrders();

    document.getElementById("selectItemIcon").innerHTML = `
        ${getBackgroundUrlByItem({
        type: type,
        item: {name: name, icon: "blueprint"},
    })}`;

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
    divCalculate.innerHTML = `<sapn>0 шт.</sapn> за <span class="holdInput cr" style="padding-right: 20px;"> 0 </span>`;
    subMenu.appendChild(divCalculate);

    divCount.inputBlock.oninput = function () {
        divCalculate.innerHTML = `<sapn>${divCount.inputBlock.value}</sapn> за <span class="holdInput cr" style="padding-right: 20px;">${divCount.inputBlock.value * divPrice.inputBlock.value}</span>`;
    };

    divPrice.inputBlock.oninput = function () {
        divCalculate.innerHTML = `<sapn>${divCount.inputBlock.value}</sapn> за <span class="holdInput cr" style="padding-right: 20px;">${divCount.inputBlock.value * divPrice.inputBlock.value}</span>`;
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

    let buttonWrapper = document.createElement("div");
    buttonWrapper.id = "BuyDialogButtonWrapper";
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
    buttonWrapper.appendChild(button);

    let close = document.createElement("input");
    close.type = "button";
    close.className = "lobbyButton inventoryTip";
    close.value = "Отмена";
    close.onclick = function () {
        document.getElementById("subMenu").remove();
    };
    buttonWrapper.appendChild(close);

    subMenu.appendChild(buttonWrapper);
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