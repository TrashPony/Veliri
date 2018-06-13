function FillingCellInventory() {
    var cells = document.getElementsByClassName('cellInventory');

    // todo будут проблемы если эквипом больше чем ячеек
    for (var i = 0; i < game.user.equip.length; i++) {
        cells[i].equip = game.user.equip[i];
        cells[i].style.backgroundImage = "url(/assets/" + cells[i].equip.type + ".png)";

        cells[i].onclick = function () {
            RemoveSelect();
            MarkEquipSelect(2, this.equip);
        };

        cells[i].onmouseover = function () {
            TipEquipOn(this.equip);
        };

        cells[i].onmouseout = function () {
            TipEquipOff();
        };
    }
}