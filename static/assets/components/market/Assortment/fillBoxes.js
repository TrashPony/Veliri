let boxes;

function fillBoxes(putBoxes) {
    boxes = putBoxes;

    let filterBlock = document.getElementById("boxCategoryItem");
    filterBlock.onclick = function () {
        openBoxesScroll.call(this);
    }
}

function openBoxesScroll() {
    this.innerText = " ▼ Ящики";

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    scroll.onclick = function () {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };
    this.appendChild(scroll);

    for (let i in boxes) {
        let box = document.createElement("span");
        box.innerText = boxes[i].name;
        box.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            selectItem(boxes[i].type_id, "boxes", boxes[i].name);
        };

        scroll.appendChild(box);
    }

    this.onclick = function () {
        // innerHTML вытерает внутрености, а то я 15 минут вспоминал почему все удаляется Х)
        this.innerHTML = " ▶ Ящики";
        this.onclick = openBoxesScroll;
        clearFilter();
    }
}