function CreatePageDialog(page) {

    console.log(page)

    let dialogBlock = document.createElement("div");
    dialogBlock.id = "dialogBlock";

    let buttons = CreateControlButtons("82px", "30px", "-3px", "29px", "", "145px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'dialogBlock')
    });
    $(buttons.close).mousedown(function () {
        dialogBlock.remove();
    });
    dialogBlock.appendChild(buttons.move);
    dialogBlock.appendChild(buttons.close);

    CreateText(dialogBlock, page);
    CreateAsk(dialogBlock, page);

    document.body.appendChild(dialogBlock);
}

function CreateText(dialogBlock, page) {
    let picture = document.createElement("div");
    picture.id = "dialogPicture";
    picture.innerHTML = "<div class='nameDialog'> Какой - то хер</div>";
    dialogBlock.appendChild(picture);

    let pictureBack = document.createElement("div");
    pictureBack.id = "pictureBack";
    pictureBack.style.backgroundImage = "url(../assets/dialogPictures/" + page.picture + ")";
    picture.appendChild(pictureBack);

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

        // TODO функция нажатия с отправкой сообщения на какую страницу перешел юзер

        dialogBlock.appendChild(ask);
    }
}