function UpdateWatchZone(watch) {
    if (watch) {
        let closeCoordinate = watch.close_coordinate;
        let openCoordinate = watch.open_coordinate;
        let openUnits = watch.open_unit;

        if (closeCoordinate) {
            // TODO дожидатся окончания анимации и только после перерисовывать туман
            CloseCoordinates(closeCoordinate);
        }

        if (openCoordinate) {
            OpenCoordinates(openCoordinate);
        }

        if (openUnits) {
            while (openUnits.length > 0) {
                let openUnit = openUnits.shift();
                CreateUnit(openUnit)
            }
        }
    }

    CreateAllFogOfWar();
    CreateMiniMap();
}

function OpenCoordinate(q, r, hide) {
    // TODO требует оптимизации, например отслеживать на все карту а только ту часть где идет бой или камера игрока, иначе фризы при движение
    for (let i in game.mapPoints) {
        if (game.mapPoints[i].q === q && game.mapPoints[i].r === r) {
            game.mapPoints[i].fogOfWar = hide;
            return {x: game.mapPoints[i].x, y: game.mapPoints[i].y}
        }
    }
}

function OpenCoordinates(coordinates) {
    while (coordinates.length > 0) {
        let coordinate = coordinates.shift();
        let xy = OpenCoordinate(coordinate.q, coordinate.r, false);

        // анимация открытия
        let fogSprite = game.fogOfWar.create(xy.x, xy.y, 'FogOfWar');
        fogSprite.anchor.setTo(0.5);
        fogSprite.scale.set(0.25);
        let tween = game.add.tween(fogSprite).to({alpha: 0}, 1500, Phaser.Easing.Linear.None, true, 0);
        tween.onComplete.add(function () {
            fogSprite.destroy();
        });
    }
}

function CloseCoordinates(coordinates) {
    while (coordinates.length > 0) {
        let coordinate = coordinates.shift();
        let xy = OpenCoordinate(coordinate.q, coordinate.r, true);

        // анимация закрытия
        let fogSprite = game.fogOfWar.create(xy.x, xy.y, 'FogOfWar');
        fogSprite.anchor.setTo(0.5);
        fogSprite.scale.set(0.25);
        fogSprite.alpha = 0;
        let tween = game.add.tween(fogSprite).to({alpha: 0.5}, 1500, Phaser.Easing.Linear.None, true, 0);
        tween.onComplete.add(function () {
            fogSprite.destroy();
        });

        // если на координате был юнит то удаляем его ибо скрылся за туманом войны
        let unit = GetGameUnitXY(coordinate.q, coordinate.r);
        if (unit) {
            UnitHide(unit);
        }
    }
}