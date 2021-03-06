let refreshSelectActiveWork;

// перменная отвечает за состояние обьекта, выбран чертеж или текущая работа
let workBenchState = null;
let bpSlot = null;
let bpCount = null;

function FillWorkbench(jsonData) {
    let bpBlock = document.getElementById("bluePrints");
    if (!bpBlock) return;

    let blueRows = $('.blueRow');
    blueRows.each(function () {
        if ($(this).data("count")) $(this).data("count").count = 0;
        if ($(this).data("time")) $(this).data("time").finishTime = 0;
        if ($(this).data("time")) $(this).data("time").startTime = Infinity;
    });

    for (let i in jsonData.storage.slots) {
        if (jsonData.storage.slots[i].type === "blueprints") {

            let bpSlot = jsonData.storage.slots[i];
            let time = getTimeWork(bpSlot.item.craft_time, true, jsonData.user_work_skill_time_percent);

            let blueRow = document.getElementById("blueRowBP" + i + bpSlot.item.id);
            if (!blueRow) {
                blueRow = document.createElement("div");
                blueRow.className = "blueRow";
                blueRow.id = "blueRowBP" + i + bpSlot.item.id;
                bpBlock.appendChild(blueRow);
                blueRow.innerHTML = `
                    <div class="bpIconItem">
                        ${getBackgroundUrlByItem(bpSlot)}
                    </div>
                    <div class='nameBP'>${bpSlot.item.name}</div>
                    <div class='countBP'>x${bpSlot.quantity}</div>
                    <div class='timerWork'>
                        <span> ${time.days}${time.hours} : ${time.minutes} : ${time.seconds} </span>
                        <div class='workTimeLine' style='width: 0'></div>
                    </div>           
                `;
            }

            $(blueRow).data("update", {update: true});
            $(blueRow).find(".countBP")[0].innerHTML = `x${bpSlot.quantity}`;
            $(blueRow).find(".timerWork")[0].innerHTML = `
                <span> ${time.days}${time.hours} : ${time.minutes} : ${time.seconds} </span>
                <div class='workTimeLine' style='width: 0'></div>
            `;

            blueRow.onclick = function () {
                clearInterval(refreshSelectActiveWork);
                refreshSelectActiveWork = null;
                lobby.send(JSON.stringify({
                    event: "SelectBP",
                    storage_slot: Number(i),
                    count: 1,
                }));
            };
        }
    }

    FillCurrentWorks(jsonData.blue_works);

    // удаляем всех кого не обновили
    blueRows.each(function () {
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
        if (!dataRow) continue;

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
                "" + works[i].mineral_tax_percentage +
                "" + works[i].time_tax_percentage +
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
            blueRow.innerHTML = dataRow.html;
        }

        $(blueRow).find(".timerWork")[0].innerHTML = `
                <span> ${dataRow.days}${dataRow.hours} : ${dataRow.minutes} : ${dataRow.seconds}</span>
                <div class='workTimeLine' style='width: ${dataRow.widthTimeLine}%'></div>
         `;

        $(blueRow).data("update", {update: true});
        $(blueRow).data("count").count++;
        $(blueRow).find(".countBP")[0].innerHTML = `x${$(blueRow).data("count").count}`;

        // формирование интервала времени для это строки
        // например в этой старке находятся 5 работ по интервалу от 05 до 15 минут
        if ($(blueRow).data("time").finishTime < new Date(works[i].finish_time).getTime()) {
            $(blueRow).data("time").finishTime = new Date(works[i].finish_time).getTime()
        }

        if ($(blueRow).data("time").startTime > new Date(works[i].finish_time).getTime()) {
            $(blueRow).data("time").startTime = new Date(works[i].finish_time).getTime()
        }

        blueRow.onclick = function () {
            clearInterval(refreshSelectActiveWork);
            refreshSelectActiveWork = null;

            lobby.send(JSON.stringify({
                event: "SelectWork",
                start_time: $(this).data("time").startTime,
                to_time: $(this).data("time").finishTime,
                blue_print_id: works[i].blueprint.id,
                mineral_saving: works[i].mineral_tax_percentage,
                time_saving: works[i].time_tax_percentage,
                count: $(this).data("count").count,
                id: works[i].id,
            }));
        }
    }
}

function innerDate(work) {

    let data = new Date();

    let finishTime = new Date(work.finish_time);
    let startTime = new Date(work.start_time);

    data.setTime(finishTime.getTime() - new Date().getTime());

    let percent = ((new Date() - startTime.getTime()) * 100) / (finishTime.getTime() - startTime.getTime());
    let widthTimeLine = 0;

    let days, hours, minutes, seconds;
    if (percent > 0) {
        widthTimeLine = Math.round(percent);

        days = (data.getUTCDate() - 1 > 0) ? data.getUTCDate() - 1 + "d: " : '';
        hours = (data.getUTCHours() > 9) ? data.getUTCHours() : "0" + data.getUTCHours();
        minutes = (data.getUTCMinutes() > 9) ? data.getUTCMinutes() : "0" + data.getUTCMinutes();
        seconds = (data.getUTCSeconds() > 9) ? data.getUTCSeconds() : "0" + data.getUTCSeconds();
    } else {
        let craftTime = (finishTime.getTime() - startTime.getTime()) / 1000;
        data.setTime(craftTime * 1000);

        days = (data.getUTCDate() - 1 > 0) ? data.getUTCDate() - 1 + "d: " : '';
        hours = (data.getUTCHours() > 9) ? data.getUTCHours() : "0" + data.getUTCHours();
        minutes = (data.getUTCMinutes() > 9) ? data.getUTCMinutes() : "0" + data.getUTCMinutes();
        seconds = (data.getUTCSeconds() > 9) ? data.getUTCSeconds() : "0" + data.getUTCSeconds();
    }

    if (widthTimeLine > 100 || data.getUTCDate() > 20) {
        return null
    }

    return {
        html: `
            <div class="bpIconItem">
                ${getBackgroundUrlByItem({type: work.blueprint.item_type, item: {name: work.item.name}})}
            </div>
            <div class='nameBP'>${work.item.name}</div>
            <div class='countBP'></div>
            <div class='timerWork'>
                <span>${days} ${hours} : ${minutes} : ${seconds}</span>
            <div class='workTimeLine' style='width: ${widthTimeLine}%'></div>
            </div>`,
        active: percent > 0,
        days: days,
        hours: hours,
        minutes: minutes,
        seconds: seconds,
        widthTimeLine: widthTimeLine,
    }
}

function SelectWork(jsonData) {

    if (jsonData.count === 0) {
        clearPreview();
        return
    }

    workBenchState = 'selectWork';
    fillHeadWorkbench(jsonData, false);

    if (jsonData.blue_work) {
        fillEfficiencyPanel(jsonData.blue_work.mineral_tax_percentage, jsonData.blue_work.time_tax_percentage);
    } else {
        fillEfficiencyPanel(jsonData.mineral_saving, jsonData.time_saving);
    }

    let bpCountWork = document.getElementById("bpCountWork");
    bpCountWork.value = jsonData.count;
    bpCountWork.max = jsonData.max_count;
    bpCountWork.oninput = function () {
        lobby.send(JSON.stringify({
            event: "SelectWork",
            id: jsonData.blue_work.id,
            count: Number(bpCountWork.value),
            start_time: jsonData.start_time,
            to_time: jsonData.to_time,
        }));
    };

    let time = getTimeWork(new Date(jsonData.blue_work.finish_time).getTime(), false);
    document.getElementById("bpCraftTime").innerHTML = time.days + time.hours + ":" + time.minutes + ":" + time.seconds + "s";

    if (!refreshSelectActiveWork) {
        refreshSelectActiveWork = setInterval(function () {
            lobby.send(JSON.stringify({
                event: "SelectWork",
                id: jsonData.blue_work.id,
                count: Number(bpCountWork.value),
                start_time: jsonData.start_time,
                to_time: jsonData.to_time,
            }));
        }, 1000);
    }


    let processButton = document.getElementById("processButton");
    processButton.value = "Отменить";
    processButton.onclick = function () {
        // TODO диалоговое окно подтверждения отмены работы с выводом инфы о том сколько хочет отменить работ юзер + сколько вернется ресурсов
        lobby.send(JSON.stringify({
            event: "CancelCraft",
            id: jsonData.id,
            count: Number(bpCountWork.value),
            start_time: jsonData.start_time,
            to_time: jsonData.to_time,
        }));
    };
}

function clearPreview() {
    document.getElementById("bpName").innerHTML = "";
    document.getElementById("bpIcon").style.backgroundImage = "none";
    document.getElementById("bpCraftTime").innerHTML = "";
    document.getElementById("processButton").value = "";
    document.getElementById("processButton").onclick = null;
    document.getElementById("itemPreview").style.backgroundImage = "none";
    document.getElementById("itemPreview").innerHTML = "";
    document.getElementById("bpCountWork").value = "";
    document.getElementById("needItems").innerHTML = "";
}

function getTimeWork(craftTime, full, tax) {
    let data = new Date();

    if (!tax) tax = 0;
    craftTime = craftTime - ((craftTime * tax) / 100);

    if (full) {
        data.setTime(craftTime * 1000);
    } else {
        data.setTime(craftTime - new Date().getTime());
    }

    let days = (data.getUTCDate() - 1 > 0) ? data.getUTCDate() - 1 + "d: " : '';
    let hours = (data.getUTCHours() > 9) ? data.getUTCHours() : "0" + data.getUTCHours();
    let minutes = (data.getUTCMinutes() > 9) ? data.getUTCMinutes() : "0" + data.getUTCMinutes();
    let seconds = (data.getUTCSeconds() > 9) ? data.getUTCSeconds() : "0" + data.getUTCSeconds();
    return {days: days, hours: hours, minutes: minutes, seconds: seconds}
}

function fillEfficiencyPanel(mineralEfficiency, timeEfficiency) {
    document.getElementById('mineralTaxSpan').innerHTML = mineralEfficiency + '%';
    if (mineralEfficiency < 0) {
        document.getElementById('mineralTaxSpan').style.color = '#00ff2d';
    } else if (mineralEfficiency === 0) {
        document.getElementById('mineralTaxSpan').style.color = '#00c2ff';
    }

    document.getElementById('timeTaxSpan').innerHTML = -timeEfficiency + '%';
    document.getElementById('timeTaxSpan').style.color = '#00ff2d';
}

function SelectBP(jsonData) {
    if (!jsonData.blue_print) {
        clearPreview();
        return
    }

    workBenchState = 'selectBP';

    bpSlot = Number(jsonData.storage_slot);
    bpCount = Number(jsonData.count);

    let mineralEfficiency = 100 + (jsonData.base.efficiency - jsonData.user_work_skill_detail_percent);
    fillEfficiencyPanel(mineralEfficiency, jsonData.user_work_skill_time_percent);
    fillHeadWorkbench(jsonData, true);
    clearInterval(refreshSelectActiveWork);
    refreshSelectActiveWork = null;

    let time = getTimeWork(jsonData.blue_print.craft_time * jsonData.count, true, jsonData.user_work_skill_time_percent);
    document.getElementById("bpCraftTime").innerHTML = time.days + time.hours + ":" + time.minutes + ":" + time.seconds + "s";

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

        // let dialog = document.createElement("div");
        // dialog.className = "veliriDialog";
        // dialog.innerHTML = `
        // <p>
        //     Произвести ${jsonData.bp_item.name} в кол-ве ${bpCountWork.value} <br>
        //     С налогом на минералы ${mineralEfficiency}?
        // </p>
        // <input type="button" value="Создать">
        // <input type="button" value="Отмена">
        // `;
        //document.body.appendChild(dialog)

        //todo диалоговое окно с указанием процентов налога и отправлять их на бек, на беке проверять если проценты
        // совпадают то крафтим если нет то выплевывается алерт что налог извенился в виде диалога
        lobby.send(JSON.stringify({
            event: "Craft",
            storage_slot: Number(jsonData.storage_slot),
            count: Number(bpCountWork.value),
            user_work_skill_detail_percent: jsonData.user_work_skill_detail_percent,
            user_work_skill_time_percent: jsonData.user_work_skill_time_percent,
            efficiency: jsonData.base.efficiency,
        }));
    };
}

function fillHeadWorkbench(jsonData, needMark) {

    if ($('#bpName').text() !== jsonData.blue_print.name) {
        document.getElementById("bpName").innerHTML = jsonData.blue_print.name;
        document.getElementById("bpIcon").style.backgroundImage = "url(/assets/blueprints/" + jsonData.blue_print.icon + ".png)";

        let itemPreview = document.getElementById("itemPreview");
        itemPreview.innerHTML = `
        ${getBackgroundUrlByItem({type: jsonData.blue_print.item_type, item: {name: jsonData.bp_item.name}})}
            <span style="z-index: 1">x${jsonData.blue_print.count}</span>
        `;
    }

    fillNeedItems(jsonData.preview_recycle_slots, needMark);
}

function fillNeedItems(items, needMark) {
    let needItemsBlock = document.getElementById("needItems");

    for (let i in items) {

        let item = Object.assign({}, items[i]);
        let cell = document.getElementById("workbenchNeedItem" + items[i].type + items[i].item.id);

        if (!cell) {

            cell = document.createElement("div");
            cell.id = "workbenchNeedItem" + items[i].type + items[i].item.id;
            CreateInventoryCell(cell, item, i, "", onclick);

            let section = CheckRecycleSection(item, needItemsBlock);
            section.appendChild(cell);

        } else {
            $(cell).find('.QuantityItems')[0].innerHTML = item.quantity;
        }

        if (!item.find && needMark) {
            cell.style.outline = "1px solid red";
            if ($(cell).find('.noAllowCell').length === 0) cell.innerHTML += "<div class='noAllowCell'></div>"
        } else {
            cell.style.outline = "unset";
            if ($(cell).find('.noAllowCell').length > 0) $(cell).find('.noAllowCell')[0].remove()
        }

        $(cell).data("update", {update: true, type: item.type});
    }

    let needItems = $(needItemsBlock).find('.InventoryCell');
    needItems.each(function () {
        if ($(this).data("update").update) {
            $(this).data("update").update = false;
        } else {
            $(this).remove()
        }
    });

    let recycleSections = $(needItemsBlock).find('.RecycleSection');
    recycleSections.each(function () {
        if (this.childElementCount === 1) $(this).remove()
    });
}