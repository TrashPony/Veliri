function CreateBuyTable(BuyOrdersBlock) {
    let buyTable = document.createElement("table");
    buyTable.className = "ordersTable";
    buyTable.id = "marketBuyTable";

    let headRow = document.createElement("tr");

    headRow.innerHTML = `
        <th onclick="sortingTableByColumn(0, 'buy', 'number', this)">Растояние <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(1, 'buy', 'number', this)">Количество <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(2, 'buy', 'number', this)">Цена <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(3, 'buy', 'text', this)">Тип <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(4, 'buy', 'text', this)">Название <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(5, 'buy', 'text', this)">Место <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(6, 'buy', 'number', this)">Мин. выкуп <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(7, 'buy', 'date', this)">Истекает через <span class="sortArrow">&#x21D5;</span></th>
    `;

    buyTable.appendChild(headRow);
    BuyOrdersBlock.appendChild(buyTable);
}

function CreateSellTable(SellOrdersBlock) {
    let sellTable = document.createElement("table");
    sellTable.className = "ordersTable";
    sellTable.id = "marketSellTable";

    let headRow = document.createElement("tr");

    headRow.innerHTML = `
        <th onclick="sortingTableByColumn(0, 'sell', 'number', this)">Растояние <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(1, 'sell', 'number', this)">Количество <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(2, 'sell', 'number', this)">Цена <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(3, 'sell', 'text', this)">Тип <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(4, 'sell', 'text', this)">Название <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(5, 'sell', 'text', this)">Место <span class="sortArrow">&#x21D5;</span></th>
        <th onclick="sortingTableByColumn(6, 'sell', 'date', this)">Истекает через <span class="sortArrow">&#x21D5;</span></th>
    `;

    sellTable.appendChild(headRow);
    SellOrdersBlock.appendChild(sellTable);
}

function sortingTableByColumn(tdNumber, table, typeData, td) {

    const setParams = function (filter) {

        if (filter.columnNumber === tdNumber) {
            filter.sorting === 'ASC' ? filter.sorting = 'DESC' : filter.sorting = 'ASC'
        } else {
            filter = {columnNumber: tdNumber, sorting: "ASC", type: typeData}
        }

        if (filter.sorting === 'ASC') {
            $(td).find('span').html('&#x25B2;');
        } else if (filter.sorting === 'DESC') {
            $(td).find('span').html('&#x25BC;');
        }

        return filter;
    };

    $('.ordersTable th').each(function (i, th) {
        $(th).find('span').html('&#x21D5;');
    });

    if (table === "sell") {
        sellSortingRules = setParams(sellSortingRules)
    } else if (table === "buy") {
        buySortingRules = setParams(buySortingRules)
    } else {
        userSortingTable = setParams(userSortingTable)
    }

    filterOrders();
}