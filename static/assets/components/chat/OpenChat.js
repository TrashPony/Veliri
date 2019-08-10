let currentChatID = 0;
let userName = ''; // текущей пользователь

function OpenChat(data) {
    console.log(data);
    if (data.user)
        userName = data.user.user_name;

    if (document.getElementById("allGroupsWindow")) document.getElementById("allGroupsWindow").remove();

    let tabs = document.getElementById('tabsGroup');

    // вкладка локального чата
    tabs.innerHTML = `<div id="chat0" onclick="ChangeCanal(0)">Локальный</div>`;
    for (let i in data.groups) {
        tabs.innerHTML += `<div id="chat${data.groups[i].id}" onclick="ChangeCanal(${data.groups[i].id})">${data.groups[i].name}</div>`
    }

    if (data.group) {
        ChangeCanal(data.group.id);
    } else {
        ChangeCanal(currentChatID);
    }
}

function ChangeCanal(id) {

    let oldChatTab = document.getElementById('chat' + currentChatID);
    if (oldChatTab) oldChatTab.className = '';

    let chatTab = document.getElementById('chat' + id);
    if (chatTab) chatTab.className = 'actionChatTab';

    currentChatID = Number(id);

    if (chat) {
        chat.send(JSON.stringify({
            event: "ChangeGroup",
            group_id: Number(id),
        }));
    }
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

            let chatUserLine = document.createElement("div");
            chatUserLine.className = "chatUserLine";
            chatUserLine.id = "chatUserLine" + users[i].user_name;
            chatUserLine.innerHTML = `<div class="chatUserName">${users[i].user_name}</div>`;
            $(chatUserLine).click(function (e) {
                ChatUserSubMenu(e, this, users[i])
            });

            let userAvatar = document.createElement("div");
            userAvatar.className = "chatUserIcon";
            $(chatUserLine).prepend(userAvatar);
            GetUserAvatar(users[i].user_id).then(function (response) {
                userAvatar.style.backgroundImage = "url('" + response.data.avatar + "')";
            });
            usersBox.appendChild(chatUserLine);

            if (users[i].user_name === userName && group.id !== 0) {
                $('#chatUserLine' + users[i].user_name).append('<div class="exitChatButton" onclick="Unsubscribe(event, ' + group.id + ')">x</div>')
            }
        }
    }
}

function Unsubscribe(event, groupID) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    chat.send(JSON.stringify({
        event: "Unsubscribe",
        group_id: Number(groupID),
    }));
}

function ChatUserSubMenu(e, context, user) {
    if (document.getElementById('chatSubMenu')) document.getElementById('chatSubMenu').remove();

    let hideWriteButton = "";
    if (user.user_name === userName) {
        hideWriteButton = `display: none;`
    }

    let subMenu = document.createElement("div");
    subMenu.id = "chatSubMenu";
    subMenu.style.left = (e.pageX) + "px";
    subMenu.style.top = (e.pageY) + "px";

    subMenu.innerHTML = `
        <div class="chatSubMenuUserAction">
            <input type="button" value="Подробнее" onclick="informationFunc('${user.user_name}', '${user.user_id}'); document.getElementById('chatSubMenu').remove()">
            <input type="button" value="Написать" style="${hideWriteButton}" onclick="CreatePrivateChatGroup('${user.user_id}'); document.getElementById('chatSubMenu').remove()">
            <input type="button" value="Закрыть" onclick="document.getElementById('chatSubMenu').remove()">
        </div>
    `;

    let chatSubMenuUserHead = document.createElement("div");
    chatSubMenuUserHead.className = "chatSubMenuUserHead";

    let chatUserLine = document.createElement("div");
    chatUserLine.className = "chatUserLine";
    chatUserLine.innerHTML = `
    <div>
        <div class="chatUserName">${user.user_name}</div>
        <div class="chatUserTitle">${user.title}</div>
    </div>
    `;

    let userAvatar = document.createElement("div");
    userAvatar.className = "chatUserIcon";
    $(chatUserLine).prepend(userAvatar);
    GetUserAvatar(user.user_id).then(function (response) {
        userAvatar.style.backgroundImage = "url('" + response.data.avatar + "')";
    });

    $(chatUserLine).prepend(userAvatar);
    chatSubMenuUserHead.appendChild(chatUserLine);
    $(subMenu).prepend(chatSubMenuUserHead);

    document.body.appendChild(subMenu);
}

function CreatePrivateChatGroup(userID) {
    chat.send(JSON.stringify({
        event: "CreateNewPrivateGroup",
        user_id: Number(userID),
    }));
}

function RemoveGroup(groupID) {
    if (document.getElementById("chat" + groupID)) document.getElementById("chat" + groupID).remove();
    if (groupID === currentChatID) {
        ChangeCanal(0)
    } else {
        ChangeCanal(currentChatID)
    }
}