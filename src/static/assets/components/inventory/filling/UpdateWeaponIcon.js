function UpdateWeaponIcon(parent, className, slotData) {
    let weaponIcon = document.createElement("div");
    weaponIcon.className = className;

    for (let i in slotData.unit.body.weapons) {
        if (slotData.unit.body.weapons.hasOwnProperty(i) && slotData.unit.body.weapons[i].weapon) {
            weaponIcon.style.backgroundImage = "url(/assets/units/weapon/" + slotData.unit.body.weapons[i].weapon.name + ".png)"
        }
    }
    parent.appendChild(weaponIcon);
}