function chatMessage(e) {
    if (e.keyCode === 13) {
        var chatInput = document.getElementById("chatInput");
        var text = chatInput.value;
        if (text !== "") {
            chatInput.value = null;

            sock.send(JSON.stringify({
                event: "NewChatMessage",
                message: text
            }));
        }
        return false;
    }
}