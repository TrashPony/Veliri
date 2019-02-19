function InitWorkbench() {
    if (document.getElementById("Workbench")) {
        document.getElementById("Workbench").remove()
    }

    let workbench = document.createElement("div");
    workbench.id = "Workbench";

    $(workbench).resizable({
        minHeight: 237,
        minWidth: 420,
        handles: "se",
        resize: function (event, ui) {

            $(this).find('#detailWork').css("height", $(this).height() - 20);
            $(this).find('#needItems').css("height", $(this).height() - 70);
            $(this).find('#ButtonWrapper').css("height", $(this).height() - 125);

            $(this).find('#bluePrints').css("height", $(this).height() / 2 + 5);
            $(this).find('#currentCrafts').css("height", $(this).height() / 2 - 12);

            $(this).find('#bluePrints').css("width", $(this).width() - 280);
            $(this).find('#currentCrafts').css("width", $(this).width() - 280);

            $('#bluePrints').resizable({
                maxHeight: $(this).height() - 70,
            });
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

    let wrapperBP = document.createElement("div");
    wrapperBP.style.cssFloat = "left";

    let bluePrints = document.createElement("div");
    bluePrints.id = "bluePrints";
    bluePrints.innerHTML = "<div class='blueHead'>Доступные чертежи:</div>";

    let currentCrafts = document.createElement("div");
    currentCrafts.id = "currentCrafts";
    currentCrafts.innerHTML = "<div class='blueHead' id='queueProduction'>Очередь производаства:</div>";

    $(bluePrints).resizable({
        alsoResizeReverse: "#currentCrafts",
        minHeight: 20,
        maxHeight: 215,
        handles: "s",
    });

    let detailWork = document.createElement("div");
    detailWork.id = "detailWork";

    let bpIcon = document.createElement("div");
    bpIcon.id = "bpIcon";

    let info = document.createElement("div");
    info.id = "bpName";

    let workStatus = document.createElement("div");
    workStatus.innerHTML = "<div style='background-image: url(../../lobby/img/mineral.png)'><span>0%</span></div><div style='background-image: url(../../lobby/img/timeIcon.png)'><span>0%</span></div>";
    workStatus.id = "workStatus";

    let itemPreview = document.createElement("div");
    itemPreview.id = "itemPreview";

    let needItems = document.createElement("div");
    needItems.id = "needItems";

    let countHead = document.createElement("div");
    countHead.id = "countHead";
    countHead.innerHTML = "Кол-во:";

    let count = document.createElement("input");
    count.id = "bpCountWork";
    count.type = "number";
    count.value = "";
    count.min = 1;

    let time = document.createElement("div");
    time.id = "bpCraftTime";
    time.innerHTML = "";

    let buttonWrapper = document.createElement("div");
    buttonWrapper.id = "ButtonWrapper";
    buttonWrapper.appendChild(countHead);
    buttonWrapper.appendChild(count);
    buttonWrapper.appendChild(time);


    detailWork.appendChild(bpIcon);
    detailWork.appendChild(info);
    detailWork.appendChild(workStatus);
    detailWork.appendChild(itemPreview);
    detailWork.appendChild(needItems);
    detailWork.appendChild(buttonWrapper);

    wrapperBP.appendChild(bluePrints);
    wrapperBP.appendChild(currentCrafts);

    workbench.appendChild(wrapperBP);
    workbench.appendChild(detailWork);


    let process = createInput("", buttonWrapper);
    process.style.bottom = "20px";
    process.id = "processButton";

    let cancel = createInput("Закрыть", buttonWrapper);
    $(cancel).click(function () {
        workbench.remove();
    });

    document.body.appendChild(workbench);

    lobby.send(JSON.stringify({
        event: "OpenWorkbench",
    }));
}