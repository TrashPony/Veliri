// движение на глобальной карте
function MoveTo(jsonData) {

    console.log(jsonData.path_unit)
    if (!game || !game.units) return;

    let unit = game.units[jsonData.short_unit.id];
    let path = jsonData.path_unit;

    if (unit) {

        if (unit.owner === game.user_name && unit.body.mother_ship) {
            let thoriumEfficiency = document.getElementById("speedBarEfficiency");
            thoriumEfficiency.innerHTML = (path.Speed * 10).toFixed(0);
        }

        unit.moveTween = game.add.tween(unit.sprite).to({
                x: path.x,
                y: path.y
            }, path.millisecond, Phaser.Easing.Linear.None, true, 0
        );

        unit.moveTween.onComplete.add(function () {
            unit.moveTween = null;
        });

        unit.speed = path.Speed * 10;
        unit.animateSpeed = path.animate;
        unit.rotate = path.rotate;

        SetAngle(unit, path.rotate, path.millisecond, true);
        CreateMiniMap();
    } else {
        CreateNewUnit(unit)
    }
}

function MoveStop(jsonData) {
    if (!game || !game.units) return;

    let unit = game.units[jsonData.short_unit.id];
    let path = jsonData.path_unit;

    if (unit) {
        if (unit.moveTween) {
            unit.moveTween.onComplete.add(function () {
                unit.speed = path.Speed * 10;
            });
        } else {
            unit.speed = path.Speed * 10;
        }
    }
}

function AnimationMove(unit) {
    if (unit.speed && unit.speed > 0 && unit.animateSpeed) {

        if (unit.speed < 30) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 70
        } else if (unit.speed >= 30 && unit.speed < 40) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 65
        } else if (unit.speed >= 40 && unit.speed < 50) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 60
        } else if (unit.speed >= 50 && unit.speed < 60) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 55
        } else if (unit.speed >= 60 && unit.speed < 70) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 50
        } else if (unit.speed >= 70 && unit.speed < 80) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 45
        } else if (unit.speed >= 80 && unit.speed < 90) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 40
        } else if (unit.speed >= 90 && unit.speed < 100) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 35
        } else if (unit.speed >= 100 && unit.speed < 110) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 30
        } else if (unit.speed >= 110 && unit.speed < 120) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 25
        } else if (unit.speed >= 120 && unit.speed < 130) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 20
        } else if (unit.speed >= 130 && unit.speed < 140) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 15
        } else if (unit.speed >= 140) {
            unit.sprite.bodyBottom.animations.getAnimation('move').delay = 10
        }

        if (unit.sprite.bodyBottom.animations.getAnimation('move').isPaused) {
            unit.sprite.bodyBottom.animations.getAnimation('move').paused = false;
            unit.sprite.bodyBottomShadow.animations.getAnimation('move').paused = false;
        }
    } else {
        unit.sprite.bodyBottom.animations.getAnimation('move').paused = true;
        unit.sprite.bodyBottomShadow.animations.getAnimation('move').paused = true;
    }
}