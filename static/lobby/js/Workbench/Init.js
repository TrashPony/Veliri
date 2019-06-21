function InitWorkbench() {
    if (document.getElementById("Workbench")) {
        let jBox = $('#Workbench');
        setState('Workbench', jBox.position().left, jBox.position().top, jBox.height(), jBox.width(), false);
        return
    }

    let workbench = document.createElement("div");
    workbench.id = "Workbench";
    document.body.appendChild(workbench);

    $(workbench).data({
        resize: function (event, ui, el) {
            el.find('#detailWork').css("height", el.height() - 20);
            el.find('#needItems').css("height", el.height() - 70);
            el.find('#ButtonWrapper').css("height", el.height() - 125);

            el.find('#wrapperBP').css("height", el.height() / 2 + 5);
            el.find('#bluePrints').css("height", el.height() / 2 + 5);
            el.find('#currentCrafts').css("height", el.height() / 2 - 12);

            el.find('#bluePrints').css("width", el.width() - 280);
            el.find('#currentCrafts').css("width", el.width() - 280);

            $('#bluePrints').resizable({
                maxHeight: $(this).height() - 70,
            });
        }
    });

    $(workbench).resizable({
        minHeight: 338,
        minWidth: 512,
        handles: "se",
        resize: function (event, ui) {
            $(this).data("resize")(event, ui, $(this))
        },
        stop: function (e, ui) {
            setState(this.id, $(this).position().left, $(this).position().top, $(this).height(), $(this).width(), true);
        }
    });

    let buttons = CreateControlButtons("0", "61px", "-3px", "29px", "ПРОИЗВОДСТВО", "145px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'Workbench')
    });
    $(buttons.close).mousedown(function () {
        workBenchState = null;
        setState(workbench.id, $(workbench).position().left, $(workbench).position().top, $(workbench).height(), $(workbench).width(), false);
    });
    workbench.appendChild(buttons.move);
    workbench.appendChild(buttons.hide);
    workbench.appendChild(buttons.close);
    workbench.appendChild(buttons.head);

    let wrapperBP = document.createElement("div");
    wrapperBP.style.cssFloat = "left";

    let wrapper2BP = document.createElement("div");
    wrapper2BP.id = "wrapperBP";

    let bluePrints = document.createElement("div");
    bluePrints.id = "bluePrints";
    bluePrints.innerHTML = "<div class='blueHead'>Доступные чертежи:</div>";

    let currentCrafts = document.createElement("div");
    currentCrafts.id = "currentCrafts";
    currentCrafts.innerHTML = "<div class='blueHead' id='queueProduction'>Очередь производаства:</div>";

    $(wrapper2BP).resizable({
        alsoResize: "#bluePrints",
        alsoResizeReverse: "#currentCrafts",
        minHeight: 20,
        maxHeight: 310,
        handles: "s",
    });

    let detailWork = document.createElement("div");
    detailWork.id = "detailWork";

    let bpIcon = document.createElement("div");
    bpIcon.id = "bpIcon";

    let info = document.createElement("div");
    info.id = "bpName";

    let workStatus = document.createElement("div");
    workStatus.innerHTML = "" +
        "<div id='mineralTax' style='background-image: url(../../lobby/img/mineral.png)'><span id='mineralTaxSpan'>0%</span></div>" +
        "<div id='timeTax' style='background-image: url(../../lobby/img/timeIcon.png)'><span id='timeTaxSpan'>0%</span></div>";
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

    wrapperBP.appendChild(wrapper2BP);
    wrapper2BP.appendChild(bluePrints);
    wrapperBP.appendChild(currentCrafts);

    workbench.appendChild(wrapperBP);
    workbench.appendChild(detailWork);

    let process = createInput("", buttonWrapper);
    process.style.bottom = "20px";
    process.id = "processButton";

    let cancel = createInput("Закрыть", buttonWrapper);
    $(cancel).click(function () {
        workBenchState = null;
        setState(workbench.id, $(workbench).position().left, $(workbench).position().top, $(workbench).height(), $(workbench).width(), false);
    });

    setTimeout(function () {
        lobby.send(JSON.stringify({
            event: "OpenWorkbench",
        }));
    }, 100);

    openWindow(workbench.id, workbench);

}