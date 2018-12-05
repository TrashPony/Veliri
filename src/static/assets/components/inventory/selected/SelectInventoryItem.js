function SelectInventoryItem(e) {

    DestroyInventoryTip();
    DestroyInventoryClickEvent();

    if (this.inventoryType === 'storage') {
        InventorySelectTip(JSON.parse(this.slotData), e.clientX, e.clientY, false, false, this.number, true);
    } else {
        InventorySelectTip(JSON.parse(this.slotData), e.clientX, e.clientY, false, true, this.number);

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
}