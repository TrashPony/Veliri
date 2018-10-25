function ThrowItems() {
    // todo активирует режим масовой переработки итемов
    // todo после нажатия кнопки, она менят свой свет обозначая режим удаления
    // todo пользователь выбирает итемы которые хочет переработать
    // todo когда он выбрал все итемы которые хочет переработать нажимает подтвеждение
    // todo появляетсмя модальное окно с подтвеждение действия
    // todo данные уежают на бекенд, происходит удаление



    // todo создат ьколекцию куда будут клатца выделеные итемы
    // todo пройтись по всему инвентарю и заменить функцию онклик, при нажатие итем добавляется в колекцию

    // todo при снятие выделения онклик для итемов должен возвращаться, колекция чиститься, выделение сниматься

    checkConfirmMenu();

    let ConfirmMenu = document.createElement("div");
    ConfirmMenu.className = "ConfirmInventoryMenu";
    ConfirmMenu.id = "ConfirmInventoryMenu";
    ConfirmMenu.typeAction = "throw";

    let equipButton = document.createElement("div");
    equipButton.innerHTML = "Выбросить";
    equipButton.onclick = function () {
        // todo отправка выделеные итемы на сервер
    };
    ConfirmMenu.appendChild(equipButton);

    let allButton = document.createElement("div");
    allButton.innerHTML = "Отмена";
    allButton.onclick = cancelThrow;
    ConfirmMenu.appendChild(allButton);

    document.getElementById("Inventory").appendChild(ConfirmMenu);

    this.className = "throwButtonActive";
    this.onclick = cancelThrow;
}

function cancelThrow() {
    document.getElementById("ConfirmInventoryMenu").remove();
    document.getElementsByClassName("throwButtonActive")[0].className = "destroyButton";
    document.getElementsByClassName("destroyButton")[0].onclick = ThrowItems;
    // todo выход из режима выделения
}