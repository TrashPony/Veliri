function ThoriumBar(thoriumSlots) {

    let countSlot = 0;
    let fullCount = 0;

    for (let i in thoriumSlots) {
        countSlot++;

        let slot = thoriumSlots[i];
        let box = document.getElementsByClassName("Thorium")[i - 1];

        if (!box.actionFlag) {
            box.className = "Thorium";
            box.innerText = slot.count + "/" + slot.max_count;
            box.count = slot.count;
            box.maxCount = slot.max_count;
            box.numberSlot = i;

            box.onclick = function () {
                global.send(JSON.stringify({
                    event: "removeThorium",
                    thorium_slot: Number(i)
                }))
            };

            let wrapper = document.createElement("div");
            wrapper.className = "wrapper";
            box.appendChild(wrapper);

            let workOutBar = document.createElement("div");
            workOutBar.className = "WorkOutBar";
            workOutBar.style.width = 100 / (100 / (100 - slot.worked_out)) + "%";
            wrapper.appendChild(workOutBar);

            if (slot.count / slot.max_count >= 0.2) {
                fullCount++;
                box.style.backgroundImage = "url(/assets/resource/enriched_thorium.png)";
                box.style.boxShadow = "inset 0 0 5px rgba(0, 0, 0, 1)";
                box.style.animation = "none";
            } else if (slot.count / slot.max_count < 0.2 && slot.count !== 0) {
                fullCount++;
                box.style.backgroundImage = "url(/assets/resource/enriched_thorium.png)";
                box.style.animation = "alertPulse 2s infinite";
            } else if (slot.count === 0) {
                workOutBar.style.visibility = "hidden";
                box.style.backgroundImage = "none";
                box.style.animation = "noCountPulse 2s infinite";
            }
        }
    }

    let speedEfficiency = document.getElementById("speedBarEfficiency");
    speedEfficiency.id = "speedBarEfficiency";

    let thoriumEfficiency = document.getElementById("thoriumBarEfficiency");
    thoriumEfficiency.id = "thoriumBarEfficiency";

    let efficiencyCalc = 0;
    let thoriumEfficiencyCalc = 0;
    if (fullCount > 0) {
        efficiencyCalc = (fullCount * 100) / countSlot;
        thoriumEfficiencyCalc = (100 - efficiencyCalc);
    } else {
        speedEfficiency.innerHTML = "0";
    }

    if (efficiencyCalc <= 33) {
        speedEfficiency.style.color = "#FF0000";
    } else if (efficiencyCalc <= 66) {
        speedEfficiency.style.color = "#FFF000";
    } else if (efficiencyCalc <= 100) {
        speedEfficiency.style.color = "#00FF00";
    }

    if(aburner) {
        thoriumEfficiencyCalc -= thoriumEfficiencyCalc * 15;
    }

    thoriumEfficiency.innerHTML = (thoriumEfficiencyCalc).toFixed(0) + "%";
}