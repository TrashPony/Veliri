let chat;

function ConnectChat() {
    chat = new WebSocket("ws://" + window.location.host + "/wsChat");
    console.log("Websocket chat - status: " + chat.readyState);

    chat.onopen = function () {
        console.log("Connection chat opened..." + this.readyState);
        chat.send(JSON.stringify({
            event: "OpenChat",
        }));
        initChatInterface();
    };

    chat.onmessage = function (msg) {
        ChatReader(JSON.parse(msg.data));
    };

    chat.onerror = function (msg) {
        console.log("Error chat occured sending..." + msg.data);
    };

    chat.onclose = function (msg) {
        console.log("Disconnected chat - status " + this.readyState);
    };
}

function initChatInterface() {
    let chat = $('#chat');
    chat.resizable({
        minHeight: 200,
        minWidth: 300,
        handles: "se, ne",
        resize: function (event, ui) {
            $(this).find('#chatBox').css("height", $(this).height() - 65);
            $(this).find('#usersBox').css("height", $(this).height() - 55);

            $(this).find('#chatBox').css("width", $(this).width() - 140);
            $(this).find('#chatInput').css("width", $(this).width() - 16);
            $(this).find('#tabsGroupWrapper').css("width", $(this).width() - 116);
            $(this).find('#chatTabs').css("width", $(this).width() - 100);
        }
    });
}

function ChatReader(data) {

    if (data.event === 'OpenChat') {
        FillNotifyBlock(data);
        OpenChat(data);
    }

    if (data.event === 'GetAllGroups') {
        AllGroups(data.groups);
    }

    if (data.event === 'ChangeGroup') {
        OpenCanal(data.group, data.users);
    }

    if (data.event === 'NewChatMessage') {
        NewChatMessage(data.message, data.group_id)
    }

    if (data.event === "UpdateUsers") {
        updateUsers(data.group, data.users)
    }

    if (data.event === "newNotify") {
        newNotify(data.notify)
    }

    if (data.event === "OpenLocalChat") {
        //systemMessage("Вы входите на территорию" + data.group.name);
    }

    // other reader
    if (data.event === "openMapMenu") {
        FillGlobalMap(data.maps, data.id)
    }

    if (data.event === "previewPath") {
        PreviewPath(data.search_maps)
    }

    if (data.event === 'OpenUserStat') {
        FillUserStatus(data.player);
    }

    if (data.event === "upSkill") {
        if (data.error) {
            if (document.getElementById('skillUpdatePanel')) {
                document.getElementById('skillUpdatePanel').style.animation = 'alert 500ms 1 ease-in-out';
                setTimeout(function () {
                    document.getElementById('skillUpdatePanel').style.animation = 'none';
                }, 500)
            }
        } else {
            FillUserStatus(data.player, data.skill)
        }
    }

    if (data.event === "openDepartmentOfEmployment") {
        FillDepartment(data.dialog_page)
    }

    if (data.event === "dialog") {
        FillDepartment(data.dialog_page, data.dialog_action, data.mission)
    }

    if (data.event === "training") {
        Training(data.count)
    }
}