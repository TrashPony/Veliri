function update() {

    dynamicMap(game.floorLayer, game.mapPoints);

    if (game && game.typeService === "battle") {
        UpdateRotateUnit(); // функция для повора юнитовский спрайтов
        MoveUnit();
        FlightBullet(); // ослеживает все летящие спрайты пуль
    }

    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    game.floorObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
}

// метод создает обьекты карты в пределах видимости камеры, и удаляет их за ее пределами, серьездно поднимает фпс
function dynamicMap(group, points) {
    points.forEach(function (point) {

        let distCam = 850/game.camera.scale.x;

        let camX = (game.camera.view.width / 2 + game.camera.view.x) / game.camera.scale.x;
        let camY = (game.camera.view.height / 2 + game.camera.view.y) / game.camera.scale.y;

        let dist = Phaser.Math.distance(point.x,point.y, camX, camY);
        if (dist < distCam && !game.map.OneLayerMap[point.q][point.r].sprite) {

            let coordinate = game.map.OneLayerMap[point.q][point.r];

            CreateTerrain(coordinate, point.x, point.y, point.q, point.r);

            if (coordinate.texture_object !== "") {
                CreateObjects(coordinate);
            }

            if (coordinate.animate_sprite_sheets !== "") {
                CreateAnimate(coordinate);
            }

            if (coordinate.effects != null && coordinate.effects.length > 0) {
                MarkZoneEffect(coordinate)
            }

        } else if(dist > distCam && game.map.OneLayerMap[point.q][point.r].sprite) {

            let coordinate = game.map.OneLayerMap[point.q][point.r];

            if (coordinate.objectSprite) {
                if (coordinate.objectSprite.shadow) {
                    coordinate.objectSprite.shadow.destroy()
                }
                coordinate.objectSprite.destroy();

                coordinate.objectSprite = null
            }

            if (coordinate.coordinateText) {
                for (let i in coordinate.coordinateText) {
                    if (coordinate.coordinateText.hasOwnProperty(i)){
                        coordinate.coordinateText[i].destroy();
                    }
                }

                coordinate.coordinateText = null
            }

            coordinate.sprite.destroy();
            coordinate.sprite = null;
        }
    });
}