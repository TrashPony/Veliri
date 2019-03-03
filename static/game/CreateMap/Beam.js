function CreateBeamLaser(xStart, yStart, xEnd, yEnd, color, inGraphics, outGraphics) {
    let laserOut, laserIn;

    if (!inGraphics || !outGraphics) {
        laserIn = game.add.graphics(0, 0, game.floorOverObjectLayer);
        laserOut = game.add.graphics(0, 0, game.floorOverObjectLayer);
    } else {
        laserIn = inGraphics;
        laserOut = outGraphics;
    }

    laserOut.lineStyle(6, color, 0.6);
    laserOut.moveTo(xStart, yStart);
    laserOut.lineTo(xEnd, yEnd);

    laserIn.lineStyle(1, 0xFFFFFF, 0.8);
    laserIn.moveTo(xStart, yStart);
    laserIn.lineTo(xEnd, yEnd);

    let blurX = game.add.filter('BlurX');
    let blurY = game.add.filter('BlurY');
    blurX.blur = 5;
    blurY.blur = 5;
    laserOut.filters = [blurX, blurY];

    game.add.tween(laserOut).to({alpha: 0.2}, 2500, "Linear").loop(true).yoyo(true).start();
    game.add.tween(laserIn).to({alpha: 0.4}, 2500, "Linear").loop(true).yoyo(true).start();
}