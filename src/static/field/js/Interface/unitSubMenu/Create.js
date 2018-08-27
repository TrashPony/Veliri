function CreateUnitSubMenu(unit) {
    if (game.Phase === "move" || game.Phase === "targeting") {

        let unitSubMenu = document.getElementById("UnitSubMenu");

        if (unitSubMenu) {
            unitSubMenu.remove();
        }

        unitSubMenu = document.createElement("div");
        unitSubMenu.id = "UnitSubMenu";

        unitSubMenu.style.left = stylePositionParams.left - 95 + 'px';
        unitSubMenu.style.top = stylePositionParams.top - 85 + 'px';
        unitSubMenu.style.display = "block";

        let equipPanel = document.createElement("div");
        equipPanel.id = "EquipPanel";
        unitSubMenu.appendChild(equipPanel);

        if (!(unit.action && unit.use_equip) && game.user.name === unit.owner) {

            FillingEquipPanel(equipPanel, unit);

            if (game.Phase === "move") {
                ActionButton(equipPanel, unit, "SkipMoveUnit", "Пропустить ход");
            }

            if (game.Phase === "targeting") {
                ActionButton(equipPanel, unit, "Defend", "Защита");
            }
        } else {
            unitSubMenu.style.animation = "none";
            unitSubMenu.style.border = "0px";
        }

        if (unit.effects !== null && unit.effects.length > 0) {
            EffectsPanel(unitSubMenu, unit);
        }

        document.body.appendChild(unitSubMenu);
    }
}