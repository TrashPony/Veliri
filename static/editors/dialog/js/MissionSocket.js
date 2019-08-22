function GetListMissions() {
    editor.send(JSON.stringify({
        event: "GetAllMissions"
    }));
}

function SaveMission(missionID) {
    let mission = getMissionByID(missionID);
    editor.send(JSON.stringify({
        event: "SaveMissions",
        mission: mission,
    }));
}