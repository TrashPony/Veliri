function InitDepartmentOfEmployment() {

    if (document.getElementById('DepartmentOfEmployment')) {
        document.getElementById('DepartmentOfEmployment').remove()
        return
    }

    let departmentOfEmployment = document.createElement('div');
    departmentOfEmployment.id = 'DepartmentOfEmployment';
    document.body.appendChild(departmentOfEmployment);

    departmentOfEmployment.innerHTML = (`
            <h3 id="missionHead">MissionName</h3>
            <div id="infoBlock">
                <div id="missionText"> Text text text text text text text text text</div>
                <div id="missionAsc">
                    <div class="asks"><div class='wrapperAsk'>Понятно</div></div>
                    <div class="asks"><div class='wrapperAsk'>Не Понятно</div></div>
                </div>
            </div>
            
            <div id="rewardBlock">
                <div id="missionFace" style="background-image: url('../assets/replics_logo.png')"></div>
               
                <h3>Награда:</h3>
                <div id="rewards">
                   <div id="rewardsCredits">Крудиты: <span id="countRewardCredits">250</span></div>
                   <div id="rewardsItems"></div>
                </div>
            </div>   
    `);

    let buttons = CreateControlButtons("2px", "31px", "-3px", "29px", "", "105px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'DepartmentOfEmployment')
    });
    $(buttons.close).mousedown(function () {
        departmentOfEmployment.remove();
    });
    departmentOfEmployment.appendChild(buttons.move);
    departmentOfEmployment.appendChild(buttons.close);

    $(departmentOfEmployment).resizable({
        minHeight: 400,
        minWidth: 400,
        handles: "se",
        resize: function (event, ui) {

        }
    });
}