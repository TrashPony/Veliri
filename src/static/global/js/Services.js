let hideInventory = false;
function OpenInventory() {
    if (document.getElementById("Inventory")) {
        document.getElementById("Inventory").remove();
        return
    }
    InitInventoryMenu(null, 'inventory');

    setTimeout(function () { // todo костыль
        document.getElementById("utilButton").remove();
        document.getElementById("destroyButton").style.left = "80px";
        document.getElementById("InventoryHead").style.margin = "2px 0px 3px 6px";
        document.getElementById("Inventory").style.marginLeft = "0";

        let hideButton = document.createElement("div");
        hideButton.className = "topButton";
        hideButton.innerText = "_";
        hideButton.style.position = "absolute";
        hideButton.style.top = "0";
        hideButton.style.left = "135px";
        hideButton.style.width = "22px";
        hideButton.style.lineHeight = "0";
        hideButton.onclick = function () {
            if (!hideInventory) {
                document.getElementById("Inventory").style.height = "14px";
                document.getElementById("Inventory").style.overflow = "hidden";
                document.getElementById("Inventory").style.bottom = "363px";
                document.getElementById("destroyButton").style.opacity = "0";
                document.getElementById("InventoryHead").style.margin = "-5px 0px 2px 0px";
                hideInventory = true;
            } else {
                document.getElementById("Inventory").style.height = "307px";
                document.getElementById("Inventory").style.overflow = "visible";
                document.getElementById("Inventory").style.bottom = "70px";
                document.getElementById("destroyButton").style.opacity = "1";
                document.getElementById("InventoryHead").style.margin = "2px 0px 3px 6px";
                hideInventory = false;
            }
        };
        document.getElementById("Inventory").appendChild(hideButton);

        let moveButton = document.createElement("div");
        moveButton.className = "topButton";
        moveButton.innerText = "⇿";
        moveButton.style.position = "absolute";
        moveButton.style.top = "0";
        moveButton.style.left = "108px";
        moveButton.style.width = "22px";
        moveButton.style.fontSize = "20px";
        moveButton.onmousedown = function (event) {
            moveWindow(event, 'Inventory')
        };
        document.getElementById("Inventory").appendChild(moveButton);

    }, 500)

}

let hideMarket = false;

function OpenMarket() {
    if (document.getElementById("marketBox")) {
        document.getElementById("marketBox").remove();
        return
    }

    InitMarketMenu(true);

    setTimeout(function () {
        document.getElementById("marketBox").style.marginLeft = "0";
        document.getElementById("marketBox").style.left = "unset";
        document.getElementById("marketBox").style.right = "15px";
        document.getElementById("marketBox").style.backgroundImage = "linear-gradient(1deg,rgba(33, 176, 255, 0.9), rgba(37, 160, 225, 0.9) 6px)"
        let hideButton = document.createElement("div");
        hideButton.className = "topButton";
        hideButton.innerText = "_";
        hideButton.style.position = "absolute";
        hideButton.style.top = "3px";
        hideButton.style.right = "0px";
        hideButton.style.lineHeight = "0";
        hideButton.onclick = function(){
            if (!hideMarket) {
                document.getElementById("marketBox").style.height = "24px";
                document.getElementById("marketBox").style.width = "200px";
                document.getElementById("marketBox").style.overflow = "hidden";
                document.getElementById("headMarket").style.opacity = "0";
                hideMarket = true;
            } else {
                document.getElementById("marketBox").style.height = "600px";
                document.getElementById("marketBox").style.width = "1000px";
                document.getElementById("marketBox").style.overflow = "visible";
                document.getElementById("headMarket").style.opacity = "1";
                hideMarket = false;
            }
        };
        document.getElementById("marketBox").appendChild(hideButton);

        let moveButton = document.createElement("div");
        moveButton.className = "topButton";
        moveButton.innerText = "⇿";
        moveButton.style.position = "absolute";
        moveButton.style.top = "3px";
        moveButton.style.right = "35px";
        moveButton.style.fontSize = "20px";
        moveButton.onmousedown = function (event) {
            moveWindow(event, 'marketBox');
        };
        document.getElementById("marketBox").appendChild(moveButton);
    }, 500)
}