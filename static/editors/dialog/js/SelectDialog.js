let selectDialog = {};

function ViewDialog(dialog) {
    selectDialog = dialog;

    document.getElementById("dialogList").style.display = "none";
    let selectDialogBlock = document.getElementById("selectDialog");
    selectDialogBlock.style.display = "block";

    /*
        access_type: "base"
        fraction: "All"
        id: 1
        mission: ""
        name: "training_1"
        pages: {1: {…}, 2: {…}, 3: {…}, 4: {…}}
        type: ""
    */

    selectDialogBlock.innerHTML = `
        <div id="headSelectDialog">
        
             <input type="button" value="Назад" style="float: left; margin-left: 20px" onclick="GetListDialogs()">

            <label> Название диалога: 
                <input type="text" title="name" value="${dialog.name}" oninput="SetNameDialog(this)">
            </label>
            
            <label> Где доступен: 
                <select id="access_type" onchange="SetAccessType(this)">
                    <option value="base">На базе</option>
                    <option value="world">Везде</option>
                </select>
            </label>
            
            <label> Кому доступен: 
                <select id="fraction" onchange="SetDialogFraction(this)">
                    <option value="All">Всем</option>
                    <option value="Replics">Replics</option>
                    <option value="Explores">Explores</option>
                    <option value="Reverses">Reverses</option>
                </select>
            </label>
            
             <label> Тип: 
                <select id="dialogType" onchange="SetTypeDialog(this)">
                    <option value="">Прост)</option>
                    <option value="greeting">Приветсвие базы</option>
                    <option value="greeting_before_mission_not_complete">Приветсвие базы, если не сдал квест</option>
                    <option value="mission">Задание</option>
                </select>
            </label>
            
            <input type="button" value="Сохранить" style="float: right; margin-right: 20px" onclick="SaveDialog()">
        </div>
        
        <div id="pagesSelectDialog">
        </div>
    `;

    setTimeout(function () {

        $('#access_type').val(dialog.access_type);
        $('#fraction').val(dialog.fraction);
        $('#dialogType').val(dialog.type);

        RenderPages(dialog)
    }, 200);
}

function RenderPages() {
    let pagesSelectDialog = document.getElementById("pagesSelectDialog");

    /*
        asc: (3) [{…}, {…}, {…}]
        id: 1
        name: ""
        number: 1
        text: "Здравствуйте прокси-разум %UserName%
        ↵Вы были отформатированы и ваша нейроматрица нуждается в новом курсе обучения."
        type: ""
    */

    for (let i in selectDialog.pages) {
        if (selectDialog.pages.hasOwnProperty(i)) {

            let page = selectDialog.pages[i];

            pagesSelectDialog.innerHTML += `
                    <div class="page" id="page${page.id}">
                    
                        <label for="onePictureCheckBox" style="float: left"> 
                            <span style="float: left">
                                Одна картинка для всех: 
                                <input type="checkbox" title="onePicture" id="onePictureCheckBox${page.id}" onclick="ChangePicOption(${page.id})">
                            </span> 
                            <span style="float: right; line-height: 22px;">
                                Номер страницы: <span style="color: #ecb416">${page.number}</span><br>
                                <input type="button" value="Удалить" onclick="RemovePage(${page.id})">
                            </span>
                        <label>

                        <input type="text" value="${page.name}" style="float: left; margin-top: 1px;" oninput="SetDialogPageHead(this, ${page.id})">
                        <textarea class="pageText" oninput="SetDialogPageText(this, ${page.id})">${page.text}</textarea>
                        
                        <div class="pictures" id="pictures${page.id}">
                            <div class="pic">
                                <label for="picMain">Загрузить фаил</label>
                                <input type="file" id="picMain">
                            </div>
                        </div>
                        
                        <div class="ascs" id="asc${page.id}">
                        
                        </div>
                    </div>
                `;

            GetDialogPicture(page.id, -1).then(function (response) {
                if (response.data.hasOwnProperty('main')) {
                    $('#onePictureCheckBox' + page.id).prop("checked", true);
                }
                ChangePicOption(page.id)
            });

            setTimeout(function () {
                RenderAsc(page);
            }, 200);
        }
    }

    pagesSelectDialog.innerHTML += `
                    <div class="page" style="height: 18px;" onclick="AddPage()"> Добавить страницу </div>
                    `
}

function RenderAsc(page) {
    /*
                 id: 1
                 name: ""
                 text: "Кто я? :D"
                 to_page: 3
                 type_action: ""
    */


    let pageAscBlock = document.getElementById("asc" + page.id);
    for (let j in page.asc) {
        if (page.asc.hasOwnProperty(j)) {

            let asc = page.asc[j];

            pageAscBlock.innerHTML += `
                <div class="asc" onmouseover="lightToPage(${asc.id}, true)" onmouseout="lightToPage(${asc.id}, false)">
                    <input type="text" value="${asc.text}" oninput="SetDialogAscText(this, ${asc.id})">
                    <input type="number" value="${asc.to_page}" oninput="SetDialogAscToPageNumber(this, ${asc.id})">
                    <select id="ascSelectAction${asc.id}" onchange="SetDialogAscTypeAction(this, ${asc.id})">
                        <option value="">Ничего</option>
                        <option value="close">Закрыть диалог</option>
                        <option value="start_training">Начать обучение</option>
                        <option value="miss_training">Пропустить обучение</option>
                        <option value="get_base_greeting"> Взять диалог привествия </option>
                        <option value="accept_mission"> Принять задание </option>
                        <option value="get_reward"> Получить награду </option>
                        <option value="get_base_mission"> Запросить задание на базе </option>
                    </select>
                    <input type="button" value="Удалить" style=" float: right; margin-top: 2px;" onclick="RemoveAsc(${page.id}, ${asc.id})">
                </div>
            `;

            setTimeout(function () {
                $('#ascSelectAction' + asc.id).val(asc.type_action)
            }, 100);
        }
    }

    pageAscBlock.innerHTML += `
    <div class="asc" onclick="AddAsc(${page.id})"> 
            Добавить ответ +
    </div>
    `
}

function lightToPage(ascID, light) {

    let asc = getAscByID(ascID);
    for (let i in selectDialog.pages) {
        if (selectDialog.pages.hasOwnProperty(i) && selectDialog.pages[i].number === Number(asc.to_page)) {

            let page = document.getElementById("page" + selectDialog.pages[i].id);

            if (light) {
                page.style.background = "#ffb700";
            } else {
                page.style.background = "cadetblue";
            }
        }
    }
}