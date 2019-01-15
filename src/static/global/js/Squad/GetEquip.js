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