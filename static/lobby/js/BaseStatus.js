function UpdateBaseStatus(base) {
    console.log(base);

    let baseLogoWrapper = document.getElementById('BaseStatus');
    let baseEfficiency = document.getElementById('efficiencyPercent');

    let logoStyle = 'animation: baseLogo 5s infinite ease-in-out;';
    if (base.fraction === 'Explores') {
        logoStyle = 'animation: baseLogoExplores 5s infinite ease-in-out;';
    } else if (base.fraction === 'Reverses') {
        logoStyle = 'animation: baseLogoReverses 5s infinite ease-in-out;';
    }

    if (!baseEfficiency) {
        baseLogoWrapper.innerHTML = `
            <div id="logoWrapper">
                <div id="BaseLogo" style="${logoStyle} background-image: url('../assets/logo/${base.fraction}.png')"></div>
            </div>

            <div id="baseEfficiency" onclick="CreateDetailMarket()">
                <div>
                    Эффективность: 
                    <span id="efficiencyPercent" style="color: ${GetStyleEfficiency(base.efficiency)}">${100 - base.efficiency}</span>
                    <div id="detailStatusBase"></div>
                </div>
            </div>
        `
    } else {
        baseEfficiency.innerHTML = 100 - base.efficiency;
        baseEfficiency.style.color = GetStyleEfficiency(base.efficiency);
    }

    UpdateDetailStatus(base);

    if (document.getElementById('processorRoot')) {
        lobby.send(JSON.stringify({
            event: "PlaceItemToProcessor",
        }));
    }

    if (document.getElementById('Workbench') && workBenchState === "selectBP") {
        lobby.send(JSON.stringify({
            event: "SelectBP",
            storage_slot: bpSlot,
            count: bpCount
        }));
    }
}

function GetStyleEfficiency(efficiency) {
    if (100 - efficiency > 75) {
        return '#00ff00';
    } else if (100 - efficiency <= 75 && 100 - efficiency > 50) {
        return '#edff00';
    } else if (100 - efficiency <= 50 && 100 - efficiency > 25) {
        return '#ffca00';
    } else if (100 - efficiency <= 25) {
        return '#ff0000';
    }
}

function UpdateDetailStatus(base) {
    let detailStatusBase = document.getElementById('detailStatusBase');
    detailStatusBase.innerHTML = ``;

    for (let i in base.current_resources) {

        let percentFull = base.current_resources[i].quantity * 100 / base.boundary_amount_of_resources;

        detailStatusBase.innerHTML += `
            <div class="baseDetailStatusRow" id="baseDetailStatusRow${base.current_resources[i].item.name}">
                <div class="baseDetailStatusIcon">
                    ${getBackgroundUrlByItem(base.current_resources[i])}
                </div>
                <div class="baseDetailStatusWrapperCount">
                    <span class="currentCount"> ${base.current_resources[i].quantity}</span>
                    <span class="currentTax"> налог: ${base.current_resources[i].tax} %</span>
                    <div class="progressBar">
                        <div style="width: ${percentFull}%"></div>
                    </div>
                </div>
            </div>
        `
    }

    detailStatusBase.innerHTML += `<input type="button" value="Продать сырье" onclick="CreateDetailMarket()">`
}

function CreateDetailMarket() {

    if (document.getElementById("detailMarketWrapper")) {
        document.getElementById("detailMarketWrapper").remove();
        return
    }

    let detailMarketWrapper = document.createElement("div");
    detailMarketWrapper.id = "detailMarketWrapper";
    document.body.appendChild(detailMarketWrapper);

    detailMarketWrapper.innerHTML = `
        <h3 style="float: left">Продажа сырья:</h3>
        <div class="topButton" onmousedown="document.getElementById('detailMarketWrapper').remove();">x</div>
        <div class="topButton" onmousedown="moveWindow(event,'detailMarketWrapper')">⇿</div>
        <div id="DetailBaseStatus"></div>
        <div id="detailMarket"></div>
    `;

    let interval = setInterval(function () {
        if (document.getElementById("detailMarketWrapper")) {
            lobby.send(JSON.stringify({
                event: "GetDetails",
            }));
        } else {
            clearInterval(interval)
        }
    }, 300);
}

let BaseDetailState = {};

function FillDetailMarket(base, inventorySlots) {
    let detailMarket = document.getElementById("detailMarket");

    let sumCountRes = 0;

    for (let i in base.current_resources) {

        sumCountRes += base.current_resources[i].quantity;
        let percentFull = base.current_resources[i].quantity * 100 / base.boundary_amount_of_resources;
        let needRes = base.boundary_amount_of_resources - base.current_resources[i].quantity;
        if (needRes < 0) needRes = 0;

        let row = document.getElementById("baseDetailMarketStatusRow" + base.current_resources[i].item.name);

        if (row) {
            // ленивое обновление данных
            // todo сюда же надо обновлять цены, но пока эта часть не продумана
            $(row).find('#baseDetailStatusNeedRes' + base.current_resources[i].item.name).text(needRes);
            $(row).find('#baseDetailStatusHave' + base.current_resources[i].item.name).text(inventorySlots[i].quantity);
            $(row).find('#baseDetailProgressBar' + base.current_resources[i].item.name).css("width", percentFull + "%");
        } else {
            detailMarket.innerHTML += `
                <div class="baseDetailStatusRow" id="baseDetailMarketStatusRow${base.current_resources[i].item.name}">
                    <div class="baseDetailStatusIcon">
                        ${getBackgroundUrlByItem(base.current_resources[i])}
                    </div>
                    <div class="baseDetailStatusWrapperCount">
                        <span class="currentCount" style="margin-top: -32px">Необходимо: <span id="baseDetailStatusNeedRes${base.current_resources[i].item.name}" style="color: #f57c00;"> ${needRes} <span style="color: white; float: none;"> ед.</span></span></span>
                        <span class="currentCount">На складе: <span id="baseDetailStatusHave${base.current_resources[i].item.name}" style="color: #11cb00;">${inventorySlots[i].quantity} <span style="color: white; float: none;"> ед.</span></span></span>
                        <div class="progressBar"><div id="baseDetailProgressBar${base.current_resources[i].item.name}" style="width: ${percentFull}%"></div></div>
                    </div>
                    <div class="sellDetailBlock">
                        <h6>Продать</h6>
                        <div class="sellDetailButton" style="margin-left: 4px" onclick="SellDetail(${base.current_resources[i].item_id}, 10, 1)">
                            <span class="sellDetailCount">10</span>
                            <span class="sellDetailPrice">10</span>
                        </div>
                        <div class="sellDetailButton" style="margin-left: 10px" onclick="SellDetail(${base.current_resources[i].item_id}, 100, 1)">
                            <span class="sellDetailCount">100</span>
                            <span class="sellDetailPrice">100</span>
                        </div>
                    </div>
                </div>
            `
        }
    }

    let percentFull = sumCountRes * 100 / base.sum_work_resources;


    let detailBaseStatus = document.getElementById("DetailBaseStatus");

    if (document.getElementById("BaseNameStatus")) {
        // ленивое обновление данных
        if (BaseDetailState.fraction !== base.fraction) {
            $(detailBaseStatus).find('#baseOwnedIcon').css("background-image", "url('../assets/logo/" + base.fraction + ".png')");
            $(detailBaseStatus).find('#baseOwnedText').text("Владелец: " + base.fraction);
        }

        $(detailBaseStatus).find('#baseSummaryEfficiency').css("color", GetStyleEfficiency(base.efficiency)).text(100 - base.efficiency);
        $(detailBaseStatus).find('#baseSummaryCountRes').text(sumCountRes + " / " + base.sum_work_resources);
        $(detailBaseStatus).find('#baseSummaryProgressBar2').css("width", percentFull + "%");
        $(detailBaseStatus).find('#baseSummaryTax').text(base.efficiency + "%");
    } else {
        detailBaseStatus.innerHTML = `
            <h4 id="BaseNameStatus">${base.name}</h4>
            <div id="baseOwnedIcon">
                <div id="baseOwnedIcon" style="background-image: url('../assets/logo/${base.fraction}.png')"></div>
                <span id="baseOwnedText">Владелец: ${base.fraction}</span>
            </div>
            <div id="baseSummary">
                <div>Эффективность: <span id="baseSummaryEfficiency" style="color: ${GetStyleEfficiency(base.efficiency)}">${100 - base.efficiency}</span></div>
                <div>Количество ресурсов: <span id="baseSummaryCountRes">${sumCountRes} / ${base.sum_work_resources}</span></div>
                <div id="baseSummaryProgressBar"><div id="baseSummaryProgressBar2" style="width: ${percentFull}%"></div></div>
                <div>Налог на верстак <span id="baseSummaryTax">${base.efficiency}%</span></div>
            </div>
        `;
    }

    BaseDetailState = base;
}

function SellDetail(detailID, count, price) {
    lobby.send(JSON.stringify({
        event: "SellDetail",
        count: count,
        id: detailID,
        price: price,
    }));
}