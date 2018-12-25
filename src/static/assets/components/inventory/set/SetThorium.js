function SetThorium(thorium, slot) {
    let thoriumCells = document.getElementsByClassName("thoriumSlots");

    for (let i = 0; thoriumCells && i < thoriumCells.length; i++) {
        if (Number(thoriumCells[i].count) < Number(thoriumCells[i].maxCount)) {
            thoriumCells[i].style.boxShadow = "inset 0 0 5px 3px rgb(255, 149, 32)";
            thoriumCells[i].style.cursor = "pointer";
            thoriumCells[i].onmouseout = null;

            thoriumCells[i].onclick = function (event) {
                event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
                inventorySocket.send(JSON.stringify({
                    event: "SetThorium",
                    inventory_slot: Number(slot),
                    thorium_slot: Number(thoriumCells[i].numberSlot),
                }));

                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    }
}