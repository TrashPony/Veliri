function MyOrdersTab(myMarket, allMarket) {
    myMarket.onclick = null;
    allMarket.onclick = function () {
        AllOrdersTab(myMarket, allMarket)
    };
    if (document.getElementById("sellOrdersBlock")) {
        document.getElementById("sellOrdersBlock").style.display = "none"
    }
    if (document.getElementById("BuyOrdersBlock")) {
        document.getElementById("BuyOrdersBlock").style.display = "none"
    }
    if (document.getElementById("selectItemIcon")) {
        document.getElementById("selectItemIcon").style.display = "none"
    }
    if (document.getElementById("selectItemName")) {
        document.getElementById("selectItemName").style.display = "none"
    }

    let headers = document.getElementsByClassName("ordersHead");
    for (let i = 0; i < headers.length; i++) {
        headers[i].style.display = "none";
    }

    myMarket.className = "activePin";
    allMarket.className ="";

    if (!document.getElementById("MyOrdersBlock")) {
        let myOrdersBlock = document.createElement("div");
        myOrdersBlock.id = "MyOrdersBlock";
        document.getElementById("ordersBlock").appendChild(myOrdersBlock);
    } else {
        document.getElementById("MyOrdersBlock").style.display = "block"
    }

    marketSocket.send(JSON.stringify({
        event: "getMyOrders"
    }));
}

function AllOrdersTab(myMarket, allMarket) {
    allMarket.onclick = null;
    myMarket.onclick = function () {
        MyOrdersTab(myMarket, allMarket)
    };

    if (document.getElementById("MyOrdersBlock")) {
        document.getElementById("MyOrdersBlock").style.display = "none"
    }
    if (document.getElementById("sellOrdersBlock")) {
        document.getElementById("sellOrdersBlock").style.display = "block"
    }
    if (document.getElementById("BuyOrdersBlock")) {
        document.getElementById("BuyOrdersBlock").style.display = "block"
    }
    if (document.getElementById("selectItemIcon")) {
        document.getElementById("selectItemIcon").style.display = "block"
    }
    if (document.getElementById("selectItemName")) {
        document.getElementById("selectItemName").style.display = "block"
    }

    let headers = document.getElementsByClassName("ordersHead");
    for (let i = 0; i < headers.length; i++) {
        headers[i].style.display = "block";
    }

    allMarket.className = "activePin";
    myMarket.className ="";

    marketSocket.send(JSON.stringify({
        event: "getMyOrders"
    }));
}