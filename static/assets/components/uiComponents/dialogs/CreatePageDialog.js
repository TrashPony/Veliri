function CreatePageDialog(id, page, action, full, needPicture, credits, items, noCreate) {

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

    let dialogText = CreateText(dialogBlock, page);

    if (full) {
        let buttons = CreateControlButtons("83px", "-8px", "-3px", "29px", "", "145px");
        $(buttons.move).mousedown(function (event) {
            moveWindow(event, id)
        });
        dialogBlock.appendChild(buttons.move);

        CreateAsk(dialogBlock, page, true);
    }

    if (credits || items) {
        dialogBlock.style.height = "140px";
        CreateItems(dialogText, credits, items);
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
    return dialogText;
}

function CreateItems(dialogBlock, items, credits) {
    let dialogItems = document.createElement("div");
    dialogItems.className = "dialogItems";
    dialogBlock.appendChild(dialogItems);
    let dialogCredits = document.createElement("div");
    dialogCredits.className = "dialogCredits";
    dialogItems.appendChild(dialogCredits);
    dialogCredits.innerHTML = `
        <div>Кредиты: </div>
        <div> ${credits} </div>
    `;

    let dialogSlots = document.createElement("div");
    dialogSlots.className = "dialogSlots";
    dialogItems.appendChild(dialogSlots);

    for (let i in items) {
        if (items.hasOwnProperty(i) && items[i].item && items[i].quantity > 0) {
            let cell = document.createElement("div");
            CreateInventoryCell(cell, items[i], i, "");
            $(cell).draggable({
                disabled: true,
            });
            $(dialogSlots).append(cell);
        }
    }
}

function CreateAsk(dialogBlock, page, deletePage) {
    for (let i in page.asc) {
        let ask = document.createElement("div");
        ask.className = "asks";
        ask.innerHTML = "<div class='wrapperAsk' id='ask" + page.asc[i].id + "'>" + page.asc[i].text + "</div>";

        $(ask).click(function () {
            chat.send(JSON.stringify({
                event: "Ask",
                to_page: page.asc[i].to_page,
                ask_id: page.asc[i].id,
            }));
            if (page.asc[i].to_page === 0 && deletePage) {
                dialogBlock.remove();
            }
        });

        dialogBlock.appendChild(ask);
    }
}