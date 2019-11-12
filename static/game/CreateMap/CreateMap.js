function CreateMap() {

    game.objects = [];
    game.mapPoints = []; // карта точек координат для динамического обнавления карты в методе Update
    game.bmdTerrain.clear();

    for (let x in game.map.OneLayerMap) {
        if (game.map.OneLayerMap.hasOwnProperty(x)) {
            for (let y in game.map.OneLayerMap[x]) {
                if (game.map.OneLayerMap[x].hasOwnProperty(y)) {

                    let coordinate = game.map.OneLayerMap[x][y];

                    CreateLabels(coordinate, Number(x), Number(y));

                    if (coordinate.effects != null && coordinate.effects.length > 0) {
                        MarkZoneEffect(coordinate, Number(x), Number(y));
                    }

                    game.mapPoints.push({
                        x: Number(x),
                        y: Number(y),
                    });
                }
            }
        }
    }

    CreateFlore().then(function () {
        CreateObjects();
    }).then(function () {
        CreateBeams();
    }).then(function () {
        CreateReservoirs()
    }).then(function () {
        // TODO CreateEmitters();
    });
}

function CreateObjects() {
    for (let x in game.map.static_objects) {
        for (let y in game.map.static_objects[x]) {
            game.objects.push(game.map.static_objects[x][y]);
        }
    }

    // сортировка по приоритету отрисовки обьектов
    game.objects.sort(function (a, b) {
        return a.object_priority - b.object_priority;
    });

    for (let i in game.objects) {
        if (game.objects[i].texture !== '') {
            CreateObject(game.objects[i], game.objects[i].x, game.objects[i].y);
        }

        if (game.objects[i].animate_sprite_sheets !== '') {
            CreateAnimate(game.objects[i], game.objects[i].x, game.objects[i].y);
        }
    }
}

function CreateEmitters() {
    for (let i = 0; i < game.map.emitters.length; i++) {
        CreateEmitter(
            game.map.emitters[i].x,
            game.map.emitters[i].y,
            game.map.emitters[i].min_scale / 100,
            game.map.emitters[i].max_scale / 100,
            game.map.emitters[i].min_speed,
            game.map.emitters[i].max_speed,
            game.map.emitters[i].ttl,
            game.map.emitters[i].width,
            game.map.emitters[i].height,
            game.map.emitters[i].color,
            game.map.emitters[i].frequency,
            game.map.emitters[i].min_alpha / 100,
            game.map.emitters[i].max_alpha / 100,
            game.map.emitters[i].animate_speed,
            game.map.emitters[i].animate,
            game.map.emitters[i].name_particle,
            game.map.emitters[i].alpha_loop_time,
            game.map.emitters[i].yoyo
        );
    }
}

function CreateBeams() {
    for (let i = 0; i < game.map.beams.length; i++) {
        CreateBeamLaser(
            game.map.beams[i].x_start,
            game.map.beams[i].y_start,
            game.map.beams[i].x_end,
            game.map.beams[i].y_end,
            game.map.beams[i].color,
        );
    }
}

function CreateFlore() {
    // сортировка по приоритету отрисовки текстур
    return new Promise(function (resolve) {

        game.flore = [];

        for (let x in game.map.flore) {
            for (let y in game.map.flore[x]) {
                game.flore.push(game.map.flore[x][y]);
            }
        }

        game.flore.sort(function (a, b) {
            return a.texture_priority - b.texture_priority;
        });

        for (let i in game.flore) {
            if (game.flore[i].texture_over_flore !== '') {
                let bmd = game.make.bitmapData(512, 512);
                bmd.alphaMask(game.flore[i].texture_over_flore, 'brush');
                game.bmdTerrain.draw(bmd, game.flore[i].x - 256, game.flore[i].y - 256);
                bmd.destroy();
            }
        }
        resolve();
    });
}