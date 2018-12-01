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
        head.innerHTML = "Выбери на что заменить";

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
        let block = document.getElementById("coordinates");
        let notification = document.createElement("div");
        notification.id = "notification";

        let head = document.createElement("h4");
        head.innerHTML = "Выуверены что ходитет удолить?";

        let ok = document.createElement("input");
        ok.value = "УДОЛИ!";
        ok.type = "submit";
        ok.onclick = function(){
            mapEditor.send(JSON.stringify({
                event: "deleteType",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                id_type: Number(typeCoordinate.id),
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

        notification.appendChild(head);
        notification.appendChild(cancel);
        notification.appendChild(ok);
        block.appendChild(notification);
    };
    menu.appendChild(del);

    let edit = document.createElement("input");
    edit.type = "submit";
    edit.value = "изменить";
    edit.onclick = function () {

        let block = document.getElementById("coordinates");

        let changeType = document.createElement("div");
        changeType.id = "changeType";

        let watchText = document.createElement("span");
        watchText.innerHTML = "Обзор";
        let watch = document.createElement("input");
        watch.id = "changeWatch";
        watch.type = "checkbox";
        watch.checked = typeCoordinate.view;

        let moveText = document.createElement("span");
        moveText.innerHTML = "Движение";
        let move = document.createElement("input");
        move.id = "changeMove";
        move.type = "checkbox";
        move.checked = typeCoordinate.move;

        let attackText = document.createElement("span");
        attackText.innerHTML = "Атака";
        let attack = document.createElement("input");
        attack.id = "changeAttack";
        attack.type = "checkbox";
        attack.checked = typeCoordinate.attack;

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
                shadow: document.getElementById("changeShadow").checked,
                move: document.getElementById("changeMove").checked,
                watch: document.getElementById("changeWatch").checked,
                attack: document.getElementById("changeAttack").checked,
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


        changeType.appendChild(watchText);
        changeType.appendChild(watch);
        changeType.appendChild(document.createElement("br"));

        changeType.appendChild(moveText);
        changeType.appendChild(move);
        changeType.appendChild(document.createElement("br"));

        changeType.appendChild(attackText);
        changeType.appendChild(attack);
        changeType.appendChild(document.createElement("br"));

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