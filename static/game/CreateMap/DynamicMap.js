function dynamicMap(group, points) {
    points.forEach(function (point) {

        let distCam;
        if (game.camera.view.width > game.camera.view.height) {
            distCam = (game.camera.view.width / 2 + 300) / game.camera.scale.x;
        } else {
            distCam = (game.camera.view.height / 2 + 300) / game.camera.scale.y;
        }

        let camX = (game.camera.view.width / 2 + game.camera.view.x) / game.camera.scale.x;
        let camY = (game.camera.view.height / 2 + game.camera.view.y) / game.camera.scale.y;

        let dist = Phaser.Math.distance(point.x, point.y, camX, camY);
        if (dist < distCam) {

            let coordinate = game.map.OneLayerMap[point.q][point.r];

            if (coordinate.fogSprite && coordinate.open) {
                game.add.tween(coordinate.fogSprite).to({alpha: 0}, 100, Phaser.Easing.Linear.None, true, 0);
            }

            // && !coordinate.base что бы не уничтожались значимые спрайты с их ивентами)
            if (coordinate.texture_object !== "" && !coordinate.base && !coordinate.objectSprite) {
                CreateObjects(coordinate, point.x, point.y);
            }

            if (coordinate.animate_sprite_sheets !== "" && !coordinate.objectSprite) {
                CreateAnimate(coordinate, point.x, point.y);
            }

            if (coordinate.dynamic_object) {
                CreateDynamicObjects(coordinate.dynamic_object, point.q, point.r, true, coordinate);
            }

            if (coordinate.effects != null && coordinate.effects.length > 0) {
                MarkZoneEffect(coordinate, point.x, point.y);
            }

        } else if (dist > distCam && game.map.OneLayerMap[point.q][point.r].sprite) {

            let coordinate = game.map.OneLayerMap[point.q][point.r];

            if (coordinate.objectSprite && !coordinate.base) {
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


            if (coordinate.dynamicObjects && coordinate.dynamicObjects.length > 0) {
                for (let i in coordinate.dynamicObjects) {
                    if (!coordinate.dynamicObjects[i]) {
                        continue
                    }

                    if (coordinate.dynamicObjects[i].object) {
                        coordinate.dynamicObjects[i].object.destroy();
                    }
                    if (coordinate.dynamicObjects[i].background) {
                        coordinate.dynamicObjects[i].background.destroy();
                    }
                    if (coordinate.dynamicObjects[i].shadow) {
                        coordinate.dynamicObjects[i].shadow.destroy();
                    }
                    coordinate.dynamicObjects[i] = null
                }
            }
        }
    });
}