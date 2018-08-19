function SelectInventoryItem(e) {

    DestroyInventoryTip();
    DestroyInventoryClickEvent();

    InventorySelectTip(JSON.parse(this.slotData).item, e.clientX, e.clientY);

    if (JSON.parse(this.slotData).type === "body") {
        SetBody(JSON.parse(this.slotData).item, this.number);
    }

    if (JSON.parse(this.slotData).type === "weapon") {
        SetWeapon(JSON.parse(this.slotData).item, this.number);
    }

    if (JSON.parse(this.slotData).type === "equip") {
        SetEquip(JSON.parse(this.slotData).item, this.number)
    }

    if (JSON.parse(this.slotData).type === "ammo") {
        SetAmmo(JSON.parse(this.slotData).item, this.number)
    }
}