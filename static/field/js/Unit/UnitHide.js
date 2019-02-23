function UnitHide(unit) {
    delete game.units[unit.q][unit.r];

    if (unit.sprite.alpha === 0 || unit.hp <= 0) {
        return
    }

    unit = unit.sprite;

    let tween = game.add.tween(unit).to({alpha: 0}, 500, Phaser.Easing.Linear.None, true, 1000);
    // функция выполняемая после завершение tween таймера в данном случае удаление спрайта анимации //
    tween.onComplete.add(function (unit) {
        unit.destroy();
    }, unit);
}