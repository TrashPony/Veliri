function CreateInventoryMenu() {
    ConnectInventory();

    setTimeout(function(){
        inventory.send(JSON.stringify({
            event: "openInventory"
        }));
    }, 2000);


}