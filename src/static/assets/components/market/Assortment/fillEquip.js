let equipTypes;

function fillEquip(types) {
    equipTypes = types;

    let filterBlock = document.getElementById("equipCategoryItem");
    filterBlock.onclick = openScroll;
}

function openScroll() {
    this.innerText = " ▼ Оборудование";

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    this.appendChild(scroll);

    for (let i = 0; i < 5; i++) {
        let sub = document.createElement("span");
        sub.innerText = " ▶ " + Number(i + 1);

        sub.onclick = function() {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepEquip(i+1, this)
        };

        scroll.appendChild(sub);
    }

    this.onclick = function () {
        this.innerHTML = " ▶ Оборудование";
        this.onclick = openScroll;
    }
}

function openDeepEquip(numberTypeSlot, parent) {
    parent.innerText = " ▼ " + numberTypeSlot;

    parent.onclick = function () {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        parent.innerHTML = " ▶ " + numberTypeSlot;
        parent.onclick = function (){
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepEquip(numberTypeSlot, parent)
        }
    };

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    parent.appendChild(scroll);

    for (let i in equipTypes) {
        if (equipTypes.hasOwnProperty(i) && Number(equipTypes[i].type_slot) === numberTypeSlot) {
            let equip = document.createElement("span");
            equip.innerText = equipTypes[i].name;
            equip.onclick = function() {
                event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

                filterKey.type = "equip";
                filterKey.id = equipTypes[i].id;

                let marketRows = document.getElementsByClassName("marketRow");
                for (let j = 0; j < marketRows.length; j++) {
                    if (marketRows[j].order.IdItem === filterKey.id && marketRows[j].order.TypeItem === filterKey.type) {
                        marketRows[j].style.display = "table-row";
                    } else {
                        marketRows[j].style.display = "none";
                    }
                }

                document.getElementById("selectItemIcon").style.background = "url(/assets/units/equip/" +
                    equipTypes[i].name + ".png) center / cover";

                let headEquip = document.getElementById("selectItemName");
                headEquip.innerHTML = "<span>" + equipTypes[i].name + "</span><br>";

                let placeBuyOrderButton = document.createElement("div");
                placeBuyOrderButton.className = "marketButton";
                placeBuyOrderButton.innerHTML = "Купить";
                placeBuyOrderButton.style.margin = "20px auto";

                // todo функция покупки итема на кнопку

                headEquip.appendChild(placeBuyOrderButton);
            };

            scroll.appendChild(equip);
        }
    }
}