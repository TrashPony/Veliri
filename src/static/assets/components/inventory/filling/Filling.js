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
        fillSquadList(JSON.parse(jsonData).base_squads);

        if (squad) {
            if (document.getElementById("Inventory")) {
                InventoryTable(squad.inventory);

                document.getElementById("inventoryStorageInventory").style.opacity = "1";
                $("#squadName span").last().text(squad.name).css('color', '#00FFFD');
                $("#renameSquadButton").removeClass("noActive");
                $("#deleteSquadButton").removeClass("noActive");

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
            }
        } else {
            if (document.getElementById("Inventory")) {
                NoActiveCell();
                SquadTable(squad);
                $("#inventoryStorageInventory").css('opacity', '0.5');
                $("#inventoryStorageInventory > .InventoryCell").css('background-image', 'none').empty();
                $("#sizeInventoryInfo").empty();
                $("#deleteSquadButton").addClass("noActive");
                $("#renameSquadButton").addClass("noActive");
                $("#squadName span").last().text(" отряд не выбран").css('color', '#00FFFD');
            }
        }

        if (event === "openInventory") {
            // склад и магазин поднимаются только тогда когда игрок на базе
            if (JSON.parse(jsonData).in_base && !document.getElementById("inventoryStorage")) {
                CreateStorage();
                ConnectMarket();
            }
        }
    }
}

function fillSquadList(squads) {
    let squadList = $('#SquadsList');
    squadList.empty();

    for (let i in squads) {
        if (squads.hasOwnProperty(i) && squads[i]) {
            let squad = $('<div id = "squadListName" ' + squads[i].id + '><span class="squadListName">' + squads[i].name + '</span></div>');
            squad.click(function () {
                inventorySocket.send(JSON.stringify({
                    event: "changeSquad",
                    squad_id: squads[i].id,
                }));
            });
            squadList.append(squad)
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