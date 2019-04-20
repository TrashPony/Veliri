function GetEquip(typeSlot, numberSlot) {
    if (typeSlot === 1) {
        return game.squad.mather_ship.body.equippingI[numberSlot]
    }

    if (typeSlot === 2) {
        return game.squad.mather_ship.body.equippingII[numberSlot]
    }

    if (typeSlot === 3) {
        return game.squad.mather_ship.body.equippingIII[numberSlot]
    }

    if (typeSlot === 4) {
        return game.squad.mather_ship.body.equippingIV[numberSlot]
    }

    if (typeSlot === 5) {
        return game.squad.mather_ship.body.equippingV[numberSlot]
    }
}

function GetSpriteEqip(typeSlot, numberSlot) {
    for (let i = 0; i < game.squad.sprite.equipSprites.length; i++) {
        let slot = game.squad.sprite.equipSprites[i];
        if (slot.slot.type_slot === typeSlot && slot.slot.number_slot === numberSlot) {
            return slot
        }
    }
}