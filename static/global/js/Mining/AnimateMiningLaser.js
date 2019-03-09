function AnimateMiningLaser() {

    function animateLaser(squad, miningLaser) {
        miningLaser.out.clear();
        miningLaser.out.lineStyle(3, 0x10EDFF, 1);
        miningLaser.out.moveTo(squad.sprite.x, squad.sprite.y);
        miningLaser.out.lineTo(miningLaser.xy.x, miningLaser.xy.y);

        miningLaser.in.clear();
        miningLaser.in.lineStyle(1, 0xFFFFFF, 1);
        miningLaser.in.moveTo(squad.sprite.x, squad.sprite.y);
        miningLaser.in.lineTo(miningLaser.xy.x, miningLaser.xy.y);
    }

    if (game.squad && game.squad.miningLaser && game.squad.miningLaser.length > 0) {
        for (let i in game.squad.miningLaser) {
            if (game.squad.miningLaser[i]) {
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