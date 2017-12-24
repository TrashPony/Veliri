var chat;

function ConnectChat() {
    chat = new WebSocket("ws://" + window.location.host + "/wsChat");
    console.log("Websocket chat - status: " + chat.readyState);

    chat.onopen = function() {
        console.log("Connection chat opened..." + this.readyState);
    };

    chat.onmessage = function(msg) {
        console.log(msg);
        NewChatMessage(msg.data);
    };

    chat.onerror = function(msg) {
        console.log("Error chat occured sending..." + msg.data);
    };

    chat.onclose = function(msg) {
        console.log("Disconnected chat - status " + this.readyState);
    };
}

function chatMessage() {
    var chatInput = document.getElementById("chatInput");
    var text = chatInput.value;
    if (text !== "") {
        chatInput.value = null;

        chat.send(JSON.stringify({
            event: "NewChatMessage",
            message: text
        }));
    }
}

function NewChatMessage(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    if (event === "NewChatMessage") {
        var chatBox = document.getElementById("chatBox");
        var UserName = document.createElement("span");
        UserName.className = "ChatUserName";
        UserName.innerHTML = JSON.parse(jsonMessage).game_user + ":";
        var TextMessage = document.createElement("span");
        TextMessage.className = "ChatText";
        TextMessage.innerHTML = JSON.parse(jsonMessage).message;
        chatBox.appendChild(UserName);
        chatBox.appendChild(TextMessage);
        chatBox.appendChild(document.createElement("br"));
    }
}
