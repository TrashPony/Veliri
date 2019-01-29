function InitProcessor() {
    if (document.getElementById("processorRoot")) {
        document.getElementById("processorRoot").remove()
    }

    lobby.send(JSON.stringify({
        event: "ClearProcessor",
    }));

    let processor = document.createElement("div");
    processor.id = "processorRoot";

    $(processor).resizable({
        minHeight: 128,
        minWidth: 461,
        handles: "se",
        resize: function (event, ui) {
            $(this).find('.itemsPools').css("width", $(this).width() / 2 - 14);
            $(this).find('.itemsPools').css("height", $(this).height() - 60);
            $(this).find('.pollHead').css("width", $(this).width() / 2 - 18);
        }
    });

    let buttons = CreateControlButtons("0", "61px", "-3px", "29px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'processorRoot')
    });
    $(buttons.close).mousedown(function () {
        lobby.send(JSON.stringify({
            event: "ClearProcessor",
        }));
        processor.remove();
    });
    processor.appendChild(buttons.move);
    processor.appendChild(buttons.hide);
    processor.appendChild(buttons.close);

    let items = document.createElement("div");
    items.className = "itemsPools";
    items.id = "itemsPool";
    items.innerHTML = "" +
        "<div class='pollHead'>" +
        "<h3>Input materials</h3>" +
        "<div>" +
        "<div id='RecyclePercent'><div class='fillBackPercent'></div><span>50%</span></div>" +
        "<div class='util'></div>" +
        "</div>" +
        "</div>";
    $(items).mousedown(function (event) {
        // это костыль что бы работали полосы прокрутки, https://bugs.jqueryui.com/ticket/4441#no1
        if (event.offsetX >= event.target.clientWidth || event.offsetY >= event.target.clientHeight) {
            event.stopImmediatePropagation();
        }
    });
    $(items).selectable({
        filter: '.InventoryCell.active',
        start: function () {
            $('.ui-selected').removeClass('ui-selected')
        }
    });

    $(items).droppable({
        drop: function (event, ui) {
            $('.ui-selected').removeClass('ui-selected');
            let draggable = ui.draggable;
            if (draggable.data("slotData") && draggable.data("slotData").parent === "storage") {
                if (draggable.data("selectedItems") !== undefined) {
                    lobby.send(JSON.stringify({
                        event: "PlaceItemsToProcessor",
                        storage_slots: draggable.data("selectedItems").slotsNumbers,
                    }));
                } else {
                    lobby.send(JSON.stringify({
                        event: "PlaceItemToProcessor",
                        storage_slot: Number(draggable.data("slotData").number),
                    }));
                }
            }
        }
    });

    let preview = document.createElement("div");
    preview.className = "itemsPools";
    preview.innerHTML = "<div class='pollHead'><h3>Output results</h3></div>";
    preview.id = "previewPool";

    processor.appendChild(items);
    processor.appendChild(preview);
    document.body.appendChild(processor);

    let cancel = createInput("Отмена", processor);
    $(cancel).click(function () {
        lobby.send(JSON.stringify({
            event: "ClearProcessor",
        }));
        processor.remove();
    });

    let process = createInput("Переработать", processor);
    $(process).click(function () {
        lobby.send(JSON.stringify({
            event: "recycle",
        }));
    });
}