let userStat = {};

function UsersStatus() {
    if (document.getElementById("UsersStatus")) {
        document.getElementById("UsersStatus").remove();
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
                    <div id="StatUserStat">Статистика</div>
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
        usersStatus.remove();
    });
    usersStatus.appendChild(buttons.move);
    usersStatus.appendChild(buttons.close);

    OpenCommonUserStat();
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
                <div id="userTitle">Срущий мимо ванной</div>
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

    lobby.send(JSON.stringify({
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
            <h3> SkillName </h3>
            <div id="skillDescription"> Описание ...</div>
        </div>
        
        <div id="listSkills">
            
            <h4 style="color: cornflowerblue;"> Наука <span id="scientific_points_points">1000</span></h4>
            <div class="ScySkills" style="float: left">
                <div class="skill">
                    <div class="miniSkillIcon"></div>
                    <div class="skillPrice">500</div>
                    <div class="skillLvl">
                        <div></div>
                        <div></div>
                        <div></div>
                        <div></div>
                        <div></div>
                    </div>
                </div>
            </div>
            
            <h4 style="color: crimson;" > Атака <span id="attack_points_points">1000</span></h4>
            <div class="AttackSkills"  style="float: left">
                <div class="skill">
                    <div class="miniSkillIcon"></div>
                    <div class="skillPrice">500</div>
                    <div class="skillLvl">
                        <div></div>
                        <div></div>
                        <div></div>
                        <div></div>
                        <div></div>
                    </div>
                </div>
            </div>
            
            <h4 style="color: chartreuse;" > Производство <span id="production_points_points">1000</span></h4>
            <div class="IndustrySkills" style="float: left">
                <div class="skill">
                    <div class="miniSkillIcon"></div>
                    <div class="skillPrice">500</div>
                    <div class="skillLvl">
                        <div></div>
                        <div></div>
                        <div></div>
                        <div></div>
                        <div></div>
                    </div>
                </div>
            </div>
        </div>`
}

function SelectAvatarFile(e) {
    let file_reader = new FileReader(e.target.files[0]);
    file_reader.readAsDataURL(e.target.files[0]);
    file_reader.onload = function (evt) {
        lobby.send(JSON.stringify({
            event: "LoadAvatar",
            file: evt.target.result
        }));
    };
}

function SetBiography() {
    lobby.send(JSON.stringify({
        event: "SetBiography",
        biography: document.getElementById("userBiography").value,
    }));
}