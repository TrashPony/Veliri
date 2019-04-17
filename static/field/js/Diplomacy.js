function OpenDiplomacy() {
    field.send(JSON.stringify({
        event: "OpenDiplomacy",
    }));
}

function CreateDiplomacyMenu(data) {
    if (document.getElementById("diplomacyBlock")) {
        document.getElementById("diplomacyBlock").remove();
    }

    let diplomacyBlock = document.createElement("div");
    diplomacyBlock.id = "diplomacyBlock";
    document.body.appendChild(diplomacyBlock);

    diplomacyBlock.innerHTML = `
    <h3>Меню дипломатии</h3>
    <div class="wrapperTableDiplomacy">
        <table id="diplomacyTable">
        
        </table>
    </div>`;

    let buttons = CreateControlButtons("0px", "31px", "-3px", "");
    buttons.close.onclick = function () {
        diplomacyBlock.remove();
    };
    diplomacyBlock.appendChild(buttons.close);
    buttons.move.onmousedown = function (event) {
        moveWindow(event, "diplomacyBlock")
    };
    diplomacyBlock.appendChild(buttons.move);

    createDiplomacyTable(data);
    fillDiplomacyTable(data);

    console.log(data)
}

function fillDiplomacyTable(data) {
    for (let i = 0; i < data.users_name.length; i++) {
        for (let j = 0; j < data.users_name.length; j++) {

            let cell = document.getElementById(data.users_name[i] + ':' + data.users_name[j]);

            if (cell && cell.className === '' && (data.users_name[i] === game.user.name || data.users_name[j] === game.user.name)) {
                let sendUser;

                if (findPack(data, data.users_name[j], data.users_name[i])) {
                    cell.innerHTML = `
                    <div style="margin: 0 auto; width: 48px">
                        <div class='DiplomacyButton' style='background-image: url(https://img.icons8.com/doodle/48/000000/filled-flag.png)'></div>
                    </div>`

                } else {

                    if (data.users_name[j] === game.user.name) sendUser = data.users_name[i];
                    if (data.users_name[i] === game.user.name) sendUser = data.users_name[j];

                    cell.innerHTML = `
                    <div style="margin: 0 auto; width: 48px">
                        <div class='DiplomacyButton' onclick="SendRequestPact(\'${sendUser}\')" style='background-image: url(https://img.icons8.com/doodle/48/000000/handshake.png)'></div>
                        <div class='DiplomacyButton' onclick="BuyOutPact(\'${sendUser}\')" style='background-image: url(https://img.icons8.com/office/16/000000/coins.png)'></div>
                    </div>`
                }
            }
        }
    }
}

function findPack(data, user1, user2) {
    for (let i = 0; i < data.diplomacy_user.length; i++) {
        if ((user1 === data.diplomacy_user[i].user_name_1 && user2 === data.diplomacy_user[i].user_name_2) ||
            (user1 === data.diplomacy_user[i].user_name_2 && user2 === data.diplomacy_user[i].user_name_1)) {
            return true
        }
    }
    return false
}

function createDiplomacyTable(data) {
    // заполняем ячейки с юзерами
    let diplomacyTable = $('#diplomacyTable');

    let row = document.createElement("tr");
    let cell = document.createElement("td");
    row.appendChild(cell);

    diplomacyTable.append(row);
    for (let i = 0; i < data.users_name.length; i++) {
        let cell = document.createElement("td");
        cell.innerHTML = data.users_name[i];
        row.appendChild(cell);
    }

    for (let i = 0; i < data.users_name.length; i++) {
        let row = document.createElement("tr");
        diplomacyTable.append(row);

        let cell = document.createElement("td");
        cell.innerHTML = data.users_name[i];
        row.appendChild(cell);

        for (let j = 0; j < data.users_name.length; j++) {
            let cell = document.createElement("td");

            if (data.users_name[i] === data.users_name[j] || j < i) {
                cell.className = "noActive"
            } else {
                cell.id = data.users_name[i] + ':' + data.users_name[j];
            }

            row.appendChild(cell);
        }
    }
}

function BuyOutPact(user) {
    field.send(JSON.stringify({
        event: "ArmisticePact",
        to_user: user,
    }));
}

function BuyOutMenu(data) {

}

function SendRequestPact(userName) {
    field.send(JSON.stringify({
        event: "ArmisticePact",
        to_user: userName,
    }));
}

function CreateDiplomacyRequests(jsonData) {
    let page = {
        text: "Игрок " + jsonData.to_user + " предлагает перемирие!",
        picture: "base.png",
        asc: [],
    };

    let dialogBlock = CreatePageDialog("DiplomacyRequestsBlock", page, null, false, true);
    dialogBlock.style.right = "calc(50% - 125px)";
    dialogBlock.style.top = "calc(50% - 300px)";
    dialogBlock.style.bottom = "unset";
    dialogBlock.style.left = "unset";

    let ask = document.createElement("div");
    ask.className = "asks";
    ask.innerHTML = "<div class='wrapperAsk'>Принять</div>";
    ask.onclick = function () {
        field.send(JSON.stringify({
            event: "AcceptArmisticePact",
            accept: true,
            to_user: jsonData.to_user,
        }));
        dialogBlock.remove();
    };

    let ask2 = document.createElement("div");
    ask2.className = "asks";
    ask2.innerHTML = "<div class='wrapperAsk'>Отказать</div>";
    ask2.onclick = function () {
        field.send(JSON.stringify({
            event: "AcceptArmisticePact",
            accept: false,
            to_user: jsonData.to_user,
        }));
        dialogBlock.remove();
    };

    dialogBlock.appendChild(ask);
    dialogBlock.appendChild(ask2);
}