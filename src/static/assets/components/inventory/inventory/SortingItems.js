function SortingItems() {
    inventorySocket.send(JSON.stringify({
        event: "SortItems",
        type: this.sort
    }));
}