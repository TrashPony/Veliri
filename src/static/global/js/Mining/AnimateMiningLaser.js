function AnimateMiningLaser() {

    function animateLaser(squad, miningLaser) {
        miningLaser.out.clear();
        miningLaser.out.lineStyle(6, 0x10EDFF, 1);
        miningLaser.out.moveTo(squad.sprite.x, squad.sprite.y);
        miningLaser.out.lineTo(miningLaser.xy.x, miningLaser.xy.y);

        miningLaser.in.clear();
        miningLaser.in.lineStyle(2, 0xFFFFFF, 1);
        miningLaser.in.moveTo(squad.sprite.x, squad.sprite.y);
        miningLaser.in.lineTo(miningLaser.xy.x, miningLaser.xy.y);
    }

    if (game.squad.miningLaser && game.squad.miningLaser.length > 0) {
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
}