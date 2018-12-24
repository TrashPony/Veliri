function CreateMiniMap(map) {
    let canvas = document.getElementById("canvasMap");
    if (canvas) {
        let ctx = canvas.getContext("2d");
        ctx.clearRect(0, 0, canvas.width, canvas.height);

        let hexagonHeight = (canvas.offsetWidth / game.map.QSize) * 0.95;
        let hexagonWidth = (canvas.offsetHeight / game.map.RSize) * 1.43;

        let verticalOffset = hexagonHeight * 3 / 4;
        let horizontalOffset = hexagonWidth;
        let startX;
        let startY;
        let startXInit = hexagonWidth / 2;
        let startYInit = hexagonHeight / 2;

        let kX = game.hexagonWidth * game.camera.scale.x / hexagonWidth;
        let kY = game.hexagonHeight * game.camera.scale.y / hexagonHeight;

        canvas.onmousedown = function (e) {
            fastMove(e, canvas, hexagonWidth, hexagonHeight)
        };

        let mapPoints = [];
        for (let r = 0; r < game.map.RSize; r++) {

            if (r % 2 !== 0) {
                startX = 2 * startXInit;
            } else {
                startX = startXInit;
            }

            startY = startYInit + (r * verticalOffset);

            for (let q = 0; q < game.map.QSize; q++) {
                mapPoints.push({x: startX, y: startY, q: q, r: r, move: game.map.OneLayerMap[q][r].move});
                startX += horizontalOffset;
            }
        }

        for (let i = 0; i < mapPoints.length; i++) {
            if (mapPoints[i].move) {
                ctx.fillStyle = "#7f8189";
            } else {
                ctx.fillStyle = "#000000";
            }
            ctx.fillRect(mapPoints[i].x, mapPoints[i].y, hexagonWidth, hexagonHeight);
        }

        if (game.squad) {
            ctx.fillStyle = "#19ff00";
            ctx.fillRect(game.squad.sprite.x / kX, game.squad.sprite.y / kY, hexagonWidth, hexagonHeight)
        }

        ctx.strokeStyle = "#fffc1f";
        ctx.strokeRect(game.camera.x / kX, game.camera.y / kY, game.camera.view.width / kX, game.camera.view.height / kY);
    }
}

function fastMove(e, canvas) {
    // TODO ошибка в расчетах, весь метод работает через жопу)

    let x;
    let y;
    if (e.pageX || e.pageY) {
        x = e.pageX;
        y = e.pageY;
    }
    else {
        x = e.clientX + document.body.scrollLeft + document.documentElement.scrollLeft;
        y = e.clientY + document.body.scrollTop + document.documentElement.scrollTop;
    }

    x -= canvas.offsetLeft;
    y -= canvas.offsetTop;

    let q = Math.round(x / (canvas.offsetWidth / game.map.QSize));
    let r = Math.round(y / (canvas.offsetHeight / game.map.RSize));

    let coordinate = game.map.OneLayerMap[q][r];
    if (coordinate) {
        let x, y;

        if (coordinate.r % 2 !== 0) {
            x = game.hexagonWidth + (game.hexagonWidth * coordinate.q)
        } else {
            x = game.hexagonWidth / 2 + (game.hexagonWidth * coordinate.q)
        }
        y = game.hexagonHeight / 2 + (coordinate.r * (game.hexagonHeight * 3 / 4));
        game.camera.focusOnXY(x, y - game.camera.view.height / 2);
        CreateMiniMap(game.map)
    }
}