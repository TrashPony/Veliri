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
                        text: `<p>Пре&shy;вос&shy;ход&shy;но, вы спра&shy;ви&shy;лись! Мно&shy;гие ИИ, уже сда&shy;ют&shy;ся
                        кон&shy;крет&shy;но на дан&shy;ном эта&shy;пе. Те&shy;перь да&shy;вай&shy;те бо&shy;лее де&shy;таль&shy;но
                         рас&shy;смот&shy;рим осо&shy;бен&shy;но&shy;сти каж&shy;дой вклад&shy;ки</p>`,
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
                    ask.innerHTML = "<div class='wrapperAsk'>Продолжить</div>";
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

                    createInfoText(
                        $('#SquadsList'), "training1SquadsBlock",
                        -30, -220, 200, 175,

                        `<p> <span style="font-weight: 900; color: #ff9400;">Ангар</span> - На дан&shy;ной вклад&shy;ке,
                        отоб&shy;ра&shy;же&shy;на чис&shy;лен&shy;ность ак&shy;тив&shy;ных от&shy;ря&shy;дов, ко&shy;то&shy;рые 
                        име&shy;ют&shy;ся в твоём рас&shy;по&shy;ря&shy;же&shy;нии. Помни, что управ&shy;лять ты мо&shy;жешь только 
                        ка&shy;ким-то од&shy;ним от&shy;ря&shy;дом. Де&shy;лай свой вы&shy;бор, с умом.</p>`,

                        false,
                        100);

                    createInfoText(
                        $('#SquadsList'), "training1ParamsBlock",
                        100, -220, 200, 175,

                        `<p><span style="font-weight: 900; color: #ff9400;">Панель</span> - Дан&shy;ная па&shy;нель сле&shy;ва,
                         отоб&shy;ра&shy;жа&shy;ет те&shy;ку&shy;щее со&shy;сто&shy;я&shy;ние от&shy;ря&shy;да. Сю&shy;да вхо&shy;дит: 
                         па&shy;ра&shy;мет&shy;ры, свойства, осо&shy;бые мо&shy;ди&shy;фи&shy;ка&shy;то&shy;ры.</p>`,

                        false, 100);

                    createInfoText(
                        $('#Inventory'), "training1InvBlock",
                        0, +155, 200, 175,

                        `<p><span style="font-weight: 900; color: #ff9400;">Трюм</span> - На дан&shy;ной вклад&shy;ке, 
                        отоб&shy;ра&shy;же&shy;ны все ве&shy;щи,что хра&shy;нят&shy;ся и на&shy;хо&shy;дят&shy;ся в твоём рас&shy;
                        по&shy;ря&shy;же&shy;нии. У них име&shy;ет&shy;ся раз&shy;мер и вес, и они не ак&shy;тив&shy;ны, ес&shy;
                        ли не за&shy;действо&shy;ва&shy;ны с ка&shy;ким-то от&shy;ря&shy;дом.</p>`,

                        false, 110, true);
                    createInfoText(
                        $('#storage'), "training1StorageBlock",
                        0, +175, 200, 175,

                        `<p><span style="font-weight: 900; color: #ff9400;">Склад</span> Стра&shy;те&shy;ги&shy;че&shy;ский 
                        склад ба&shy;зы, где хра&shy;нит&shy;ся всё куп&shy;лен&shy;ное, най&shy;ден&shy;ное, со&shy;з&shy;дан&shy;ное
                         и пе&shy;ре&shy;ра&shy;бо&shy;тан&shy;ное. Не&shy;за&shy;ви&shy;си&shy;мо от по&shy;лез&shy;но&shy;сти 
                         или спе&shy;ци&shy;а&shy;ли&shy;за&shy;ции ве&shy;щей.<p>`,

                        false, 100, true);
                    createInfoText(
                        $('#Squad'), "training1SquadBlock",
                        80, 20, 300, 275,

                        `<p><span style="font-weight: 900; color: #ff9400;">Отсеки</span> - На дан&shy;ной вклад&shy;ке,
                         отоб&shy;ра&shy;же&shy;ны твои вой&shy;ска - ДУ ро&shy;бо&shy;ты</p>`,

                        false, 100);
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
                    createInfoText(
                        storage, "training1SquadBlock",
                        0, +175, 175, 150,
                        `<p>
                                От&shy;лич&shy;но, ты узнал пер&shy;вич&shy;ную по&shy;лез&shy;ную ин&shy;фор&shy;ма&shy;цию,
                                 и те&shy;перь го&shy;тов ак&shy;ти&shy;ви&shy;ро&shy;вать свою мо&shy;биль&shy;ную плат&shy;фор&shy;му.
                              </p>

                              <p>
                              Для это&shy;го действия, те&shy;бе при по&shy;мо&shy;щи кур&shy;со&shy;ра по&shy;тре&shy;бу&shy;ет&shy;ся 
                              пе&shy;ре&shy;та&shy;щить «икон&shy;ку» кор&shy;пу&shy;са мо&shy;биль&shy;ной плат&shy;фор&shy;мы из вклад&shy;ки 
                              «склад», в цен&shy;траль&shy;ный эле&shy;мент гра&shy;фи&shy;че&shy;ско&shy;го ин&shy;тер&shy;фей&shy;са – 
                              «ме&shy;сто для кор&shy;пу&shy;са»
                              </p>`,
                        true, 200, true);
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
            let thoriumPanel = $('#ConstructorMS');
            let thoriumSlots = $('.thoriumSlots');

            if (document.getElementById("ConstructorBackGround") && document.getElementById("wrapperInventoryAndStorage")) {
                if (document.getElementById("training1IntoHangar")) document.getElementById("training1IntoHangar").remove();

                if (!document.getElementById("training1SquadBlock") && document.getElementById("storage") && document.getElementById("thorium")) {

                    createInfoText(
                        storage, "training1SquadBlock",
                        0, 0, 175, 150,

                        `<p>За&shy;меть, ко&shy;г&shy;да мо&shy;биль&shy;ная плат&shy;фор&shy;ма ак&shy;ти&shy;ви&shy;
                                ро&shy;ва&shy;на, она не спо&shy;соб&shy;на пе&shy;ре&shy;дви&shy;гать&shy;ся без топ&shy;ли&shy;ва.
                                 По&shy;это&shy;му, да&shy;вай обо&shy;ру&shy;ду&shy;ем на&shy;шу мо&shy;биль&shy;ную плат&shy;фор&shy;му 
                                 ядер&shy;ным топ&shy;ли&shy;вом - То&shy;ри&shy;ем</p>
                             `,

                        true, 205, true);

                    createInfoText(
                        thoriumPanel, "training1ThoriumBlock",
                        30, 15, 300, 275,

                        `<p><span style="font-weight: 900; color: #ff9400;">Слоты для топлива</span>- Об&shy;ра&shy;ти
                                вни&shy;ма&shy;ние на за&shy;го&shy;рев&shy;ши&shy;е&shy;ся ячейки под мо&shy;биль&shy;ной 
                                плат&shy;фор&shy;мой. Это – сло&shy;ты для топ&shy;ли&shy;ва</p>
                              <p>Ты мо&shy;жешь за&shy;действо&shy;вать их все (1 – 3), та&shy;ким об&shy;ра&shy;зом 
                              мак&shy;си&shy;маль&shy;но уве&shy;ли&shy;чив ско&shy;рость хо&shy;да. А мо&shy;жешь ис&shy;
                              поль&shy;зо&shy;вать все&shy;го од&shy;ну из трёх яче&shy;ек. Та&shy;ким об&shy;ра&shy;зом,
                               умень&shy;шив ско&shy;рость хо&shy;да, но сэко&shy;но&shy;мив топ&shy;ли&shy;во.</p>`,

                        false, 160);
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

                    createInfoText(
                        storage, "training1SquadBlock",
                        0, 0, 175, 150,

                        `<p>Дви&shy;га&shy;ем&shy;ся даль&shy;ше. На&shy;ста&shy;ла оче&shy;редь сна&shy;ря&shy;дить 
                        на&shy;шу мо&shy;биль&shy;ную плат&shy;фор&shy;му раз&shy;лич&shy;ным обо&shy;ру&shy;до&shy;ва&shy;ни&shy;ем. 
                        В за&shy;ви&shy;си&shy;мо&shy;сти от сво&shy;ей спе&shy;ци&shy;фи&shy;ки, обо&shy;ру&shy;до&shy;ва&shy;ние, 
                        мо&shy;жет быть: до&shy;бы&shy;ва&shy;ю&shy;щее, обо&shy;ро&shy;ни&shy;тель&shy;ное, или да&shy;же 
                        ра&shy;бо&shy;та&shy;ю&shy;щее в ка&shy;че&shy;стве средств на&shy;па&shy;де&shy;ния<p>`,

                        true, 205, true);

                    createInfoText(
                        powerPanel, "training1PowerBlock",
                        35, -210, 275, 250,

                        `<p><span style="font-weight: 900; color: #ff9400;">Реактор</span> - Учти, те&shy;бе сле&shy;ду&shy;ет
                        сле&shy;дить за на&shy;груз&shy;кой уста&shy;нов&shy;лен&shy;но&shy;го обо&shy;ру&shy;до&shy;
                        ва&shy;ния на ре&shy;ак&shy;тор. В про&shy;тив&shy;ном слу&shy;чае, ес&shy;ли уста&shy;нов&shy;лен&shy;
                        ное обо&shy;ру&shy;до&shy;ва&shy;ние бу&shy;дет пре&shy;вы&shy;шать ли&shy;мит энер&shy;гии ре&shy;ак&shy;
                        то&shy;ра, то оно по&shy;про&shy;сту не ста&shy;нет функ&shy;ци&shy;о&shy;ни&shy;ро&shy;вать.</p>`,

                        false, 105);

                    createInfoText(
                        MSIcon, "training1EquipBlock",
                        -70, -255, 235, 210,

                        `<p><span style="font-weight: 900; color: #ff9400;">Ячейки для обо&shy;ру&shy;до&shy;ва&shy;ния</span> – Ячейки 
                        обо&shy;ру&shy;до&shy;ва&shy;ния де&shy;лят&shy;ся на три тех&shy;но&shy;ло&shy;ги&shy;че&shy;ских ти&shy;па. Это 
                        озна&shy;ча&shy;ет, что ты не су&shy;ме&shy;ешь вло&shy;жить обо&shy;ру&shy;до&shy;ва&shy;ние I тех&shy;но&shy;ло&shy;
                        ги&shy;че&shy;ско&shy;го ти&shy;па, в ячейку II тех&shy;но&shy;ло&shy;ги&shy;че&shy;ско&shy;го ти&shy;па.</p>

                        <p>
                        При&shy;ме&shy;няя со&shy;о&shy;т&shy;вет&shy;ству&shy;ю&shy;щее тех&shy;но&shy;ло&shy;ги&shy;че&shy;ско&shy;му 
                        уров&shy;ню обо&shy;ру&shy;до&shy;ва&shy;ние, на ис&shy;клю&shy;чи&shy;тель&shy;но под&shy;хо&shy;дя&shy;щей ячей&shy;ке.
                        </p>
                        `,

                        false, 220);

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

                    createInfoText(
                        storage, "training1SquadBlock",
                        0, 0, 175, 150,

                        `<p>Преды&shy;ду&shy;щая за&shy;да&shy;ча бы&shy;ла слож&shy;ной, но к сча&shy;стью, мы с ней спра&shy;ви&shy;лись!</p>

                              <p>
                              При&shy;шло вре&shy;мя на&shy;учить&shy;ся уста&shy;нав&shy;ли&shy;вать во&shy;ору&shy;же&shy;ние. 
                              Для это&shy;го, ис&shy;поль&shy;зуй спе&shy;ци&shy;аль&shy;ные сло&shy;ты, что под&shy;све&shy;чи&shy;ва&shy;ют&shy;ся 
                              “крас&shy;ным” цве&shy;том в ин&shy;тер&shy;фей&shy;се.
                              </p>

                              <p>
                              Не за&shy;бы&shy;вая и про бо&shy;е&shy;при&shy;па&shy;сы, без ко&shy;их ни од&shy;но во&shy;ору&shy;же&shy;ние не
                              ста&shy;нет ве&shy;сти огонь. Уста&shy;нав&shy;ли&shy;вая их в ма&shy;лый слот, над сло&shy;том во&shy;ору&shy;же&shy;ния
                              </p>`,

                        true, 215, true);

                    createInfoText(
                        weaponPanel, "training1WeaponBlock",
                        0, -210, 205, 180,

                        `<p><span style="font-weight: 900; color: #ff9400;">За&shy;мет&shy;ка о раз&shy;ме&shy;ре ору&shy;жия</span>
                        – Важ&shy;но пом&shy;нить, что у лю&shy;бо&shy;го
                        во&shy;ору&shy;же&shy;ния име&shy;ют&shy;ся свои раз&shy;ме&shy;ры. А это озна&shy;ча&shy;ет, что ес&shy;ли 
                        кор&shy;пус мо&shy;биль&shy;ной плат&shy;фор&shy;мы не под&shy;хо&shy;дит для то&shy;го или ино&shy;го ти&shy;па 
                        во&shy;ору&shy;же&shy;ния, то “ору&shy;жие” на дан&shy;ный кор&shy;пус, по&shy;ста&shy;вить бу&shy;дет нель&shy;зя.<p>`,

                        false, 185);

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

                    createInfoText(
                        storage, "training1SquadBlock",
                        -80, 175, 195, 170,

                        `<p>На&shy;ко&shy;нец, при&shy;шло вре&shy;мя для са&shy;мо&shy;го ин&shy;те&shy;рес&shy;но&shy;го - сна&shy;ря&shy;же&shy;ния войск.</p>
                              <p>Ар&shy;мии, ко&shy;то&shy;рая де&shy;лит&shy;ся на три груп&shy;пы: лёг&shy;кие вой&shy;ска, сред&shy;ние вой&shy;ска, тя&shy;жё&shy;лые 
                              вой&shy;ска. За&shy;ни&shy;мая ме&shy;сто в ан&shy;га&shy;ре мо&shy;биль&shy;ной плат&shy;фор&shy;мы и раз&shy;ня&shy;ща&shy;я&shy;ся по 
                              ко&shy;ли&shy;че&shy;ству сло&shy;тов, стро&shy;го за&shy;ви&shy;ся от раз&shy;ме&shy;ра ан&shy;га&shy;ра той или иной мо&shy;биль&shy;ной 
                              плат&shy;фор&shy;мы.</p>
                              <p>Не за&shy;будь об&shy;ра&shy;тить вни&shy;ма&shy;ние и на то, что твои «юни&shy;ты» име&shy;ют соб&shy;ствен&shy;ные источ&shy;ни&shy;ки 
                              пи&shy;та&shy;ния – ре&shy;ак&shy;то&shy;ры, с опре&shy;делён&shy;ны&shy;ми огра&shy;ни&shy;че&shy;ни&shy;я&shy;ми.</p>
                              <p>Вдо&shy;ба&shy;вок, учи&shy;ты&shy;вая и вес обо&shy;ру&shy;до&shy;ва&shy;ния, что твои юни&shy;ты не&shy;сут с со&shy;бой.</p>`,

                        true, 260, true);

                    createInfoText(
                        squad, "training1Squad2Block",
                        -205, -160, 160, 135,
                        `<p>Вз&shy;г&shy;ля&shy;ни на пик&shy;то&shy;грам&shy;му под ак&shy;тив&shy;ным сло&shy;том. 
                              Она сви&shy;де&shy;тельству&shy;ет о том, под ка&shy;кой кон&shy;крет&shy;но тип “юни&shy;та” 
                              пред&shy;на&shy;зна&shy;чен тот или иной слот.</p>
                              <p>В не&shy;ко&shy;то&shy;рых слу&shy;ча&shy;ях, не&shy;за&shy;действо&shy;ван&shy;ные сло&shy;ты 
                              мож&shy;но ком&shy;пен&shy;си&shy;ро&shy;вать. По&shy;ме&shy;стив в тя&shy;жё&shy;лые и сред&shy;ние 
                              сло&shy;ты - лёг&shy;кие вой&shy;ска.</p>`,

                        false, 250);
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
            text: `<p>Ну вот и всё. Ты осво&shy;ил азы управ&shy;ле&shy;ния и те&shy;перь го&shy;тов по&shy;ко&shy;рить Veliri-5.</p>
                   <p>Зайди в штаб у главнокомандующего есть для тебя задание.</p>`,
            picture: "training.png",
        };
        let dialogBlock = CreatePageDialog("training1Block", page, null, false, true);
        dialogBlock.style.right = "calc(50% - 125px)";
        dialogBlock.style.top = "calc(50% - 300px)";
        dialogBlock.style.left = "auto";
        dialogBlock.className += " Training";

        let ask = document.createElement("div");
        ask.className = "asks";
        ask.innerHTML = "<div class='wrapperAsk'>Продолжить.</div>";
        $(ask).click(function () {
            dialogBlock.remove();
            progressTraining(lvl);
        });

        dialogBlock.appendChild(ask);
    }

    if (lvl === 9) {
        let page = {
            text: `
                <p>Что бы войти в штаб нажми желтую пиктограмму.</p>
            `,
            picture: "training.png",
        };

        let dialogBlock = CreatePageDialog("training1IntoDepartmentOfEmployment", page, null, false, true);
        dialogBlock.style.left = "15px";
        dialogBlock.style.top = "60px";
        dialogBlock.className += " Training";

        let DepartmentOfEmploymentButton = $('#DepartmentOfEmploymentButton');
        DepartmentOfEmploymentButton.css("animation", "selectMenu 1500ms infinite");

        let afterOpen = false;
        let intoDOE = setInterval(function () {
            if (document.getElementById('DepartmentOfEmployment')) {
                afterOpen = true;
                dialogBlock.remove();
                document.getElementById('DepartmentOfEmploymentButton').style.animation = "none";
            } else {
                if (afterOpen) {
                    afterOpen = false;
                    dialogBlock.remove();
                    clearInterval(intoDOE);
                    Training(lvl)
                }
            }
        }, 200);
    }
}

function IntoToHangar() {

    if (document.getElementById("training1IntoHangar")) {
        return;
    }

    let page = {
        text: `
        <p>Начнём свой путь с незаменимых вещей: ангара и инвентаря.</p>
        <p>Чтобы открыть и ознакомиться с меню “ангара”, нажмите на “жёлтую” кнопку интерфейса с соответствующим названием.</p>
        <p>Если вы желаете ознакомиться с меню “инвентаря”, нажмите на “зелёную” кнопку интерфейса</p>
        `,
        picture: "training.png",
    };

    let dialogBlock = CreatePageDialog("training1IntoHangar", page, null, false, true);
    dialogBlock.style.left = "15px";
    dialogBlock.style.top = "60px";
    dialogBlock.className += " Training";
    dialogBlock.style.height = 183 + "px";
    $(dialogBlock).find('.wrapperText').css('height', 160);

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