function OptionSprite() {
    let callBack = function (q, r) {
        ChangeOptionSprite(q, r)
    };
    SelectedSprite(event, 0, callBack, true)
}

function ChangeOptionSprite(q, r) {
    let coordinate = game.map.OneLayerMap[q][r];

    let block = document.getElementById("coordinates");

    let rotate = document.createElement("div");
    rotate.id = "rotateBlock";

    let rotateRange = createRange("rotateRange", 0, 360, 1, coordinate.objectSprite.angle);
    rotateRange.oninput = function () {
        document.getElementById("rotateOutput").innerHTML = rotateRange.value;
        coordinate.objectSprite.angle = rotateRange.value;
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.angle = rotateRange.value;
        }
    };
    let outputRotate = document.createElement("div");
    outputRotate.innerHTML = "<span> Градусы: </span> <span id='rotateOutput'> " + coordinate.objectSprite.angle + " </span>";

    if (coordinate.animate_sprite_sheets !== "") {
        let rangeAnimateSpeed = createRange("speedRange", 0, 300, 1, coordinate.objectSprite.animation_speed);
        rangeAnimateSpeed.oninput = function () {
            document.getElementById("speedOutput").innerHTML = rangeAnimateSpeed.value;
            coordinate.objectSprite.animations.getAnimation('objAnimate').stop();
            coordinate.objectSprite.animations.play('objAnimate', Number(rangeAnimateSpeed.value), true);
            if (coordinate.objectSprite.shadow) {
                coordinate.objectSprite.shadow.animations.getAnimation('objAnimate').stop();
                coordinate.objectSprite.shadow.animations.play('objAnimate', Number(rangeAnimateSpeed.value), true);
            }
        };
        let outputSpeed = document.createElement("div");
        outputSpeed.innerHTML = "<span> Кадров в сек: </span> <span id='speedOutput'> " + coordinate.animation_speed + " </span>";
        rotate.appendChild(rangeAnimateSpeed);
        rotate.appendChild(outputSpeed);
    }

    let rangeXOffset = createRange("rangeXOffset", -100, 100, 1, coordinate.x_offset);
    rangeXOffset.oninput = function () {
        document.getElementById("XOutput").innerHTML = rangeXOffset.value;
        coordinate.objectSprite.x = coordinate.sprite.x + Number(rangeXOffset.value);
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.x = coordinate.sprite.x + game.shadowXOffset + Number(rangeXOffset.value);
        }
    };

    let outputXOffset = document.createElement("div");
    outputXOffset.innerHTML = "<span> Смещение по Х: </span> <span id='XOutput'> " + coordinate.x_offset + " </span>";

    let rangeYOffset = createRange("rangeYOffset", -100, 100, 1, coordinate.y_offset);
    rangeYOffset.oninput = function () {
        document.getElementById("YOutput").innerHTML = rangeYOffset.value;
        coordinate.objectSprite.y = coordinate.sprite.y + Number(rangeYOffset.value);
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.y = coordinate.sprite.y + game.shadowYOffset + Number(rangeYOffset.value);
        }
    };

    let outputYOffset = document.createElement("div");
    outputYOffset.innerHTML = "<span> Смещение по Y: </span> <span id='YOutput'> " + coordinate.y_offset + " </span>";

    let apply = document.createElement("input");
    apply.value = "Применить";
    apply.type = "submit";
    apply.onclick = function () {
        let speed = 0;
        if (document.getElementById("speedRange")) {
            speed = Number(document.getElementById("speedRange").value);
        }

        mapEditor.send(JSON.stringify({
            event: "rotateObject",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(coordinate.q),
            r: Number(coordinate.r),
            rotate: Number(document.getElementById("rotateRange").value),
            speed: speed,
            x_offset: Number(document.getElementById("rangeXOffset").value),
            y_offset: Number(document.getElementById("rangeYOffset").value)
        }));

        mapEditor.send(JSON.stringify({
            event: "getAllTypeCoordinate"
        }));
    };

    let cancel = document.createElement("input");
    cancel.value = "Отменить";
    cancel.type = "submit";
    cancel.onclick = function () {
        mapEditor.send(JSON.stringify({
            event: "SelectMap",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value)
        }));
    };

    rotate.appendChild(rangeYOffset);
    rotate.appendChild(outputYOffset);

    rotate.appendChild(rangeXOffset);
    rotate.appendChild(outputXOffset);

    rotate.appendChild(rotateRange);
    rotate.appendChild(outputRotate);

    rotate.appendChild(apply);
    rotate.appendChild(cancel);

    block.appendChild(rotate);
}

function createRange(id, min, max, step, startValue) {

    let range = document.createElement("input");
    range.id = id;
    range.type = "range";
    range.min = min;
    range.max = max;
    range.step = step;
    range.value = startValue;

    return range
}