function GetEquip(unitID, typeSlot, numberSlot) {

    let unit = game.units[unitID];

    if (unit) {
        if (typeSlot === 1) {
            return unit.body.equippingI[numberSlot]
        }

        if (typeSlot === 2) {
            return unit.body.equippingII[numberSlot]
        }

        if (typeSlot === 3) {
            return unit.body.equippingIII[numberSlot]
        }

        if (typeSlot === 4) {
            return unit.body.equippingIV[numberSlot]
        }

        if (typeSlot === 5) {
            return unit.body.equippingV[numberSlot]
        }
    }
}

function GetSpriteEqip(unitID, typeSlot, numberSlot) {
    let unit = game.units[unitID];
    if (unit) {
        for (let i = 0; i < unit.sprite.equipSprites.length; i++) {
            let slot = unit.sprite.equipSprites[i];
            if (slot.slot.type_slot === typeSlot && slot.slot.number_slot === numberSlot) {
                return slot
            }
        }
    }
}