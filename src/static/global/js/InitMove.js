function initMove(e) {

    if (game.squad.toBase){
        game.squad.toBase.into = false
    }

    global.send(JSON.stringify({
        event: "MoveTo",
        to_x: e.worldX,
        to_y: e.worldY
    }));
}