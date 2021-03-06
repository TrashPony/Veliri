let positionsInterval = {};
let checksTimeOut = {};

function mouseBodyOver(body, unit, unitBox) {
    body.events.onInputOver.add(function () {

        for (let j in unit.body.weapons) {
            if (unit.body.weapons.hasOwnProperty(j) && unit.body.weapons[j].weapon) {
                let graphics = game.add.graphics(0, 0);
                unit.AttackLine = {
                    graphics: graphics,
                    minRadius: unit.body.weapons[j].weapon.min_attack_range,
                    maxRadius: unit.body.weapons[j].weapon.range
                };
                game.floorObjectLayer.add(graphics);
            }
        }

        //unitInfo(unit, unitBox)
    }, this);

    body.events.onInputOut.add(function () {
        if (unit.AttackLine) {
            unit.AttackLine.graphics.destroy();
            unit.AttackLine = null;
        }
        unitRemoveInfo(unit, unitBox)
    }, this);
}

function unitInfo(unit, unitBox) {
    if (game.user_id === unit.owner_id) {
        unitBox.frame = 1;
    }

    clearTimeout(checksTimeOut[unit.id]);
    checksTimeOut[unit.id] = null;

    if (document.getElementById("UserLabel" + unit.owner + unit.id)) {
        return;
    }

    // let userLabel = document.createElement('div');
    // userLabel.id = "UserLabel" + unit.owner + unit.id;
    // userLabel.className = "UserLabel";
    // document.body.appendChild(userLabel);
    // userLabel.innerHTML = `
    //         <div>
    //             <div>
    //                 <div class="logo" id="userAvatar${unit.owner_id}${unit.id}" ></div>
    //                 <h4>${unit.owner}</h4>
    //                 <div class="detailUser" onmousedown="informationFunc('${unit.owner}', '${unit.owner_id}')">i</div>
    //             </div>
    //         </div>
    //     `;

    positionsInterval[unit.id] = setInterval(function () {
        // userLabel.style.left = unitBox.worldPosition.x - 50 + "px";
        // userLabel.style.top = unitBox.worldPosition.y - 70 + "px";
        // userLabel.style.display = "block";
    }, 10);

    GetUserAvatar(unit.owner_id).then(function (response) {
        $("#userAvatar" + unit.owner_id + unit.id).css('background-image', "url('" + response.data.avatar + "')");
    });
}

function unitRemoveInfo(unit, unitBox) {
    if (!GetSelectUnitByID(unit.id)) {
        unitBox.frame = 0;
    }

    if (!checksTimeOut[unit.id]) {
        checksTimeOut[unit.id] = setTimeout(function () {
            if (document.getElementById("UserLabel" + unit.owner + unit.id)) document.getElementById("UserLabel" + unit.owner + unit.id).remove();
            clearInterval(positionsInterval[unit.id]);
            clearTimeout(checksTimeOut[unit.id]);
            checksTimeOut[unit.id] = null;
        }, 2000);
    }
}