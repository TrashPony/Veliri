function CreateObject(coordinate, x, y) {

    // todo некоторая логика растения слишком сложна что бы мне было не лень ее вносить в бд поэтому я буду писать ее тут
    if (coordinate.texture === 'plant_4' && coordinate.scale <= 10) {
        // маленькие кусты под машиной
        coordinate.objectSprite = gameObjectCreate(x, y, coordinate.texture, coordinate.scale, coordinate.shadow, coordinate.rotate,
            game.floorObjectLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset, coordinate.shadow_intensity);

    } else if (coordinate.texture === 'plant_5') {
        // деревья всегда вышле всех у них свой особый уровень
        coordinate.objectSprite = gameObjectCreate(x, y, coordinate.texture, coordinate.scale, coordinate.shadow, coordinate.rotate,
            game.rootLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset, coordinate.shadow_intensity);

    } else if (coordinate.unit_overlap) {

        coordinate.objectSprite = gameObjectCreate(x, y, coordinate.texture, coordinate.scale, coordinate.shadow, coordinate.rotate,
            game.floorOverObjectLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset, coordinate.shadow_intensity);

    } else {

        coordinate.objectSprite = gameObjectCreate(x, y, coordinate.texture, coordinate.scale, coordinate.shadow, coordinate.rotate,
            game.floorObjectLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset, coordinate.shadow_intensity);

    }

    if (game.typeService !== "mapEditor") {
        ObjectEvents(coordinate, coordinate.objectSprite, x, y);
    }
}

function gameObjectCreate(x, y, texture, scale, needShadow, rotate, group, xShadowOffset, yShadowOffset, shadowIntensity) {
    let shadow;

    if (needShadow) {
        shadow = group.create(x + xShadowOffset, y + yShadowOffset, texture);
        shadow.anchor.setTo(0.5);
        shadow.scale.set((scale / 100) / 2);
        shadow.tint = 0x000000;
        shadow.angle = rotate;
        shadow.alpha = shadowIntensity / 100;
    }

    let object = group.create(x, y, texture);
    object.anchor.setTo(0.5);
    object.scale.set((scale / 100) / 2);
    object.angle = rotate;

    object.shadow = shadow;

    return object
}

function ObjectEvents(coordinate, object, x, y) {

    if (coordinate.name !== "") {

        object.inputEnabled = true;
        object.input.pixelPerfectOver = true;
        object.input.pixelPerfectClick = true;
        object.input.pixelPerfectAlpha = 1;
        object.input.priorityID = 2;

        // некоторые обьекты имеют имя, но являются проходимыми например как базы, тунели, респауны
        // поэтому вешаем на них эвент движения
        AllowObjectMoveUnit(coordinate, object);

        let tip;
        let posInterval;
        object.events.onInputOver.add(function () {
            // из за того что эта очень ресурсоемкая операция приходится вот так извращатся
            if (!object.border) {
                object.border = CreateBorder(x, y, object.key, coordinate.scale, coordinate.rotate, game[object.parent.name]);
                game[object.parent.name].swap(object, object.border);
            } else {
                object.border.visible = true;
            }

            tip = document.createElement("div");
            tip.id = "reservoirTip" + coordinate.x + "" + coordinate.y;
            tip.className = "reservoirTip";
            tip.style.left = stylePositionParams.left + "px";
            tip.style.top = stylePositionParams.top + "px";
            document.body.appendChild(tip);

            tip.innerHTML = `
            <h3>${coordinate.name}</h3>
            <div class="Description" style="margin-bottom: 5px"> ${coordinate.description}</div>
            `;

            posInterval = setInterval(function () {
                tip.style.left = stylePositionParams.left + "px";
                tip.style.top = stylePositionParams.top + "px";
            }, 10)
        });

        object.events.onInputOut.add(function () {
            if (object.border) object.border.visible = false;
            setInterval(posInterval);
            if (tip) tip.remove();
        });
    }

    if (coordinate.object_inventory) {
        object.events.onInputDown.add(function () {
            game.squad.toBox = {
                boxID: coordinate.box_id,
                to: true,
                x: object.x,
                y: object.y
            }
        })
    }

    if (coordinate.texture.indexOf('base') + 1) {
        // todo выводить окошо с мин информацией по базе
    }
}

function CreateBorder(x, y, texture, scale, rotate, group) {
    let borderSprite = createSillhouette(texture);
    let border = group.create(x, y, borderSprite);
    border.scale.set((scale / 100) / 1.9);
    border.anchor.setTo(0.5);
    border.tint = 0xffffff;
    border.angle = rotate;

    return border
}

function createSillhouette(srcKey) {
    let bmd = game.make.bitmapData();
    // load our texture into the bitmap
    bmd.load(srcKey);
    bmd.processPixelRGB(forEachPixel, this);
    return bmd
}

function forEachPixel(pixel) {
    // processPixelRGB won't take an argument, so we've set our sillhouetteColor globally
    pixel.r = 255;
    pixel.g = 255;
    pixel.b = 255;
    return pixel
}

function AllowObjectMoveUnit(coordinate, object) {
    if (object.key.indexOf('base') > -1 || object.key.indexOf('tunel_out') > -1 || object.key.indexOf('tunel') > -1) {
        object.events.onInputUp.add(initMove);
    }
}