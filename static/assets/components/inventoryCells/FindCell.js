function FindCell(name, type, storage) {
    let slots = [];
    if (name && name !== "") {

    }

    if (type && type !== "") {
        if (type === "MS") {
            storage.find('.InventoryCell.active').each(function () {
                console.log();
                if ($(this).data("slotData").data.type === "body" && $(this).data("slotData").data.item.mother_ship) {
                    slots.push(this);
                }
            })
        }

        if (type === "thorium") {
            storage.find('.InventoryCell.active').each(function () {
                console.log();
                if ($(this).data("slotData").data.type === "recycle" && $(this).data("slotData").data.item.name === "enriched_thorium") {
                    slots.push(this);
                }
            })
        }

        if (type === "equips") {
            storage.find('.InventoryCell.active').each(function () {
                console.log();
                if ($(this).data("slotData").data.type === "equip") {
                    slots.push(this);
                }
            })
        }

        if (type === "weapons") {
            storage.find('.InventoryCell.active').each(function () {
                console.log();
                if ($(this).data("slotData").data.type === "weapon") {
                    slots.push(this);
                }
            })
        }

        if (type === "ammo") {
            storage.find('.InventoryCell.active').each(function () {
                console.log();
                if ($(this).data("slotData").data.type === "ammo") {
                    slots.push(this);
                }
            })
        }
    }

    return slots
}