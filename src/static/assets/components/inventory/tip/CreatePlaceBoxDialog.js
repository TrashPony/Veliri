let pass;

function CreatePlaceBoxDialog(x, y, numberSlot, slot) {
    if (slot.item.protect) {
        pass = {
            one: {number: 0, use: false},
            two: {number: 0, use: false},
            three: {number: 0, use: false},
            four: {number: 0, use: false},
        };

        let sendNewNumber = function (number) {
            let value = "";
            if (pass.four.use) {
                for (let i in pass) {
                    pass[i].use = false;
                }
            }
            for (let i in pass) {
                if (!pass[i].use) {
                    pass[i].number = number;
                    pass[i].use = true;
                    break
                }
            }
            for (let i in pass) {
                value += pass[i].number;
            }
            document.getElementById('passPlaceBox' ).value = value;
        };

        // защищеный ящик просит пароль
        let passBlock = document.createElement("div");
        passBlock.className = "passBlock";
        passBlock.innerHTML = "<h2>Введите пароль</h2>" +
            "<div><input id='passPlaceBox' type='number' disabled min='0' value='0000' max='9999'></div>";

        for (let i = 1; i < 10; i++) {
            let numberBlock = document.createElement("div");
            numberBlock.innerHTML = i;
            numberBlock.className = "numberBlock";
            numberBlock.onclick = function(){
                sendNewNumber(i)
            };
            passBlock.appendChild(numberBlock);
        }

        let numberBlock = document.createElement("div");
        numberBlock.innerHTML = 0;
        numberBlock.className = "numberBlock";
        numberBlock.style.marginLeft = "45px";
        numberBlock.onclick = function(){
            sendNewNumber(0)
        };
        passBlock.appendChild(numberBlock);

        let closeButton = createInput("Отменить", passBlock);
        closeButton.onclick = function () {
            passBlock.remove();
        };

        let placeButton = createInput("Ок", passBlock);
        placeButton.onclick = function () {
            global.send(JSON.stringify({
                event: "placeNewBox",
                slot: numberSlot,
                box_password: Number(document.getElementById("passPlaceBox").value),
            }));
            passBlock.remove();
        };
        document.body.appendChild(passBlock)
    } else {
        global.send(JSON.stringify({
            event: "placeNewBox",
            slot: numberSlot,
        }));
    }
}
