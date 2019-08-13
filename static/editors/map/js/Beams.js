let lasers = [];
let laser;

function RemoveBeam() {
    if (!game) return;
}

function AddBeam() {

    if (!game) return;

    laser = {
        xStart: 0,
        yStart: 0,
        xEnd: 0,
        yEnd: 0,
        color: '0x000000',
        twoClick: false,
        in: game.add.graphics(0, 0, game.floorOverObjectLayer),
        out: game.add.graphics(0, 0, game.floorOverObjectLayer)
    };

    if (document.getElementById("rotateBlock")) document.getElementById("rotateBlock").remove();
    let rotate = document.createElement("div");
    rotate.style.height = "50px";
    rotate.id = "rotateBlock";

    let color = document.createElement("input");
    color.type = "color";
    color.value = "#000000";
    color.oninput = function () {
        laser.color = '0x' + this.value.split('#')[1];
        laser.in.clear();
        laser.out.clear();
        CreateBeamLaser(laser.xStart, laser.yStart, laser.xEnd, laser.yEnd, laser.color, laser.in, laser.out)
    };

    let apply = document.createElement("input");
    apply.value = "Применить";
    apply.type = "submit";
    apply.onclick = function () {
        for (let i = 0; i < lasers.length; i++) {
            if (!lasers[i].twoClick) continue;

            mapEditor.send(JSON.stringify({
                event: "addBeam",
                id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
                x: Number(lasers[i].xStart),
                y: Number(lasers[i].yStart),
                to_x: Number(lasers[i].xEnd),
                to_y: Number(lasers[i].yEnd),
                color: lasers[i].color,
            }));
        }

        mapEditor.send(JSON.stringify({
            event: "SelectMap",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
        }));

        lasers = [];
        laser = {xStart: 0, yStart: 0, xEnd: 0, yEnd: 0, color: '0x000000', twoClick: false};
    };

    let cancel = document.createElement("input");
    cancel.value = "Отмена";
    cancel.type = "submit";
    cancel.onclick = function () {
        lasers = [];
        laser = {xStart: 0, yStart: 0, xEnd: 0, yEnd: 0, color: '0x000000', twoClick: false};
        mapEditor.send(JSON.stringify({
            event: "SelectMap",
            id: Number(document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value),
        }));
    };

    rotate.appendChild(color);
    rotate.appendChild(cancel);
    rotate.appendChild(apply);
    document.getElementById("coordinates").appendChild(rotate);

    function move(e) {
        if (laser.twoClick) {
            laser.xEnd = e.worldX / game.camera.scale.x;
            laser.yEnd = e.worldY / game.camera.scale.y;

            laser.in.clear();
            laser.out.clear();
            CreateBeamLaser(laser.xStart, laser.yStart, laser.xEnd, laser.yEnd, laser.color, laser.in, laser.out)
        }
    }

    game.input.onDown.add(placeLaser, game);
    game.input.addMoveCallback(move, this);
}

function placeLaser(e) {
    if (game.input.activePointer.leftButton.isDown) {
        if (!laser.twoClick) {
            laser.twoClick = true;
            laser.xStart = e.worldX / game.camera.scale.x;
            laser.yStart = e.worldY / game.camera.scale.y;
        } else {
            lasers.push(laser);

            laser = {
                xStart: 0,
                yStart: 0,
                xEnd: 0,
                yEnd: 0,
                color: '0x000000',
                twoClick: false,
                in: game.add.graphics(0, 0, game.floorOverObjectLayer),
                out: game.add.graphics(0, 0, game.floorOverObjectLayer)
            };
        }
    }
}

