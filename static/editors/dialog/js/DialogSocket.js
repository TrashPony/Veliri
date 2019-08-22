function DeleteDialog(dialogId) {
    editor.send(JSON.stringify({
        event: "DeleteDialog",
        id: Number(dialogId),
    }));
}

function SelectDialog(dialogId) {
    editor.send(JSON.stringify({
        event: "GetDialog",
        id: Number(dialogId),
    }));
}

function GetListDialogs() {
    editor.send(JSON.stringify({
        event: "OpenEditor"
    }));
}

function SaveDialog() {
    editor.send(JSON.stringify({
        event: "SaveDialog",
        dialog: selectDialog,
    }));
}

function AddPage() {
    editor.send(JSON.stringify({
        event: "AddPage",
        dialog: selectDialog,
    }));
}

function RemovePage(pageID) {

    for (let i in selectDialog.pages) {
        if (selectDialog.pages[i].id === pageID) {
            delete selectDialog.pages[i];
        }
    }

    SaveDialog();
}

function AddAsc(pageID) {
    let page = getPageByID(pageID);
    if (!page.asc) page.asc = [];

    page.asc.push({
        id: 0,
        name: "",
        text: "",
        to_page: 0,
        type_action: "",
    });

    SaveDialog();
}

function RemoveAsc(pageID, ascID) {
    let page = getPageByID(pageID);
    for (let i in page.asc) {
        if (page.asc[i].id === Number(ascID)) {
            page.asc.splice(i, 1);
        }
    }

    SaveDialog();
}