function chatMessage() {
    let chatInput = document.getElementById("chatInput");
    let text = chatInput.value;

    if (text.indexOf("/color squad:") >= 0) {
        if (game.squad.sprite.weaponColorMask) {
            //game.squad.sprite.weaponColorMask.tint = Phaser.WHITE;
            //game.squad.sprite.weaponColorMask.tint = '0x' + text.split(':')[1];
        }

        game.squad.sprite.bodyMask2.tint = Phaser.WHITE;
        game.squad.sprite.bodyMask2.tint = '0x' + text.split(':')[1];
    }

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

        if (!message.system) {
            let chatMessage = document.createElement("div");
            chatMessage.className = "chatMessage";
            chatMessage.innerHTML += `
                <span class="ChatUserName">${message.user_name} > </span>
                <span class="ChatText">${message.message}</span>
        `;
            chatBox.appendChild(chatMessage);

            let userAvatar = document.createElement("div");
            userAvatar.className = "chatUserIcon";
            $(chatMessage).prepend(userAvatar);
            GetUserAvatar(message.user_id).then(function (response) {
                userAvatar.style.backgroundImage = "url('" + response.data.avatar + "')";
            });
        } else {
            systemMessage(message.message)
        }
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

function systemMessage(text) {
    let chatBox = document.getElementById("chatBox");

    chatBox.innerHTML += `
            <div class="chatMessage">
                <span class="chatSystem">${text}</span>
            </div>
        `;
}