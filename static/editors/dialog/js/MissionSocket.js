function GetListMissions() {
    editor.send(JSON.stringify({
        event: "GetAllMissions"
    }));
}