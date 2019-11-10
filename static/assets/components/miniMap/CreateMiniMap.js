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

    if ((!game.map) && window.location.pathname !== "/editors/map/") return;

    if (canvas) {

        let ctx = canvas.getContext("2d");
        let offsetX = game.map.XSize / canvas.width;
        let offsetY = game.map.YSize / canvas.height;
        canvas.onmousedown = function (e) {
            fastMove(e, canvas)
        };

        ctx.clearRect(0, 0, canvas.width, canvas.height);
        ctx.fillStyle = "#4e4e4e";
        ctx.fillRect(0, 0, canvas.width, canvas.height);

        // дальность видимости
        for (let id in game.units) {
            if (game.units[id].sprite) {
                if (game.units[id].owner_id === game.user_id) {
                    ctx.beginPath();
                    ctx.strokeStyle = "rgb(127, 127, 127)";
                    ctx.fillStyle = "rgb(127, 127, 127)";
                    ctx.ellipse(game.units[id].sprite.x / offsetX, game.units[id].sprite.y / offsetY,
                        game.units[id].body.range_view / offsetX, game.units[id].body.range_view / offsetY,
                        0, 0, 2 * Math.PI, true);
                    ctx.fill();
                    ctx.stroke();
                }
            }
        }

        // линии радара
        for (let id in game.units) {
            if (game.units[id].sprite) {
                if (game.units[id].owner_id === game.user_id && game.units[id].body.range_radar > 0) {
                    ctx.beginPath();
                    ctx.strokeStyle = "rgba(0, 0, 255, 1)";
                    ctx.ellipse(game.units[id].sprite.x / offsetX, game.units[id].sprite.y / offsetY,
                        game.units[id].body.range_radar / offsetX, game.units[id].body.range_radar / offsetY,
                        0, 0, 2 * Math.PI, true);
                    ctx.stroke();
                }
            }
        }

        // todo довольно тормазнуто отрисовывать всю геодату каждый раз, возможно есть способ это закешировать)
        for (let i in game.map.geo_data) {

            let obstacle = game.map.geo_data[i];

            if (obstacle) {
                ctx.beginPath();
                ctx.fillStyle = "#000000";
                ctx.strokeStyle = "rgba(255, 255, 255, 0)";
                ctx.ellipse(obstacle.x / offsetX, obstacle.y / offsetY,
                    obstacle.radius / offsetX, obstacle.radius / offsetY,
                    0, 0, 2 * Math.PI, true);
                ctx.fill();
                ctx.stroke();
            }
        }

        for (let i in game.map.handlers_coordinates) {
            let outSector = game.map.handlers_coordinates[i];
            if (outSector.handler === "sector") {
                // TODO стрелочка тут будет уместнее
                ctx.beginPath();
                ctx.fillStyle = "rgba(0, 155, 0, 1)";
                ctx.strokeStyle = "#00fff5";
                ctx.ellipse(outSector.x / offsetX, outSector.y / offsetY,
                    60 / offsetX, 60 / offsetY,
                    0, 0, 2 * Math.PI, true);
                ctx.fill();
                ctx.stroke();
            }
        }

        for (let i in game.map.entry_points) {
            for (let j in game.map.entry_points[i].positions) {
                let position = game.map.entry_points[i].positions[j];
                ctx.beginPath();
                ctx.fillStyle = "rgba(155, 155, 0, 1)";
                ctx.strokeStyle = "#00fff9";
                ctx.ellipse(position.x / offsetX, position.y / offsetY,
                    30 / offsetX, 30 / offsetY,
                    0, 0, 2 * Math.PI, true);
                ctx.fill();
                ctx.stroke();
            }
        }

        for (let id in game.units) {
            if (game.units[id].sprite) {
                if (game.units[id].owner_id === game.user_id) {

                    ctx.fillStyle = "#19ff00"; // свои юниты

                    if (game.units[id].moveTo) {
                        ctx.beginPath();
                        ctx.strokeStyle = "#00fcff";
                        ctx.moveTo(game.units[id].sprite.x / offsetX - 1, game.units[id].sprite.y / offsetY - 1);
                        ctx.lineTo(game.units[id].moveTo.x / offsetX - 1, game.units[id].moveTo.y / offsetY - 1);
                        ctx.stroke();
                    }

                } else {
                    // TODO союзные юниты (союзные игроки) ctx.fillStyle = "#00F7FF"
                    //todo враг красны
                    ctx.fillStyle = "#ff7a00"; // нейтрал
                }

                ctx.fillRect(game.units[id].sprite.x / offsetX - 3, game.units[id].sprite.y / offsetY - 2, 6, 4);
            }
        }

        for (let q in game.map.reservoir) {
            for (let r in game.map.reservoir[q]) {
                let reservoir = game.map.reservoir[q][r];

                ctx.beginPath();
                ctx.fillStyle = "#ffd700";
                ctx.strokeStyle = "rgba(255, 255, 255, 0)";
                ctx.ellipse(reservoir.x / offsetX, reservoir.y / offsetY,
                    15 / offsetX, 15 / offsetY,
                    0, 0, 2 * Math.PI, true);
                ctx.fill();
                ctx.stroke();
            }
        }

        for (let i in game.bases) {
            ctx.fillStyle = "#1efcff";
            for (let j in game.bases[i].transports) {
                if (game.bases[i].transports[j].sprite) {
                    ctx.fillRect(game.bases[i].transports[j].sprite.x / offsetX, game.bases[i].transports[j].sprite.y / offsetY,
                        3, 3)
                }
            }
        }

        for (let i in game.boxes) {
            ctx.fillStyle = "#e8e6ff";
            if (game.boxes[i].sprite) {
                ctx.fillRect(game.boxes[i].sprite.x / offsetX, game.boxes[i].sprite.y / offsetY, 6, 3)
            }
        }

        for (let i in game.bases) {

            ctx.beginPath();
            ctx.strokeStyle = "rgba(0, 243, 255, 1)";
            ctx.fillStyle = "rgba(0, 243, 255, 1)";
            ctx.ellipse(game.bases[i].x / offsetX, game.bases[i].y / offsetY,
                4, 2,
                0, 0, 2 * Math.PI, true);
            ctx.fill();
            ctx.stroke();

            ctx.beginPath();
            ctx.strokeStyle = "rgba(0, 243, 255, 0.5)";
            ctx.fillStyle = "rgba(0, 243, 255, 0.1)";
            ctx.ellipse(game.bases[i].x / offsetX, game.bases[i].y / offsetY,
                game.bases[i].gravity_radius / offsetX, game.bases[i].gravity_radius / offsetY,
                0, 0, 2 * Math.PI, true);
            ctx.fill();
            ctx.stroke();
        }

        // // todo
        // if (game.squad) {
        //     if (game.squad.missionMove) {
        //         ctx.beginPath();
        //         ctx.strokeStyle = "#00ff03";
        //         ctx.moveTo(game.squad.sprite.x / kX + hexagonWidth / 2, game.squad.sprite.y / kY + hexagonHeight / 2);
        //         ctx.lineTo(game.squad.missionMove.x / kX, game.squad.missionMove.y / kY);
        //         ctx.stroke();
        //
        //         ctx.beginPath();
        //         ctx.strokeStyle = "rgba(0, 255, 0, 0.5)";
        //         ctx.fillStyle = "rgba(0, 255, 0, 0.1)";
        //         ctx.ellipse(game.squad.missionMove.x / kX + hexagonWidth / 2, game.squad.missionMove.y / kY + hexagonHeight / 2,
        //             game.squad.missionMove.radius / kX, game.squad.missionMove.radius / kY,
        //             0, 0, 2 * Math.PI, true);
        //         ctx.fill();
        //         ctx.stroke();
        //     }
        //
        //     ctx.fillStyle = "#19ff00";
        //     ctx.fillRect(game.squad.sprite.x / kX, game.squad.sprite.y / kY, hexagonWidth, hexagonHeight);
        // }

        let kXCam = (game.camera.scale.x * offsetX);
        let kYCam = (game.camera.scale.y * offsetY);
        ctx.strokeStyle = "#fffc1f";
        ctx.strokeRect(game.camera.x / kXCam, game.camera.y / kYCam, game.camera.view.width / kXCam, game.camera.view.height / kYCam);
    }
}

function fastMove(e, canvas) {
    // TODO неправильный расчет
    let offsetX = game.map.XSize / canvas.width;
    let offsetY = game.map.YSize / canvas.height;


    game.camera.focusOnXY(
        (e.offsetX * offsetX) + (game.camera.scale.x * game.camera.width) / 2,
        (e.offsetY * offsetY) + (game.camera.scale.y * game.camera.height) / 2,
    );
    CreateMiniMap(game.map)
}