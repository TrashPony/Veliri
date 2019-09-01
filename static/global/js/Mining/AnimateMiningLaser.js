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

    for (let i in game.units) {
        let unit = game.units[i];

        if (unit.selectMiningLine) {
            unit.selectMiningLine.graphics.clear();
            unit.selectMiningLine.graphics.lineStyle(3, 0xb74213, 0.2);
            unit.selectMiningLine.graphics.drawCircle(unit.sprite.x, unit.sprite.y, unit.selectMiningLine.radius);
            unit.selectMiningLine.graphics.lineStyle(1, 0xff0000, 1);
            unit.selectMiningLine.graphics.drawCircle(unit.sprite.x, unit.sprite.y, unit.selectMiningLine.radius);
        }

        if (unit.miningLaser && unit.miningLaser.length > 0) {
            for (let j in unit.miningLaser) {
                if (unit.miningLaser[j]) {
                    animateLaser(unit, unit.miningLaser[j])
                }
            }
        }
    }
}