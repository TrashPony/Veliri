function ViewAllGroups() {

    if (document.getElementById("allGroupsWindow")) {
        document.getElementById("allGroupsWindow").remove();
        return
    }

    let allGroupsWindow = document.createElement("div");
    allGroupsWindow.id = "allGroupsWindow";

    chat.send(JSON.stringify({
        event: "GetAllGroups",
    }));

    allGroupsWindow.innerHTML = `
        <h3 style="float: left">Чат Каналы:</h3>
        <div class="topButton" onmousedown="document.getElementById('allGroupsWindow').remove()">x</div>
        <div class="topButton" onmousedown="moveWindow(event,'allGroupsWindow')">⇿</div>
        <div id="ViewChatGroups"></div>
        <div>
            <input type="button" value="Создать" style="float: right; margin-right: 25px" onclick="CreateNewGroup()">
            <input type="button" value="Закрыть" style="margin-left: 25px" onclick="document.getElementById('allGroupsWindow').remove()">
        </div>
    `;

    document.body.appendChild(allGroupsWindow);
}

function AllGroups(groups) {
    let ViewChatGroups = document.getElementById("ViewChatGroups");
    ViewChatGroups.innerHTML = '';

    let commonGroups = document.createElement("div");
    commonGroups.innerHTML = '<h4>Основные группы: </h4>';
    ViewChatGroups.appendChild(commonGroups);

    let userGroups = document.createElement("div");
    userGroups.id = "userGroups";
    userGroups.innerHTML = `
        <h4>Пользовательские группы: </h4>
        <input id="filterChatGroup" type="text" placeholder="поиск группы..." oninput="FilterChatGroup(event, this)">
        `;
    ViewChatGroups.appendChild(userGroups);

    for (let i in groups) {

        let playerCount = 0;
        let playerOnline = 0;

        for (let j in groups[i].users) {
            playerCount++;
            if (groups[i].users[j]) playerOnline++;
        }

        let passIcon = "";
        if (!groups[i].secure) {
            passIcon = "visibility: hidden;"
        }

        if (groups[i].user_create) {
            userGroups.innerHTML += `
            <div class="chatGroup" onclick="SubscribeGroup('${groups[i].id}', '${groups[i].secure}')" id="chatGroup${groups[i].id}">
                <div class="chatName"> ${groups[i].name}</div>
                <div class="chatPlayerCount">${playerOnline} / ${playerCount}</div>
                <div class="public" style="${passIcon}"></div>
            </div>
            `;
        } else {
            commonGroups.innerHTML += `
            <div class="chatGroup" onclick="SubscribeGroup('${groups[i].id}', '${groups[i].secure}')" id="chatGroup${groups[i].id}">
                <div class="chatName"> ${groups[i].name}</div>
                <div class="chatPlayerCount">${playerOnline} / ${playerCount}</div>
                <div class="public" style="${passIcon}"></div>
            </div>
            `;
        }
    }

    for (let i in groups) {
        GetGroupDivAvatar($(ViewChatGroups).find('#chatGroup' + groups[i].id), groups[i].id)
    }
}

function FilterChatGroup(e, context) {
    $('#userGroups').find('.chatGroup').each(function (i, row) {
        let groupName = $(row).find('.chatName')[0].innerHTML;

        if (groupName.indexOf(context.value) + 1) {
            row.style.display = "block";
        } else {
            row.style.display = "none";
        }
    })
}

let newGroup = {};

function CreateNewGroup() {
    let NewChatCreateWrapper = document.createElement("div");
    NewChatCreateWrapper.id = "NewChatCreateWrapper";
    document.body.appendChild(NewChatCreateWrapper);
    NewChatCreateWrapper.innerHTML = `
        <h3 style="float: left">Создание нового канала:</h3>
        <div class="topButton" onmousedown="document.getElementById('NewChatCreateWrapper').remove(); newGroup = {}">x</div>
        <div class="topButton" onmousedown="moveWindow(event,'NewChatCreateWrapper')">⇿</div>
        <div id="NewChatCreate">
        
            <input id="nameNewChatGroup" type="text" placeholder="Имя канала...">
            <input id="passNewChatGroup" type="text" placeholder="Пароль (если пусто то без пароля)">

            <div id="avatarNewChatGroupWrapper">
                <div id="avatarNewChatGroup"></div>
                <input style="position: absolute; left: -9999px;" type="file" name="uploadFile" id="file" onchange="previewAvatarGroup(event)"/>
                <label for="file" id="labelFile"> Загрузить</label>
            </div>
            <div id="greetingsNewChatGroupWrapper">
                <h4> Приветственное сообщение </h4>
                <textarea id="greetingsNewChatGroup"></textarea>
            </div>
            <input type="button" value="Создать" style="float: right; margin-right: 25px" onclick="SendCreateNewGroup()">
        </div>
        <div>
            <input type="button" value="Закрыть" style="float: left; margin-left: 25px" onclick="document.getElementById('NewChatCreateWrapper').remove(); newGroup = {}">
        </div>
    `
}

function previewAvatarGroup(e) {
    let file_reader = new FileReader(e.target.files[0]);
    file_reader.readAsDataURL(e.target.files[0]);
    file_reader.onload = function (evt) {
        document.getElementById("avatarNewChatGroup").style.backgroundImage = "url('" + evt.target.result + "')";
        newGroup.avatar = evt.target.result;

    };
}

function SendCreateNewGroup() {
    let name = document.getElementById("nameNewChatGroup").value;
    if (name === "" || name === " ") { // todo нужно парсить строку на небезопасные символы
        alert("Не введено имя канала");
        return;
    }

    chat.send(JSON.stringify({
        event: "CreateNewChatGroup",
        password: document.getElementById("passNewChatGroup").value,
        file: newGroup.avatar,
        name: name,
        greetings: document.getElementById("greetingsNewChatGroup").value
    }));

    newGroup = {}
}

function GetGroupDivAvatar(parent, groupId) {
    let groupAvatar = document.createElement("div");
    groupAvatar.className = "chatLogo";
    $(parent).prepend(groupAvatar);
    GetGroupAvatar(groupId).then(function (response) {
        groupAvatar.style.backgroundImage = "url('" + response.data.avatar + "')";
    });
}

function SubscribeGroup(id, secure, pass) {

    if (secure === "true" && !pass) {

        if (document.getElementById("chatPassBlock")) document.getElementById("chatPassBlock").remove();

        let chatPassBlock = document.createElement("chatPassBlock");
        chatPassBlock.id = "chatPassBlock";
        chatPassBlock.innerHTML = `
            <input id="enterChatPassword" type="text" placeholder="введите пароль" 
                onchange="SubscribeGroup('${id}', '${secure}', document.getElementById('enterChatPassword').value);
                document.getElementById('chatPassBlock').remove()">
            <input type="button" value="Отмена" onclick="document.getElementById('chatPassBlock').remove()">
            <input type="button" value="Войти" 
                onclick="SubscribeGroup('${id}', '${secure}', document.getElementById('enterChatPassword').value);
                document.getElementById('chatPassBlock').remove()">
        `;

        document.body.appendChild(chatPassBlock);
    } else {
        chat.send(JSON.stringify({
            event: "SubscribeGroup",
            group_id: Number(id),
            password: pass,
        }));
    }
}