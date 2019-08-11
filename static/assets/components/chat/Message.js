function chatMessage() {
    let chatInput = document.getElementById("chatInput");
    let text = chatInput.value;

    if (text.indexOf("/color squad:") >= 0) {
        // if (game.squad.sprite.weaponColorMask) {
        //     //game.squad.sprite.weaponColorMask.tint = Phaser.WHITE;
        //     //game.squad.sprite.weaponColorMask.tint = '0x' + text.split(':')[1];
        // }
        //
        // game.squad.sprite.bodyMask2.tint = Phaser.WHITE;
        // game.squad.sprite.bodyMask2.tint = '0x' + text.split(':')[1];
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

        // смотрим находится ли скрол в самом низу
        let chatIsDown = chatBox.scrollTop === chatBox.scrollHeight - chatBox.clientHeight;

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

        // если скрол чата в самом низу то прокручивать его при новых сообщениях, если нет то без автопрокрутки
        if (chatIsDown) chatBox.scrollTop = chatBox.scrollHeight - chatBox.clientHeight;

    } else {
        let chatTab = document.getElementById('chat' + id);
        if (chatTab) chatTab.className = 'alertChatTab';
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