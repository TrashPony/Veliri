function CreatePageDialog(id, page, action, full, needPicture) {

    DialogAction(action);

    if (!page) {
        return
    }

    let dialogBlock = document.getElementById(id);

    if (!dialogBlock) {
        dialogBlock = document.createElement("div");
        dialogBlock.id = id;
        dialogBlock.className = "dialogBlock";
    } else {
        $(dialogBlock).empty()
    }

    if (needPicture) {
        CreatePic(dialogBlock, page)
    }

    CreateText(dialogBlock, page);

    if (full) {
        let buttons = CreateControlButtons("83px", "30px", "-3px", "29px", "", "145px");
        $(buttons.move).mousedown(function (event) {
            moveWindow(event, id)
        });
        $(buttons.close).mousedown(function () {
            dialogBlock.remove();
        });
        dialogBlock.appendChild(buttons.move);
        dialogBlock.appendChild(buttons.close);

        CreateAsk(dialogBlock, page);
    }

    document.body.appendChild(dialogBlock);
    return dialogBlock;
}

function CreatePic(dialogBlock, page) {
    let picture = document.createElement("div");
    picture.id = "dialogPicture";
    picture.innerHTML = "<div class='nameDialog'> Какой - то хер</div>";
    dialogBlock.appendChild(picture);

    let pictureBack = document.createElement("div");
    pictureBack.id = "pictureBack";
    pictureBack.style.backgroundImage = "url(../assets/dialogPictures/" + page.picture + ")";
    picture.appendChild(pictureBack);
}

function CreateText(dialogBlock, page) {
    let dialogText = document.createElement("div");
    dialogText.className = "dialogText";
    dialogText.innerHTML = "<div class='wrapperText'>" + page.text + "</div>";
    dialogBlock.appendChild(dialogText);
}

function CreateAsk(dialogBlock, page) {
    for (let i in page.asc) {
        let ask = document.createElement("div");
        ask.className = "asks";
        ask.innerHTML = "<div class='wrapperAsk'>" + page.asc[i].text + "</div>";

        $(ask).click(function () {
            lobby.send(JSON.stringify({
                event: "Ask",
                to_page: page.asc[i].to_page,
                ask_id: page.asc[i].id,
            }));
            if (page.asc[i].to_page === 0) {
                dialogBlock.remove();
            }
        });

        dialogBlock.appendChild(ask);
    }
}