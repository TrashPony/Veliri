function FillingInventory(jsonData) {
    let event = JSON.parse(jsonData).event;

    if (event === "openInventory" || event === "UpdateSquad") {
        let squad = JSON.parse(jsonData).squad;
        InventoryTable(squad.inventory);
        SquadTable(squad);
        if (squad.mather_ship != null && squad.mather_ship.body != null) {
            ConstructorTable(squad.mather_ship);
            FillPowerPanel(squad.mather_ship.body, "powerPanel")
        } else {
            NoActiveCell();
        }

        inventoryMetaInfo(JSON.parse(jsonData))

    } else if (event === "ms error") {

        let powerPanel = document.getElementById("powerPanel");

        let start = Date.now();

        let timer = setInterval(function () {
            let timePassed = Date.now() - start;
            if (timePassed >= 600) {
                clearInterval(timer);
                powerPanel.style.border = "1px solid #25a0e1";
                powerPanel.style.boxShadow = "none";
                return;
            }
            powerPanel.style.boxShadow = "inset 1px 1px 25px 1px rgba(255,0,0,1)";
            powerPanel.style.border = "1px solid #e10006";
        }, 20);

    } else if (event === "unit error") {

        let panel;

        if (JSON.parse(jsonData).error === "lacking size") {
            panel = document.getElementById("unitCubePanel");
        } else if (JSON.parse(jsonData).error === "lacking power") {
            panel = document.getElementById("unitPowerPanel");
        }

        if (panel) {
            let start = Date.now();
            let timer = setInterval(function () {
                let timePassed = Date.now() - start;
                if (timePassed >= 600) {
                    clearInterval(timer);
                    panel.style.border = "1px solid #25a0e1";
                    panel.style.boxShadow = "none";
                    return;
                }
                panel.style.boxShadow = "inset 1px 1px 25px 1px rgba(255,0,0,1)";
                panel.style.border = "1px solid #e10006";
            }, 20);
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

    sizeBlock.innerHTML = "<div style='width:" + percentFill + "%'>" +
        "<span>" + data.inventory_size + " / " + data.squad.mather_ship.body.capacity_size + "</span>" +
        "</div>";
    sizeBlock.style.color = textColor;
}