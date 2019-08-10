function InitMarketMenu(noMask) {
    // обнуляем прошлый фильтр поиска
    filterKey = {type: '', id: 0};
    radiusFilter = 2;
    buySortingRules = {columnNumber: "", sorting: "", type: ""};
    sellSortingRules = {columnNumber: "", sorting: "", type: ""};
    userSortingTable = {columnNumber: "", sorting: "", type: ""};
    searchFilter = '';

    ConnectMarket();
    CreateMarketMenu(noMask);
}