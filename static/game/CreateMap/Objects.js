function CreateObject(coordinate, x, y) {
    let object;

    if (coordinate.unit_overlap) {
        object = gameObjectCreate(x, y, coordinate.texture, coordinate.scale, coordinate.shadow, coordinate.rotate,
            game.floorOverObjectLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset,
            coordinate.shadow_intensity);
    } else {
        object = gameObjectCreate(x, y, coordinate.texture, coordinate.scale, coordinate.shadow, coordinate.rotate,
            game.floorObjectLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset,
            coordinate.shadow_intensity);
    }

    if (game.typeService !== "mapEditor") {
        // TODO метод вызывающий фризы
        //ObjectEvents(coordinate, object, x, y);
    }

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture, scale, needShadow, rotate, group, xShadowOffset, yShadowOffset, shadowIntensity) {
    let shadow;

    if (needShadow) {
        shadow = group.create(x + game.shadowXOffset + xShadowOffset, y + game.shadowYOffset + yShadowOffset, texture);
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

    object.inputEnabled = true;
    object.input.pixelPerfectOver = true;
    object.input.pixelPerfectClick = true;
    object.input.pixelPerfectAlpha = 1;

    if (coordinate.object_name !== "") {

        let tip;
        let posInterval;
        object.events.onInputOver.add(function () {

            // из за того что эта очень ресурсоемкая операция приходится вот так извращатся
            if (!object.border) {
                if (coordinate.unit_overlap) {
                    object.border = CreateBorder(x, y, coordinate.texture, coordinate.scale, coordinate.rotate, game.floorOverObjectLayer);
                    game.floorOverObjectLayer.swap(object, object.border);
                } else {
                    object.border = CreateBorder(x, y, coordinate.texture, coordinate.scale, coordinate.rotate, game.floorObjectLayer);
                    game.floorObjectLayer.swap(object, object.border);
                }
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
