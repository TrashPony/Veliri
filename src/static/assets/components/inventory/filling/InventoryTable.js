function InventoryTable(inventoryItems) {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);

        if (inventoryItems.hasOwnProperty(i) && inventoryItems[i].item !== null) {

            cell.slotData = JSON.stringify(inventoryItems[i]);
            cell.number = i;

            cell.style.backgroundImage = "url(/assets/" + JSON.parse(cell.slotData).item.name + ".png)";
            cell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(cell.slotData).quantity + "</span>";

            CreateHealBar(cell);

            cell.onclick = SelectInventoryItem;

            cell.addEventListener("mousemove", InventoryOverTip);
            cell.addEventListener("mouseout", function () {
                let inventoryTip = document.getElementById("InventoryTipOver");
                if (inventoryTip) {
                    inventoryTip.remove()
                }
            });
        } else {

            cell.slotData = null;

            cell.style.backgroundImage = null;
            cell.innerHTML = "";

            cell.onclick = function () {
                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            };
        }
    }
}

function CreateHealBar(cell) {
    let cellData = JSON.parse(cell.slotData);

    if (cellData.type !== "ammo") {
        let backHealBar = document.createElement("div");
        backHealBar.className = "backHealBar";

        let healBar = document.createElement("div");
        healBar.className = "healBar";

        let percentHP = 100 / (cellData.item.max_hp / cellData.hp);
        healBar.style.width = percentHP + "%";

        if (percentHP === 100) {
            backHealBar.style.opacity = "0"
        } else if (percentHP < 90 && percentHP > 75) {
            healBar.style.backgroundColor = "#fff326"
        } else if (percentHP < 75 && percentHP > 50) {
            healBar.style.backgroundColor = "#fac227"
        } else if (percentHP < 50 && percentHP > 25) {
            healBar.style.backgroundColor = "#fa7b31"
        } else if (percentHP < 25) {
            healBar.style.backgroundColor = "#fa2e26"
        }

        backHealBar.appendChild(healBar);
        cell.appendChild(backHealBar);

        console.log(percentHP);
    }
}

function InventoryOverTip(e) {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.style.top = e.clientY + "px";
        inventoryTip.style.left = e.clientX + "px";
    } else {
        InventorySelectTip(JSON.parse(this.slotData), e.clientX, e.clientY, true);
    }
}