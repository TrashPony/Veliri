function StopActions() {
    global.send(JSON.stringify({
        event: "StopAll",
        units_id: getIDsSelectUnits(),
    }));
}