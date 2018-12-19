function ReadResponse(jsonData) {
    if (jsonData.event === "InitGame") {
        LoadGame(jsonData);
    }

    if (jsonData.event === "Error") {
        alert(jsonData.error);
    }

    if (jsonData.event === "MoveTo") {
        console.log(jsonData)
        game.add.tween(game.squad.sprite).to(
            {x: jsonData.path.x, y: jsonData.path.y},
            jsonData.path.millisecond,
            Phaser.Easing.Linear.None, true, 0
        );

        SetAngle(game.squad, jsonData.path.rotate + 90)
    }
}

function SetMSAngle(unit, angle, time) {
    // TODO эти твины не выбирают минимульный путь поворота Ю_Б
    game.add.tween(unit.sprite.unitBody).to({angle: angle + 90}, time, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(unit.sprite.bodyShadow).to({angle: angle + 90}, time, Phaser.Easing.Linear.None, true, 0);

    if (unit.sprite.weapon) {
        game.add.tween(unit.sprite.weapon).to({angle: angle + 90}, time, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(unit.sprite.weaponShadow).to({angle: angle + 90}, time, Phaser.Easing.Linear.None, true, 0);
    }
}