function FillRecycler(jsonData) {

    let itemsPool = document.getElementById("itemsPool");
    if (!itemsPool) return;

    $("#itemsPool .InventoryCell").remove();
    for (let i in jsonData.recycle_slots) {

        let cell = document.createElement("div");
        CreateInventoryCell(cell, jsonData.recycle_slots[i].slot, i, "recycler", onclick);

        // TODO заполнение по разделам

        cell.onclick = function () {
            lobby.send(JSON.stringify({
                event: "RemoveItemFromProcessor",
                recycler_slot: Number(i),
            }));
        };
        itemsPool.appendChild(cell);
    }
}