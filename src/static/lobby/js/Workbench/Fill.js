function FillWorkbench(jsonData) {
    console.log(jsonData)
    let bpBlock = document.getElementById("bluePrints");
    for (let i in jsonData.storage.slots) {
        if (jsonData.storage.slots[i].type === "blueprints") {
            let blueRow = document.createElement("div");
            blueRow.className = "blueRow";
            blueRow.innerHTML = "" +
                "<div class='nameBP'>" + jsonData.storage.slots[i].item.name + "</div>" +
                "<div class='countBP'>x" + jsonData.storage.slots[i].quantity + "</div>";
            bpBlock.appendChild(blueRow);

            console.log(jsonData.storage.slots[i])
        }
    }
}