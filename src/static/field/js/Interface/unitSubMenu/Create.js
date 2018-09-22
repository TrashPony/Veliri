function CreateUnitSubMenu(unit) {
    if (game.Phase === "move" || game.Phase === "targeting") {

        let BoxUnitSubMenu = document.getElementById("BoxUnitSubMenu");
        let unitSubMenu = document.getElementById("UnitSubMenu");

        if (unitSubMenu) {
            unitSubMenu.remove();
        }

        unitSubMenu = document.createElement("div");
        unitSubMenu.id = "UnitSubMenu";

        unitSubMenu.style.display = "block";

        let equipPanel = document.createElement("div");
        equipPanel.id = "EquipPanel";
        unitSubMenu.appendChild(equipPanel);

        if (game.user.name === unit.owner) {

            FillingEquipPanel(equipPanel, unit);

            if (game.Phase === "move") {
                ActionButton(equipPanel, unit, "SkipMoveUnit", "Пропустить ход");
            }

            if (game.Phase === "targeting" && !unit.defend) {
                ActionButton(equipPanel, unit, "Defend", "Защита");
            }
        } else {
            unitSubMenu.style.boxShadow = "none";
            unitSubMenu.style.animation = "none";
            unitSubMenu.style.border = "0px";
        }

        if (unit.effects !== null && unit.effects.length > 0) {
            EffectsPanel(unitSubMenu, unit);
        }

        BoxUnitSubMenu.appendChild(unitSubMenu);
    }
}