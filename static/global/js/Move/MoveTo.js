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

        game.add.tween(game.squad.sprite).to({
                x: jsonData.path_unit.x,
                y: jsonData.path_unit.y
            }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
        );
        SetMSAngle(game.squad, jsonData.path_unit.rotate, jsonData.path_unit.millisecond);
        game.squad.mather_ship.rotate = jsonData.path_unit.rotate;
    } else {
        MoveOther(jsonData)
    }
}

function MoveOther(jsonData) {
    for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
        if (game.otherUsers[i].squad_id === jsonData.other_user.squad_id) {

            if (!game.otherUsers[i].sprite) {
                console.log("1")
                CreateOtherUser(game.otherUsers[i]);
            }  else  {
                game.add.tween(game.otherUsers[i].sprite).to({
                        x: jsonData.path_unit.x,
                        y: jsonData.path_unit.y
                    }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
                );
                SetMSAngle(game.otherUsers[i], jsonData.path_unit.rotate, jsonData.path_unit.millisecond);
                game.otherUsers[i].rotate = jsonData.path_unit.rotate;
            }
        }
    }
}

function SetMSAngle(unit, angle, time) {
    if (angle > 180) {
        angle -= 360
    }

    ShortDirectionRotateTween(unit.sprite.unitBody, Phaser.Math.degToRad(angle), time);
    ShortDirectionRotateTween(unit.sprite.bodyShadow, Phaser.Math.degToRad(angle), time);
    if (unit.sprite.weapon) {
        ShortDirectionRotateTween(unit.sprite.weaponShadow, Phaser.Math.degToRad(angle), time);
        ShortDirectionRotateTween(unit.sprite.weapon, Phaser.Math.degToRad(angle), time);
    }
}