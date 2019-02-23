function CreateSquadHead() {
    let squadHead = document.getElementById("SquadHead");

    $(squadHead).append('<div id="squadName"> <span>Отряд: </span> <span></span> </div> ' +
        '<div id="deleteSquadButton" class="deleteButton"></div>' +
        '<div id="renameSquadButton" class="renameButton"></div>');

    $("#renameSquadButton").click(function () {
        let unitIcon = document.getElementById("MSIcon");
        if (!unitIcon || !unitIcon.shipBody) {
            return
        }

        $('<div><input id="renameSquad" type="text" placeholder="Новое имя отряда"></div>').appendTo('body').dialog({
            open: function (event, ui) {
                $(".ui-dialog-titlebar-close", ui.dialog | ui).hide();
            },
            classes: {"ui-dialog": "renameSquad"},
            modal: true, title: 'Переименовать отряд?', autoOpen: true,
            width: '250', minHeight: 20, height: 'auto', resizable: false,
            buttons: [{
                text: "Переименовать",
                click: function () {

                    let input = $('#renameSquad');
                    inventorySocket.send(JSON.stringify({
                        event: "RenameSquad",
                        name: input.val(),
                    }));

                    input.remove();
                    DestroyInventoryClickEvent();
                    DestroyInventoryTip();
                    $(this).dialog("close");
                }
            }, {
                text: "Отмена",
                click: function () {
                    $(this).dialog("close");
                }
            }],
        });
    });


    $("#deleteSquadButton").click(function () {
        let unitIcon = document.getElementById("MSIcon");
        if (!unitIcon || !unitIcon.shipBody) {
            return
        }
        $('<div></div>').appendTo('body').dialog({
            open: function (event, ui) {
                $(".ui-dialog-titlebar-close", ui.dialog | ui).hide();
            },
            classes: {"ui-dialog": "deleteSquad"},
            modal: true, title: 'Удалить отряд?', autoOpen: true,
            width: 'auto', height: '0', resizable: false,
            buttons: [{
                text: "Удалить",
                click: function () {
                    inventorySocket.send(JSON.stringify({
                        event: "RemoveMotherShipBody",
                        destination: "storage",
                    }));

                    DestroyInventoryClickEvent();
                    DestroyInventoryTip();
                    $(this).dialog("close");
                }
            }, {
                text: "Отмена",
                click: function () {
                    $(this).dialog("close");
                }
            }],
        });
    });
}