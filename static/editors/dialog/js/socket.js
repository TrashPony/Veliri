let editor;

function OpenMapEditorSocket() {
    editor = new WebSocket("ws://" + window.location.host + "/wsDialogEditor");

    editor.onopen = function () {
        console.log("open socket...");
        GetListDialogs();
    };
    editor.onmessage = function (msg) {
        reader(JSON.parse(msg.data));
    };
    editor.onerror = function (msg) {
        console.log("Error mapEditor occured sending..." + msg.data);
    };
    editor.onclose = function () {
        console.log("Disconnected mapEditor - status " + this.readyState);
        location.href = "../../login";
    };
}

function reader(data) {
    if (data.event === "OpenEditor") {
        CreateDialogList(data.dialogs);
    }

    if (data.event === "GetDialog") {
        ViewDialog(data.dialog);
    }
}

function CreateDialog() {
    if (document.getElementById("nameNewDialog").value === "") {
        alert("Имя не может быть пустым");
        return
    }

    editor.send(JSON.stringify({
        event: "CreateDialog",
        name: document.getElementById("nameNewDialog").value,
    }));

    document.getElementById("nameNewDialog").value = "";
}

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