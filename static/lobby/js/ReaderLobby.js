function ReaderLobby(jsonMessage) {

    let event = JSON.parse(jsonMessage).event;
    if (event === "InitLobby") {
        let login = document.getElementById('login');
        let userName = JSON.parse(jsonMessage).user_name;
        login.innerHTML = "Вы зашли как: " + userName;
    }

    if (event === "DisconnectLobby") {
        location.reload();
    }

    if (event === "DelUser") {
        let userTr = document.getElementById(JSON.parse(jsonMessage).game_user);
        if (userTr !== null) {
            userTr.remove();
        }
    }

    if (event === "Ready") {
        Ready(jsonMessage);
    }

    if (event === "GetSquad") {
        FillSquadBlock(jsonMessage)
    }

    if (event === "StartOutBase") {
        document.getElementById('OutDialog').style.visibility = "visible";
    }

    if (event === "OutBase") {
        location.href = "http://" + window.location.host + "/global";
    }

    if (event === "Error") {
        alert(JSON.parse(jsonMessage).error)
    }

    if (event === "updateRecycler") {
        FillRecycler(JSON.parse(jsonMessage));
    }

    if (event === "WorkbenchStorage") {
        FillWorkbench(JSON.parse(jsonMessage))
    }

    if (event === "SelectBP") {
        SelectBP(JSON.parse(jsonMessage))
    }

    if (event === "SelectWork") {
        SelectWork(JSON.parse(jsonMessage))
    }

    if (event === "dialog") {
        // этот костыль тут из за диалога обучения который идет не через стандартный путь
        setTimeout(function () {
            FillDepartment(JSON.parse(jsonMessage).dialog_page, JSON.parse(jsonMessage).dialog_action, JSON.parse(jsonMessage).mission, JSON.parse(jsonMessage).user_id)
        }, 200);
    }

    if (event === "training") {
        Training(JSON.parse(jsonMessage).count)
    }

    if (event === 'choiceFractionComplete') {
        document.location.reload(true);
    }
    if (event === "choiceFraction") {
        choiceFraction()
    }

    if (event === "BaseStatus") {
        UpdateBaseStatus(JSON.parse(jsonMessage).base)
    }

    if (event === "GetDetails") {
        FillDetailMarket(JSON.parse(jsonMessage).base, JSON.parse(jsonMessage).inventory_slots)
    }
}