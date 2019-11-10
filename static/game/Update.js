function update() {
    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */

    // if (game && game.mapPoints) {
    //     // todo идея хорошая реализация нет dynamicMap(game.floorLayer, game.mapPoints);
    //     // todo идея хорошая реализация нет DynamicShadowMap();
    // }

    if (game && game.typeService === "global") {
        StartSelectableUnits();
        ClearFog();

        game.UnitStatusLayer.bmd.clear();

        for (let i in game.units) {
            let unit = game.units[i];
            AnimationMove(unit);

            if (unit && unit.toBox && unit.toBox.to) {
                let dist = game.physics.arcade.distanceToXY(unit.sprite, unit.toBox.x, unit.toBox.y);
                if (dist < 100) {
                    global.send(JSON.stringify({
                        event: "openBox",
                        box_id: unit.toBox.boxID
                    }));
                }
            }

            AnimateMiningLaser(unit);
            AnimateDigger(unit);
            CreateMapHealBar(unit.sprite, unit.body.max_hp, unit.hp);
            UpdateFogOfWar(unit);

            if (unit.AttackLine) {
                CreateAttackLine(unit);
            }
        }

        for (let i in game.boxes) {
            CreateMapHealBar(game.boxes[i].sprite, game.boxes[i].max_hp, game.boxes[i].hp);
        }
    }
}

function CreateAttackLine(unit) {
    // TODO сделать общий метод для отрисовки линий
    unit.AttackLine.graphics.clear();
    unit.AttackLine.graphics.lineStyle(3, 0xb74213, 0.2);
    unit.AttackLine.graphics.drawCircle(unit.sprite.x, unit.sprite.y, unit.AttackLine.minRadius);
    unit.AttackLine.graphics.lineStyle(1, 0xff0000, 1);
    unit.AttackLine.graphics.drawCircle(unit.sprite.x, unit.sprite.y, unit.AttackLine.minRadius);

    unit.AttackLine.graphics.lineStyle(3, 0xb74213, 0.2);
    unit.AttackLine.graphics.drawCircle(unit.sprite.x, unit.sprite.y, unit.AttackLine.maxRadius);
    unit.AttackLine.graphics.lineStyle(1, 0xff0000, 1);
    unit.AttackLine.graphics.drawCircle(unit.sprite.x, unit.sprite.y, unit.AttackLine.maxRadius);
}

function CreateMapHealBar(sprite, maxHP, hp) {

    if (!game.UnitStatusLayer) return;

    let hpInBox = 5;
    let sizeBox = 4;
    let interval = 2; // промеж уток между квадратиками

    let centerX = sprite.x - (game.camera.x / game.camera.scale.x);
    let centerY = sprite.y - (game.camera.y / game.camera.scale.y);

    let countBoxes = Math.ceil(maxHP / hpInBox);
    // для особо жирных
    if (countBoxes > 10) {
        hpInBox = 10;
        countBoxes = Math.ceil(maxHP / hpInBox);
    }

    let startX = Math.round(centerX - ((countBoxes / 2) * (sizeBox + interval)));

    let percentHP = 100 / (maxHP / hp);

    for (let i = 0; i < countBoxes; i++) {

        game.UnitStatusLayer.bmd.ctx.beginPath();
        game.UnitStatusLayer.bmd.ctx.rect(startX, centerY + sprite.offsetY / 1.5, sizeBox, sizeBox);

        if (hp > 0) {
            game.UnitStatusLayer.bmd.ctx.fillStyle = GetColorDamage(percentHP);
        } else {
            game.UnitStatusLayer.bmd.ctx.fillStyle = '#999b9f';
        }

        game.UnitStatusLayer.bmd.ctx.strokeStyle = '#000000';
        game.UnitStatusLayer.bmd.ctx.fill();
        game.UnitStatusLayer.bmd.ctx.stroke();

        hp -= hpInBox;
        startX += sizeBox + interval
    }
}


function ClearFog() {
    game.FogOfWar.bmd.clear();
    game.FogOfWar.bmd.context.fillStyle = 'rgba(0,0,0,0.3)';
    game.FogOfWar.bmd.context.rect(0, 0, game.camera.width, game.camera.height);
    game.FogOfWar.bmd.context.fill();
}

function UpdateFogOfWar(unit) {
    if (!game.FogOfWar || unit.owner_id !== game.user_id) return;

    let centerX = unit.sprite.x - (game.camera.x / game.camera.scale.x);
    let centerY = unit.sprite.y - (game.camera.y / game.camera.scale.y);

    //http://stackoverflow.com/a/12895687/1250044 рисование очищающих кругов
    game.FogOfWar.bmd.context.save();
    game.FogOfWar.bmd.context.globalCompositeOperation = 'destination-out';
    game.FogOfWar.bmd.context.beginPath();
    game.FogOfWar.bmd.context.fillStyle = 'rgba(0,0,0,1)';
    game.FogOfWar.bmd.context.arc(centerX, centerY, unit.body.range_view, 0, Math.PI * 2, true);
    game.FogOfWar.bmd.context.fill();
    game.FogOfWar.bmd.context.restore();
}