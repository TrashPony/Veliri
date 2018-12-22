let chat;

function ConnectChat() {
    chat = new WebSocket("ws://" + window.location.host + "/wsChat");
    console.log("Websocket chat - status: " + chat.readyState);

    chat.onopen = function () {
        console.log("Connection chat opened..." + this.readyState);
    };

    chat.onmessage = function (msg) {
        NewChatMessage(msg.data);
    };

    chat.onerror = function (msg) {
        console.log("Error chat occured sending..." + msg.data);
    };

    chat.onclose = function (msg) {
        console.log("Disconnected chat - status " + this.readyState);
    };
}

function chatMessage() {
    let chatInput = document.getElementById("chatInput");
    let text = chatInput.value;
    if (text !== "") {
        chatInput.value = null;

        chat.send(JSON.stringify({
            event: "NewChatMessage",
            message: text
        }));
    }
}

function NewChatMessage(jsonMessage) {
    let event = JSON.parse(jsonMessage).event;

    if (event === "NewChatMessage") {
        let chatBox = document.getElementById("chatBox");
        let UserName = document.createElement("span");
        UserName.className = "ChatUserName";
        UserName.innerHTML = JSON.parse(jsonMessage).game_user + ":";
        let TextMessage = document.createElement("span");
        TextMessage.className = "ChatText";
        TextMessage.innerHTML = JSON.parse(jsonMessage).message;
        chatBox.appendChild(UserName);
        chatBox.appendChild(TextMessage);
        chatBox.appendChild(document.createElement("br"));
    }
}

let chatHide = false;

function HideChat() {
    let chat = document.getElementById("chat");
    let chatBox = document.getElementById("chatBox");
    let chatInput = document.getElementById("chatInput");
    let usersBox = document.getElementById("usersBox");

    function transform(el, height, opacity, sec) {
        el.style.transition = sec + "s";
        el.style.height = height + "px";
        el.style.opacity = opacity;
        setTimeout(function () {
            el.style.transition = "0s";
        }, 1000);
    }

    if (!chatHide) {
        chat.style.bottom = "180px";
        transform(chat, 20, 1, 1);
        transform(chatBox, 0, 0, 0.5);
        transform(chatInput, 0, 0, 0.5);
        transform(usersBox, 0, 0, 0.5);

        chatHide = true;
    } else {
        chat.style.bottom = "20px";
        transform(chat, 200, 1, 1);
        transform(chatBox, 125, 1, 1.5);
        transform(chatInput, 30, 1, 1.5);
        transform(usersBox, 125, 1, 1.5);
        chatHide = false;
    }
}