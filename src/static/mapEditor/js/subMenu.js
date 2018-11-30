function CreateSubMenu(typeCoordinate) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

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

        removeSubMenus();

        let types = document.getElementsByClassName("coordinateType");
        let menus = document.getElementsByClassName("menuButton");

        while (menus.length > 0) {
            menus[0].parentNode.removeChild(menus[0]);
        }

        let block = document.getElementById("coordinates");
        let notification = document.createElement("div");
        notification.id = "notification";

        let head = document.createElement("h4");

        let error = document.createElement("h5");

        let cancel = document.createElement("input");
        cancel.value = "Отменить";
        cancel.type = "submit";
        cancel.onclick = function(){
            mapEditor.send(JSON.stringify({
                event: "getAllTypeCoordinate"
            }));
        };

        notification.appendChild(head);
        notification.appendChild(error);
        notification.appendChild(cancel);
        block.appendChild(notification);

        for (let i = 0; i < types.length; i++) {
            types[i].onmousemove = null;
            types[i].onclick = function () {

                if (Number(Number(types[i].coordinateType.id)) === Number(typeCoordinate.id)) {
                    error.innerHTML = "Выбран тот же тип!";
                    return
                }

                mapEditor.send(JSON.stringify({
                    event: "replaceCoordinateType",
                    new_id_type: Number(Number(types[i].coordinateType.id)),
                    old_id_type: Number(typeCoordinate.id),
                    id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value)
                }));

                mapEditor.send(JSON.stringify({
                    event: "getAllTypeCoordinate"
                }));
            }
        }
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

        mapEditor.send(JSON.stringify({
            event: "getAllTypeCoordinate"
        }));
    };
    menu.appendChild(del);

    let edit = document.createElement("input");
    edit.type = "submit";
    edit.value = "изменить";
    edit.onclick = function () {
        // todo добавить изменение move, attack, watch
        let block = document.getElementById("coordinates");

        let changeType = document.createElement("div");
        changeType.id = "changeType";

        let shadowText = document.createElement("span");
        shadowText.innerHTML = "Тень";
        let shadow = document.createElement("input");
        shadow.id = "changeShadow";
        shadow.type = "checkbox";
        shadow.checked = typeCoordinate.shadow;

        let scaleText = document.createElement("span");
        scaleText.innerHTML = "Размер";
        let scale = document.createElement("input");
        scale.id = "changeScale";
        scale.type = "number";
        scale.value = typeCoordinate.scale;

        let apply = document.createElement("input");
        apply.value = "Применить";
        apply.type = "submit";
        apply.onclick = function(){
            mapEditor.send(JSON.stringify({
                event: "changeType",
                id_type: Number(typeCoordinate.id),
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                scale: Number(document.getElementById("changeScale").value),
                shadow: document.getElementById("changeShadow").checked
            }));

            mapEditor.send(JSON.stringify({
                event: "getAllTypeCoordinate"
            }));
        };

        let cancel = document.createElement("input");
        cancel.value = "Отменить";
        cancel.type = "submit";
        cancel.onclick = function(){
            mapEditor.send(JSON.stringify({
                event: "getAllTypeCoordinate"
            }));
        };

        changeType.appendChild(shadowText);
        changeType.appendChild(shadow);

        changeType.appendChild(document.createElement("br"));

        changeType.appendChild(scaleText);
        changeType.appendChild(scale);

        changeType.appendChild(cancel);
        changeType.appendChild(apply);
        block.appendChild(changeType);
    };
    menu.appendChild(edit);

    let close = document.createElement("input");
    close.type = "submit";
    close.value = "закрыть";
    close.onclick = function () {
        removeSubMenus();
        menu.remove();
    };
    menu.appendChild(close);

    document.body.appendChild(menu)
}