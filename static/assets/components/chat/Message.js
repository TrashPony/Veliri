function chatMessage() {
    let chatInput = document.getElementById("chatInput");
    let text = chatInput.value;
    if (text !== "") {
        chatInput.value = null;

        chat.send(JSON.stringify({
            event: "NewChatMessage",
            message_text: text,
            group_id: Number(currentChatID),
        }));
    }
}

function NewChatMessage(message, id) {

    if (id === currentChatID) {
        let chatBox = document.getElementById("chatBox");

        chatBox.innerHTML += `
            <div class="chatMessage">
            
                <div class="chatUserIcon"></div>
                <span class="ChatUserName">${message.user_name} > </span>
                <span class="ChatText">${message.message}</span>
            
            </div>
        `;
    } else {
        let chatTab = document.getElementById('chat' + id);
        if (chatTab) chatTab.className = 'alertChatTab';
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