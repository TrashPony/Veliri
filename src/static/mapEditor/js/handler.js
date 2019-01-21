function RemoveHandler() {
    let callBack = function (q, r) {
        if (game.input.activePointer.leftButton.isDown) {
            mapEditor.send(JSON.stringify({
                event: "removeHandler",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                q: Number(q),
                r: Number(r)
            }));
        }
    };
    SelectedSprite(event, 0, callBack)
}

function AddHandler() {
    let callBack = function (q, r) {
        ChangeHandlerOption(q, r);
    };
    SelectedSprite(event, 0, callBack)
}

function ChangeHandlerOption(q, r) {
    if (document.getElementById("handlerBlockOption")) document.getElementById("handlerBlockOption").remove();

    let handlerBlockOption = document.createElement("div");
    handlerBlockOption.id = "handlerBlockOption";

    let typeSelect = document.createElement("select");
    typeSelect.id = "typeSelect";
    typeSelect.innerHTML = "" +
        "<option value disabled selected> Выберите тип перехода </option>" +
        "<option value='base'>Base</option>" +
        "<option value='sector'>Sector</option>";

    let baseSelect = document.createElement("select");
    baseSelect.id = "baseSelect";
    baseSelect.style.opacity = 0;
    baseSelect.innerHTML = "<option value disabled selected> Выберите базу </option>";
    for (let i in game.bases) {
        baseSelect.innerHTML += "<option value='" + game.bases[i].id + "'> " + game.bases[i].name + " </option>"
    }

    let mapOptions = document.getElementById("mapSelector").innerHTML;
    let mapSelect = document.createElement("select");
    mapSelect.style.opacity = 0;
    mapSelect.id = "mapSelect";
    mapSelect.innerHTML = mapOptions;

    let position = document.createElement("div");
    position.style.opacity = 0;
    position.innerHTML = "" +
        "<span>to Q</span><input type='number' id='toQ' value='0'><br>" +
        "<span>to R</span><input type='number' id='toR' value='0'>";

    let apply = document.createElement("input");
    apply.value = "Отменить";
    apply.type = "submit";
    apply.onclick = function () {
        if (document.getElementById('typeSelect').value === 'base') {
            mapEditor.send(JSON.stringify({
                event: "addHandler",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                q: Number(q),
                r: Number(r),
                to_base_id: Number(document.getElementById('baseSelect').value),
                type_handler: document.getElementById('typeSelect').value,
            }));
        } else if (document.getElementById('typeSelect').value === 'sector') {
            mapEditor.send(JSON.stringify({
                event: "addHandler",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                q: Number(q),
                r: Number(r),
                to_q: Number(document.getElementById('toQ').value),
                to_r: Number(document.getElementById('toR').value),
                to_map_id: Number(document.getElementById('mapSelect').value),
                type_handler: document.getElementById('typeSelect').value,
            }));
        }
        handlerBlockOption.remove();
    };


    typeSelect.onchange = function () {
        if (this.value === "base") {
            apply.value = "Применить";
            position.style.opacity = 0;
            mapSelect.style.opacity = 0;
            baseSelect.style.opacity = 1;
        } else if (this.value === "sector") {
            apply.value = "Применить";
            position.style.opacity = 1;
            mapSelect.style.opacity = 1;
            baseSelect.style.opacity = 0;
        } else {
            apply.value = "отменить";
        }
    };

    handlerBlockOption.appendChild(typeSelect);
    handlerBlockOption.appendChild(baseSelect);
    handlerBlockOption.appendChild(mapSelect);
    handlerBlockOption.appendChild(position);
    handlerBlockOption.appendChild(apply);
    document.getElementById("coordinates").appendChild(handlerBlockOption);
}