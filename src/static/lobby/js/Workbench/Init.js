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

    workbench.appendChild(bluePrints);
    workbench.appendChild(detailWork);

    document.body.appendChild(workbench);
}