let size = 0;

function FillingInventory(jsonData) {
    let event = JSON.parse(jsonData).event;
    let error = JSON.parse(jsonData).error;

    if (error) {
        alertError(jsonData);
        return
    }

    if (event === "openInventory" || event === "UpdateSquad") {
        let squad = JSON.parse(jsonData).squad;
        InventoryTable(squad.inventory);

        if (squad.mather_ship != null && squad.mather_ship.body != null) {

            size = squad.mather_ship.body.capacity_size;
            inventoryMetaInfo(JSON.parse(jsonData));

            if (document.getElementById("inventoryBox")) {
                SquadTable(squad);
                ConstructorTable(squad.mather_ship);
                FillPowerPanel(squad.mather_ship.body, "powerPanel");
                FillMSWeaponTypePanel(squad.mather_ship.body, "MSWeaponPanel");
            }
        } else {
            NoActiveCell();
            SquadTable(squad);
        }

        if (event === "openInventory") {
            // склад и магазин поднимаются только тогда когда игрок на базе
            if (JSON.parse(jsonData).in_base) {
                CreateStorage();
                ConnectMarket();
            }
        }
    }
}

function inventoryMetaInfo(data) {

    let percentFill = 100 / (data.squad.mather_ship.body.capacity_size / data.inventory_size);

    let sizeBlock = document.getElementById("sizeInventoryInfo");
    let textColor = "";
    if (data.inventory_size > data.squad.mather_ship.body.capacity_size) {
        textColor = "#b9281d"
    } else {
        textColor = "#decbcb"
    }

    sizeBlock.innerHTML = "<div id='realSize' style='width:" + percentFill + "%'>" +
        "<span>" + data.inventory_size + " / " + data.squad.mather_ship.body.capacity_size + "</span>" +
        "</div>";
    sizeBlock.style.color = textColor;

    if (document.getElementById("InventoryTip")) {
        document.getElementById("InventoryTip").remove();
    }
}