function Stop() {
    global.send(JSON.stringify({
        event: "StopMove",
    }));
}