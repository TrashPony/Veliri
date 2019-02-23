function SetThorium(thorium, slot, source) {
    let thoriumCells;
    if (game && game.typeService === "global") {
        thoriumCells = document.getElementsByClassName("Thorium");
    } else {
        thoriumCells = document.getElementsByClassName("thoriumSlots");
    }

    for (let i = 0; thoriumCells && i < thoriumCells.length; i++) {
        if (Number(thoriumCells[i].count) < Number(thoriumCells[i].maxCount)) {

            thoriumCells[i].style.boxShadow = "inset 0 0 5px 3px rgb(255, 149, 32)";
            thoriumCells[i].style.cursor = "pointer";
            thoriumCells[i].onmouseout = null;
            thoriumCells[i].style.animation = "none";

            thoriumCells[i].actionFlag = true;

            thoriumCells[i].onclick = function (event) {
                event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

                for (let j = 0; thoriumCells && j < thoriumCells.length; j++) {
                    thoriumCells[j].actionFlag = false;
                }

                if (game && game.typeService === "global") {
                    global.send(JSON.stringify({
                        event: "updateThorium",
                        inventory_slot: Number(slot),
                        thorium_slot: Number(thoriumCells[i].numberSlot),
                        source: source,
                    }));
                } else {
                    inventorySocket.send(JSON.stringify({
                        event: "SetThorium",
                        inventory_slot: Number(slot),
                        thorium_slot: Number(thoriumCells[i].numberSlot),
                        source: source,
                    }));
                }

                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    }
}