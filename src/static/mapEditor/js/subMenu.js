function CreateSubMenu(typeCoordinate) {
    console.log(typeCoordinate);

    if (document.getElementById("menuBlock")) {
        document.getElementById("menuBlock").remove();
    }

    let menu = document.createElement("div");
    menu.id = "menuBlock";

    menu.style.top = stylePositionParams.top + "px";
    menu.style.left = stylePositionParams.left + "px";

    let replace = document.createElement("input");
    replace.type = "submit";
    replace.value = "Заменить";
    replace.onclick = function () {
        // todo
    };
    menu.appendChild(replace);

    let del = document.createElement("input");
    del.type = "submit";
    del.value = "удалить";
    del.onclick = function () {
        mapEditor.send(JSON.stringify({
            event: "deleteType",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            id_type: Number(typeCoordinate.id),
        }));
    };
    menu.appendChild(del);

    let edit = document.createElement("input");
    edit.type = "submit";
    edit.value = "изменить";
    edit.onclick = function () {
        // todo
    };
    menu.appendChild(edit);

    let close = document.createElement("input");
    close.type = "submit";
    close.value = "закрыть";
    close.onclick = function () {
        menu.remove();
    };
    menu.appendChild(close);

    document.body.appendChild(menu)
}