function ReadResponse(jsonData) {
    if (jsonData.event === "InitGame"){
        LoadGame(jsonData);
    }

    if (jsonData.event === "Error"){
        alert(jsonData.error);
    }

    if (jsonData.event === "Move"){

    }
}