function CreateHealBar(cell, type, append) {
    let cellData = JSON.parse(cell.slotData);

    if (cellData.type !== "ammo") {
        let backHealBar = document.createElement("div");

        let percentHP = 0;

        if (type === "inventory") {
            backHealBar.className = "backInventoryHealBar";
            percentHP = 100 / (cellData.item.max_hp / cellData.hp);
        } else if (type === "equip") {
            backHealBar.className = "backEquipHealBar";
            percentHP = 100 / (cellData.equip.max_hp / cellData.hp);
        } else if (type === "weapon") {
            backHealBar.className = "backWeaponHealBar";
            percentHP = 100 / (cellData.weapon.max_hp / cellData.hp);
        } else if (type === "body") {
            backHealBar.className = "backBodyHealBar";
            percentHP = 100 / (cellData.body.max_hp / cellData.hp);
        }

        let healBar = document.createElement("div");
        healBar.className = "healBar";

        healBar.style.width = percentHP + "%";

        if (percentHP === 100) {
            backHealBar.style.opacity = "0"
        } else if (percentHP < 90 && percentHP > 75) {
            healBar.style.backgroundColor = "#fff326"
        } else if (percentHP < 75 && percentHP > 50) {
            healBar.style.backgroundColor = "#fac227"
        } else if (percentHP < 50 && percentHP > 25) {
            healBar.style.backgroundColor = "#fa7b31"
        } else if (percentHP < 25 && cellData.hp > 1) {
            backHealBar.style.opacity = "0";
            // todo показывать что предмет сломан например box-shadow insert red
        }

        if (append) {
            backHealBar.appendChild(healBar);
            cell.appendChild(backHealBar);
        }

        return percentHP;
    }
}