function FillWorkbench(jsonData) {
    let bpBlock = document.getElementById("bluePrints");
    if (!bpBlock) return;
    $('#bluePrints .blueRow').remove();

    for (let i in jsonData.storage.slots) {
        if (jsonData.storage.slots[i].type === "blueprints") {

            let blueRow = document.getElementById("blueRowBP" + i);
            if (!blueRow) {
                blueRow = document.createElement("div");
                blueRow.className = "blueRow";
                blueRow.id = "blueRowBP" + i;
            }

            $(blueRow).data("update", {update: true});

            blueRow.innerHTML = "" +
                "<div class='nameBP'>" + jsonData.storage.slots[i].item.name + "</div>" +
                "<div class='countBP'>x" + jsonData.storage.slots[i].quantity + "</div>";
            bpBlock.appendChild(blueRow);

            $(blueRow).click(function () {
                lobby.send(JSON.stringify({
                    event: "SelectBP",
                    storage_slot: Number(i),
                    count: 1,
                }));
            });
        }
    }

    FillCurrentWorks(jsonData.blue_works);

    // удаляем всех кого не обновили
    $('.blueRow').each(function () {
        if ($(this).data("update").update) {
            $(this).data("update").update = false;
            if ($(this).data("count")) $(this).data("count").count = 0;
        } else {
            $(this).remove()
        }
    })
}

function FillCurrentWorks(works) {
    let workBlock = document.getElementById("currentCrafts");
    if (!workBlock) return;

    for (let i in works) {
        let dataRow = innerDate(works[i]);

        // сложный ид для того что бы групировать одинаковые итемы а не добавлять каждый раз новый
        let idRow = "blueRow" +
            works[i].blueprint.id +
            "" + works[i].mineral_saving_percentage +
            "" + works[i].time_saving_percentage;

        if (dataRow.active) {
            idRow = "blueRow" + works[i].id
        }

        let blueRow = document.getElementById(idRow);
        if (!blueRow) {
            blueRow = document.createElement("div");
            blueRow.className = "blueRow";
            blueRow.id = idRow;

            if (dataRow.active) {
                $(blueRow).insertAfter('#queueProduction');// активный крафт всегда первый
            } else {
                workBlock.appendChild(blueRow);
            }

            $(blueRow).data("count", {count: 0});
        }

        blueRow.innerHTML = dataRow.html;
        $(blueRow).data("update", {update: true});

        if (!dataRow.active) {
            $(blueRow).data("count").count++;
            blueRow.innerHTML = dataRow.html;
            blueRow.innerHTML += "<div class='countBP' style='margin-right: 4px'>x" + $(blueRow).data("count").count + "</div>";
        }

        $(blueRow).click(function () {
            // TODO заполнять привью,
        })
    }
}

function innerDate(work) {

    let data = new Date();
    let finishTime = new Date(work.finish_time);
    data.setTime(finishTime.getTime() - new Date().getTime());

    let realTimeCraft = work.blueprint.craft_time - (work.blueprint.craft_time * work.time_saving_percentage / 100);
    let startTime = new Date().setTime(finishTime.getTime() - realTimeCraft * 1000);
    let diffTime = (new Date() - startTime) / 1000;

    let percent = (diffTime * 100) / realTimeCraft;
    let widthTimeLine = 0;

    let days, hours, minutes, seconds;
    if (percent > 0) {
        widthTimeLine = Math.round(percent);

        days = (data.getUTCDate() - 1 > 0) ? data.getUTCDate() - 1 + "d: " : '';
        hours = (data.getUTCHours() > 9) ? data.getUTCHours() : "0" + data.getUTCHours();
        minutes = (data.getUTCMinutes() > 9) ? data.getUTCMinutes() : "0" + data.getUTCMinutes();
        seconds = (data.getUTCSeconds() > 9) ? data.getUTCSeconds() : "0" + data.getUTCSeconds();
    } else {

        data.setTime(work.blueprint.craft_time * 1000);

        days = (data.getUTCDate() - 1 > 0) ? data.getUTCDate() - 1 + "d: " : '';
        hours = (data.getUTCHours() > 9) ? data.getUTCHours() : "0" + data.getUTCHours();
        minutes = (data.getUTCMinutes() > 9) ? data.getUTCMinutes() : "0" + data.getUTCMinutes();
        seconds = (data.getUTCSeconds() > 9) ? data.getUTCSeconds() : "0" + data.getUTCSeconds();
    }

    return {
        html: "" +
            "<div class='nameBP'>" + work.item.name + "</div>" +
            "<div class='timerWork'><span>"
            + days
            + hours
            + " : "
            + minutes
            + " : "
            + seconds
            + "</span>" +
            "<div class='workTimeLine' style='width: " + widthTimeLine + "%'></div></div>", active: percent > 0
    }
}

function SelectBP(jsonData) {
    document.getElementById("bpName").innerHTML = jsonData.blue_print.name;
    document.getElementById("bpIcon").style.backgroundImage = "url(/assets/blueprints/" + jsonData.blue_print.name + ".png)";
    document.getElementById("bpCraftTime").innerHTML = jsonData.blue_print.craft_time * jsonData.count + "s";

    let itemPreview = document.getElementById("itemPreview");

    if (jsonData.blue_print.item_type === "resource" || jsonData.blue_print.item_type === "recycle") {
        itemPreview.style.backgroundImage = "url(/assets/resource/" + jsonData.bp_item.name + ".png)";
    } else if (jsonData.blue_print.item_type === "boxes") {
        itemPreview.style.backgroundImage = "url(/assets/" + jsonData.blue_print.item_type + "/" + jsonData.bp_item.name + ".png)";
    } else if (jsonData.blue_print.item_type === "detail") {
        itemPreview.style.backgroundImage = "url(/assets/resource/detail/" + jsonData.bp_item.name + ".png)";
    } else if (jsonData.blue_print.item_type === "blueprints") {
        itemPreview.style.backgroundImage = "url(/assets/blueprints/" + jsonData.bp_item.name + ".png)";
    } else {
        itemPreview.style.backgroundImage = "url(/assets/units/" + jsonData.blue_print.item_type + "/" + jsonData.bp_item.name + ".png)";
    }
    itemPreview.innerHTML = "<span>x" + jsonData.count + "</span>";

    fillNeedItems(jsonData.preview_recycle_slots);

    let bpCountWork = document.getElementById("bpCountWork");
    bpCountWork.value = jsonData.count;
    bpCountWork.max = jsonData.max_count;
    bpCountWork.oninput = function () {
        lobby.send(JSON.stringify({
            event: "SelectBP",
            storage_slot: Number(jsonData.storage_slot),
            count: Number(this.value)
        }));
    };

    let processButton = document.getElementById("processButton");
    $(processButton).off('click'); // обязательно удаляем прошлое событие
    $(processButton).click(function () {
        lobby.send(JSON.stringify({
            event: "Craft",
            storage_slot: Number(jsonData.storage_slot),
            count: Number(bpCountWork.value)
        }));
    });
}

function fillNeedItems(items) {
    let needItems = document.getElementById("needItems");
    $(needItems).empty();

    for (let i in items) {
        let cell = document.createElement("div");

        let item = Object.assign({}, items[i]);

        CreateInventoryCell(cell, item, i, "", onclick);
        let section = CheckRecycleSection(item, needItems);

        if (!item.find) {
            cell.style.border = "1px solid red";
            cell.innerHTML += "<div class='noAllowCell'></div>"
        }

        section.appendChild(cell);
    }
}