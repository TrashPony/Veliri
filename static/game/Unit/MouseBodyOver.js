let positionInterval = null;
let checkTimeOut = null;

function mouseBodyOver(body, unit, unitBox, userID) {

    body.events.onInputOver.add(function () {
        unitInfo(unit, unitBox, userID)
    }, this);

    body.events.onInputOut.add(function () {
        unitRemoveInfo(unit, unitBox, userID)
    }, this);
}

function unitInfo(unit, unitBox, userID) {
    if (Data.squad.user_id === unit.user_id) {
        unitBox.frame = 1;
    } else {
        //todo враг красны, нейтрал белый
        unitBox.frame = 2;
    }

    clearTimeout(checkTimeOut);
    checkTimeOut = null;

    if (document.getElementById("UserLabel" + unit.owner + unit.id)) {
        return;
    }

    let userLabel = document.createElement('div');
    userLabel.id = "UserLabel" + unit.owner + unit.id;
    userLabel.className = "UserLabel";
    document.body.appendChild(userLabel);

    userLabel.innerHTML = `
            <div>
                <div>
                    <div class="logo" id="userAvatar${userID}${unit.id}" ></div>
                    <h4>${unit.owner}</h4>
                    <div class="detailUser" onmousedown="informationFunc('${unit.owner}', '${unit.owner_id}')">i</div>
                </div>
            </div>
        `;

    positionInterval = setInterval(function () {
        userLabel.style.left = unitBox.worldPosition.x - 50 + "px";
        userLabel.style.top = unitBox.worldPosition.y - 70 + "px";
        userLabel.style.display = "block";
    }, 10);

    GetUserAvatar(userID).then(function (response) {
        $("#userAvatar" + userID + unit.id).css('background-image', "url('" + response.data.avatar + "')");
    });
}

function unitRemoveInfo(unit, unitBox) {
    if (!GetSelectUnitByID(unit.id)) {
        unitBox.frame = 0;
    }

    if (!checkTimeOut) {
        checkTimeOut = setTimeout(function () {
            if (document.getElementById("UserLabel" + unit.owner + unit.id)) document.getElementById("UserLabel" + unit.owner + unit.id).remove();
            clearInterval(positionInterval);
            clearTimeout(checkTimeOut);
            checkTimeOut = null;
        }, 2000);
    }
}