function InitLeave() {
    field.send(JSON.stringify({
        event: "InitLeave",
    }));
}

function LeaveBattle() {

    let page = {
        text: "Уйти из боя?",
        picture: "base.png",
        asc: [],
    };

    let dialogBlock = CreatePageDialog("LeaveBlock", page, null, false, true);
    dialogBlock.style.right = "calc(50% - 125px)";
    dialogBlock.style.top = "calc(50% - 300px)";
    dialogBlock.style.bottom = "unset";
    dialogBlock.style.left = "unset";

    let ask = document.createElement("div");
    ask.className = "asks";
    ask.innerHTML = "<div class='wrapperAsk'>Уходим! (теряем не погруженых юнитов)</div>";
    ask.onclick = function () {
        field.send(JSON.stringify({
            event: "FleeBattle",
        }));
        dialogBlock.remove();
    };

    let ask2 = document.createElement("div");
    ask2.className = "asks";
    ask2.innerHTML = "<div class='wrapperAsk'>Уходим! но медленно... (ждем 15 сек)</div>";
    ask2.onclick = function () {
        field.send(JSON.stringify({
            event: "softFlee",
        }));
        dialogBlock.remove();
    };

    let ask3 = document.createElement("div");
    ask3.className = "asks";
    ask3.innerHTML = "<div class='wrapperAsk'>Не уходим!</div>";
    ask3.onclick = function () {
        dialogBlock.remove();
    };

    dialogBlock.appendChild(ask);
    dialogBlock.appendChild(ask2);
    dialogBlock.appendChild(ask3);
}

function LeaveTimer(sec) {
    console.log(sec)
    // TODO создание окошка по центру экрана с текстом выйдете из боя через ...  sec sec
}