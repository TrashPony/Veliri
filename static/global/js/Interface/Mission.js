let hideMission = false;
let hideStory = true;

let allUserMission = {};
let selectedMission = "";

setInterval(function () {
    global.send(JSON.stringify({
        event: "GetMissions",
    }));
}, 1000);

function HideMission() {
    if (hideMission) {
        hideMission = false;
        document.getElementById("MissionPoints").style.display = "block";
        document.getElementById("MissionInfoBlock").style.height = "200px";
        document.getElementById("hideMissionButton").value = "Скрыть";
    } else {
        hideMission = true;
        hideStory = true;
        document.getElementById("MissionPoints").style.display = "none";
        document.getElementById("MissionStory").style.display = "none";
        document.getElementById("MissionInfoBlock").style.height = "41px";
        document.getElementById("hideMissionButton").value = "Показать";
    }
}

function DetailMission() {
    if (hideStory) {
        hideStory = false;
        document.getElementById("MissionStory").style.display = "block";
    } else {
        hideStory = true;
        document.getElementById("MissionStory").style.display = "none";
    }
}

function SelectMission(context) {
    global.send(JSON.stringify({
        event: "SelectMission",
        mission_uuid: context.value,
    }));
}

function FillMissionsSelect(allMissions, selectUUIDMission) {
    let selectMission = $('#SelectMission');

    // список поменялся, обновляем селект
    if (allMissions !== allUserMission) {
        allUserMission = allMissions;
        selectMission.html(` <option value="">Выбраное задание</option>`);

        for (let key in allMissions) {
            if (allMissions.hasOwnProperty(key)) {
                selectMission.append(`<option value="${key}">${allMissions[key].name}</option>`);
            }
        }
        selectMission.val(selectUUIDMission);
    }

    // выбраное задание изменилось проверяем селект
    if (selectedMission !== selectUUIDMission) {
        selectedMission = selectUUIDMission;
        selectMission.val(selectUUIDMission);
    }

    FillMissionPoints(selectUUIDMission)
}

function FillMissionPoints(selectUUIDMission) {
    let missionPointsBlock = document.getElementById('MissionPoints');
    let userMission = allUserMission[selectUUIDMission];

    missionPointsBlock.innerHTML = ``;
    let tableTask = document.createElement("table");
    missionPointsBlock.appendChild(tableTask);

    if (!userMission) {
        if (game && game.squad) game.squad.missionMove = null;
        return;
    }

    let actionSort = [];
    for (let i in userMission.actions) {
        // сортируем так что бы в масив не попали не открытые участки задания,
        // понятно что от сюда их можно наснифить но это не критично
        let append = false;
        if (userMission.actions[i].number === 1 || userMission.actions[i].complete) {
            actionSort[userMission.actions[i].number] = userMission.actions[i];
        } else {
            for (let j in userMission.actions) {
                // если впереди есть выполненые таски то прошлые даже не выполненые можно добавить
                if (userMission.actions[j].number > userMission.actions[i].number && userMission.actions[j].complete) {
                    append = true
                }

                // если текущая не выполнена, а позади да, то добавляем
                if (userMission.actions[j].number === userMission.actions[i].number - 1 && userMission.actions[j].complete) {
                    append = true
                }
            }
        }

        if (append) actionSort[userMission.actions[i].number] = userMission.actions[i];
    }

    // актуальное действие всегда самое ближнее к 0лю не выполненое дейтсиве, от этого зависит отображение подсказок (например линия  на мине карте)
    let actionAction = null;
    for (let i in actionSort) {
        if (actionSort[i]) {

            if (!actionSort[i].complete && !actionAction) actionAction = actionSort[i];
            // TODO подробное описние под каждым пунктом шрифтом поменьше 
            tableTask.innerHTML += `
                    <tr class="missionActions"> 
                        <td class="actionNumber">${actionSort[i].number}.</td>
                        <td class="actionShortDescription">${actionSort[i].short_description}</td>
                        <td><div class="actionComplete${actionSort[i].complete}"></div></td>
                    </tr>
                `;
        }
    }

    MissionHelpers(actionAction)
}

function MissionHelpers(actionAction) {
    if (actionAction.type_func_monitor === "to_sector") {
        // запрашиваем у сервера координату тунеля куда надо двигатся, отрисуется на мине карте
        global.send(JSON.stringify({
            event: "GetPortalPointToGlobalPath",
            name: "mission",
            map_id: actionAction.map_id,
        }));
    } else if (actionAction.type_func_monitor === "to_q_r") {
        let {x, y} = GetXYCenterHex(actionAction.q, actionAction.r);
        if (game && game.squad) game.squad.missionMove = {x: x, y: y, radius: actionAction.radius}
    } else {
        if (game && game.squad) game.squad.missionMove = null;
    }
}

function FillMissionDetail() {

}