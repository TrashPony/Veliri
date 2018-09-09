function FillingInventory(jsonData) {
    let event = JSON.parse(jsonData).event;

    if (event === "openInventory" || event === "UpdateSquad") {
        let squad = JSON.parse(jsonData).squad;
        InventoryTable(squad.inventory);
        SquadTable(squad);
        if (squad.mather_ship != null && squad.mather_ship.body != null) {
            ConstructorTable(squad.mather_ship.body);
            FillPowerPanel(squad.mather_ship.body, "powerPanel")
        } else {
            NoActiveCell();
        }
    } else if (event === "ms error") {
        console.log(":dfdfd");
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

        let powerPanel = document.getElementById("unitPowerPanel");
        if (powerPanel) {

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
        }
    }
}