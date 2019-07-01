function MoveTo(jsonData) {

    if (!game) return;

    CreateMiniMap();

    if (game.squad && Number(jsonData.other_user.squad_id) === game.squad.id) {
        game.floorSelectLineLayer.forEach(function (sprite) {
            sprite.visible = false;
        });

        let thoriumEfficiency = document.getElementById("speedBarEfficiency");
        thoriumEfficiency.innerHTML = (jsonData.path_unit.Speed * 10).toFixed(0);

        game.squad.q = jsonData.path_unit.q;
        game.squad.r = jsonData.path_unit.r;
        game.squad.speed = jsonData.path_unit.Speed * 10;

        game.add.tween(game.squad.sprite).to({
                x: jsonData.path_unit.x,
                y: jsonData.path_unit.y
            }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
        );

        SetMSAngle(game.squad, jsonData.path_unit.rotate, jsonData.path_unit.millisecond);
        //AnimationMove(game.squad);

        game.squad.mather_ship.rotate = jsonData.path_unit.rotate;
    } else {
        MoveOther(jsonData)
    }
}

function MoveOther(jsonData) {
    for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
        if (game.otherUsers[i].squad_id === jsonData.other_user.squad_id) {

            if (!game.otherUsers[i].sprite) {
                CreateOtherUser(game.otherUsers[i]);
            } else {
                game.add.tween(game.otherUsers[i].sprite).to({
                        x: jsonData.path_unit.x,
                        y: jsonData.path_unit.y
                    }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
                );

                game.otherUsers[i].speed = jsonData.path_unit.Speed * 10;

                SetMSAngle(game.otherUsers[i], jsonData.path_unit.rotate, jsonData.path_unit.millisecond);
                //AnimationMove(game.otherUsers[i]);

                game.otherUsers[i].rotate = jsonData.path_unit.rotate;
            }
        }
    }
}

function SetMSAngle(unit, angle, time) {


    SetShadowAngle(unit, angle);
    if (angle > 180) {
        angle -= 360
    }

    ShortDirectionRotateTween(unit.sprite, Phaser.Math.degToRad(angle), time);
}

function AnimationMove(unit) {


    if (unit.speed && unit.speed > 0) {

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