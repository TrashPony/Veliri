function FillAssortment(assortment) {
    fillAmmo(assortment.ammo);
    fillWeapon(assortment.weapons);
    fillCabs(assortment.bodies);
    fillEquip(assortment.equips);
    fillRes(assortment.resources, assortment.recycles, assortment.details);
    fillBlueprint(assortment.blueprints);
    fillBoxes(assortment.boxes);
    fillTrash(assortment.trash);
}