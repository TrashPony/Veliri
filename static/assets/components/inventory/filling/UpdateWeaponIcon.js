function UpdateWeaponIcon(parent, className, slotData) {

    let weaponIcon = document.createElement("div");
    weaponIcon.className = className;

    for (let i in slotData.unit.body.weapons) {
        if (slotData.unit.body.weapons.hasOwnProperty(i) && slotData.unit.body.weapons[i].weapon) {
            weaponIcon.style.backgroundImage = "url(/assets/units/weapon/" + slotData.unit.body.weapons[i].weapon.name + ".png)";
            weaponIcon.style.top = (slotData.unit.body.weapons[i].y_attach - slotData.unit.body.weapons[i].weapon.y_attach) / 2 + "px";
            weaponIcon.style.left = (slotData.unit.body.weapons[i].x_attach - slotData.unit.body.weapons[i].weapon.x_attach) / 2 + "px";
        }
    }
    parent.appendChild(weaponIcon);
}