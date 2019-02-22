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
        CreatePageDialog(JSON.parse(jsonMessage).dialog_page, JSON.parse(jsonMessage).dialog_action)
    }
}