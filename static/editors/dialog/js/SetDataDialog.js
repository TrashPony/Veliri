/**
 * Свойства диалога
 **/
function SetNameDialog(context) {
    selectDialog.name = context.value;
}

function SetAccessType(context) {
    selectDialog.access_type = context.value;
}

function SetDialogFraction(context) {
    selectDialog.fraction = context.value;

}

function SetTypeDialog(context) {
    selectDialog.type = context.value;
}

/**
 * Свойства страницы
 **/

function getPageByID(id) {
    for (let i in selectDialog.pages) {
        if (selectDialog.pages.hasOwnProperty(i) && selectDialog.pages[i].id === Number(id)) {
            return selectDialog.pages[i];
        }
    }
}

function SetDialogPageHead(context, idPage) {
    let page = getPageByID(idPage);
    page.name = context.value;
}

function SetDialogPageText(context, idPage) {
    let page = getPageByID(idPage);
    page.text = context.value;
}

/**
 * Свойства ответов
 **/

function getAscByID(id) {
    for (let i in selectDialog.pages) {

        if (selectDialog.pages.hasOwnProperty(i)) {

            for (let j in selectDialog.pages[i].asc) {

                if (selectDialog.pages[i].asc.hasOwnProperty(j) && selectDialog.pages[i].asc[j].id === Number(id)) {
                    return selectDialog.pages[i].asc[j];
                }
            }
        }
    }
}

function SetDialogAscText(context, idAsc) {
    let asc = getAscByID(idAsc);
    asc.text = context.value;
}

function SetDialogAscToPageNumber(context, idAsc) {
    let asc = getAscByID(idAsc);
    lightToPage(asc.id, false);
    asc.to_page = Number(context.value);
}

function SetDialogAscTypeAction(context, idAsc) {
    let asc = getAscByID(idAsc);
    asc.type_action = context.value;
}