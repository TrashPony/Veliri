function CreateCloud(jsonData) {
    if (game && game.cloudsLayer && jsonData) {

        let find = false;
        game && game.cloudsLayer.forEach(function (cloud) {
            if (cloud.uuid === jsonData.cloud.uuid) {
                find = true;

                game.add.tween(cloud).to({
                        x: jsonData.cloud.x,
                        y: jsonData.cloud.y
                    }, 1000, Phaser.Easing.Linear.None, true, 0
                );
            }
        });

        if (!find) {

            let cloud = game.cloudsLayer.create(jsonData.cloud.x, jsonData.cloud.y, jsonData.cloud.name);
            cloud.scale.setTo(0.5);
            game.physics.enable(cloud, Phaser.Physics.ARCADE);

            cloud.alpha = jsonData.cloud.alpha;
            cloud.angle = jsonData.cloud.angle;

            setTimeout(function () {
                // что бы механизм удаление включатся когда облако уже в мире, иначе будут лаги на границе
                cloud.checkWorldBounds = true;
                cloud.events.onOutOfBounds.add(function () {
                    cloud.destroy()
                });
            }, 120000);


            cloud.uuid = jsonData.cloud.uuid;

            game.add.tween(cloud).to({
                    x: jsonData.cloud.x,
                    y: jsonData.cloud.y
                }, 1000, Phaser.Easing.Linear.None, true, 0
            );
        }
    }
}