function InitWorkbench() {
    if (document.getElementById("Workbench")) {
        document.getElementById("Workbench").remove()
    }

    let workbench = document.createElement("div");
    workbench.id = "Workbench";

    $(workbench).resizable({
        minHeight: 211,
        minWidth: 420,
        handles: "se",
        resize: function (event, ui) {

            $(this).find('#detailWork').css("height", $(this).height() - 20);
            $(this).find('#needItems').css("height", $(this).height() - 70);
            $(this).find('#ButtonWrapper').css("height", $(this).height() - 125);

            $(this).find('#bluePrints').css("height", $(this).height() / 2 - 15);
            $(this).find('#currentCrafts').css("height", $(this).height() / 2 - 32);

            $(this).find('#bluePrints').css("width", $(this).width() - 300);
            $(this).find('#currentCrafts').css("width", $(this).width() - 300);

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
    bluePrints.innerHTML = "<div class='blueHead'>Доступные Чертежи:</div>";

    let currentCrafts = document.createElement("div");
    currentCrafts.id = "currentCrafts";
    currentCrafts.innerHTML = "<div class='blueHead'>Текущие проекты:</div>";

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
    workStatus.innerHTML = "<div style='background-image: url(../../lobby/img/mineral.png)'><span>10%</span></div><div style='background-image: url(../../lobby/img/timeIcon.png)'><span>10%</span></div>";
    workStatus.id = "workStatus";

    let itemPreview = document.createElement("div");
    itemPreview.id = "itemPreview";

    let needItems = document.createElement("div");
    needItems.id = "needItems";

    let countHead = document.createElement("div");
    countHead.id = "countHead";
    countHead.innerHTML = "Кол-во:";

    let count = document.createElement("input");
    count.type = "number";
    count.value = "1";
    count.min = 1;

    let buttonWrapper = document.createElement("div");
    buttonWrapper.id = "ButtonWrapper";
    buttonWrapper.appendChild(countHead);
    buttonWrapper.appendChild(count);

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


    let process = createInput("Создать", buttonWrapper);
    process.style.bottom = "20px";
    $(process).click(function () {
        // TODO
    });

    let cancel = createInput("Отмена", buttonWrapper);
    $(cancel).click(function () {
        workbench.remove();
    });

    document.body.appendChild(workbench);

    lobby.send(JSON.stringify({
        event: "OpenWorkbench",
    }));
}