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

    // todo функция покупки итема на кнопку

    headEquip.appendChild(placeBuyOrderButton);
}