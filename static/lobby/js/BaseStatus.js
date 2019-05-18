function UpdateBaseStatus(base) {

    let baseLogoWrapper = document.getElementById('BaseStatus');
    let baseEfficiency = document.getElementById('efficiencyPercent');

    let styleEfficiency;

    if (100 - base.efficiency > 75) {
        styleEfficiency = 'green';
    } else if (100 - base.efficiency <= 75 && 100 - base.efficiency > 50) {
        styleEfficiency = '#edff00';
    } else if (100 - base.efficiency <= 50 && 100 - base.efficiency > 25) {
        styleEfficiency = '#ffca00';
    } else if (100 - base.efficiency <= 25) {
        styleEfficiency = '#ff0000';
    }

    if (!baseEfficiency) {
        baseLogoWrapper.innerHTML = `
            <div id="logoWrapper">
                <div id="BaseLogo" style="background-image: url('../assets/logo/${base.fraction}.png')"></div>
            </div>

            <div id="baseEfficiency">
                <div>
                    Эффективность: 
                    <span id="efficiencyPercent" style="color: ${styleEfficiency}">${100 - base.efficiency}</span>
                    <div id="detailStatusBase"></div>
                </div>
            </div>
        `
    } else {
        baseEfficiency.innerHTML = 100 - base.efficiency;
        baseEfficiency.style.color = styleEfficiency;
    }

    UpdateDetailStatus(base)

    if (document.getElementById('processorRoot')) {
        lobby.send(JSON.stringify({
            event: "PlaceItemToProcessor",
        }));
    }
}

function UpdateDetailStatus(base) {
    let detailStatusBase = document.getElementById('detailStatusBase');
    detailStatusBase.innerHTML = ``;

    for (let i in base.current_resources) {

        let percentFull = base.current_resources[i].quantity * 100 / base.boundary_amount_of_resources;
        detailStatusBase.innerHTML += `
            <div class="baseDetailStatusRow" id="baseDetailStatusRow${base.current_resources[i].item.name}">
                <div class="baseDetailStatusIcon" style="background-image: ${getBackgroundUrlByItem(base.current_resources[i])}"></div>
                <div class="baseDetailStatusWrapperCount">
                    <span class="currentCount"> ${base.current_resources[i].quantity}</span>
                    <span class="currentTax"> налог: ${base.current_resources[i].tax} %</span>
                    <div class="progressBar" style="width: ${percentFull}%"></div>
                </div>
            </div>
        `
    }
}