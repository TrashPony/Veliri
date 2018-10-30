function createNavInfo(menu) {
    let info = document.createElement("div");
    info.className = "infoParams";
    info.id = "infoParamsNav";
    info.style.display = "none";
    menu.appendChild(info);
}