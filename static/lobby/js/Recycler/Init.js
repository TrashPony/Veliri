function InitProcessorRoot() {
    if (document.getElementById("processorRoot")) {
        let jBox = $('#processorRoot');
        setState('processorRoot', jBox.position().left, jBox.position().top, jBox.height(), jBox.width(), false);
        return
    }

    setTimeout(function () {
        lobby.send(JSON.stringify({
            event: "ClearProcessor",
        }));
    }, 100);

    let processor = document.createElement("div");
    processor.id = "processorRoot";
    document.body.appendChild(processor);


    $(processor).data({
        resize: function (event, ui, el) {
            el.find('.itemsPools').css("width", el.width() / 2 - 14);
            el.find('.itemsPools').css("height", el.height() - 60);
            el.find('.pollHead').css("width", el.width() / 2 - 18);
        }
    });

    $(processor).resizable({
        minHeight: 128,
        minWidth: 461,
        handles: "se",
        resize: function (event, ui) {
            $(this).data("resize")(event, ui, $(this))
        },
        stop: function (e, ui) {
            setState(this.id, $(this).position().left, $(this).position().top, $(this).height(), $(this).width(), true);
        }
    });

    let buttons = CreateControlButtons("0", "61px", "-3px", "29px", "ПЕРЕРАБОТЧИК", "145px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'processorRoot')
    });
    $(buttons.close).mousedown(function () {
        lobby.send(JSON.stringify({
            event: "ClearProcessor",
        }));
        setState(processor.id, $(processor).position().left, $(processor).position().top, $(processor).height(), $(processor).width(), false);
    });
    processor.appendChild(buttons.move);
    processor.appendChild(buttons.hide);
    processor.appendChild(buttons.close);
    processor.appendChild(buttons.head);

    let items = document.createElement("div");
    items.className = "itemsPools";
    items.id = "itemsPool";
    items.innerHTML = "" +
        "<div class='pollHead'>" +
        "<h3>Input materials</h3>" +
        "<div>" +
        "<div id='RecyclePercent'><div id='fillBackPercent'></div><span id='UserRecyclePercent'></span></div>" +
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

            if (draggable.data("slotData") && (draggable.data("slotData").parent === "storage" || draggable.data("slotData").parent === "squadInventory")) {
                if (draggable.data("selectedItems") !== undefined) {
                    lobby.send(JSON.stringify({
                        event: "PlaceItemsToProcessor",
                        storage_slots: draggable.data("selectedItems").slotsNumbers,
                        item_source: draggable.data("slotData").parent,
                    }));
                } else {
                    lobby.send(JSON.stringify({
                        event: "PlaceItemToProcessor",
                        storage_slot: Number(draggable.data("slotData").number),
                        item_source: draggable.data("slotData").parent,
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

    let cancel = createInput("Отмена", processor);
    $(cancel).click(function () {
        lobby.send(JSON.stringify({
            event: "ClearProcessor",
        }));
        setState(processor.id, $(processor).position().left, $(processor).position().top, $(processor).height(), $(processor).width(), false);
    });

    let process = createInput("Переработать", processor);
    $(process).click(function () {
        lobby.send(JSON.stringify({
            event: "recycle",
        }));
    });

    openWindow(processor.id, processor)
}