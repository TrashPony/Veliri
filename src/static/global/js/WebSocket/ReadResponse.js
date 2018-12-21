function ReadResponse(jsonData) {
    if (jsonData.event === "InitGame") {
        LoadGame(jsonData);
    }

    if (jsonData.event === "Error") {
        alert(jsonData.error);
    }

    if (jsonData.event === "PreviewPath") {
        PreviewPath(jsonData);
    }

    if (jsonData.event === "MoveTo") {
        MoveTo(jsonData);
    }

    if (jsonData.event === "IntoToBase") {
        location.href = "http://" + window.location.host + "/lobby";
    }
}