//todo ультра говнокод, но иного пути я не нашел, код сценария обучения игрока
let interval;

function Training(lvl) {
    if (lvl === 1) {
        let hangarButton = IntoToHangar();
        hangarButton.click(function () {
            progressTraining(lvl);
        })
    }
    if (lvl === 2) {
        let hangarButton = $('#hangarButton');

        interval = setInterval(function () {

            if (document.getElementById("ConstructorBackGround") && document.getElementById("wrapperInventoryAndStorage")) {

                if (document.getElementById("training1IntoHangar")) document.getElementById("training1IntoHangar").remove();
                if (!document.getElementById("training1Block")) {
                    let page = {
                        text: "Отлично перед табой находится меню ангара которое поделен на разделы",
                        picture: "training.png",
                    };
                    let dialogBlock = CreatePageDialog("training1Block", page, null, false, true);
                    dialogBlock.style.right = "calc(50% - 125px)";
                    dialogBlock.style.top = "calc(50% - 90px)";
                    dialogBlock.style.left = "auto";
                    dialogBlock.className += " Training";

                    $('#inventoryBox').append(dialogBlock);

                    hangarButton.css("animation", "none");

                    let ask = document.createElement("div");
                    ask.className = "asks";
                    ask.innerHTML = "<div class='wrapperAsk'>Понятно</div>";
                    $(ask).click(function () {
                        clearInterval(interval);
                        if (document.getElementById("training1Block")) document.getElementById("training1Block").remove();
                        if (document.getElementById("training1SquadsBlock")) document.getElementById("training1SquadsBlock").remove();
                        if (document.getElementById("training1ParamsBlock")) document.getElementById("training1ParamsBlock").remove();
                        if (document.getElementById("training1InvBlock")) document.getElementById("training1InvBlock").remove();
                        if (document.getElementById("training1StorageBlock")) document.getElementById("training1StorageBlock").remove();
                        if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();
                        progressTraining(lvl);
                    });

                    dialogBlock.appendChild(ask);

                    createInfoText($('#SquadsList'), "training1SquadsBlock", -30, -220, 200, 175, "Ангар: Тут находятся " +
                        "активные отряды которыми ты можешь управлять, но одновременно может быть под упралвение только 1", false, 100);
                    createInfoText($('#SquadsList'), "training1ParamsBlock", 100, -220, 200, 175, "Панель которые " +
                        "отображает текущие параметры отряда", false, 100);
                    createInfoText($('#Inventory'), "training1InvBlock", 0, +155, 200, 175, "Трюм: тут хранятся все " +
                        "вещи который находятся в трюме твоего мазершипа, у него есть размер и он не активен если не выбран отряд", false, 110, true);
                    createInfoText($('#storage'), "training1StorageBlock", 0, +175, 200, 175, "Склад: Тут хранятся все вещи" +
                        " которые находятся на базе, так же сюда попадают все купленые созданные и переработаные вещи", false, 100, true);
                    createInfoText($('#Squad'), "training1SquadBlock", 80, 20, 300, 275, "Отсеки для юнитов: тут хранятся" +
                        " роботы на ду которые будут помогать тебе в бою", false, 100);
                }

            } else {
                if (document.getElementById("training1Block")) document.getElementById("training1Block").remove();
                if (document.getElementById("training1SquadsBlock")) document.getElementById("training1SquadsBlock").remove();
                if (document.getElementById("training1ParamsBlock")) document.getElementById("training1ParamsBlock").remove();
                if (document.getElementById("training1InvBlock")) document.getElementById("training1InvBlock").remove();
                if (document.getElementById("training1StorageBlock")) document.getElementById("training1StorageBlock").remove();
                if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();

                IntoToHangar();
            }
        }, 10);
    }
    if (lvl === 3) {

        let cellMS = null;

        interval = setInterval(function () {
            if (document.getElementById("ConstructorBackGround") && document.getElementById("wrapperInventoryAndStorage")) {

                if (document.getElementById("training1IntoHangar")) document.getElementById("training1IntoHangar").remove();
                let MSIcon = $('#MSIcon');
                let storage = $('#storage');

                if (!document.getElementById("training1SquadBlock")) {
                    createInfoText(storage, "training1SquadBlock", 0, +175, 175, 150, "Давай активируем первый " +
                        "мазершип, для этого выдели его, или перетяни в \"место для корпуса\"", true, 100, true);
                    MSIcon.css("animation", "selectMenu 1500ms infinite");
                }

                if ((!cellMS || cellMS.length === 0) && storage) {
                    cellMS = FindCell('', 'MS', storage);
                    for (let i = 0; i < cellMS.length; i++) {
                        $(cellMS[i]).css("animation", "selectMenu 1500ms infinite");
                    }
                }

                if (MSIcon.css("background-image") !== "none") {
                    clearInterval(interval);
                    $(cellMS).css("animation", "none");
                    MSIcon.css("animation", "none");
                    if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();
                    progressTraining(lvl);
                }
            } else {
                if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();
                IntoToHangar();
            }
        }, 10)
    }

    if (lvl === 4) {

        let thorium = null;

        interval = setInterval(function () {
            let storage = $('#storage');
            let thoriumPanel = $('#thorium');
            let thoriumSlots = $('.thoriumSlots');

            if (document.getElementById("ConstructorBackGround") && document.getElementById("wrapperInventoryAndStorage")) {
                if (document.getElementById("training1IntoHangar")) document.getElementById("training1IntoHangar").remove();

                if (!document.getElementById("training1SquadBlock") && document.getElementById("storage") && document.getElementById("thorium")) {

                    createInfoText(storage, "training1SquadBlock", 0, 0, 175, 150, "Когда активирован мазершип то " +
                        "его можно конфигурировать, так же мс не может передвигатся без топлива, давай снарядим наш мс " +
                        "топливом. Топливо - это обогащенный торий, добывается из руд или покупается у других мехов", true, 205, true);

                    createInfoText(thoriumPanel, "training1ThoriumBlock", -160, -100, 300, 275, "Это слоты для топлива," +
                        " топливо может хранится сразу в 3х ячейках а может и в 1. Cверху показатели максимальной " +
                        "скорости и экономии топлива, чем меньше ячеек задействовано тем меньше скорость но больше экономия топлива",
                        false, 110);
                }

                thorium = FindCell('', 'thorium', storage)[0];
                $(thorium).css("animation", "selectMenu2 1500ms infinite");
                thoriumSlots.css("animation", "selectMenu2 1500ms infinite");
                thoriumSlots.each(function () {
                    if ($(this).css("background-image") !== "none") {
                        if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();
                        if (document.getElementById("training1ThoriumBlock")) document.getElementById("training1ThoriumBlock").remove();
                        clearInterval(interval);
                        thoriumSlots.css("animation", "none");
                        $(thorium).css("animation", "none");
                        progressTraining(lvl);
                    }
                })

            } else {
                IntoToHangar();
            }
        }, 10)
    }

    if (lvl === 5) {
        let equips = [];

        interval = setInterval(function () {
            if (document.getElementById("ConstructorBackGround") && document.getElementById("wrapperInventoryAndStorage")) {
                if (document.getElementById("training1IntoHangar")) document.getElementById("training1IntoHangar").remove();

                let storage = $('#storage');
                let powerPanel = $('#powerPanel');
                let MSIcon = $('#MSIcon');
                let inventoryEquipping = $('.inventoryEquipping.active');

                if (!document.getElementById("training1SquadBlock") && document.getElementById("storage")
                    && document.getElementById("powerPanel") && document.getElementById("MSIcon")) {

                    createInfoText(storage, "training1SquadBlock", 0, 0, 175, 150, "Теперь надо снарядить переносимое " +
                        "оборудование, оно служит для самых развличных целей добыча, защита или даже для атаки противников.",
                        true, 135, true);

                    createInfoText(powerPanel, "training1PowerBlock", -110, -145, 205, 180, "Учти что ты не сможет установить " +
                        "оборудования больше чем сможет выдержать реактор",
                        false, 85);

                    createInfoText(MSIcon, "training1EquipBlock", -70, -255, 235, 210, "Ячейки для оборудования " +
                        "делятся на стандарты I, II, III и в ячейку I можно положить только оборудование с " +
                        "стандартом I, оборудование II только в ячеку II и тд.",
                        false, 110);

                    powerPanel.css("animation", "selectMenu2 1500ms infinite");
                }

                if (equips.length === 0) {
                    equips = FindCell('', 'equips', storage);
                    $(equips).each(function () {
                        $(this).css("animation", "selectMenu2 1500ms infinite");
                    })
                }

                inventoryEquipping.each(function () {
                    if ($(this).css("background-image") !== "none") {
                        if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();
                        if (document.getElementById("training1PowerBlock")) document.getElementById("training1PowerBlock").remove();
                        if (document.getElementById("training1EquipBlock")) document.getElementById("training1EquipBlock").remove();
                        clearInterval(interval);
                        powerPanel.css("animation", "none");
                        $(equips).each(function () {
                            $(this).css("animation", "none");
                        });
                        progressTraining(lvl);
                    }
                })

            } else {
                IntoToHangar();
            }
        }, 10)
    }

    if (lvl === 6) {

        let weapons = [];
        let ammo = [];

        interval = setInterval(function () {
            let storage = $('#storage');
            let weaponPanel = $('#MSWeaponPanel');

            if (document.getElementById("ConstructorBackGround") && document.getElementById("wrapperInventoryAndStorage")) {
                if (document.getElementById("training1IntoHangar")) document.getElementById("training1IntoHangar").remove();

                if (!document.getElementById("training1SquadBlock") && document.getElementById("storage") && document.getElementById("MSWeaponPanel")) {

                    createInfoText(storage, "training1SquadBlock", 0, 0, 175, 150, "Установим оружие." +
                        " Оружие устанавливаются в специальные слоты, они подсвечены в интерфейсе красным цветом, а так же" +
                        " для каждого типа оружие есть свои боеприпасы они устанавливаются в малый слот над слотом оружия.",
                        true, 195, true);

                    createInfoText(weaponPanel, "training1WeaponBlock", 0, -210, 205, 180, "у оружия есть стандарт " +
                        "размера, он определяется корпусом некоторые корпуса могут носить любой тип оружия, а другие нет",
                        false, 95);

                    weaponPanel.css("animation", "selectMenu3 1500ms infinite");
                }

                if (weapons.length === 0) {
                    weapons = FindCell('', 'weapons', storage);
                    $(weapons).each(function () {
                        $(this).css("animation", "selectMenu3 1500ms infinite");
                    })
                }

                if (ammo.length === 0) {
                    ammo = FindCell('', 'ammo', storage);
                    $(ammo).each(function () {
                        $(this).css("animation", "selectMenu3 1500ms infinite");
                    })
                }

                $('.inventoryAmmoCell.inventoryEquipping').each(function () {
                    if ($(this).css("background-image") !== "none") {
                        clearInterval(interval);
                        progressTraining(lvl);
                        if (weapons.length > 0) {
                            $(weapons).each(function () {
                                this.style.animation = "none";
                            })
                        }
                        if (ammo.length > 0) {
                            $(ammo).each(function () {
                                this.style.animation = "none";
                            })
                        }
                        weaponPanel.css("animation", "none");
                        if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();
                        if (document.getElementById("training1WeaponBlock")) document.getElementById("training1WeaponBlock").remove();
                    }
                })
            } else {
                if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();
                if (document.getElementById("training1WeaponBlock")) document.getElementById("training1WeaponBlock").remove();
                IntoToHangar();
            }
        }, 10)
    }

    if (lvl === 7) {
        interval = setInterval(function () {
            let storage = $('#storage');
            let unitSlot = $('.inventoryUnit.active');
            let squad = $('#Squad');

            if (document.getElementById("ConstructorBackGround")) {
                if (document.getElementById("training1IntoHangar")) document.getElementById("training1IntoHangar").remove();

                if (!document.getElementById("training1SquadBlock") && document.getElementById("storage") && document.getElementById("Squad")) {

                    createInfoText(storage, "training1SquadBlock", -80, 175, 195, 170, "Отлично теперь осталось" +
                        " снарядить отряд юнитов, они собираются аналогично но имеют намного меньше энергии," +
                        " а так же учитывается обьем снаряжения который они могут в себе переносить. У каждого Мазершипа" +
                        " свой ангар они отличаются как количеством слотов, так и вмещаемыми размерами юнитов, в свою " +
                        "очередь юниты делятся на 3 класса легкие, средние и тяжелые.",
                        true, 260, true);

                    createInfoText(squad, "training1Squad2Block", 70, 15, 300, 275, "Пиктограмма под активным слотом " +
                        "говорит какой тип юнита может быть помещен в слот, можно помещать более легких юнитов в " +
                        "тяжелые и средние слоты.",
                        false, 80);
                }

                unitSlot.css('animation', 'selectMenu2 1500ms infinite');
                let passed = false;

                $('.inventoryUnit.select').each(function () {
                    if ($(this).css("background-image") !== "none") {
                        passed = true;
                        clearInterval(interval);
                        $('.inventoryUnit.select').css('animation', 'none');
                        unitSlot.css('animation', 'none');
                        if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();
                        if (document.getElementById("training1Squad2Block")) document.getElementById("training1Squad2Block").remove();
                    }
                });

                unitSlot.each(function () {
                    if ($(this).css("background-image") !== "none") {
                        clearInterval(interval);
                        passed = true;
                        unitSlot.css('animation', 'none');
                        if (document.getElementById("training1SquadBlock")) document.getElementById("training1SquadBlock").remove();
                        if (document.getElementById("training1Squad2Block")) document.getElementById("training1Squad2Block").remove();
                    }
                });

                if (passed) {
                    progressTraining(lvl);
                }
            } else {
                IntoToHangar();
            }
        }, 10)
    }

    if (lvl === 8) {
        let page = {
            text: "Очень хорошо, надеюсь я тебе смог помочь с освоение инвентаря.",
            picture: "training.png",
        };
        let dialogBlock = CreatePageDialog("training1Block", page, null, false, true);
        dialogBlock.style.right = "calc(50% - 125px)";
        dialogBlock.style.top = "calc(50% - 300px)";
        dialogBlock.style.left = "auto";
        dialogBlock.className += " Training";

        let ask = document.createElement("div");
        ask.className = "asks";
        ask.innerHTML = "<div class='wrapperAsk'>И это все?!</div>";
        $(ask).click(function () {
            dialogBlock.remove();
            progressTraining(lvl);
        });

        dialogBlock.appendChild(ask);
    }
}

function IntoToHangar() {

    if (document.getElementById("training1IntoHangar")) {
        return;
    }

    let page = {
        text: "Для начала надо научится использовать ангар и инвентарь, что бы открыть меню ангара нажми желтую пиктограмму на интерфейсе, а для инвентаря зеленую.",
        picture: "training.png",
    };

    let dialogBlock = CreatePageDialog("training1IntoHangar", page, null, false, true);
    dialogBlock.style.left = "15px";
    dialogBlock.style.top = "60px";
    dialogBlock.className += " Training";

    let hangarButton = $('#hangarButton');
    hangarButton.css("animation", "selectMenu 1500ms infinite");

    let inventoryButton = $('#inventoryButton');
    inventoryButton.css("animation", "selectMenu2 1500ms infinite");

    let intoHangar = setInterval(function () {
        if (document.getElementById('wrapperInventoryAndStorage') && document.getElementById('inventoryBox')) {
            $("#wrapperInventoryAndStorage").css('left', $('#inventoryBox').position().left + 200);

            dialogBlock.remove();
            document.getElementById('inventoryButton').style.animation = "none";
            document.getElementById('hangarButton').style.animation = "none";

            clearInterval(intoHangar);
        }
    }, 200);

    return hangarButton
}

function createInfoText(toInfoBlock, id, offsetY, offsetX, width, widthText, text, pic, height, bottom) {
    let squadPage = {
        text: text,
        picture: "training.png",
    };
    let SquadsBlock = CreatePageDialog(id, squadPage, null, false, pic);
    SquadsBlock.style.width = width + "px";
    SquadsBlock.style.height = height + "px";
    SquadsBlock.className += " Training";

    let interval = setInterval(function () {
        if (document.getElementById(id)) {

            if (bottom) {
                if (!toInfoBlock.find(SquadsBlock).length) toInfoBlock.append(SquadsBlock);
                SquadsBlock.style.right = -width - 10 + "px";
                SquadsBlock.style.bottom = "0";
                SquadsBlock.style.left = "unset";
                SquadsBlock.style.top = "unset";
            } else {
                SquadsBlock.style.top = Number(toInfoBlock.offset().top + offsetY) + "px";
                SquadsBlock.style.left = Number(toInfoBlock.offset().left + offsetX) + "px";
            }

            $(SquadsBlock).find('.wrapperText').css('width', widthText + "px");
            $(SquadsBlock).find('.wrapperText').css('height', height - 23 + "px");
        } else {
            clearInterval(interval)
        }
    }, 100);


    return $(SquadsBlock).find('.wrapperText')[0]
}

function progressTraining(lvl) {
    clearInterval(interval);

    lvl++;
    chat.send(JSON.stringify({
        event: "training",
        count: lvl,
    }));
    setTimeout(function () {
        Training(lvl);
    }, 500);
}