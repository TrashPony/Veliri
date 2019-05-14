function FillAttackPanel(unit) {
    let attackParams = document.getElementById("infoParamsAttack");

    let weapon;

    for (let i in unit.body.weapons) {
        if (unit.body.weapons.hasOwnProperty(i)) {
            weapon = unit.body.weapons[i]
        }
    }

    if (weapon) {
        let artilleryIcon = '/assets/components/inventory/img/artillery-off.png';
        if (weapon.artillery) {
            artilleryIcon = '/assets/components/inventory/img/artillery-on.png';
        }

        attackParams.innerHTML = `

        <div> <!-- дамаг мин-макс -->
            <div class="iconParams"></div> <span class="valueParams"> ${weapon.min_damage}-${weapon.max_damage} </span>
        </div> 
        
        <div> <!-- дальность атаки мин-макс -->
            <div class="iconParams"></div> <span class="valueParams"> ${weapon.min_attack_range}-${weapon.max_attack_range} </span>
        </div> 
        
        <!-- инициатива -->
        <div><div class="iconParams"></div> <span class="valueParams">${weapon.initiative}</span></div> 
        
        <!-- площадь атаки -->
        <div><div class="iconParams"></div> <span class="valueParams">${weapon.area_covers}</span></div> 
        
        <!-- урон по эквипу -->
        <div><div class="iconParams"></div> <span class="valueParams">${weapon.equip_damage}</span></div> 
        
        <!-- шанс сломать эквип -->
        <div><div class="iconParams"></div> <span class="valueParams">${weapon.equip_critical_damage}</span></div> 
        
        <!-- иконка типа атаки -->
        <div  style="background-image: url(/assets/components/inventory/img/${weapon.type_attack}.png); margin-left: 15px"></div> 
        
        <!-- иконка артилерии -->
        <div style="background-image: url(${artilleryIcon}); margin-left: 20px"></div> 
        
        <!-- пустая клетка -->
        <div></div> 
    `
    } else {
        attackParams.innerHTML = ``;
    }
}

function FillDefendPanel(unit) {
    let defendParams = document.getElementById("infoParamsDefend");

    console.log(unit);

    defendParams.innerHTML = `

        <!-- текущие хп/макс хп -->
        <div style="grid-column-start: 1; grid-column-end: 3;">
            <div class="iconParams"></div> <span class="valueParams"> ${unit.hp} / ${unit.max_hp} </span>
        </div>
        <!-- востановление хп -->
        <div>
            <div class="iconParams"></div> <span class="valueParams"> ${unit.recovery_HP} </span>        
        </div>
        
        <!-- макс энергии -->
        <div style="grid-column-start: 1; grid-column-end: 3;">
            <div class="iconParams"></div> <span class="valueParams"> ${unit.max_power} </span>
        </div>
        <!-- востановление энергии -->
        <div>
            <div class="iconParams"></div> <span class="valueParams"> ${unit.recovery_power} </span>
        </div>

        <!-- броня -->
        <div>
            <div class="iconParams"></div> <span class="valueParams"> ${unit.armor} </span>
        </div>
        <!-- уязвимости -->
        <div style="grid-column-start: 2; grid-column-end: 4;">
            <div class="iconParams"></div> <span class="valueParams"> ${unit.vul_to_explosion}\\${unit.vul_to_kinetics}\\${unit.vul_to_thermo} </span>
        </div>
    `
}