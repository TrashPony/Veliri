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

function DeleteMission(missionID) {
    let mission = getMissionByID(missionID);
    editor.send(JSON.stringify({
        event: "DeleteMission",
        mission: mission,
    }));
}

function AddMission() {
    let name = document.getElementById("nameNewMission").value;
    if (name === "") {
        alert("Укажите имя");
        return
    }

    editor.send(JSON.stringify({
        event: "AddMission",
        name: name,
    }));
}