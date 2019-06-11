let resource, recycles, detail;

function fillRes(putResource, putRecycles, putDetail) {

    resource = putResource;
    recycles = putRecycles;
    detail = putDetail;

    let filterBlock = document.getElementById("resCategoryItem");
    filterBlock.onclick = function () {
        openResourceScroll.call(this);
    }
}

function openResourceScroll() {

    this.innerText = " ▼ Ресурсы";

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    scroll.onclick = function () {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };
    this.appendChild(scroll);

    const types = [
        {name: "Ископаемые", assortment: resource, type: "resource"},
        {name: "Сырье", assortment: recycles, type: "recycle"},
        {name: "Детали", assortment: detail, type: "detail"}
    ];

    for (let i = 0; i < types.length; i++) {
        let sub = document.createElement("span");
        sub.innerText = " ▶ " + types[i].name;

        sub.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepResource.call(this, types[i])
        };

        scroll.appendChild(sub);
    }

    this.onclick = function () {
        // innerHTML вытерает внутрености, а то я 15 минут вспоминал почему все удаляется Х)
        this.innerHTML = " ▶ Ресурсы";
        this.onclick = openResourceScroll;
        clearFilter();
    }
}

function openDeepResource(assortment) {
    this.innerText = " ▼ " + assortment.name;

    this.onclick = function () {
        clearFilter();
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        this.innerHTML = " ▶ " + assortment.name;
        this.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            openDeepResource.call(this, assortment)
        }
    };

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    scroll.onclick = function () {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };
    this.appendChild(scroll);

    for (let i in assortment.assortment) {

        let res = document.createElement("span");
        res.innerText = assortment.assortment[i].name;
        res.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            selectItem(assortment.assortment[i].id, assortment.type, assortment.assortment[i].name);
        };

        scroll.appendChild(res);
    }
}