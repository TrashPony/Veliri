function initMove(e) {
    //game.squad.toMove = {x: e.worldX, y: e.worldY};
    //if (game.squad.tweenTo) {
    //    game.squad.tweenTo.stop();
    //    game.squad.tweenTo = null;
    //}

    global.send(JSON.stringify({
        event: "MoveTo",
        to_x: e.worldX,
        to_y: e.worldY
    }));
}