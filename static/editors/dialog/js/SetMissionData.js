/**
 *   Основная информация по заданию
 **/

function getMissionByID(id) {
    for (let i in allMission) {
        if (allMission.hasOwnProperty(i) && allMission[i].id === id) {
            return allMission[i]
        }
    }
}

function SetMissionName(context, id) {
    let mission = getMissionByID(id);
    mission.name = context.value;
}

function SetFraction(context, id) {
    let mission = getMissionByID(id);
    mission.fraction = context.value;
}

function SetTypeMission(context, id) {
    let mission = getMissionByID(id);
    mission.type = context.value;
}

function SetRewardCr(context, id) {
    let mission = getMissionByID(id);
    mission.reward_cr = Number(context.value);
}

function SetBaseStart(context, id) {
    let mission = getMissionByID(id);
    mission.start_base_id = Number(context.value);
}

function SetStartDialog(context, id) {
    let mission = getMissionByID(id);
    mission.start_dialog_id = Number(context.value);
}

function SetNotFinishedDialog(context, id) {
    let mission = getMissionByID(id);
    mission.not_finished_dialog_id = Number(context.value);
}

function AddAction(id) {
    let mission = getMissionByID(id);
    let number = 1;
    if (!mission.actions) mission.actions = [];

    for (let i in mission.actions) {
        if (mission.actions[i].number > number) {
            number = mission.actions[i].number
        }
    }
    number++;
    mission.actions.push({
        alternative_dialog_id: 0,
        async: false,
        count: 0,
        current_count: 0,
        description: "",
        dialog: null,
        dialog_id: 0,
        need_items: null,
        number: number,
        player_id: 0,
        q: 0,
        r: 0,
        radius: 0,
        sec: 0,
        short_description: "",
        type_func_monitor: "",
    });

    SaveMission(id);
}

function removeAction(missID, actID) {
    let mission = getMissionByID(missID);
    if (!mission.actions) return;

    for (let i in mission.actions) {
        if (mission.actions[i].id === actID) {
            mission.actions.splice(i, 1);
        }
    }

    SaveMission(missID);
}

/**
 *   информация по действиям
 **/

function getActionByID(id) {
    for (let i in allMission) {
        if (allMission.hasOwnProperty(i)) {
            for (let j in allMission[i].actions) {
                if (allMission[i].actions[j].id === id) {
                    return allMission[i].actions[j];
                }
            }
        }
    }
}

function SetNumberAction(context, id) {
    let action = getActionByID(id);
    action.number = Number(context.value);
}

function SetDescriptionAction(context, id) {
    let action = getActionByID(id);
    action.description = context.value;
}

function SetShortDescriptionAction(context, id) {
    let action = getActionByID(id);
    action.short_description = context.value;
}

function SetTypeMonitorAction(context, id) {
    let action = getActionByID(id);
    action.type_func_monitor = context.value;
}

function SetEndTextAction(context, id) {
    let action = getActionByID(id);
    action.end_text = context.value;
}

function SetAsyncAction(context, id) {
    let action = getActionByID(id);
    action.async = $(context).is(':checked');
}

/**
 *   мета-информация по действиям
 **/

function SetOwnerPlace(context, id) {
    let action = getActionByID(id);
    action.owner_place = $(context).is(':checked');
}

function SetActionBaseID(context, id) {
    let action = getActionByID(id);
    action.base_id = Number(context.value);
}

function SetActionMapID(context, id) {
    let action = getActionByID(id);
    action.map_id = Number(context.value);
}

function SetActionQ(context, id) {
    let action = getActionByID(id);
    action.q = Number(context.value);
}

function SetActionR(context, id) {
    let action = getActionByID(id);
    action.r = Number(context.value);
}

function SetActionRadius(context, id) {
    let action = getActionByID(id);
    action.radius = Number(context.value);
}

function SetActionSec(context, id) {
    let action = getActionByID(id);
    action.sec = Number(context.value);
}

function SetActionCount(context, id) {
    let action = getActionByID(id);
    action.count = Number(context.value);
}

function SetActionDialogID(context, id) {
    let action = getActionByID(id);
    action.dialog_id = Number(context.value);
}

function SetAltDialogID(context, id) {
    let action = getActionByID(id);
    action.alternative_dialog_id = Number(context.value);
}