function CreateColorInputs(unitIcon, unit, unitSlot, idPrefix) {

    let createColorInput = function (id, bColor, unitSlot) {
        let labelColor = document.createElement('label');
        labelColor.id = id;
        labelColor.className = 'changeColorUnit';
        labelColor.style.background = bColor;
        labelColor.setAttribute("hex-color", bColor);
        unitIcon.appendChild(labelColor);

        labelColor.onmouseover = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        };

        labelColor.onmousemove = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        };

        labelColor.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

            let colorUnitPicker = $('#colorUnitPicker');

            $('.colorpicker').fadeToggle("fast", "linear");
            initColorPicker();

            colorUnitPicker.off("click");
            colorUnitPicker.click(function () {
                $('.colorpicker').css("display", "none");

                if (Number(unitSlot) === 0) {
                    inventorySocket.send(JSON.stringify({
                        event: "changeColor",
                        unit_slot: Number(unitSlot),
                        body_color_1: '0x' + document.getElementById('msbodyColor1').getAttribute("hex-color").split('#')[1],
                        body_color_2: '0x' + document.getElementById('msbodyColor2').getAttribute("hex-color").split('#')[1],
                        weapon_color_1: '0x' + document.getElementById('msweaponColor1').getAttribute("hex-color").split('#')[1],
                        weapon_color_2: '0x' + document.getElementById('msweaponColor2').getAttribute("hex-color").split('#')[1],
                    }));
                } else {
                    inventorySocket.send(JSON.stringify({
                        event: "changeColor",
                        unit_slot: Number(unitSlot),
                        body_color_1: '0x' + document.getElementById('unitbodyColor1').getAttribute("hex-color").split('#')[1],
                        body_color_2: '0x' + document.getElementById('unitbodyColor2').getAttribute("hex-color").split('#')[1],
                        weapon_color_1: '0x' + document.getElementById('unitweaponColor1').getAttribute("hex-color").split('#')[1],
                        weapon_color_2: '0x' + document.getElementById('unitweaponColor2').getAttribute("hex-color").split('#')[1],
                    }));
                }
            });

            colorUnitPicker.off("mousemove");
            colorUnitPicker.mousemove(function (e) { // mouse move handler
                event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

                let canvas = document.getElementById('colorUnitPicker');
                let ctx = canvas.getContext('2d');

                let canvasOffset = $(canvas).offset();
                let canvasX = Math.floor(e.clientX - canvasOffset.left);
                let canvasY = Math.floor(e.clientY - canvasOffset.top);
                // get current pixel
                let imageData = ctx.getImageData(canvasX, canvasY, 1, 1);
                let pixel = imageData.data;
                let dColor = pixel[2] + 256 * pixel[1] + 65536 * pixel[0];
                let hexColor = '#' + ('0000' + dColor.toString(16)).substr(-6);
                labelColor.style.background = hexColor;
                labelColor.setAttribute("hex-color", hexColor);


                if (unitSlot === 0) {
                    if (id === 'msbodyColor1') {
                        document.getElementById('msBodyMask1').style.background = hexColor
                    }
                    if (id === 'msbodyColor2') {
                        document.getElementById('msBodyMask2').style.background = hexColor
                    }
                    if (id === 'msweaponColor1') {
                        document.getElementById('msWeaponMask1').style.background = hexColor
                    }
                    if (id === 'msweaponColor2') {
                        document.getElementById('msWeaponMask2').style.background = hexColor
                    }
                } else {
                    if (id === 'unitbodyColor1') {
                        document.getElementById('unitMaskBody1').style.background = hexColor
                    }
                    if (id === 'unitbodyColor2') {
                        document.getElementById('unitMaskBody2').style.background = hexColor
                    }
                    if (id === 'unitweaponColor1') {
                        document.getElementById('unitWeaponMask1').style.background = hexColor
                    }
                    if (id === 'unitweaponColor2') {
                        document.getElementById('unitWeaponMask2').style.background = hexColor
                    }
                }
            });
        };
    };

    createColorInput(idPrefix + 'bodyColor1', "#" + unit.body_color_1.split('x')[1], unitSlot);
    createColorInput(idPrefix + 'bodyColor2', "#" + unit.body_color_2.split('x')[1], unitSlot);
    createColorInput(idPrefix + 'weaponColor1', "#" + unit.weapon_color_1.split('x')[1], unitSlot);
    createColorInput(idPrefix + 'weaponColor2', "#" + unit.weapon_color_2.split('x')[1], unitSlot);
}

function initColorPicker() {

    let brightnessColorPicker = document.getElementById('brightnessColorPicker');

    let canvas = document.getElementById('colorUnitPicker');
    let ctx = canvas.getContext('2d');
    let image = new Image();
    image.src = '/assets/components/inventory/img/colorwheel2.png';
    image.onload = function () {
        ctx.filter = "brightness(" + brightnessColorPicker.value + "%)";
        ctx.drawImage(image, 0, 0);
        ctx.drawImage(image, 0, 0, image.width, image.height);
    };

    brightnessColorPicker.oninput = function () {
        initColorPicker(this.value)
    }
}