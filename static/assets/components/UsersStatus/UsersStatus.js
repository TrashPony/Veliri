let userStat = {};

function UsersStatus(noMessage) {
    if (document.getElementById("UsersStatus")) {
        let jBox = $('#UsersStatus');
        setState('UsersStatus', jBox.position().left, jBox.position().top, jBox.height(), jBox.width(), false);
        return
    }

    let usersStatus = document.createElement("div");
    document.body.appendChild(usersStatus);

    usersStatus.id = "UsersStatus";
    usersStatus.innerHTML = `
        <div id="usersStatusTabs">
            <div id="TabsLeftArrow" onclick="document.getElementById('usersStatusTabsGroup').scrollLeft -= 20;"><</div>
            <div id="tabsWrapper">
                <div id="usersStatusTabsGroup">
                    <div id="commonUserStat" class="actionChatTab" onclick="OpenCommonUserStat()">Общие</div>
                    <div id="skillUserStat" onclick="OpenSkillsUserStat()">Навыки</div>
                    <div id="StatUserStat">Задания</div>
                </div>
            </div>
            <div id="TabsRightArrow" onclick="document.getElementById('usersStatusTabsGroup').scrollLeft += 20;">></div>
        </div>
        <div id="usersStatusWrapper">
        </div>
    `;

    let buttons = CreateControlButtons("2px", "31px", "-3px", "29px", "", "145px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'UsersStatus')
    });
    $(buttons.close).mousedown(function () {
        setState(usersStatus.id, $(usersStatus).position().left, $(usersStatus).position().top, $(usersStatus).height(), $(usersStatus).width(), false);
    });
    usersStatus.appendChild(buttons.move);
    usersStatus.appendChild(buttons.close);

    if (noMessage) {
        OpenCommonUserStat();
    }
    openWindow(usersStatus.id, usersStatus)
}

function OpenCommonUserStat() {
    document.getElementById('skillUserStat').className = '';
    document.getElementById('StatUserStat').className = '';
    document.getElementById('commonUserStat').className = 'actionChatTab';

    let usersStatus = document.getElementById("usersStatusWrapper");
    usersStatus.innerHTML = `
            <h3 id="userName"> UserName </h3>
                <div id="userAvatarWrapper">
                <div id="userAvatar"></div> 
                <input style="position: absolute; left: -9999px;" type="file" name="uploadFile" id="file" onchange="SelectAvatarFile(event)"/>
                <label for="file" id="labelFile"> Загрузить</label>
            </div>
            
            <div id="UserStatusPanel">
                <div id="userTitle"></div>
                <div id="scientific_points">очки науки <span id="scientific_points_points">1000</span></div>
                <div id="attack_points">очки атаки <span id="attack_points_points">1000</span></div>
                <div id="production_points">очки произвосдва <span id="production_points_points">1000</span></div>
            </div>
             
            <div id="biography">
                <h3>Биография:</h3>
                <textarea id="userBiography"></textarea>
                <input type="button" value="Сохранить" onclick="SetBiography()">
            </div>
    `;

    chat.send(JSON.stringify({
        event: "OpenUserStat",
    }));
}

function OpenSkillsUserStat() {
    document.getElementById('commonUserStat').className = '';
    document.getElementById('StatUserStat').className = '';
    document.getElementById('skillUserStat').className = 'actionChatTab';

    let usersStatus = document.getElementById("usersStatusWrapper");
    usersStatus.innerHTML = `
        <div id="skillTip">
            <div id="skillIcon">
                <div id="skillUpdatePanel">
                    <div id="skillLvl">
                        <div></div>
                        <div></div>
                        <div></div>
                        <div></div>
                        <div></div>
                    </div>
                    <div id="upperSkill">+</div>
                </div>
            </div>
            <h3><span id="skillName"> SkillName</span> <span id="needPrice" style="float: right; color: #00efff"></span></h3>
            <div id="skillDescription"> Описание ...</div>
        </div>
        
        <div id="listSkills">
            
            <h4 style="color: cornflowerblue;"> Наука <span id="scientific_points_points">1000</span></h4>
            <div id="ScySkills" style="float: left">
                
            </div>
            
            <h4 style="color: crimson;" > Атака <span id="attack_points_points">1000</span></h4>
            <div id="AttackSkills"  style="float: left">
                
            </div>
            
            <h4 style="color: chartreuse;" > Производство <span id="production_points_points">1000</span></h4>
            <div id="IndustrySkills" style="float: left">
                
            </div>
        </div>`;

    chat.send(JSON.stringify({
        event: "OpenUserStat",
    }));
}

function SelectAvatarFile(e) {
    let file_reader = new FileReader(e.target.files[0]);
    file_reader.readAsDataURL(e.target.files[0]);
    file_reader.onload = function (evt) {
        chat.send(JSON.stringify({
            event: "LoadAvatar",
            file: evt.target.result
        }));
    };
}

function SetBiography() {
    chat.send(JSON.stringify({
        event: "SetBiography",
        biography: document.getElementById("userBiography").value,
    }));
}

function OtherUserStatus(user) {

    // создание минималистичного статут юзера из 1й страницы, добавить кнопки типо добавить во френды, отправить сообщение, начать чат;
    // TODO почти повторяющийся код с функцией UsersStatus(), но я решил не трогать ту функцию а создать новую
    if (document.getElementById("UsersStatus")) {
        document.getElementById("UsersStatus").remove();
    }

    let usersStatus = document.createElement("div");
    usersStatus.id = "UsersStatus";
    document.body.appendChild(usersStatus);

    usersStatus.innerHTML = `
        <div id="usersStatusTabs">
            <div id="TabsLeftArrow" onclick="document.getElementById('usersStatusTabsGroup').scrollLeft -= 20;"><</div>
            <div id="tabsWrapper">
                <div id="usersStatusTabsGroup">
                    <div id="commonUserStat" class="actionChatTab" onclick="OpenCommonUserStat()">Общие</div>
                </div>
            </div>
            <div id="TabsRightArrow" onclick="document.getElementById('usersStatusTabsGroup').scrollLeft += 20;">></div>
        </div>
        <div id="usersStatusWrapper">
        </div>
    `;

    let buttons = CreateControlButtons("2px", "31px", "-3px", "29px", "", "145px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'UsersStatus')
    });
    $(buttons.close).mousedown(function () {
        setState(usersStatus.id, $(usersStatus).position().left, $(usersStatus).position().top, $(usersStatus).height(), $(usersStatus).width(), false);
    });
    usersStatus.appendChild(buttons.move);
    usersStatus.appendChild(buttons.close);
    openWindow(usersStatus.id, usersStatus);

    let usersStatusWrap = document.getElementById("usersStatusWrapper");
    usersStatusWrap.innerHTML = `
            <h3 id="userName"> ${userName} </h3>
                <div id="userAvatarWrapper">
                <div id="userAvatar"></div> 
            </div>
            
            <div id="UserStatusPanel">
                <div id="userTitle"></div>
                <div id="actionPanel">
                    <input type="button" value="Отправить сообщение">
                    <input type="button" value="Предложить дружбу">
                    <input type="button" value="что то еще">
                </div>
            </div>
             
            <div id="biography">
                <h3>Биография:</h3>
                <textarea id="userBiography" disabled></textarea>
            </div>
    `;

    FillOtherUserStat(user);
}