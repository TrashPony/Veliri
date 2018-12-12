let equipTypes;

function fillEquip(types) {
    equipTypes = types;

    let filterBlock = document.getElementById("equipCategoryItem");
    filterBlock.onclick = openEquipScroll;
}

function openEquipScroll() {
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
        this.onclick = openEquipScroll;
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
                selectItem(equipTypes[i].id, "equip", equipTypes[i].name, "url(/assets/units/equip/");
            };

            scroll.appendChild(equip);
        }
    }
}