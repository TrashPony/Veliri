function CreateMap() {
    return new Promise((resolve) => {
        let verticalOffset = game.hexagonHeight * 3 / 4;
        let horizontalOffset = game.hexagonWidth;
        let startX;
        let startY;
        let startXInit = game.hexagonWidth / 2;
        let startYInit = game.hexagonHeight / 2;

        game.mapPoints = []; // карта точек координат для динамического обнавления карты в методе Update

        game.bmdTerrain.clear();

        for (let r = 0; r < game.map.RSize; r++) {

            if (r % 2 !== 0) {
                startX = 2 * startXInit;
            } else {
                startX = startXInit;
            }

            startY = startYInit + (r * verticalOffset);

            for (let q = 0; q < game.map.QSize; q++) {
                let coordinate = game.map.OneLayerMap[q][r];

                CreateTerrain(coordinate, startX, startY, q, r);

                if (coordinate.texture_object !== "") {
                    CreateObjects(coordinate, startX, startY);
                }

                if (coordinate.animate_sprite_sheets !== "") {
                    CreateAnimate(coordinate, startX, startY);
                }

                if (coordinate.dynamic_object) {
                    CreateDynamicObjects(coordinate.dynamic_object, q, r, true, coordinate);
                }

                if (coordinate.effects != null && coordinate.effects.length > 0) {
                    MarkZoneEffect(coordinate, startX, startY);
                }

                game.mapPoints.push({
                    x: startX,
                    y: startY,
                    q: q,
                    r: r,
                    coordinate: coordinate
                }); // x y - пиксельная координата положения, q r гексовая сеть
                startX += horizontalOffset;

                CreateFowOfWar(coordinate, startX, startY);
            }
        }

        game.fogOfWar.add(game.add.sprite(-50, 0, game.bmdFogOfWar));

        CreateTexture();
        CreateBeams();
        // TODO CreateEmitters();

        resolve()
    });
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

function CreateTexture() {
    // сортировка по приоритету отрисовки текстур
    game.mapPoints.sort(function (a, b) {
        return a.coordinate.texture_priority - b.coordinate.texture_priority;
    });

    for (let i in game.mapPoints) {
        if (game.mapPoints[i].coordinate.texture_over_flore !== '') {
            let bmd = game.make.bitmapData(512, 512);
            bmd.alphaMask(game.mapPoints[i].coordinate.texture_over_flore, 'brush');
            game.bmdTerrain.draw(bmd, game.mapPoints[i].x - 256, game.mapPoints[i].y - 256);
            bmd.destroy();
        }
    }
}