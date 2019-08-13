function SaveMap() {
    for (let i in responseChangeOption) {
        mapEditor.send(JSON.stringify(responseChangeOption[i]));
    }
    responseChangeOption = []
}