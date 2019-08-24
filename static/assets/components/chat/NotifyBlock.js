function FillNotifyBlock(data) {
    for (let i in data.missions) {
        let tip = NotifyPin(data.missions[i].uuid);
        addMission(tip, data.missions[i]);
    }

    for (let i in data.notifys) {
        if (!document.getElementById(data.notifys[i].uuid)) {
            if (data.notifys[i].event === "new") {
                let tip = NotifyPin(data.missions[i].uuid);
                addMission(tip, data.notifys[i].data);
            }
        }
    }
}

function newNotify(notify) {

    if (document.getElementById(notify.uuid)) document.getElementById(notify.uuid).remove();

    if (notify && notify.name === "mission") {

        if (notify.event === "new") {
            let tip = NotifyPin(notify.uuid);
            addMission(tip, notify.data);
        }

        if (notify.event === "complete") {
            let tip = NotifyPin(notify.uuid);
            addMission(tip, notify.data);
            tip.style.background = "#3dff00";
            tip.onclick = function () {
                // todo запрос на бекенд для удаления.
                this.remove()
            };
        }
    }

    if (notify && notify.name === "craft") {
        if (notify.event === "complete") {
            let tip = NotifyPin(notify.uuid);
            document.getElementById(notify.uuid).style.background = "#00fffe";
            addCompleteCraft(tip, notify)
        }
    }

    if (notify && (notify.name === "sell" || notify.name === "buy")) {
        let tip = NotifyPin(notify.uuid);
        addDeal(tip, notify)
    }
}

function NotifyPin(uuid) {
    let notifyWrap = document.createElement("div");
    document.getElementById('notifyBlock').appendChild(notifyWrap);

    notifyWrap.id = uuid;
    notifyWrap.style.animation = "new 1s linear 1";

    notifyWrap.onmouseover = function () {
        if (document.getElementById(uuid + '_tip')) document.getElementById(uuid + '_tip').style.visibility = 'visible';
    };
    notifyWrap.onmouseout = function () {
        if (document.getElementById(uuid + '_tip')) document.getElementById(uuid + '_tip').style.visibility = 'hidden';
    };
    notifyWrap.onclick = function () {
        chat.send(JSON.stringify({
            event: "DeleteNotify",
            uuid: uuid,
        }));
    };

    setTimeout(function () {
        notifyWrap.style.animation = "none"
    }, 1000);

    let notifyTip = document.createElement("div");
    notifyTip.className = "missionNotify";
    notifyTip.id = uuid + '_tip';
    notifyWrap.appendChild(notifyTip);

    return notifyTip;
}

function addMission(notifyTip, mission) {
    notifyTip.innerHTML = ` 
                <h3>${mission.name}</h3>
                <h4>Заказчик глава ${mission.start_base.name} из сектора ${mission.start_map.Name}</h4>
                <p>${mission.start_dialog.pages[1].text}</p>
`;
}

function addCompleteCraft(notifyTip, notify) {
    notifyTip.innerHTML = `
        <h3>Завершено производство</h3>
        <div class="notifyParagraph">
            На базе <span class="importantly">${notify.base.name}</span> 
            что в секторе <span class="importantly">${notify.map.Name}</span>
            завершено производство <div class="notifyIconItem">${getBackgroundUrlByItem(notify.item)}</div><span class="importantly">${notify.item.item.name}</span>
            в количестве <span class="importantly">${notify.item.quantity}</span> единиц.
            <p> Созданные вещи ожидают вас на складе базы <span class="importantly">${notify.base.name}</span> </p>
        </div>
    `
}

function addDeal(notifyTip, notify) {

    let type = [];

    if (notify.name === 'sell') {
        type = ['продажа', 'продано', ''];
    } else if (notify.name === "buy") {
        type = ['покупка', 'куплено', 'Купленные вещи ожидают вас на складе базы <span class="importantly">' + notify.base.name + '</span>'];
    }

    notifyTip.innerHTML = `
        <h3>Совершена ${type[0]}</h3>
        <div class="notifyParagraph">
            На базе <span class="importantly">${notify.base.name}</span> 
            что в секторе <span class="importantly">${notify.map.Name}</span>
            было ${type[1]} <div class="notifyIconItem">${getBackgroundUrlByItem(notify.item)}</div><span class="importantly">${notify.item.item.name}</span>
            в количестве <span class="importantly">${notify.item.quantity}</span> единиц, по цене <span class="importantly">${notify.price}</span><span class="cr">cr</span> за шт. 
            <p>Общая стоимость сделки составила  <span class="importantly">${notify.item.quantity * notify.price}</span> <span class="cr">cr</span>.</p>
            <p>${type[2]}</p>  
        </div>
    `;

    document.getElementById(notify.uuid).style.background = "#cd00ff";
}