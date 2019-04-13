function OutFogFire(start, target, weapon, targetType) {

    let targetX;
    let targetY;

    let startX = start.x;
    let startY = start.y;

    if (targetType === "coordinate") {
        targetX = target.x + 50;
        targetY = target.y + 40;
    } else {
        targetX = target.sprite.x;
        targetY = target.sprite.y;
    }

    let rotate = Math.atan2(targetY - startY, targetX - startX);
    let angle = (rotate * 180 / 3.14) + 90;

    return new Promise((resolve) => {

        if (weapon.type === "missile") {
            let connectPoints = PositionAttachSprite(angle, 60);

            launchRocket(startX - connectPoints.x / 2, startY - connectPoints.y / 2, angle, targetX, targetY, weapon.artillery, "outFog");
            setTimeout(function () {
                launchRocket(startX + connectPoints.x / 2, startY + connectPoints.y / 2, angle, targetX, targetY, weapon.artillery, "outFog")
                    .then(function () {
                        resolve();
                    });
            }, 500);
        }

        if (weapon.type === "firearms") {
            if (weapon.artillery) {
                let connectPointsOne = PositionAttachSprite(angle - 85, 60);
                let connectPointsTwo = PositionAttachSprite(angle - 95, 60);

                LaunchArtilleryBallistics(startX + connectPointsOne.x, startY + connectPointsOne.y, angle, targetX, targetY, "outFog");
                setTimeout(function () {
                    LaunchArtilleryBallistics(startX + connectPointsTwo.x, startY + connectPointsTwo.y, angle, targetX - 9, targetY - 7, "outFog")
                        .then(function () {
                            resolve();
                        });
                }, 500);
            } else {
                let connectPoints = PositionAttachSprite(angle - 90, 60 / 1.1);

                LaunchSmallBallistics(startX + connectPoints.x, startY + connectPoints.y, angle, targetX, targetY, "outFog")
                    .then(function () {
                        resolve();
                    });
            }
        }

        if (weapon.type === "laser") {
            if (weapon.name === "big_laser") {
                let connectPointsOne = PositionAttachSprite(angle - 85, 60 / 1.5);
                let connectPointsTwo = PositionAttachSprite(angle - 95, 60 / 1.5);

                LaunchLaser(startX + connectPointsOne.x, startY + connectPointsOne.y, angle, targetX, targetY, "outFog");
                LaunchLaser(startX + connectPointsTwo.x, startY + connectPointsTwo.y, angle, targetX, targetY, "outFog")
                    .then(function () {
                        resolve();
                    })
            }
            if (weapon.name === "small_laser") {
                let connectPoints = PositionAttachSprite(angle - 90, 60 / 1.5);
                LaunchLaser(startX + connectPoints.x, startY + connectPoints.y, angle, targetX, targetY, "outFog")
                    .then(function () {
                        resolve();
                    })
            }
        }
    });
}