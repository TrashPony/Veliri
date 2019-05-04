function UpdateWeaponIcon(parent, className, slotData, scale) {
    let scaleSize = 1;
    if (scale) {
        scaleSize = 2;
    }

    let weaponIcon = document.createElement("div");
    weaponIcon.className = className;

    let weaponIconMask = document.createElement('div');
    weaponIconMask.className = 'mask';

    let weaponIconMask2 = document.createElement('div');
    weaponIconMask2.className = 'mask';

    if (slotData.unit.body.mother_ship) {
        weaponIconMask.id = "msWeaponMask1";
    } else if (!scale) {
        weaponIconMask.id = "unitWeaponMask1";
    }

    if (slotData.unit.body.mother_ship) {
        weaponIconMask2.id = "msWeaponMask2";
    } else if (!scale) {
        weaponIconMask2.id = "unitWeaponMask2";
    }

    weaponIcon.appendChild(weaponIconMask2);
    weaponIcon.appendChild(weaponIconMask);

    for (let i in slotData.unit.body.weapons) {

        if (slotData.unit.body.weapons.hasOwnProperty(i) && slotData.unit.body.weapons[i].weapon) {

            $(weaponIconMask).css("-webkit-mask-image", "url(/assets/units/weapon/" + slotData.unit.body.weapons[i].weapon.name + "_mask.png)");
            weaponIconMask.style.background = "#" + slotData.unit.weapon_color_1.split('x')[1];

            $(weaponIconMask2).css("-webkit-mask-image", "url(/assets/units/weapon/" + slotData.unit.body.weapons[i].weapon.name + "_mask2.png)");
            weaponIconMask2.style.background = "#" + slotData.unit.weapon_color_2.split('x')[1];
            weaponIconMask2.style.opacity = '0.3';

            weaponIcon.style.backgroundImage = "url(/assets/units/weapon/" + slotData.unit.body.weapons[i].weapon.name + ".png)";
            weaponIcon.style.top = (slotData.unit.body.weapons[i].y_attach - slotData.unit.body.weapons[i].weapon.y_attach) / (2 * scaleSize) + "px";
            weaponIcon.style.left = (slotData.unit.body.weapons[i].x_attach - slotData.unit.body.weapons[i].weapon.x_attach) / (2 * scaleSize) + "px";
        }
    }
    parent.appendChild(weaponIcon);
}