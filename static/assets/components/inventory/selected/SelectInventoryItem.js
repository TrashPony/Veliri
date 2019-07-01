function SelectInventoryItem(e) {

    DestroyInventoryTip();
    DestroyInventoryClickEvent();

    if (this.source === 'storage') {
        InventorySelectTip(JSON.parse(this.slotData), false, false, this.number, true);
    } else {
        InventorySelectTip(JSON.parse(this.slotData), false, true, this.number);
    }

    if (JSON.parse(this.slotData).type === "body") {
        SetBody(JSON.parse(this.slotData).item, this.number, this.source);
    }

    if (JSON.parse(this.slotData).type === "weapon") {
        SetWeapon(JSON.parse(this.slotData).item, this.number, this.source);
    }

    if (JSON.parse(this.slotData).type === "equip") {
        SetEquip(JSON.parse(this.slotData).item, this.number, this.source)
    }

    if (JSON.parse(this.slotData).type === "ammo") {
        SetAmmo(JSON.parse(this.slotData).item, this.number, this.source)
    }

    if (JSON.parse(this.slotData).type === "resource") {
    }

    if (JSON.parse(this.slotData).type === "recycle" && JSON.parse(this.slotData).item.name === "enriched_thorium") {
        SetThorium(JSON.parse(this.slotData).item, this.number, this.source)
    }
}