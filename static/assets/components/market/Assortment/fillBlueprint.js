let blueprints;

function fillBlueprint(putBlueprints) {

    blueprints = putBlueprints;

    let filterBlock = document.getElementById("bpCategoryItem");
    filterBlock.onclick = function () {
        openBlueprintsScroll.call(this);
    }
}

function openBlueprintsScroll() {
    this.innerText = " ▼ Чертежи";

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    scroll.onclick = function () {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };
    this.appendChild(scroll);

    const types = [
        {name: "Оружие", categories: ["Лазерное", "Ракетное", "Балистическое"], type: "weapon"},
        {name: "Корпуса", categories: ["MS", "Тяжелые", "Средние", "Легкие"], type: "body"},
        {name: "Оборудование", categories: ["1", "2", "3", "4", "5"], type: "equip"},
        {name: "Боеприпасы", categories: ["Лазерное", "Ракетное", "Балистическое"], type: "ammo"},
        {name: "Детали", categories: [], type: "detail"},
        {name: "Ящики", categories: [], type: "boxes"},
    ];

    for (let i = 0; i < types.length; i++) {
        let sub = document.createElement("span");
        sub.innerText = " ▶ " + types[i].name;

        sub.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepBlueprints.call(this, types[i])
        };

        scroll.appendChild(sub);
    }

    this.onclick = function () {
        // innerHTML вытерает внутрености, а то я 15 минут вспоминал почему все удаляется Х)
        clearFilter();
        this.innerHTML = " ▶ Чертежи";
        this.onclick = openBlueprintsScroll;
    }
}

function openDeepBlueprints(assortment) {

    this.innerText = " ▼ " + assortment.name;

    this.onclick = function () {
        clearFilter();
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        this.innerHTML = " ▶ " + assortment.name;
        this.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepBlueprints.call(this, assortment)
        }
    };

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    scroll.onclick = function () {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };
    this.appendChild(scroll);

    for (let i in blueprints) {
        if (blueprints.hasOwnProperty(i) && blueprints[i].item_type === assortment.type) {
            let bp = document.createElement("span");
            bp.innerText = blueprints[i].name;
            bp.onclick = function () {
                event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
                selectItem(blueprints[i].id, "blueprints", blueprints[i].name);
            };

            scroll.appendChild(bp);
        }
    }
}