function OpenDiplomacy() {
    field.send(JSON.stringify({
        event: "OpenDiplomacy",
    }));
}

function CreateDiplomacyMenu(data) {
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
            // todo проверка на наличие пакта, если да ставить влажек
            //  https://img.icons8.com/doodle/48/000000/filled-flag.png
            let cell = document.getElementById(data.users_name[i] + data.users_name[j]);
            if (cell && data.users_name[i] === game.user.name) {
                cell.innerHTML = `
                    <div style="margin: 0 auto; width: 48px">
                        <div class='DiplomacyButton' onclick="SendRequestPact(\'${data.users_name[j]}\')" style='background-image: url(https://img.icons8.com/doodle/48/000000/handshake.png)'></div>
                        <div class='DiplomacyButton' onclick="BuyOutPact(\'${data.users_name[j]}\')" style='background-image: url(https://img.icons8.com/office/16/000000/coins.png)'></div>
                    </div>`
            }
        }
    }
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
                cell.id = data.users_name[i] + data.users_name[j];
            }

            row.appendChild(cell);
        }
    }
}

function BuyOutPact(user) {
    console.log(user)
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