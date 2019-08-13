let responseChangeOption = [];

function OptionSprite() {
    let callBack = function (q, r) {
        ChangeOptionSprite(q, r)
    };
    SelectedSprite(event, 0, callBack, true)
}

function ChangeOptionSprite(q, r) {
    let coordinate = game.map.OneLayerMap[q][r];

    let block = document.getElementById("coordinates");
    if (document.getElementById("rotateBlock")) document.getElementById("rotateBlock").remove();
    let rotate = document.createElement("div");
    block.appendChild(rotate);

    let rotateSprite = coordinate.obj_rotate;
    if (rotateSprite < 0) rotateSprite += 360;

    rotate.style.height = "420px";
    rotate.id = "rotateBlock";
    rotate.innerHTML = `
            <div><span> Кадров в сек: </span> <span id="speedOutput"> ${coordinate.animation_speed} </span></div>
            <input id="speedRange" value="${coordinate.animation_speed}" type="range" min="0" max="300" step="1">
            <div><span> Смещение Тени по Y: </span> <span id="YShadowOutput"> ${coordinate.y_shadow_offset} </span></div>
            <input id="rangeShadowYOffset"  value="${coordinate.y_shadow_offset}" type="range" min="-100" max="200" step="1">
            <div><span> Смещение Тени по Х: </span> <span id="XShadowOutput"> ${coordinate.x_shadow_offset} </span></div>
            <input id="rangeShadowXOffset" value="${coordinate.x_shadow_offset}" type="range" min="-100" max="200" step="1">
            <div><span> Интенсивность тени: </span> <span id="YShadowIOutput"> ${coordinate.shadow_intensity} </span></div>
            <input id="rangeShadowIntensity" value="${coordinate.shadow_intensity}" type="range" min="0" max="100" step="1">
            <div><span> Смещение по Y: </span> <span id="YOutput"> ${coordinate.y_offset} </span></div>
            <input id="rangeYOffset" value="${coordinate.y_offset}" type="range" min="-100" max="100" step="1">
            <div><span> Смещение по Х: </span> <span id="XOutput"> ${coordinate.x_offset} </span></div>
            <input id="rangeXOffset" value="${coordinate.x_offset}" type="range" min="-100" max="100" step="1">
            <div><span> Градусы: </span> <span id="rotateOutput"> ${rotateSprite} </span></div>
            <input id="rotateRange" value="${rotateSprite}" type="range" min="0" max="360" step="1">
            
            <div><span> Размер: </span> <span id="scaleOutput"> ${coordinate.scale} </span></div>
            <input id="scaleRange" value="${coordinate.scale}" type="range" min="0" max="200" step="1">
            
            <div><span> Тень: </span> <input type="checkbox" id="needShadow"></div>
            
            <input type="submit" value="Вперед" id="frontSprite">
            <input type="submit" value="Назад" id="backSprite">
            
            <br>
            <input type="submit" value="Применить" id="applyRotate">
            <input type="submit" value="Отменить" id="cancelRotate">

`;

    // TODO включение и отключение теней при изменение состояние чекбокса для теней

    let xy = GetXYCenterHex(q, r);

    // отправить спрайт на задний план
    $('#backSprite').click(function () {
        mapEditor.send(JSON.stringify({
            event: "toBack",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(coordinate.q),
            r: Number(coordinate.r),
        }));

        rotate.remove();
    });

    $('#frontSprite').click(function () {
        mapEditor.send(JSON.stringify({
            event: "toFront",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
            q: Number(coordinate.q),
            r: Number(coordinate.r),
        }));

        for (let i in game.mapPoints) {
            if (game.mapPoints[i].q === Number(coordinate.q) && game.mapPoints[i].r === Number(coordinate.r)) {
                ReloadCoordinate(game.mapPoints[i]);
            }
        }

        rotate.remove();
    });

    // Угол поворота
    $('#rotateRange').on('input', function () {
        $('#rotateOutput').text(this.value);
        coordinate.objectSprite.angle = this.value;
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.angle = this.value;
        }
    });

    // Размер
    $('#scaleRange').on('input', function () {
        $('#scaleOutput').text(this.value);
        coordinate.objectSprite.scale.set((this.value / 100) / 2);
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.scale.set((this.value / 100) / 2);
        }
    });

    // Скорость анимации
    $('#speedRange').on('input', function () {
        $('#speedOutput').text(this.value);
        coordinate.objectSprite.animations.getAnimation('objAnimate').stop();
        coordinate.objectSprite.animations.play('objAnimate', Number(this.value), true);
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.animations.getAnimation('objAnimate').stop();
            coordinate.objectSprite.shadow.animations.play('objAnimate', Number(this.value), true);
        }
    });

    // Сдвиг по Х
    $('#rangeXOffset').on('input', function () {
        $('#XOutput').text(this.value);
        coordinate.objectSprite.x = xy.x + Number(this.value);
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.x = xy.x + game.shadowXOffset + Number($('#rangeShadowXOffset').val()) + Number($('#rangeXOffset').val());
        }
    });

    // Сдвиг по У
    $('#rangeYOffset').on('input', function () {
        $('#YOutput').text(this.value);
        coordinate.objectSprite.y = xy.y + Number(this.value);
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.y = xy.y + game.shadowYOffset + Number($('#rangeShadowYOffset').val()) + Number($('#rangeYOffset').val());
        }
    });


    // Тень
    $('#needShadow').prop('checked', coordinate.shadow);
    // Сдвиг тени по Х
    $('#rangeShadowXOffset').on('input', function () {
        $('#XShadowOutput').text(this.value);
        coordinate.objectSprite.shadow.x = xy.x + game.shadowXOffset + Number($('#rangeShadowXOffset').val()) + Number($('#rangeXOffset').val());
    });

    // Сдвиг тени по Y
    $('#rangeShadowYOffset').on('input', function () {
        $('#YShadowOutput').text(this.value);
        coordinate.objectSprite.shadow.y = xy.y + game.shadowYOffset + Number($('#rangeShadowYOffset').val()) + Number($('#rangeYOffset').val());
    });

    // Интенсивность тени
    $('#rangeShadowIntensity').on('input', function () {
        $('#YShadowIOutput').text(this.value);
        coordinate.objectSprite.shadow.alpha = Number(rangeShadowIntensity.value) / 100;
    });

    $('#applyRotate').click(function () {
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
            y_offset: Number(document.getElementById("rangeYOffset").value),
            x_shadow_offset: Number(document.getElementById("rangeShadowXOffset").value),
            y_shadow_offset: Number(document.getElementById("rangeShadowYOffset").value),
            shadow_intensity: Number(document.getElementById("rangeShadowIntensity").value),
            scale: Number(document.getElementById("scaleRange").value),
            shadow: $('#needShadow').prop('checked'),
        }));

        for (let i in game.mapPoints) {
            if (game.mapPoints[i].q === Number(coordinate.q) && game.mapPoints[i].r === Number(coordinate.r)) {

                game.mapPoints[i].coordinate.obj_rotate = Number(document.getElementById("rotateRange").value);
                game.mapPoints[i].coordinate.animation_speed = speed;
                game.mapPoints[i].coordinate.x_offset = Number(document.getElementById("rangeXOffset").value);
                game.mapPoints[i].coordinate.y_offset = Number(document.getElementById("rangeYOffset").value);
                game.mapPoints[i].coordinate.x_shadow_offset = Number(document.getElementById("rangeShadowXOffset").value);
                game.mapPoints[i].coordinate.y_shadow_offset = Number(document.getElementById("rangeShadowYOffset").value);
                game.mapPoints[i].coordinate.shadow_intensity = Number(document.getElementById("rangeShadowIntensity").value);
                game.mapPoints[i].coordinate.scale = Number(document.getElementById("scaleRange").value);
                game.mapPoints[i].coordinate.shadow = $('#needShadow').prop('checked');

                ReloadCoordinate(game.mapPoints[i]);
            }
        }

        rotate.remove();
    });

    $('#cancelRotate').click(function () {
        mapEditor.send(JSON.stringify({
            event: "SelectMap",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value)
        }));
    });
}