function LeaveBattle() {
    //диалогове окно о том что вы потеряете всех юнитов что не в трюме с ок и отмена

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
    ask.innerHTML = "<div class='wrapperAsk'>Уходим!</div>";
    ask.onclick = function () {
        field.send(JSON.stringify({
            event: "FleeBattle",
        }));
    };

    let ask2 = document.createElement("div");
    ask2.className = "asks";
    ask2.innerHTML = "<div class='wrapperAsk'>Не уходим!</div>";
    ask2.onclick = function () {
        dialogBlock.remove();
    };

    dialogBlock.appendChild(ask);
    dialogBlock.appendChild(ask2);
}