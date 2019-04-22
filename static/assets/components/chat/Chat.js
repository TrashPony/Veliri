let chat;

function ConnectChat() {
    chat = new WebSocket("ws://" + window.location.host + "/wsChat");
    console.log("Websocket chat - status: " + chat.readyState);

    chat.onopen = function () {
        console.log("Connection chat opened..." + this.readyState);
        chat.send(JSON.stringify({
            event: "OpenChat",
        }));
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

function ChatReader(data) {
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
}