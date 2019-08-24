let filters = {
    id: 0,
    name: "",
    fraction: "",
    type: "",
    access: "",
};

function filterID(context) {
    filters.id = Number(context.value);
    GetListDialogs();
}

function filterName(context) {
    filters.name = context.value;
    GetListDialogs();
}

function filterFraction(context) {
    filters.fraction = context.value;
    GetListDialogs();
}

function filterType(context) {
    filters.type = context.value;
    GetListDialogs();
}

function filterAccess(context) {
    filters.access = context.value;
    GetListDialogs();
}

function getMissionDialogs(context) {
    // особая функция которая запрашивает у сервера ток те диалоги которые принадлежат к мисии
    editor.send(JSON.stringify({
        event: "GetMissionDialog",
        id: Number(context.value),
    }));
}

function CreateDialogList(dialogs) {
    selectDialog = {};

    let dialogList = document.getElementById("dialogList");
    dialogList.style.display = "block";

    document.getElementById("selectDialog").style.display = "none";
    document.getElementById("MissionList").style.display = "none";

    let dialogList2 = document.getElementById("dialogList2");
    dialogList2.innerHTML = '';

    for (let i in dialogs) {

        let dialog = dialogs[i];

        // проверка на фильтры
        if (!(dialog.id === filters.id || filters.id === 0)) {
            continue
        }

        if (!(dialog.name.indexOf(filters.name) + 1 || filters.name === '')) {
            continue
        }

        if (!(dialog.fraction === filters.fraction || filters.fraction === '')) {
            continue
        }

        if (!(dialog.type === filters.type || filters.type === '')) {
            continue
        }

        if (!(dialog.access_type === filters.access || filters.access === '')) {
            continue
        }

        if (!dialog.pages) {
            continue
        }

        let startPage = dialog.pages[1];

        dialogList2.innerHTML += `
            <div id="${dialog.id}" class="DepartmentOfEmployment">
                
                <h3 style="position: absolute;right: 7px;top: -17px;color: red;">
                    ID: ${dialog.id}
                </h3>
                <h3 class="missionHead" id="missionHead${dialog.id}" style="width: 315px;">
                    ${dialog.name} (${dialog.fraction}, ${dialog.access_type})
                </h3>
                
                <div class="infoBlock" id="infoBlock${dialog.id}">
                    <div class="missionText" id="missionText${dialog.id}">${startPage.text}</div>
                    <div class="missionAsc" id="missionAsc${dialog.id}"></div>
                </div>
    
                <div class="rewardBlock" id="rewardBlock${dialog.id}">
                    <div class="missionFace" id="missionFace${dialog.id}"></div>
    
                    <div class="rewardBlock2" id="rewardBlock2${dialog.id}">
                        <h3>Награда:</h3>
                        <div class="rewards" id="rewards${dialog.id}">
                           <div class="rewardsCredits" id="rewardsCredits${dialog.id}">Крудиты: <span id="countRewardCredits${dialog.id}">250</span></div>
                           <div class="rewardsItems" id="rewardsItems${dialog.id}"></div>
                        </div>
                    </div>
                </div>
                
                <div class="actions">
                    <input type="button" value="Изменить" style="margin-left: 5px; float: left" onclick="SelectDialog(${dialog.id})"> 
                    <input type="button" value="Удалить" style="margin-right: 5px; float: right" onclick="DeleteDialog(${dialog.id})">
                </div> 
            </div>
        `;

        setTimeout(function () {

            let missionAsc = document.getElementById("missionAsc" + dialog.id);
            for (let j in startPage.asc) {
                let asc = startPage.asc[j];
                missionAsc.innerHTML += `<div class='asks'><div class='wrapperAsk' id='ask"${asc.id}'>${asc.text}</div></div>`
            }

            let missionFace = document.getElementById("missionFace" + dialog.id);

            GetDialogPicture(startPage.id, -1).then(function (response) {
                // console.log(response.data)
                if (response.data.hasOwnProperty('main')) {
                    missionFace.style.backgroundImage = "url('" + response.data['main'] + "')";
                } else {
                    for (let i in response.data) {
                        let picture = response.data[i];

                        missionFace.innerHTML += `
                            <div class="innerPicture ${i}" style="background-image: url('${picture}')"></div>
                        `

                    }
                }
            });
        }, 200)
    }
}