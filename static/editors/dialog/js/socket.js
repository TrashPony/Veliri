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

    if (data.event === "GetAllMissions") {
        MissionList(data.missions)
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