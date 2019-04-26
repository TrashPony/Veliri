let currentChatID = 0;
let userName = ''; // текущей пользователь

function OpenChat(data) {

    if (data.user)
        userName = data.user.user_name;


    if (document.getElementById("allGroupsWindow")) document.getElementById("allGroupsWindow").remove();

    let tabs = document.getElementById('tabsGroup');

    // вкладка локального чата
    tabs.innerHTML = `<div id="chat0" onclick="ChangeCanal(0)">Локальный</div>`;
    for (let i in data.groups) {
        tabs.innerHTML += `<div id="chat${data.groups[i].id}" onclick="ChangeCanal(${data.groups[i].id})">${data.groups[i].name}</div>`
    }

    ChangeCanal(0); // открываем по умолчанию локальный чат
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

    updateUsers(group, users);

    let chatBox = document.getElementById("chatBox");
    chatBox.innerHTML = '';

    if (currentChatID === 0)
        systemMessage("Вы входите на территорию " + group.name);

    for (let i = 0; group.history && i < group.history.length; i++) {
        NewChatMessage(group.history[i], group.id)
    }

    chatBox.scrollTop = chatBox.scrollHeight;
}

function updateUsers(group, users) {

    if (currentChatID !== group.id) return;

    let usersBox = document.getElementById('usersBox');
    usersBox.innerHTML = '';
    for (let i in users) {
        if (users.hasOwnProperty(i) && users[i]) {
            usersBox.innerHTML += `<div class="chatUserLine" id="${users[i].user_name}">
                                        <div class="chatUserIcon" style="background-image: url('${users[i].avatar_icon}')"></div>
                                        <div class="chatUserName">${users[i].user_name}</div>
                                   </div>`;

            if (users[i].user_name === userName && group.id !== 0) {
                $('#' + users[i].user_name).append('<div class="exitChatButton" onclick="Unsubscribe(' + group.id + ')">x</div>')
            }
        }
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

function Unsubscribe(groupID) {
    chat.send(JSON.stringify({
        event: "Unsubscribe",
        group_id: Number(groupID),
    }));
}