let chat;

// тут уже не только чат, такие дела..
function ConnectChat() {
    $(document).ready(function () {
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
    });
}

function initChatInterface() {
    let chat = $('#chat');

    chat.data({
        resize: function (event, ui, el) {
            el.find('#chatBox').css("height", el.height() - 65);
            el.find('#usersBox').css("height", el.height() - 55);
            el.find('#chatBox').css("width", el.width() - 140);
            el.find('#chatInput').css("width", el.width() - 16);
            el.find('#tabsGroupWrapper').css("width", el.width() - 116);
            el.find('#chatTabs').css("width", el.width() - 100);
        }
    });

    chat.resizable({
        minHeight: 200,
        minWidth: 300,
        handles: "se, ne",
        resize: function (event, ui) {
            $(this).data("resize")(event, ui, $(this))
        },
        stop: function (e, ui) {
            setState(this.id, $(this).position().left, $(this).position().top, $(this).height(), $(this).width(), true);
        }
    });
}

function ChatReader(data) {

    if (data.event === 'OpenChat') {
        FillNotifyBlock(data);
        OpenChat(data);
    }

    if (data.event === 'RemoveGroup') {
        RemoveGroup(data.group_id)
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
        PreviewMapPath(data.search_maps)
    }

    if (data.event === 'OpenUserStat') {
        if (!document.getElementById('UsersStatus')) {
            UsersStatus(true);
            setTimeout(function () {
                FillUserStatus(data.player, null, data.user_id);
            }, 300)
        } else {
            FillUserStatus(data.player, null, data.user_id);
        }
    }

    if (data.event === 'OpenOtherUserStat') {
        OtherUserStatus(data.user);
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
            FillUserStatus(data.player, data.skill, data.user_id)
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

    if (data.event === "setWindowsState") {
        SetWindowsState(data.user_interface)
    }

    if (data.event === "DeleteNotify") {
        if (document.getElementById(data.uuid)) document.getElementById(data.uuid).remove();
    }

    if (data.event === "Error") {
        alert("ошиюка: " + data.error)
    }
}