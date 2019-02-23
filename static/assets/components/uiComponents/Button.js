function createInput(value, parrent) {
    let button = document.createElement("input");
    button.type = "button";
    button.className = "lobbyButton inventoryTip";
    button.value = value;
    button.style.pointerEvents = "auto";
    parrent.appendChild(button);
    return button
}