function InitWorkbench() {
    if (document.getElementById("Workbench")) {
        document.getElementById("Workbench").remove()
    }

    let workbench = document.createElement("div");
    workbench.id = "Workbench";

    $(workbench).resizable({
        minHeight: 128,
        minWidth: 461,
        handles: "se",
        resize: function (event, ui) {
            // TODO
        }
    });

    let buttons = CreateControlButtons("0", "61px", "-3px", "29px", "ПРОИЗВОДСТВО", "145px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'Workbench')
    });
    $(buttons.close).mousedown(function () {
        workbench.remove();
    });
    workbench.appendChild(buttons.move);
    workbench.appendChild(buttons.hide);
    workbench.appendChild(buttons.close);
    workbench.appendChild(buttons.head);

    let bluePrints = document.createElement("div");
    bluePrints.id = "bluePrints";
    bluePrints.innerHTML = "<div class='blueHead'>Доступные Чертежи:</div>";

    let detailWork = document.createElement("div");
    detailWork.id = "detailWork";

    let bpIcon = document.createElement("div");
    bpIcon.id = "bpIcon";

    let info = document.createElement("div");
    info.id = "bpName";

    let itemPreview = document.createElement("div");
    itemPreview.id = "itemPreview";

    let needItems = document.createElement("div");
    needItems.id = "needItems";

    detailWork.appendChild(bpIcon);
    detailWork.appendChild(info);
    detailWork.appendChild(itemPreview);
    detailWork.appendChild(needItems);

    workbench.appendChild(bluePrints);
    workbench.appendChild(detailWork);

    document.body.appendChild(workbench);
}