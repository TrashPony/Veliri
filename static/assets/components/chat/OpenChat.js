let currentChatID = 0;

function OpenChat(data) {
    if (document.getElementById("allGroupsWindow")) document.getElementById("allGroupsWindow").remove();

    let tabs = document.getElementById('tabsGroup');
    tabs.innerHTML = '';

    for (let i in data.groups) {

        if (currentChatID === 0) {
            currentChatID = data.groups[i].id;
        }

        tabs.innerHTML += `<div id="chat${data.groups[i].id}" onclick="ChangeCanal(${data.groups[i].id})">${data.groups[i].name}</div>`
    }

    ChangeCanal(currentChatID);
}

function ChangeCanal(id) {

    let oldChatTab = document.getElementById('chat' + currentChatID);
    if (oldChatTab) oldChatTab.className = '';

    let chatTab = document.getElementById('chat' + id);
    if (chatTab) chatTab.className = 'actionChatTab';

    currentChatID = Number(id);

    chat.send(JSON.stringify({
        event: "ChangeGroup",
        group_id: Number(id),
    }));
}

function OpenCanal(group, users) {
    //загрузка юзеров, загрузка истории сообщений

    updateUsers(users);

    let chatBox = document.getElementById("chatBox");
    chatBox.innerHTML = '';
    for (let i = 0; group.history && i < group.history.length; i++) {
        NewChatMessage(group.history[i], group.id)
    }

    chatBox.scrollTop = chatBox.scrollHeight;
}

function updateUsers(users) {
    let usersBox = document.getElementById('usersBox');
    usersBox.innerHTML = '';
    for (let i in users) {
        if (users.hasOwnProperty(i) && users[i]) {
            usersBox.innerHTML += `<div class="chatUserLine"><div class="chatUserIcon"></div><div class="chatUserName">${users[i].user_name}</div></div>`;
        }
    }
}