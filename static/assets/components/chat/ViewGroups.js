function ViewAllGroups() {
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
        <input type="button" value="Создать">
    `;

    document.body.appendChild(allGroupsWindow);
}

function AllGroups(groups) {
    let ViewChatGroups = document.getElementById("ViewChatGroups");
    ViewChatGroups.innerHTML = '';

    for (let i in groups) {

        let playerCount = 0;
        let playerOnline = 0;

        for (let j in groups[i].users) {
            playerCount++;
            if (groups[i].users[j]) playerOnline++;
        }


        ViewChatGroups.innerHTML += `
            <div class="chatGroup" onclick="SubscribeGroup(${groups[i].id})">
                <div class="chatLogo" style="background-image: url('../assets/logo/${groups[i].fraction}.png')"></div>
                <div class="chatName"> ${groups[i].name}</div>
                <div class="chatPlayerCount">${playerOnline} / ${playerCount}</div>
                <div class="public"></div>
            </div>
        `
    }
}

function SubscribeGroup(id) {
    chat.send(JSON.stringify({
        event: "SubscribeGroup",
        group_id: Number(id),
    }));
}