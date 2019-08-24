let allMission = {};

let missFilters = {
    id: 0,
    name: "",
    fraction: "",
    type: "",
    access: "",
};

function filterMissID(context) {
    missFilters.id = Number(context.value);
    GetListMissions();
}

function filterMissName(context) {
    missFilters.name = context.value;
    GetListMissions();
}

function MissionList(missions) {
    allMission = missions;

    let MissionList = document.getElementById("MissionList");
    MissionList.style.display = "block";
    let MissionList2 = document.getElementById("MissionList2");

    document.getElementById("selectDialog").style.display = "none";
    document.getElementById("dialogList").style.display = "none";

    for (let i in missions) {
        if (missions.hasOwnProperty(i)) {

            let mission = missions[i];
            let missBlock = document.getElementById("mission" + mission.id);

            // проверка на фильтры
            if (!(mission.id === missFilters.id || missFilters.id === 0)) {
                if (missBlock) missBlock.remove();
                continue
            }

            if (!(mission.name.indexOf(missFilters.name) + 1 || missFilters.name === '')) {
                if (missBlock) missBlock.remove();
                continue
            }


            if (document.getElementById("mission" + mission.id)) {
                $('#missionName' + mission.id).val(mission.name);
                $('#fractionMiss' + mission.id).val(mission.fraction);
                $('#typeMiss' + mission.id).val(mission.type);
                $('#missionRewardCR' + mission.id).val(mission.reward_cr);
                $('#missionStartBase' + mission.id).val(mission.start_base_id);
                $('#missionStartDialogID' + mission.id).val(mission.start_dialog_id);

                if (mission.reward_items) ItemFill('rewardItems' + mission.id, mission.reward_items.slots, false, mission.id);

                ActionFill(mission);
            } else {
                MissionList2.innerHTML += `
                    <div class="mission" id="mission${mission.id}">
                        <div class="missionProp">
                            <input id="missionName${mission.id}" type="text" value="${mission.name}" style="display: block;width: 320px; margin-bottom: 5px;" oninput="SetMissionName(this, ${mission.id})">
                            <input type="button" value="Сохранить" onclick="SaveMission(${mission.id})" style="width: 85px;">  
                            <input type="button" value="Удалить" onclick="DeleteMission(${mission.id})" style="width: 85px;">  

                            <label> ID: 
                                <input type="text" value="${mission.id}" disabled>
                            </label>
                            
                            <label> Доступно фракции:
                                <select id="fractionMiss${mission.id}" onchange="SetFraction(this, ${mission.id})">
                                            <option value="All">Всем</option>
                                            <option value="Replics">Replics</option>
                                            <option value="Explores">Explores</option>
                                            <option value="Reverses">Reverses</option>
                                </select>
                            </label>
                            
                            <label> Тип:
                                 <select id="typeMiss${mission.id}" onchange="SetTypeMission(this, ${mission.id})">
                                            <option value="">-</option>
                                            <option value="delivery">delivery</option>
                                </select>
                            </label>
                            
                            <label> Награда кредитов: 
                                <input id="missionRewardCR${mission.id}" type="number" value="${mission.reward_cr}" oninput="SetRewardCr(this, ${mission.id})">
                            </label>
                            
                            <label> Награда предметы:
                                <div class="rewardItems" id="rewardItems${mission.id}"></div>
                            </label>
                            
                            <label> Ид базы начала квеста (0 - на всех): 
                                <input id="missionStartBase${mission.id}" type="number" value="${mission.start_base_id}" oninput="SetBaseStart(this, ${mission.id})">
                            </label>
                            
                            <label> Ид диалога для старта задания: 
                                <input id="missionStartDialogID${mission.id}" type="number" value="${mission.start_dialog_id}" oninput="SetStartDialog(this, ${mission.id})">
                            </label>
                            
                            <label> Ид диалога не выполненого задания: 
                                <input id="not_finished_dialog_id${mission.id}" type="number" value="${mission.not_finished_dialog_id}" oninput="SetNotFinishedDialog(this, ${mission.id})">
                            </label>
                          
                        </div>
                        
                        <div class="actionsProp" id="actionsProp${mission.id}">
                        
                        </div>
                    </div>
                    `;

                setTimeout(function () {
                    $('#fractionMiss' + mission.id).val(mission.fraction);
                    $('#typeMiss' + mission.id).val(mission.type);
                    if (mission.reward_items) ItemFill('rewardItems' + mission.id, mission.reward_items.slots, false, mission.id);
                    ActionFill(mission);
                }, 100)
            }
        }
    }
}

function ActionFill(mission) {

    let actionsBlock = document.getElementById('actionsProp' + mission.id);
    actionsBlock.innerHTML = "";


    for (let i in mission.actions) {
        let action = mission.actions[i];

        actionsBlock.innerHTML += `
            <div class="rowAction" data-mission = "${mission.id}">
            
                <input type="number" class="count" value="${action.number}" style="width: 40px;" oninput="SetNumberAction(this, ${action.id})">
                <input type="button" value="Удалить" onclick="removeAction(${mission.id}, ${action.id})" style="position: absolute;left: 9px;top: 61px;">
               
                <div style="text-align: left;padding-left: 60px;">
                    <label> Описание:
                        <input style="width: 300px;" type="text" value="${action.description}" oninput="SetDescriptionAction(this, ${action.id})">
                    </label>
                    
                    <label> Краткое:
                        <input type="text" value="${action.short_description}" oninput="SetShortDescriptionAction(this, ${action.id})">
                    </label>
                    <br>
                    <label> Действие:
                        <select id="actionType${action.id}" onchange="SetTypeMonitorAction(this, ${action.id})">
                            <option value="delivery_item">Доставить предмет на базу</option>
                            <option value="get_item_on_base">Взять предмет на базе</option>
                            <option value="to_q_r">Достигнуть точки Q R</option>
                            <option value="talk_with_base">Поговорить с базой</option>
                            <option value="extract_item">Добыть предметы</option>
                            <option value="get_item_on_obj">Взять предмет из объекта</option>
                            <option value="place_item_in_obj">Положить предмет в обьект на карте</option>
                            <option value="attack_map_obj">Стрельнуть по объекту на карте</option>
                        </select>
                    </label>
                    <br>
                    <label id="async${action.id}"> Async: 
                        <input type="checkbox" onchange="SetAsyncAction(this, ${action.id})">
                    </label>
                    <textarea placeholder="Сообщение по завершению" oninput="SetEndTextAction(this, ${action.id})"
                    style="position: absolute;right: 0;top: 0;width: 460px;height: 80px;">${action.end_text}</textarea>
                </div>
                
                <div class="metaBlock" id="metaBlock${action.id}">
                    <h5>Мета информация действия</h5>
                    
                    <table>
                        <tr>
                            <td>Ид базы</td>
                            <td>Ид Карты</td>
                            <td>Q</td>
                            <td>R</td>
                            <td>Радиус</td>
                            <td>Секунды</td>
                            <td>Количество</td>
                            <td>Ид Диалога</td>
                            <td>Ид Алтернативного диалога</td>
                            <td>Необходимые предметы</td>
                            <td>Должен положить игрок</td>
                        </tr>
                        <tr>
                            <td><input type="number" value="${action.base_id}" oninput="SetActionBaseID(this, ${action.id})"></td>
                            <td><input type="number" value="${action.map_id}" oninput="SetActionMapID(this, ${action.id})"></td>
                            <td><input type="number" value="${action.q}" oninput="SetActionQ(this, ${action.id})"></td>
                            <td><input type="number" value="${action.r}" oninput="SetActionR(this, ${action.id})"></td>
                            <td><input type="number" value="${action.radius}" oninput="SetActionRadius(this, ${action.id})"></td>
                            <td><input type="number" value="${action.sec}" oninput="SetActionSec(this, ${action.id})"></td>
                            <td><input type="number" value="${action.count}" oninput="SetActionCount(this, ${action.id})"></td>
                            <td><input type="number" value="${action.dialog_id}" oninput="SetActionDialogID(this, ${action.id})"></td>
                            <td><input type="number" value="${action.alternative_dialog_id}" oninput="SetAltDialogID(this, ${action.id})"></td>
                            <td>
                                <div class="needItems" id="needItemsAction${action.id}">
                                </div>
                            </td>
                            <td><input id="ownerPlace${action.id}" type="checkbox" onchange="SetOwnerPlace(this, ${action.id})"></td>
                        </tr>
                    </table>
                </div>
            </div>
        `;

        setTimeout(function () {
            $('#actionType' + action.id).val(action.type_func_monitor);
            $('#async' + action.id).prop('checked', action.async);
            $('#ownerPlace' + action.id).prop('checked', action.owner_place);
            if (action.need_items) ItemFill('needItemsAction' + action.id, action.need_items.slots, true, action.id);
        }, 100)
    }

    actionsBlock.innerHTML += `
            <div class="rowAction" onclick="AddAction(${mission.id})">
            Добавить действие +
            </div>
    `;

    setTimeout(function () {
        let rows = $('*[data-mission="' + mission.id + '"]');

        rows.sort(function (a, b) {
            a = $(a).find('.count').val();
            b = $(b).find('.count').val();

            return a - b;
        });

        rows.appendTo(actionsBlock);
    }, 100)
}

function ItemFill(parentID, slots, need, id) {
    let parent = document.getElementById(parentID);
    if (parent) {

        parent.innerHTML = '';

        for (let i in slots) {

            let cell = document.createElement("div");
            CreateInventoryCell(cell, slots[i], i, "");
            parent.appendChild(cell);

            cell.onclick = function () {
                if (document.getElementById('RemoveItemDialog')) document.getElementById('RemoveItemDialog').remove();

                let dialog = document.createElement("div");
                dialog.id = 'RemoveItemDialog';
                dialog.innerHTML = `
                    <h4> Удалить предмет? </h4>
                    <input type="button" value="Да" onclick="RemoveItem(${i}, ${id}, ${need})">
                    <input type="button" value="Нет" onclick="document.getElementById('RemoveItemDialog').remove();">
                `;
                document.body.appendChild(dialog);
            };
        }
    }

    let addItem = document.createElement("div");
    addItem.className = "addItemButton";
    addItem.innerHTML = "+";
    addItem.onclick = function () {
        if (document.getElementById('AddItemDialog')) document.getElementById('AddItemDialog').remove();

        let dialog = document.createElement("div");
        dialog.id = 'AddItemDialog';
        dialog.innerHTML = `
                    <h4> Добавить предмет. </h4>
                    
                    <select id="addItemType">
                        <option value="weapon"> Оружие </option>
                        <option value="ammo"> Патроны </option>
                        <option value="equip"> Снаряжение </option>
                        <option value="body"> Корпус </option>
                        <option value="resource"> Ресурс </option>
                        <option value="recycle"> Полуфабрикат </option>
                        <option value="detail"> Деталь </option>
                        <option value="boxes"> Ящик </option>
                        <option value="blueprints"> Чертеж </option>
                        <option value="trash"> Хлам </option>
                    </select>
                    <br>
                        <input type="number" placeholder="ID" id="addItemID">
                    <br>
                        <input type="number" placeholder="quantity" id="addItemQuantity">
                    <br>
                    <input type="button" value="Да" onclick="AddItem(${id}, ${need})">
                    <input type="button" value="Нет" onclick="document.getElementById('AddItemDialog').remove();">
                `;
        document.body.appendChild(dialog);
    };

    setTimeout(function () {
        parent.appendChild(addItem);
    }, 200)
}

function AddItem(id, need) {
    //если need то это needItems иначе это награда миссии
    if (Boolean(need)) {
        editor.send(JSON.stringify({
            event: "AddActionItem",
            id: Number(id),
            item_type: document.getElementById("addItemType").value,
            item_id: Number(document.getElementById("addItemID").value),
            item_quantity: Number(document.getElementById("addItemQuantity").value),
        }));
    } else {
        editor.send(JSON.stringify({
            event: "AddMissionRewardItem",
            id: Number(id),
            item_type: document.getElementById("addItemType").value,
            item_id: Number(document.getElementById("addItemID").value),
            item_quantity: Number(document.getElementById("addItemQuantity").value),
        }));
    }

    if (document.getElementById('AddItemDialog')) document.getElementById('AddItemDialog').remove();
}

function RemoveItem(itemSlot, id, need) {
    if (document.getElementById('RemoveItemDialog')) document.getElementById('RemoveItemDialog').remove();

    //если need то это needItems иначе это награда миссии
    if (Boolean(need)) {
        editor.send(JSON.stringify({
            event: "RemoveActionItem",
            slot: Number(itemSlot),
            id: Number(id),
        }));
    } else {
        editor.send(JSON.stringify({
            event: "RemoveMissionRewardItem",
            slot: Number(itemSlot),
            id: Number(id),
        }));
    }
}