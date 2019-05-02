function FillNotifyBlock(data) {
    let notifyBlock = document.getElementById('notifyBlock');

    for (let i in data.missions) {
        addMission(notifyBlock, data.missions[i], null, 0, data.missions[i].uuid)
    }

    for (let i in data.notifys) {
        if (!document.getElementById(data.notifys[i].uuid)) {
            if (data.notifys[i].event === "new") {
                addMission(notifyBlock, data.notifys[i].data, "animation: new 1s linear 1", 1000, data.notifys[i].uuid)
            }
        }
    }
}

function addMission(notifyBlock, mission, animation, time, uuid) {
    notifyBlock.innerHTML += (
        `<div id="${uuid}" style="${animation}" onmouseover="document.getElementById('${mission.uuid}_tip').style.visibility = 'visible'"
                                                onmouseout="document.getElementById('${mission.uuid}_tip').style.visibility = 'hidden'">
                                                
           <div class="missionNotify" id="${mission.uuid}_tip">
                <h3>${mission.name}</h3>
                <h4>Заказчик глава ${mission.start_base.name} из сектора ${mission.start_map.Name}</h4>
                <p>${mission.start_dialog.pages[1].text}</p>
           </div> 
         </div>`
    );

    setTimeout(function () {
        document.getElementById(uuid).style.animation = "none"
    }, time)
}

function newNotify(notify) {
    if (notify && notify.name === "mission") {
        if (notify.event === "new") {
            addMission(document.getElementById('notifyBlock'), notify.data, "animation: new 1s linear 1", 1000, notify.uuid)
        }
        if (notify.event === "complete") {
            if (document.getElementById(notify.uuid)) {
                document.getElementById(notify.uuid).style.background = "#3dff00";
                document.getElementById(notify.uuid).onclick = function () {
                    this.remove()
                }
            }
        }
    }
}