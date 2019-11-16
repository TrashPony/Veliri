function CreateGeoData() {
    for (let i = 0; i < game.map.geo_data.length; i++) {
        if (game.map.geo_data[i]) {

            let centerX = game.map.geo_data[i].x - (game.camera.x / game.camera.scale.x);
            let centerY = game.map.geo_data[i].y - (game.camera.y / game.camera.scale.y);

            game.geoData.bmd.context.beginPath();
            game.geoData.bmd.context.arc(centerX, centerY, game.map.geo_data[i].radius, 0, Math.PI * 2, true);
            game.geoData.bmd.context.fill();
        }
    }
}

function CreateDynamicObjGeo() {
    for (let i in game.objects) {
        let obj = game.objects[i];
        if (obj && obj.geo_data && obj.geo_data.length > 0) {
            for (let i = 0; i < obj.geo_data.length; i++) {

                let centerX = obj.geo_data[i].x - (game.camera.x / game.camera.scale.x);
                let centerY = obj.geo_data[i].y - (game.camera.y / game.camera.scale.y);

                game.geoData.bmd.context.beginPath();
                game.geoData.bmd.context.arc(centerX, centerY, obj.geo_data[i].radius, 0, Math.PI * 2, true);
                game.geoData.bmd.context.fill();
            }
        }
    }
}

function DrawGeoData() {

    if (game.geoData) {
        game.geoData.bmd.clear();
    } else {
        let bmd = game.make.bitmapData(game.camera.width, game.camera.height);
        let dmbSprite = bmd.addToWorld();
        dmbSprite.fixedToCamera = true;
        game.geoData = {
            bmd: bmd,
            sprite: dmbSprite,
        };
    }

    game.geoData.bmd.context.fillStyle = 'rgba(255,0,0,0.3)';
    CreateGeoData();
    CreateDynamicObjGeo();
    game.geoData.bmd.context.fill();
}