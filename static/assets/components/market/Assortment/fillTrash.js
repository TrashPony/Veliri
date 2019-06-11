let trash;

function fillTrash(putTrash) {
    trash = putTrash;

    let filterBlock = document.getElementById("trashCategoryItem");
    filterBlock.onclick = function () {
        openTrashScroll.call(this);
    }
}

function openTrashScroll() {
    this.innerText = " ▼ Хлам";

    let scroll = document.createElement("div");
    scroll.className = "scrollFilter";
    scroll.onclick = function () {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };
    this.appendChild(scroll);

    for (let i in trash) {
        let trashBlock = document.createElement("span");
        trashBlock.innerText = trash[i].name;
        trashBlock.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            selectItem(trash[i].id, "trash", trash[i].name);
        };

        scroll.appendChild(trashBlock);
    }

    this.onclick = function () {
        // innerHTML вытерает внутрености, а то я 15 минут вспоминал почему все удаляется Х)
        this.innerHTML = " ▶ Хлам";
        this.onclick = openTrashScroll;
        clearFilter();
    }
}