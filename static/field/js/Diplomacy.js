let slots = [];

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
    $(buttons.close).click(function () {
        diplomacyBlock.remove();
    });

    diplomacyBlock.appendChild(buttons.close);
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, "diplomacyBlock")
    });
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
    slots = {};
    field.send(JSON.stringify({
        event: "initBuyOut",
        to_user: user,
    }));
}

function BuyOutMenu(data, buyOuySlots) {
    slots = buyOuySlots;
    let diplomacyBlock = document.getElementById("diplomacyBlock");

    if (document.getElementById("BuyOutMenu")) document.getElementById("BuyOutMenu").remove();
    let BuyOutMenuBlock = document.createElement("div");
    BuyOutMenuBlock.id = "BuyOutMenu";
    BuyOutMenuBlock.innerHTML = `
        <div>
            <input id="buyOutCredits" max="${data.credits}" type="number" placeholder="Кредиты">
            <div id="buyOutItems"></div>
        </div>
        
        <div style="margin-left: 44px;">
            <div id="buyOutUserCredits"> Кредиты: ${data.credits}</div>
            <div id="buyOutInventory"></div>
        </div>
        
        <input type="button" value="Предложить" onclick="SendRequestPact(\'${data.to_user}\')" style="width: 75px; float: left; margin-left: 25px; margin-top: 2px;">
        <input type="button" value="Отмена" onclick="removeBuyOutBlock()" style="width: 75px; float: right; margin-right: 25px; margin-top: 2px;">`;

    diplomacyBlock.appendChild(BuyOutMenuBlock);

    document.getElementById("buyOutCredits").oninput = function () {
        if (this.value > data.credits) {
            this.value = data.credits;
        }
    };

    for (let i in data.inventory.slots) {
        if (data.inventory.slots.hasOwnProperty(i) && data.inventory.slots[i].item && data.inventory.slots[i].quantity > 0) {
            let cell = document.createElement("div");
            CreateInventoryCell(cell, data.inventory.slots[i], i, "");

            $(cell).draggable({
                disabled: true,
            });

            cell.onclick = function () {

                let mask = document.createElement("span");
                mask.id = "buyOutMask";
                $(BuyOutMenuBlock).append(mask);
                // todo немножк говнокода
                let quantityRange = document.createElement("span");
                quantityRange.id = "QuantityRange";
                quantityRange.innerHTML = `
                    <form name="quantityForm" oninput="quantityOut.value = quantity.value">
                        <div class="iconItem" style='background-image: ${cell.style.backgroundImage}'></div>
                        <input name="quantity" id="quantityRangeValue" type="range" min="0" max="${data.inventory.slots[i].quantity}" value="${data.inventory.slots[i].quantity}"> 
                        <output name="quantityOut">${data.inventory.slots[i].quantity}</output>
                    </form>
                    <input type="button" id="BuyOutButton" value="Предложить" style="width: 75px; float: left; margin-left: 20px; margin-top: 2px;">
                    <input type="button" id="BuyOutCancelButton" value="Отмена" style="width: 75px; float: right; margin-right: 20px; margin-top: 2px;">
                `;
                $(BuyOutMenuBlock).append(quantityRange);

                $('#BuyOutButton').click(function () {
                    let quantity = document.getElementById("quantityRangeValue").value;
                    data.inventory.slots[i].quantity -= quantity;

                    let buySlot = Object.assign({}, data.inventory.slots[i]);
                    buySlot.quantity = Number(quantity);

                    if (buyOuySlots[i]) {
                        buyOuySlots[i].quantity = Number(buyOuySlots[i].quantity) + Number(quantity)
                    } else {
                        buyOuySlots[i] = buySlot;
                    }

                    quantityRange.remove();
                    mask.remove();

                    BuyOutMenu(data, buyOuySlots)
                });

                $('#BuyOutCancelButton').click(function () {
                    if (document.getElementById("buyOutMask")) document.getElementById("buyOutMask").remove();
                    if (document.getElementById("QuantityRange")) document.getElementById("QuantityRange").remove();
                })
            };

            $('#buyOutInventory').append(cell);
        }
    }

    for (let i in buyOuySlots) {
        if (buyOuySlots.hasOwnProperty(i) && buyOuySlots[i].item && buyOuySlots[i].quantity > 0) {
            let cell = document.createElement("div");
            CreateInventoryCell(cell, buyOuySlots[i], i, "");

            $(cell).draggable({
                disabled: true,
            });

            cell.onclick = function () {

                let mask = document.createElement("span");
                mask.id = "buyOutMask";
                $(BuyOutMenuBlock).append(mask);

                let quantityRange = document.createElement("span");
                quantityRange.id = "QuantityRange";
                quantityRange.innerHTML = `
                    <form name="quantityForm" oninput="quantityOut.value = quantity.value">
                        <div class="iconItem" style='background-image: ${cell.style.backgroundImage}'></div>
                        <input name="quantity" id="quantityRangeValue" type="range" min="0" max="${buyOuySlots[i].quantity}" value="${buyOuySlots[i].quantity}"> 
                        <output name="quantityOut">${buyOuySlots[i].quantity}</output>
                    </form>
                    <input type="button" id="BuyOutButton" value="Предложить" style="width: 75px; float: left; margin-left: 20px; margin-top: 2px;">
                    <input type="button" id="BuyOutCancelButton" value="Отмена" style="width: 75px; float: right; margin-right: 20px; margin-top: 2px;">
                `;
                $(BuyOutMenuBlock).append(quantityRange);

                $('#BuyOutButton').click(function () {
                    let quantity = document.getElementById("quantityRangeValue").value;
                    buyOuySlots[i].quantity -= quantity;
                    data.inventory.slots[i].quantity = Number(data.inventory.slots[i].quantity) + Number(quantity);

                    quantityRange.remove();
                    mask.remove();

                    BuyOutMenu(data, buyOuySlots)
                });

                $('#BuyOutCancelButton').click(function () {
                    if (document.getElementById("buyOutMask")) document.getElementById("buyOutMask").remove();
                    if (document.getElementById("QuantityRange")) document.getElementById("QuantityRange").remove();
                })
            };

            $('#buyOutItems').append(cell);
        }
    }
}

function removeBuyOutBlock() {
    slots = {};
    if (document.getElementById("BuyOutMenu")) document.getElementById("BuyOutMenu").remove();
}

function SendRequestPact(userName) {

    let credits = 0;
    if (document.getElementById('buyOutCredits'))
        credits = Number(document.getElementById('buyOutCredits').value);

    let sendSlots;
    for (let key in slots) {
        sendSlots = true;
    }

    if (!sendSlots) {
        slots = null;
    }

    field.send(JSON.stringify({
        event: "ArmisticePact",
        to_user: userName,
        slots: slots,
        credits: credits
    }));


    removeBuyOutBlock();
}

function CreateDiplomacyRequests(jsonData) {
    // TODO отрисовка итемов и кредитов если он есть
    console.log(jsonData.diplomacy_request);
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