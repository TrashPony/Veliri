let flagMiniMap = true;

function CreateMiniMap() {

    // отпмизация мини карты что бы не ресовалось чаще чем раз в 100 мс
    if (!flagMiniMap) {
        return
    }
    flagMiniMap = false;
    setTimeout(function () {
        flagMiniMap = true;
    }, 100);


    let canvas = document.getElementById("canvasMap");

    if (!game.map) return;

    if (canvas) {
        let ctx = canvas.getContext("2d");
        ctx.clearRect(0, 0, canvas.width, canvas.height);

        let hexagonHeight = (canvas.offsetWidth / game.map.QSize) * 0.90;
        let hexagonWidth = (canvas.offsetHeight / game.map.RSize) * 1.59;

        let kX = game.hexagonWidth / hexagonWidth;
        let kY = game.hexagonHeight / hexagonHeight;

        canvas.onmousedown = function (e) {
            fastMove(e, canvas, hexagonWidth, hexagonHeight)
        };

        for (let i in game.mapPoints) {
            if (game.mapPoints[i].coordinate.move) {
                ctx.fillStyle = "#7f8189";
            } else {
                ctx.fillStyle = "#000000";
            }
            ctx.fillRect(game.mapPoints[i].x / kX, game.mapPoints[i].y / kY, hexagonWidth, hexagonHeight);

            if (game.mapPoints[i].fogOfWar && game.typeService === "battle") {
                ctx.fillStyle = "#4e4e4e";
                ctx.fillRect(game.mapPoints[i].x / kX, game.mapPoints[i].y / kY, hexagonWidth, hexagonHeight);
            }
        }

        for (let id in game.units) {
            if (game.units[id].sprite) {
                if (game.units[id].owner === game.user_name) {
                    ctx.fillStyle = "#19ff00";

                    if (game.units[id].moveTo) {
                        ctx.beginPath();
                        ctx.strokeStyle = "#00fcff";
                        ctx.moveTo(game.units[id].sprite.x / kX + hexagonWidth / 2, game.units[id].sprite.y / kY + hexagonHeight / 2);
                        ctx.lineTo(game.units[id].moveTo.x / kX, game.units[id].moveTo.y / kY);
                        ctx.stroke();
                    }

                } else {
                    // TODO союзные юниты  ctx.fillStyle = "#00F7FF"
                    ctx.fillStyle = "#ff000e";
                }

                ctx.fillRect(game.units[id].sprite.x / kX, game.units[id].sprite.y / kY, hexagonWidth, hexagonHeight);
            }
        }

        // todo
        if (game.squad) {
            if (game.squad.missionMove) {
                ctx.beginPath();
                ctx.strokeStyle = "#00ff03";
                ctx.moveTo(game.squad.sprite.x / kX + hexagonWidth / 2, game.squad.sprite.y / kY + hexagonHeight / 2);
                ctx.lineTo(game.squad.missionMove.x / kX, game.squad.missionMove.y / kY);
                ctx.stroke();

                ctx.beginPath();
                ctx.strokeStyle = "rgba(0, 255, 0, 0.5)";
                ctx.fillStyle = "rgba(0, 255, 0, 0.1)";
                ctx.ellipse(game.squad.missionMove.x / kX + hexagonWidth / 2, game.squad.missionMove.y / kY + hexagonHeight / 2,
                    game.squad.missionMove.radius / kX, game.squad.missionMove.radius / kY,
                    0, 0, 2 * Math.PI, true);
                ctx.fill();
                ctx.stroke();
            }

            ctx.fillStyle = "#19ff00";
            ctx.fillRect(game.squad.sprite.x / kX, game.squad.sprite.y / kY, hexagonWidth, hexagonHeight);
        }

        if (game.otherUsers) {
            ctx.fillStyle = "#ff7a00";
            for (let i = 0; i < game.otherUsers.length; i++) {
                if (game.otherUsers[i].sprite) {
                    ctx.fillRect(game.otherUsers[i].sprite.x / kX, game.otherUsers[i].sprite.y / kY,
                        hexagonWidth, hexagonHeight)
                }
            }
        }

        for (let i in game.bases) {
            ctx.fillStyle = "#1efcff";
            for (let j in game.bases[i].transports) {
                if (game.bases[i].transports[j].sprite) {
                    ctx.fillRect(game.bases[i].transports[j].sprite.x / kX, game.bases[i].transports[j].sprite.y / kY,
                        hexagonWidth, hexagonHeight)
                }
            }
        }

        for (let i in game.boxes) {
            ctx.fillStyle = "#aba9bc";
            if (game.boxes[i].sprite) {
                ctx.fillRect(game.boxes[i].sprite.x / kX, game.boxes[i].sprite.y / kY, hexagonWidth, hexagonHeight)
            }
        }

        for (let i in game.bases) {
            ctx.fillStyle = "#0babff";
            let xy = GetXYCenterHex(game.bases[i].q, game.bases[i].r);
            ctx.fillRect(xy.x / kX, xy.y / kY, hexagonWidth, hexagonHeight);

            ctx.beginPath();
            ctx.strokeStyle = "rgba(0, 243, 255, 0.5)";
            ctx.fillStyle = "rgba(0, 243, 255, 0.1)";
            ctx.ellipse(xy.x / kX + hexagonWidth / 2, xy.y / kY + hexagonHeight / 2,
                game.bases[i].gravity_radius / kX, game.bases[i].gravity_radius / kY,
                0, 0, 2 * Math.PI, true);
            ctx.fill();
            ctx.stroke();
        }

        let kXCam = game.hexagonWidth * game.camera.scale.x / hexagonWidth;
        let kYCam = game.hexagonHeight * game.camera.scale.y / hexagonHeight;
        ctx.strokeStyle = "#fffc1f";
        ctx.strokeRect(game.camera.x / kXCam, game.camera.y / kYCam, game.camera.view.width / kXCam, game.camera.view.height / kYCam);
    }
}

function fastMove(e, canvas) {
    // TODO ошибка в расчетах, весь метод работает через жопу)

    let x;
    let y;
    if (e.pageX || e.pageY) {
        x = e.pageX;
        y = e.pageY;
    } else {
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