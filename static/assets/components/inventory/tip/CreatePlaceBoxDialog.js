function CreatePlaceBoxDialog(x, y, numberSlot, slot) {
    if (slot.item.protect) {
        let func = function () {
            global.send(JSON.stringify({
                event: "placeNewBox",
                slot: Number(numberSlot),
                box_password: Number(document.getElementById("passPlaceBox").value),
            }));
        };
        PassBlock(func);
    } else {
        global.send(JSON.stringify({
            event: "placeNewBox",
            slot: Number(numberSlot),
        }));
    }
}