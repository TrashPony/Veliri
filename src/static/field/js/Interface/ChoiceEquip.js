function ChoiceEquip() {
    var inventory = document.getElementById("inventory");

    if (!inventory) {
        inventory = document.createElement("div");
        inventory.id = "inventory";
    } else {
        inventory.style.display = "inline-block";
    }
    //todo заголовок, кнопка отмены, заполнение
    inventory.style.top = stylePositionParams.top - 87 + 'px';
    inventory.style.left = 70 + stylePositionParams.left + 'px';

    document.body.appendChild(inventory);
}