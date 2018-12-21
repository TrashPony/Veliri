// метод создает обьекты карты в пределах видимости камеры, и удаляет их за ее пределами, серьездно поднимает фпс
function dynamicMap(group, points) {
    points.forEach(function (point) {

        let distCam;
        if (game.camera.view.width > game.camera.view.height) {
            distCam = game.camera.view.width / 2 + 300;
        } else {
            distCam = game.camera.view.height / 2 + 300;
        }

        let camX = (game.camera.view.width / 2 + game.camera.view.x) / game.camera.scale.x;
        let camY = (game.camera.view.height / 2 + game.camera.view.y) / game.camera.scale.y;

        let dist = Phaser.Math.distance(point.x, point.y, camX, camY);
        if (dist < distCam && !game.map.OneLayerMap[point.q][point.r].sprite) {

            let coordinate = game.map.OneLayerMap[point.q][point.r];

            CreateTerrain(coordinate, point.x, point.y, point.q, point.r);

            if (coordinate.fogSprite && coordinate.open) {
                game.add.tween(coordinate.fogSprite).to({alpha: 0}, 100, Phaser.Easing.Linear.None, true, 0);
            }

            if (coordinate.texture_object !== "") {
                CreateObjects(coordinate);
            }

            if (coordinate.animate_sprite_sheets !== "") {
                CreateAnimate(coordinate);
            }

            if (coordinate.effects != null && coordinate.effects.length > 0) {
                MarkZoneEffect(coordinate)
            }

        } else if (dist > distCam && game.map.OneLayerMap[point.q][point.r].sprite) {

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
                    if (coordinate.coordinateText.hasOwnProperty(i)) {
                        coordinate.coordinateText[i].destroy();
                    }
                }
                coordinate.coordinateText = null
            }

            if (coordinate.buttons) {
                for (let i in coordinate.buttons) {
                    coordinate.buttons[i].destroy();
                }
                coordinate.coordinateText = null
            }

            if (coordinate.fogSprite) {
                coordinate.fogSprite.destroy();
            }

            coordinate.sprite.destroy();
            coordinate.sprite = null;
        }
    });
}