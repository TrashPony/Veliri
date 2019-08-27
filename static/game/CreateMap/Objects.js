function CreateObject(coordinate, x, y) {
    let object;

    if (coordinate.impact) {
        return
    }

    if (coordinate.unit_overlap) {
        object = gameObjectCreate(x, y, coordinate.texture_object, coordinate.scale, coordinate.shadow, coordinate.obj_rotate,
            coordinate.x_offset, coordinate.y_offset, game.floorOverObjectLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset,
            coordinate.shadow_intensity);
    } else {
        object = gameObjectCreate(x, y, coordinate.texture_object, coordinate.scale, coordinate.shadow, coordinate.obj_rotate,
            coordinate.x_offset, coordinate.y_offset, game.floorObjectLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset,
            coordinate.shadow_intensity);
    }

    ObjectEvents(coordinate, object);

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture, scale, needShadow, rotate, xOffset, yOffset, group, xShadowOffset, yShadowOffset, shadowIntensity) {
    let shadow;

    if (needShadow) {
        shadow = group.create(x + xOffset + game.shadowXOffset + xShadowOffset, y + yOffset + game.shadowYOffset + yShadowOffset, texture);
        shadow.anchor.setTo(0.5);
        shadow.scale.set((scale / 100) / 2);
        shadow.tint = 0x000000;
        shadow.angle = rotate;
        shadow.alpha = shadowIntensity / 100;
    }

    // TODO это происходит ебать как долго
    // но работает

    // let borderSprite = createSillhouette(texture);
    // let border = group.create(x + xOffset, y + yOffset, borderSprite);
    // border.scale.set(((scale / 100) / 1.9));
    // border.anchor.setTo(0.5);
    // border.tint = 0xfffab0;
    // border.visible = false;
    // border.angle = rotate;

    let object = group.create(x + xOffset, y + yOffset, texture);
    object.anchor.setTo(0.5);
    object.scale.set((scale / 100) / 2);
    object.angle = rotate;

    object.shadow = shadow;
    //object.border = border;

    return object
}

function ObjectEvents(coordinate, object) {

    object.inputEnabled = true;
    object.input.pixelPerfectOver = true;
    object.input.pixelPerfectClick = true;
    object.input.pixelPerfectAlpha = 1;

    if (coordinate.object_name !== "") {

        let tip;
        let posInterval;
        object.events.onInputOver.add(function () {

            //object.border.visible = true;
            tip = document.createElement("div");
            tip.id = "reservoirTip" + coordinate.q + "" + coordinate.r;
            tip.className = "reservoirTip";
            tip.style.left = stylePositionParams.left + "px";
            tip.style.top = stylePositionParams.top + "px";
            document.body.appendChild(tip);

            tip.innerHTML = `
            <h3>${coordinate.object_name}</h3>
            <div class="Description" style="margin-bottom: 5px"> ${coordinate.object_description}</div>
            `;

            posInterval = setInterval(function () {
                tip.style.left = stylePositionParams.left + "px";
                tip.style.top = stylePositionParams.top + "px";
            }, 10)
        });

        object.events.onInputOut.add(function () {
            //object.border.visible = false;
            setInterval(posInterval);
            tip.remove();
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

    if (coordinate.texture_object.indexOf('base') + 1) {
        // todo выводить окошо с мин информацией по базе
    }
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
