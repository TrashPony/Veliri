let hideInventory = false;
function OpenInventory() {
    if (document.getElementById("Inventory")) {
        document.getElementById("Inventory").remove();
        return
    }
    InitInventoryMenu(null, 'inventory');

    setTimeout(function () { // todo костыль
        document.getElementById("utilButton").remove();
        document.getElementById("destroyButton").style.left = "78px";
        document.getElementById("InventoryHead").style.margin = "2px 0px 3px 6px";
        document.getElementById("Inventory").style.marginLeft = "0";

        let buttons = CreateControlButtons("0", "23px", "-3px", "-3px");
        buttons.hide.style.width = "22px";
        buttons.hide.onclick = function () {
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
        document.getElementById("Inventory").appendChild(buttons.hide);

        buttons.move.style.width = "22px";
        buttons.move.onmousedown = function (event) {
            moveWindow(event, 'Inventory')
        };
        document.getElementById("Inventory").appendChild(buttons.move);

    }, 500)

}

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

        let buttons = CreateControlButtons("3px", "35px", "-3px", "-3px");
        buttons.move.onmousedown = function (event) {
            moveWindow(event, 'marketBox');
        };
        document.getElementById("marketBox").appendChild(buttons.move);
    }, 500)
}