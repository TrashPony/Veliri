let mapEditor;

function OpenMapEditorSocket() {
    mapEditor = new WebSocket("ws://" + window.location.host + "/wsMapEditor");

    mapEditor.onopen = function () {
        console.log("open socket...");
        GetMapList();
    };
    mapEditor.onmessage = function (msg) {
        ReaderMapEditor(msg.data);
    };
    mapEditor.onerror = function (msg) {
        console.log("Error mapEditor occured sending..." + msg.data);
    };
    mapEditor.onclose = function () {
        console.log("Disconnected mapEditor - status " + this.readyState);
        location.href = "../../login";
    };
}