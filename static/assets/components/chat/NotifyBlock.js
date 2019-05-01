function FillNotifyBlock(data) {
    let notifyBlock = document.getElementById('notifyBlock');

    for (let i in data.missions) {
        addMission(notifyBlock, data.missions[i])
    }
}

function addMission(notifyBlock, mission) {
    notifyBlock.innerHTML += (
        `<div onmouseover="document.getElementById('${mission.uuid}').style.visibility = 'visible'" onmouseout="document.getElementById('${mission.uuid}').style.visibility = 'hidden'">
           <div class="missionNotify" id="${mission.uuid}">
                <h3>${mission.name}</h3>
                <h4>Заказчик глава ${mission.start_base.name} из сектора ${mission.start_map.Name}</h4>
                <p>${mission.start_dialog.pages[1].text}</p>
           </div> 
         </div>`
    );
    console.log(mission)
}