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

    rotate.style.height = "350px";
    rotate.id = "rotateBlock";
    rotate.innerHTML = `
            <div><span> Кадров в сек: </span> <span id="speedOutput"> 60 </span></div>
            <input id="speedRange" type="range" min="0" max="300" step="1">
            <div><span> Смещение Тени по Y: </span> <span id="YShadowOutput"> 0 </span></div>
            <input id="rangeShadowYOffset" type="range" min="-100" max="200" step="1">
            <div><span> Смещение Тени по Х: </span> <span id="XShadowOutput"> 0 </span></div>
            <input id="rangeShadowXOffset" type="range" min="-100" max="200" step="1">
            <div><span> Интенсивность тени: </span> <span id="YShadowIOutput"> 40 </span></div>
            <input id="rangeShadowIntensity" type="range" min="0" max="100" step="1">
            <div><span> Смещение по Y: </span> <span id="YOutput"> 0 </span></div>
            <input id="rangeYOffset" type="range" min="-100" max="100" step="1">
            <div><span> Смещение по Х: </span> <span id="XOutput"> 0 </span></div>
            <input id="rangeXOffset" type="range" min="-100" max="100" step="1">
            <div><span> Градусы: </span> <span id="rotateOutput"> 0 </span></div>
            <input id="rotateRange" type="range" min="0" max="360" step="1">
            
            <div><span> Размер: </span> <span id="scaleOutput"> 0 </span></div>
            <input id="scaleRange" type="range" min="0" max="200" step="1">
            
            <input type="submit" value="Применить" id="applyRotate">
            <input type="submit" value="Отменить" id="cancelRotate">`;

    // TODO отрицательное значение поворота
    // TODO правильное позиционирование тени
    // TODO включение и отключение теней

    let xy = GetXYCenterHex(q, r);

    // Угол поворота
    $('#rotateRange').on('input', function () {
        $('#rotateRange').text(this.value);
        coordinate.objectSprite.angle = this.value;
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.angle = this.value;
        }
    });
    $('#rotateOutput').text(coordinate.objectSprite.angle);

    // Размер
    $('#scaleRange').on('input', function () {
        $('#scaleOutput').text(this.value);
        coordinate.objectSprite.scale.set((this.value / 100) / 2);
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.scale.set((this.value / 100) / 2);
        }
    });
    $('#scaleOutput').text(coordinate.scale);

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
    $('#speedOutput').text(coordinate.animation_speed);

    // Сдвиг по Х
    $('#rangeXOffset').on('input', function () {
        $('#XOutput').text(this.value);
        coordinate.objectSprite.x = xy.x + Number(this.value);
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.x = xy.x + game.shadowXOffset + Number(this.value);
        }
    });
    $('#XOutput').text(coordinate.x_offset);

    // Сдвиг по У
    $('#rangeYOffset').on('input', function () {
        $('#YOutput').text(this.value);
        coordinate.objectSprite.y = xy.y + Number(this.value);
        if (coordinate.objectSprite.shadow) {
            coordinate.objectSprite.shadow.y = xy.y + game.shadowYOffset + Number(this.value);
        }
    });
    $('#YOutput').text(coordinate.y_offset);


    // Сдвиг тени по Х
    $('#rangeShadowXOffset').on('input', function () {
        $('#XShadowOutput').text(this.value);
        coordinate.objectSprite.shadow.x = xy.x + game.shadowXOffset + Number(this.value);
    });
    $('#XShadowOutput').text(coordinate.x_shadow_offset);

    // Сдвиг тени по Y
    $('#rangeShadowYOffset').on('input', function () {
        $('#YShadowOutput').text(this.value);
        coordinate.objectSprite.shadow.x = xy.x + game.shadowXOffset + Number(this.value);
    });
    $('#YShadowOutput').text(coordinate.y_shadow_offset);

    // Интенсивность тени
    $('#rangeShadowIntensity').on('input', function () {
        $('#YShadowIOutput').text(this.value);
        coordinate.objectSprite.shadow.alpha = Number(rangeShadowIntensity.value) / 100;
    });
    $('#YShadowIOutput').text(coordinate.shadow_intensity);

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
        }));
    });

    $('#cancelRotate').click(function () {
        mapEditor.send(JSON.stringify({
            event: "SelectMap",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value)
        }));
    });
}