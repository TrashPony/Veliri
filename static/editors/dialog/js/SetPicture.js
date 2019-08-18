/**
 * т.к. картинки имеют приватное поле то они обновляются отдельно
 **/

function ChangePicOption(idPage) {
    let onePictureCheckBox = document.getElementById("onePictureCheckBox" + idPage);
    let pictures = document.getElementById("pictures" + idPage);

    if ($(onePictureCheckBox).prop("checked")) {
        pictures.innerHTML = `
            <div class="pic" id="picMainBlock${idPage}">
            
                <label for="picMain${idPage}"> Загрузить фаил</label>
                <input type="file" id="picMain${idPage}" onchange="SelectDialogFile(event, 'main', ${idPage})">
            </div>
        `;
    } else {
        SelectDialogFile(null, 'main', idPage);
        pictures.innerHTML = `
            <div class="pic Explores" id="picExploresBlock${idPage}">
            
                <label for="picExplores${idPage}"> Загрузить фаил</label>
                <input type="file" id="picExplores${idPage}" onchange="SelectDialogFile(event, 'Explores', ${idPage})">
            </div>
             <div class="pic Replics" id="picReplicsBlock${idPage}">
            
                <label for="picReplics${idPage}"> Загрузить фаил</label>
                <input type="file" id="picReplics${idPage}" onchange="SelectDialogFile(event, 'Replics', ${idPage})">
            </div>
             <div class="pic Reverses" id="picReversesBlock${idPage}">
            
                <label for="picReverses${idPage}"> Загрузить фаил</label>
                <input type="file" id="picReverses${idPage}" onchange="SelectDialogFile(event, 'Reverses', ${idPage})">
            </div>
        `
    }

    GetDialogPic(idPage);
}

function SelectDialogFile(e, type, id_page) {
    if (e) {
        let file_reader = new FileReader(e.target.files[0]);
        file_reader.readAsDataURL(e.target.files[0]);
        file_reader.onload = function (evt) {
            editor.send(JSON.stringify({
                event: "SetPicture",
                file: evt.target.result,
                fraction: type,
                id_page: id_page,
                id: selectDialog.id,
            }));
        };
    } else {
        editor.send(JSON.stringify({
            event: "SetPicture",
            file: "",
            fraction: type,
            id_page: id_page,
            id: selectDialog.id,
        }));
    }

    setTimeout(function () {
        GetDialogPic(id_page)
    }, 300);
}

function GetDialogPic(idPage) {

    let onePictureCheckBox = document.getElementById("onePictureCheckBox" + idPage);

    GetDialogPicture(idPage, -1).then(function (response) {
        if ($(onePictureCheckBox).prop("checked")) {
            document.getElementById("picMainBlock" + idPage).style.backgroundImage = "url('" + response.data['main'] + "')";
        } else {
            document.getElementById("picExploresBlock" + idPage).style.backgroundImage = "url('" + response.data['Explores'] + "')";
            document.getElementById("picReplicsBlock" + idPage).style.backgroundImage = "url('" + response.data['Replics'] + "')";
            document.getElementById("picReversesBlock" + idPage).style.backgroundImage = "url('" + response.data['Reverses'] + "')";
        }
    });
}