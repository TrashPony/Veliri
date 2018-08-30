
function ParseEffect(effect, equip) {
    // todo жажда рефакторинга и дополнения для других эфектов

    let effectsTR = document.createElement("tr");
    let effectsTD = document.createElement("td");

    effectsTD.colSpan = 2;
    effectsTD.style.fontSize = "8pt";
    effectsTD.style.backgroundColor = "#4c4c4c";
    effectsTD.style.borderRadius = "5px";

    let type = "";
    let quantity = "";
    let time = "";
    let region = "";

    if (equip !== undefined && equip.region > 0) {
        if (equip.region === 1) {
            region = "<br> в радиусе <span class='Value'>" + equip.region + " клетки</span>"
        } else {
            region = "<br> в радиусе <span class='Value'>" + equip.region + " клеток</span>"
        }
    }

    if (equip !== undefined) {
        if (effect.forever) {
            if (equip.steps_time === 1) {
                time = ""
            } else {
                time = "<br> в течение <span class='Value'>" + equip.steps_time + " ходов</span>";
            }
        } else {
            if (equip.steps_time > 4) {
                time = "<br> на <span class='Value'>" + equip.steps_time + " ходов</span>";
            } else {
                time = "<br> на <span class='Value'>" + equip.steps_time + " хода</span>";
            }
        }
    } else {
        if (effect.steps_time === 1) {
            time = "<br> остался <span class='Value'>" + effect.steps_time + " ход</span>";
        }
        if (effect.steps_time > 1 && 5 > effect.steps_time) {
            time = "<br> осталось <span class='Value'>" + effect.steps_time + " хода</span>";
        }
        if (5 <= effect.steps_time) {
            time = "<br> осталось <span class='Value'>" + effect.steps_time + " ходов</span>";
        }
    }

    if (effect.percentages) {
        quantity = effect.quantity + "%";
    } else {
        quantity = effect.quantity;
    }

    if (effect.type === "enhances") {
        type = "+"
    }

    if (effect.type === "takes_away") {
        type = "-"
    }

    if (effect.type === "replenishes") {
        type = "++"
    }

    effectsTD.innerHTML = "<span class='Value'>" + type + quantity + " " + effect.parameter + "</span>" + time + region;

    effectsTR.appendChild(effectsTD);
    return effectsTR
}