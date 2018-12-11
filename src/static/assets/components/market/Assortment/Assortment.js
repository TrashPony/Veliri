let filterKey = {type: '', id: 0};

function FillAssortment(assortment) {
    fillAmmo(assortment.ammo);
    fillWeapon(assortment.weapons);
    fillCabs(assortment.bodies);
    fillEquip(assortment.equips);
}