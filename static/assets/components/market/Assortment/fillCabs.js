let bodyTypes;

function fillCabs(types) {
    bodyTypes = types;

    let filterBlock = document.getElementById("cabsCategoryItem");
    filterBlock.onclick = openCabsScroll;
}

function openCabsScroll() {
    this.innerText = " ▼ Корпуса";

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    scroll.onclick = function(){
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };
    this.appendChild(scroll);

    let names = ['MS', 'Тяжелые', 'Средние', 'Легкие'];

    for (let i = 0; i < names.length; i++) {
        let sub = document.createElement("span");
        sub.innerText = " ▶ " + names[i];

        sub.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepCabs(names[i], this)
        };

        scroll.appendChild(sub);
    }

    this.onclick = function () {
        this.innerHTML = " ▶ Корпуса";
        this.onclick = openCabsScroll;
        clearFilter();
    }
}

function openDeepCabs(nameType, parent) {
    parent.innerText = " ▼ " + nameType;

    parent.onclick = function () {
        clearFilter();
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        parent.innerHTML = " ▶ " + nameType;
        parent.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepCabs(nameType, parent)
        }
    };

    let filter;

    // ['MS', 'Тяжелые', 'Средние', 'Легкие'];
    if (nameType === 'MS') {
        filter = 4
    } else if (nameType === 'Тяжелые') {
        filter = 3
    } else if (nameType === 'Средние') {
        filter = 2
    } else if (nameType === 'Легкие') {
        filter = 1
    }

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    scroll.onclick = function(){
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };
    parent.appendChild(scroll);

    for (let i in bodyTypes) {
        if (bodyTypes.hasOwnProperty(i) && ((bodyTypes[i].standard_size === filter) ||
            filter === 4 && bodyTypes[i].mother_ship)) {

            if (filter !== 4 && bodyTypes[i].mother_ship) {
                continue;
            }

            let cab = document.createElement("span");
            cab.innerText = bodyTypes[i].name;

            cab.onclick = function () {
                event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
                selectItem(bodyTypes[i].id, "body", bodyTypes[i].name, "url(/assets/units/body/");
            };

            scroll.appendChild(cab);
        }
    }
}