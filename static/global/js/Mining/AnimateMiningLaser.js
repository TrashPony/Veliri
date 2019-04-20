function AnimateMiningLaser() {

    function animateLaser(squad, miningLaser) {

        let targetX = (miningLaser.equipSprite.world.x / game.camera.scale.x - miningLaser.xy.x) * -1;
        let targetY = (miningLaser.equipSprite.world.y / game.camera.scale.y - miningLaser.xy.y) * -1;

        let angle = Phaser.Math.angleBetween(miningLaser.equipSprite.world.x / game.camera.scale.x, miningLaser.equipSprite.world.y / game.camera.scale.y, miningLaser.xy.x, miningLaser.xy.y) * Phaser.Math.RAD_TO_DEG;

        miningLaser.attachPoint.angle = angle - squad.sprite.angle;
        miningLaser.equipSprite.angle = angle - squad.sprite.angle;

        miningLaser.out.angle = (360 - squad.sprite.angle) - miningLaser.attachPoint.angle;
        miningLaser.in.angle = (360 - squad.sprite.angle) - miningLaser.attachPoint.angle;

        miningLaser.out.clear();
        miningLaser.out.lineStyle(9, 0x10EDFF, 1);
        miningLaser.out.moveTo(0, 0);
        miningLaser.out.lineTo(targetX * 4, targetY * 4);

        miningLaser.in.clear();
        miningLaser.in.lineStyle(3, 0xFFFFFF, 1);
        miningLaser.in.moveTo(0, 0);
        miningLaser.in.lineTo(targetX * 4, targetY * 4);
    }

    if (game.squad && game.squad.miningLaser && game.squad.miningLaser.length > 0) {
        for (let i in game.squad.miningLaser) {
            if (game.squad.miningLaser[i]) {

                //game.squad.miningLaser[i].equipSprite
                //console.log(game.squad.miningLaser[i].equipSprite);
                animateLaser(game.squad, game.squad.miningLaser[i])
            }
        }
    }

    for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
        if (game.otherUsers[i].miningLaser && game.otherUsers[i].miningLaser.length > 0) {
            for (let j in game.otherUsers[i].miningLaser) {
                if (game.otherUsers[i].miningLaser[j]) {
                    animateLaser(game.otherUsers[i], game.otherUsers[i].miningLaser[j])
                }
            }
        }
    }

    if (game.squad && game.squad.selectMiningLine) {
        game.squad.selectMiningLine.graphics.clear();
        game.squad.selectMiningLine.graphics.lineStyle(3, 0xb74213, 0.2);
        game.squad.selectMiningLine.graphics.drawCircle(game.squad.sprite.x, game.squad.sprite.y, game.squad.selectMiningLine.radius);
        game.squad.selectMiningLine.graphics.lineStyle(1, 0xff0000, 1);
        game.squad.selectMiningLine.graphics.drawCircle(game.squad.sprite.x, game.squad.sprite.y, game.squad.selectMiningLine.radius);
    }
}