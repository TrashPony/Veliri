function CreateUnitSubMenu(unit) {
    if (game.Phase === "move" || game.Phase === "targeting") {

        let BoxUnitSubMenu = document.getElementById("BoxUnitSubMenu");
        let unitSubMenu = document.getElementById("UnitSubMenu");
        BoxUnitSubMenu.style.display = "block";

        if (unitSubMenu) {
            BoxUnitSubMenu.innerHTML = '';
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
                BoxUnitSubMenu.appendChild(unitSubMenu);
                ActionButton(BoxUnitSubMenu, unit, "SkipMoveUnit", "Пропустить ход");
            }

            if (game.Phase === "targeting" && !unit.defend) {
                BoxUnitSubMenu.appendChild(unitSubMenu);
                ActionButton(BoxUnitSubMenu, unit, "Defend", "Защита");
                ActionButton(BoxUnitSubMenu, unit, "initReload", "Перезарядка");
            }

        } else {
            unitSubMenu.style.boxShadow = "none";
            unitSubMenu.style.animation = "none";
            unitSubMenu.style.border = "0px";
            unitSubMenu.style.visibility = "hidden";
            BoxUnitSubMenu.style.display = "none";
        }

        if (unit.effects !== null && unit.effects.length > 0) {
            // TODO EffectsPanel(unitSubMenu, unit);
        }
    }
}