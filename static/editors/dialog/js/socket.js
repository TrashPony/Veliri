let editor;

function OpenMapEditorSocket() {
    editor = new WebSocket("ws://" + window.location.host + "/wsDialogEditor");

    editor.onopen = function () {
        console.log("open socket...");
        editor.send(JSON.stringify({
            event: "OpenEditor"
        }));
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
}