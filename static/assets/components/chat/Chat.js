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
    console.log(data)
    if (data.event === 'OpenChat') {
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

    if (data.event === "OpenLocalChat") {
        //systemMessage("Вы входите на территорию" + data.group.name);
    }
}