function InitGame() {
    field.send(JSON.stringify({
        event: "InitGame",
    }));
}