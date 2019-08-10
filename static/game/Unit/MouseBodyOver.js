function mouseBodyOver(body, unit, unitBox, userID) {
    let positionInterval = null;
    let checkTimeOut = null;

    body.events.onInputOver.add(function () {

        clearTimeout(checkTimeOut);

        if (document.getElementById("UserLabel" + unit.user_name + unit.id)) {
            return;
        }

        let userLabel = document.createElement('div');
        userLabel.id = "UserLabel" + unit.user_name + unit.id;
        userLabel.className = "UserLabel";
        document.body.appendChild(userLabel);

        userLabel.innerHTML = `
            <div>
                <div>
                    <div class="logo" id="userAvatar${userID}${unit.id}" ></div>
                    <h4>${unit.user_name}</h4>
                    <div class="detailUser" onmousedown="informationFunc('${unit.user_name}', '${unit.user_id}')">i</div>
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
    }, this);

    body.events.onInputOut.add(function () {
        checkTimeOut = setTimeout(function () {
            if (document.getElementById("UserLabel" + unit.user_name + unit.id)) document.getElementById("UserLabel" + unit.user_name + unit.id).remove();
            clearInterval(positionInterval);
            clearTimeout(checkTimeOut);
        }, 2000);
    }, this);
}