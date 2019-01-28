function OutBase() {
    lobby.send(JSON.stringify({
        event: "OutBase"
    }));
}

function Logout() {
    lobby.send(JSON.stringify({
        event: "Logout"
    }));
}