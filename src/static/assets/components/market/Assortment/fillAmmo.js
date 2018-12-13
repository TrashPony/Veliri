let ammoTypes;

function fillAmmo(types) {
    ammoTypes = types;
    let filterBlock = document.getElementById("ammoCategoryItem");
    filterBlock.onclick = function () {
        openAmmoOneScroll("Боеприпасы", this)
    }
}

function openAmmoOneScroll(type, parent) {
    parent.innerText = " ▼ " + type;

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    parent.appendChild(scroll);

    let names = ['Лазерные', 'Ракетные', 'Балистические'];

    for (let i = 0; i < names.length; i++) {
        let sub = document.createElement("span");
        sub.innerText = " ▶ " + names[i];

        sub.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openAmmoTwoScroll(type, names[i], this)
        };

        scroll.appendChild(sub);
    }

    parent.onclick = function () {
        parent.innerHTML = " ▶ " + type;
        parent.onclick = function (){
            openAmmoOneScroll(type, parent);
        }
    }
}

function openAmmoTwoScroll(type, name, parent) {
    parent.innerText = " ▼ " + name;

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    parent.appendChild(scroll);

    let size = ['Малые', 'Средние', 'Большие'];

    for (let i = 0; i < size.length; i++) {
        let sub = document.createElement("span");
        sub.innerText = " ▶ " + size[i];

        sub.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepAmmo(type, name, size[i], this)
        };

        scroll.appendChild(sub);
    }

    parent.onclick = function () {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        parent.innerHTML = " ▶ " + name;
        parent.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openAmmoTwoScroll(type, name, parent);
        }
    }
}

function openDeepAmmo(type, nameType, sizeType, parent) {
    parent.innerText = " ▼ " + sizeType;
    parent.onclick = function () {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        parent.innerHTML = " ▶ " + sizeType;
        parent.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepAmmo(type, nameType, sizeType, parent)
        }
    };

    let filter1;
    //['Малые', 'Средние', 'Большие'];
    if (sizeType === 'Малые') {
        filter1 = 1
    } else if (sizeType === 'Средние') {
        filter1 = 2
    } else if (sizeType === 'Большие') {
        filter1 = 3
    }

    let filter2;
    //['Лазерные', 'Ракетные', 'Балистические'];
    if (nameType === 'Балистические') {
        filter2 = "firearms"
    } else if (nameType === 'Лазерные') {
        filter2 = "laser"
    } else if (nameType === 'Ракетные') {
        filter2 = "missile"
    }

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    parent.appendChild(scroll);

    let types;
    let url;
    if (type === "Оружие"){
        types = weaponTypes;
        url = "weapon"
    } else {
        types = ammoTypes;
        url = "ammo"
    }

    for (let i in types) {
        if (types.hasOwnProperty(i) && types[i].standard_size === filter1 && types[i].type === filter2) {

            let ammo = document.createElement("span");
            ammo.innerText = types[i].name;
    
            ammo.onclick = function () {
                event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
                selectItem(types[i].id, "ammo", types[i].name, "url(/assets/units/" + url + "/");
            };

            scroll.appendChild(ammo);
        }
    }
}