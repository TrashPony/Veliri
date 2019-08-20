function MissionList(missions) {

    let MissionList = document.getElementById("MissionList");
    MissionList.style.display = "block";
    let MissionList2 = document.getElementById("MissionList2");

    document.getElementById("selectDialog").style.display = "none";
    document.getElementById("dialogList").style.display = "none";

    MissionList2.innerHTML = '';
    console.log(missions);

    //          миссиям не нужена доп страница для редактирования т.к. там мало информации
    //             -- название, тип

    for (let i in missions) {
        if (missions.hasOwnProperty(i)) {

            let mission = missions[i];

            MissionList2.innerHTML += `
<div class="mission">
    <div class="missionProp">
        <h4>${mission.name}</h4>
        
        <label> Доступно фракции:
            <select id="fractionMiss${mission.id}" onchange="">
                        <option value="All">Всем</option>
                        <option value="Replics">Replics</option>
                        <option value="Explores">Explores</option>
                        <option value="Reverses">Reverses</option>
            </select>
        </label>
        
        <label> Тип:
             <select id="typeMiss${mission.id}" onchange="">
                        <option value="">-</option>
                        <option value="delivery">delivery</option>
            </select>
        </label>
        
        <label> Награда кредитов: 
            <input type="number" value="${mission.reward_cr}">
        </label>
        
        <label> Награда предметы:
            <div class="rewardItems" id="rewardItems${mission.id}"></div>
        </label>
        
        <label> Ид базы начала квеста (0 - на всех): 
            <input type="number" value="${mission.start_base_id}">
        </label>
        
        <label> Ид диалога для старта задания: 
            <input type="number" value="${mission.start_dialog_id}">
        </label>
        
    </div>
    
    <div class="actionsProp" id="actionsProp${mission.id}">
    
    </div>
</div>
`;

            setTimeout(function () {
                $('#fractionMiss' + mission.id).val(mission.fraction);
                $('#typeMiss' + mission.id).val(mission.type);
                ItemFill('rewardItems' + mission.id, mission.reward_items.slots);
                ActionFill(mission);
            }, 100)
        }
    }
}

function ActionFill(mission) {

    let actionsBlock = document.getElementById('actionsProp' + mission.id);
    actionsBlock.innerHTML = "<h4>Действия</h4>";

    let count = 0;
    for (let i in mission.actions) {
        let action = mission.actions[i];
        count++;

        console.log(action);
        actionsBlock.innerHTML += `
            <div class="rowAction">
            
                <span class="count">${count}.</span>
                <label> Описание:
                    <input type="text" value="${action.description}">
                </label>
                
                <label> Краткое:
                    <input type="text" value="${action.short_description}">
                </label>
                
                <label> Действие:
                    <select id="actionType${action.id}" onchange="">
                        <option value="delivery_item">Доставить предмет</option>
                        <option value="get_item_on_base">Взять предмет на базе</option>
                        <option value="to_q_r">Достигнуть точки Q R</option>
                        <option value="dialog">Поговорить</option>
                    </select>
                </label>
                
                <label id="async${action.id}"> Не последовательное: 
                    <input type="checkbox">
                </label>
                
                <div class="metaBlock" id="metaBlock${action.id}">
                    <h5>Мета информация действия</h5>
                </div>
            </div>
        `;

        setTimeout(function () {
            $('#actionType' + action.id).val(action.type_func_monitor);
            $('#async' + action.id).prop('checked', action.async);
        }, 100)
    }
}

function ItemFill(parentID, slots) {
    for (let i in slots) {
        let cell = document.createElement("div");
        CreateInventoryCell(cell, slots[i], i, "");
        document.getElementById(parentID).appendChild(cell);
    }
}