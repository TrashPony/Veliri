function initReload(msg) {

    if (msg.error) {
        return
    }

    if (document.getElementById("AmmoSlots")) {
        document.getElementById("AmmoSlots").remove();
    }

    let inventory = document.createElement("div");
    inventory.id = "AmmoSlots";

    let inventoryStorage = document.createElement("div");
    inventoryStorage.id = "AmmoSlotsStorage";

    inventory.appendChild(inventoryStorage);

    document.getElementById("BoxUnitSubMenu").appendChild(inventory);

    for (let i in msg.ammo_slots) {
        if (msg.ammo_slots.hasOwnProperty(i) && msg.ammo_slots[i].item) {
            let cell = document.createElement("div");
            CreateInventoryCell(cell, msg.ammo_slots[i], i, "");

            $(cell).draggable({
                disabled: true,
            });

            cell.onclick = function () {
                field.send(JSON.stringify({
                    event: "Reload",
                    slot: Number(i),
                    q: msg.q,
                    r: msg.r,
                }))
            };
            inventoryStorage.appendChild(cell);
        }
    }
}

function ReloadMark(msg) {
    if (document.getElementById("AmmoSlots")) {
        document.getElementById("AmmoSlots").remove();
    }

    let unit = GetGameUnitXY(msg.q, msg.r);

    let ReloadEquip = ReloadAnimation(true, false);
    unit.sprite.addChild(ReloadEquip);
    unit.reloadIcon = ReloadEquip;
}

function ReloadAnimation(loop, kill) {
    let ReloadEquip = game.make.sprite(0, 0, 'ReloadEquip', 0);
    ReloadEquip.animations.add('RepairKit');
    ReloadEquip.animations.play('RepairKit', 30, loop, kill);
    ReloadEquip.anchor.set(0.5);
    ReloadEquip.scale.set(0.2);

    return ReloadEquip
}