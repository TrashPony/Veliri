function CreateControlButtons(top, moveRight, closeRight, hideRight, headText, width) {
    let move = document.createElement("div");
    move.className = "topButton";
    move.innerText = "⇿";
    move.style.position = "absolute";
    move.style.top = top;
    move.style.right = moveRight;
    move.style.fontSize = "20px";

    let close = document.createElement("div");
    close.className = "topButton";
    close.innerHTML = "&#10006;";
    close.style.position = "absolute";
    close.style.top = top;
    close.style.right = closeRight;
    close.style.lineHeight = "16px";
    close.style.fontSize = "16px";

    let hide = document.createElement("div");
    hide.className = "topButton";
    hide.innerText = "_";
    hide.style.position = "absolute";
    hide.style.top = top;
    hide.style.right = hideRight;
    hide.style.lineHeight = "0";

    let head = document.createElement("div");
    head.className = "windowsHead";
    head.innerText = headText;
    head.style.width = width;

    return {move: move, close: close, hide: hide, head: head}
}