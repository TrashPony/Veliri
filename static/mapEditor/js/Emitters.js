function CreateEmittersZone(emitters) {
    if (!game.emittersDebagZone) game.emittersDebagZone = game.add.graphics(0, 0);

    game.emittersDebagZone.clear();

    for (let i = 0; i < emitters.length; i++) {
        game.emittersDebagZone.lineStyle(2, 0x00ff07, 0.6);
        game.emittersDebagZone.beginFill(0x00ff07, 0.1);
        game.emittersDebagZone.drawRect(emitters[i].x - emitters[i].width / 2, emitters[i].y - emitters[i].height / 2, emitters[i].width, emitters[i].height);
        game.emittersDebagZone.endFill();
    }
}

function RemoveEmitter() {

}

function ChangeEmitter() {

}

function AddEmitter() {

}