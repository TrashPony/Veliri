function UsersStatus() {
    if (document.getElementById("UsersStatus")) {
        document.getElementById("UsersStatus").remove();
        return
    }

    let usersStatus = document.createElement("div");
    usersStatus.id = "UsersStatus";
    usersStatus.innerHTML = `
        <div id="usersStatusTabs">
            <div id="TabsLeftArrow" onclick="document.getElementById('usersStatusTabsGroup').scrollLeft -= 20;"><</div>
            <div id="tabsWrapper">
                <div id="usersStatusTabsGroup">
                    <div class="actionChatTab">Общие</div>
                    <div>Навыки</div>
                    <div>Статистика</div>
                </div>
            </div>
            <div id="TabsRightArrow" onclick="document.getElementById('usersStatusTabsGroup').scrollLeft += 20;">></div>
        </div>
        <div id="usersStatusWrapper">
            <h3 id="userName"> UserName </h3>
            <div id="userAvatarWrapper">
                <div id="userAvatar"></div> 
                <input style="position: absolute; left: -9999px;" type="file" name="uploadFile" id="file" onchange="SelectAvatarFile(event)"/>
                <label for="file" id="labelFile"> Загрузить</label>
            </div>
            
            <div id="UserStatusPanel">
            </div>
             
            <div id="biography">
                <h3>Биография:</h3>
                <textarea></textarea>
                <input type="button" value="Сохранить">
            </div>
        </div>
    `;

    let buttons = CreateControlButtons("2px", "31px", "-3px", "29px", "", "145px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'UsersStatus')
    });
    $(buttons.close).mousedown(function () {
        usersStatus.remove();
    });
    usersStatus.appendChild(buttons.move);
    usersStatus.appendChild(buttons.close);

    document.body.appendChild(usersStatus);
}

function SelectAvatarFile(e) {
    let file_reader = new FileReader(e.target.files[0]);
    file_reader.readAsDataURL(e.target.files[0]);
    file_reader.onload = function (evt) {
        document.getElementById('userAvatar').style.backgroundImage = "url(" + evt.target.result + ")";

        lobby.send(JSON.stringify({
            event: "LoadAvatar",
            file: evt.target.result
        }));
    };
}