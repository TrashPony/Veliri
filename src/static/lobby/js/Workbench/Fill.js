function FillWorkbench(jsonData) {
    let bpBlock = document.getElementById("bluePrints");
    if (!bpBlock) return;
    $('#bluePrints .blueRow').remove();

    $('.blueRow').each(function () {
        if ($(this).data("count")) $(this).data("count").count = 0;
        if ($(this).data("time")) $(this).data("time").finishTime = 0;
        if ($(this).data("time")) $(this).data("time").startTime = Infinity;
    });

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

            blueRow.onclick = function () {
                lobby.send(JSON.stringify({
                    event: "SelectBP",
                    storage_slot: Number(i),
                    count: 1,
                }));
            };
        }
    }

    FillCurrentWorks(jsonData.blue_works);
    // TODO обновление работ если выбрана

    // удаляем всех кого не обновили
    $('.blueRow').each(function () {
        if ($(this).data("update").update) {
            $(this).data("update").update = false;
        } else {
            $(this).remove()
        }
    })
}

function FillCurrentWorks(works) {
    let workBlock = document.getElementById("currentCrafts");
    if (!workBlock) return;

    // эти переменные определяют порядок крафта по бд, и т.к. крафты идут попорядку это делит крафты на пачки по времени
    let previousBP = 0;
    let row = 0;

    for (let i in works) {
        let dataRow = innerDate(works[i]);

        // сложный ид для того что бы групировать одинаковые итемы а не добавлять каждый раз новый
        let idRow;

        if (previousBP !== works[i].blueprint.id) {
            row++
        }
        previousBP = works[i].blueprint.id;

        if (dataRow.active) {
            idRow = "blueRow" + works[i].id
        } else {
            idRow = "blueRow" +
                works[i].blueprint.id +
                "" + works[i].mineral_saving_percentage +
                "" + works[i].time_saving_percentage +
                "" + row;
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
            $(blueRow).data("time", {finishTime: 0, startTime: Infinity});
        }

        blueRow.innerHTML = dataRow.html;
        $(blueRow).data("update", {update: true});

        if (!dataRow.active) {
            $(blueRow).data("count").count++;

            if ($(blueRow).data("time").finishTime < new Date(works[i].finish_time).getTime()) {
                $(blueRow).data("time").finishTime = new Date(works[i].finish_time).getTime()
            }

            if ($(blueRow).data("time").startTime > new Date(works[i].finish_time).getTime()) {
                $(blueRow).data("time").startTime = new Date(works[i].finish_time).getTime()
            }

            blueRow.innerHTML = dataRow.html;
            blueRow.innerHTML += "<div class='countBP' style='margin-right: 4px'>x" + $(blueRow).data("count").count + "</div>";

            blueRow.onclick = function () {
                lobby.send(JSON.stringify({
                    event: "SelectWork",
                    start_time: $(this).data("time").startTime,
                    to_time: $(this).data("time").finishTime,
                    blue_print_id: works[i].blueprint.id,
                    mineral_saving: works[i].mineral_saving_percentage,
                    time_saving: works[i].time_saving_percentage,
                    count: $(this).data("count").count,
                }));
            }
        } else {
            blueRow.onclick = function () {
                lobby.send(JSON.stringify({
                    event: "SelectWork",
                    id: works[i].id,
                }));
            }
        }
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

function SelectWork(jsonData) {
    fillHeadWorkbench(jsonData, false);

    // если jsonData.id > 0 то значит игрок выбрал активный крафт, и он модет быть только 1 и работаем по ид
    // TODO блок со временем отображать оставшееся время крафта

    let bpCountWork = document.getElementById("bpCountWork");
    bpCountWork.value = jsonData.count;
    if (jsonData.id > 0) {
        bpCountWork.max = 1;
    } else {
        bpCountWork.max = jsonData.max_count;
        bpCountWork.oninput = function () {
            lobby.send(JSON.stringify({
                event: "SelectWork",
                start_time: jsonData.start_time,
                to_time: jsonData.to_time,
                blue_print_id: jsonData.blue_print_id,
                mineral_saving: jsonData.mineral_saving,
                time_saving: jsonData.time_saving,
                count: Number(bpCountWork.value),
            }));
        };
    }

    let processButton = document.getElementById("processButton");
    processButton.value = "Отменить";
    processButton.onclick = function () {
        if (jsonData.id > 0) {
            lobby.send(JSON.stringify({
                event: "CancelCraft",
                id: jsonData.id,
            }));
        } else {
            lobby.send(JSON.stringify({
                event: "CancelCraft",
                start_time: jsonData.start_time,
                to_time: jsonData.to_time,
                blue_print_id: jsonData.blue_print_id,
                mineral_saving: jsonData.mineral_saving,
                time_saving: jsonData.time_saving,
                count: Number(bpCountWork.value),
            }));
        }
    };
}

function SelectBP(jsonData) {

    fillHeadWorkbench(jsonData, true);

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
    processButton.value = "Создать";

    processButton.onclick = function () {
        lobby.send(JSON.stringify({
            event: "Craft",
            storage_slot: Number(jsonData.storage_slot),
            count: Number(bpCountWork.value)
        }));
    };
}

function fillHeadWorkbench(jsonData, needMark) {
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

    fillNeedItems(jsonData.preview_recycle_slots, needMark);
}

function fillNeedItems(items, needMark) {
    let needItems = document.getElementById("needItems");
    $(needItems).empty();

    for (let i in items) {
        let cell = document.createElement("div");

        let item = Object.assign({}, items[i]);

        CreateInventoryCell(cell, item, i, "", onclick);
        let section = CheckRecycleSection(item, needItems);

        if (!item.find && needMark) {
            cell.style.border = "1px solid red";
            cell.innerHTML += "<div class='noAllowCell'></div>"
        }

        section.appendChild(cell);
    }
}