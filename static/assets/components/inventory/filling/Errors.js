function alertError(jsonData) {
    let event = JSON.parse(jsonData).event;
    let error = JSON.parse(jsonData).error;

    console.log(event);
    console.log(error);

    if (event === 'SetUnitEquip' || event === 'SetUnitWeapon') {
        if (error === 'lacking size') {
            animateError(document.getElementById("unitCubePanel"))
        }

        if (error === 'lacking power') {
            animateError(document.getElementById("unitPowerPanel"))
        }

        if (error === 'wrong standard size') {
            animateError(document.getElementById("weaponTypePanel"))
        }
    }

    if (event === 'SetMotherShipEquip' || event === 'SetMotherShipWeapon') {
        if (error === 'lacking power') {
            animateError(document.getElementById("powerPanel"))
        }
        if (error === 'wrong standard size') {
            animateError(document.getElementById("MSWeaponPanel"))
        }
    }

    if (event === 'itemToInventory') {
        if (error === 'weight exceeded') {
            animateError(document.getElementById("sizeInventoryInfo"))
        }
    }
}

function animateError(panel) {
    if (!panel) return;

    let oldStyle = getComputedStyle(panel);
    let oldShadow = oldStyle.boxShadow;
    let oldBorder = oldStyle.border;

    let start = Date.now();
    let timer = setInterval(function () {
        let timePassed = Date.now() - start;
        if (timePassed >= 600) {
            clearInterval(timer);
            panel.style.border = oldBorder;
            panel.style.boxShadow = oldShadow;
            return;
        }
        panel.style.boxShadow = "inset 1px 1px 25px 1px rgba(255,0,0,1)";
        panel.style.border = "1px solid #e10006";
    }, 20);
}