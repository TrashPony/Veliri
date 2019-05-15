function FillParams(unit) {
    console.log(unit)
    if (!unit) {
        document.getElementById("infoParamsAttack").innerHTML = '';
        document.getElementById("infoParamsDefend").innerHTML = '';
        document.getElementById("infoParamsNav").innerHTML = '';

        return
    }

    FillAttackPanel(unit);
    FillDefendPanel(unit);
    FillNavPanel(unit);
}

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
        let noArtillery = 'не';
        if (weapon.artillery) {
            artilleryIcon = '/assets/components/inventory/img/artillery-on.png';
            noArtillery = ''
        }

        attackParams.innerHTML = `

        <div> <!-- дамаг мин-макс -->
            <div class="iconParams"></div> <span class="valueParams"> ${weapon.min_damage}-${weapon.max_damage} </span>
            <div class="paramsTip"> Базовый урон оружия </div>
        </div> 
        
        <div> <!-- дальность атаки мин-макс -->
            <div class="iconParams"></div> <span class="valueParams"> ${weapon.min_attack_range}-${weapon.max_attack_range} </span>
            <div class="paramsTip"> Минимальная и максимальная дальность оружия </div>
        </div> 
        
        <div> <!-- инициатива -->
            <div class="iconParams"></div> <span class="valueParams">${weapon.initiative}</span>
            <div class="paramsTip"> Инициатива оружия (у кого выше тот стреляет первей) </div>
        </div> 
        
        <!-- площадь атаки -->
        <div>
            <div class="iconParams"></div> <span class="valueParams">${weapon.area_covers}</span>
            <div class="paramsTip"> Радиус поражения снаряда </div>
        </div> 
        
        <!-- урон по эквипу -->
        <div>
            <div class="iconParams"></div> <span class="valueParams">${weapon.equip_damage}</span>
            <div class="paramsTip"> Максимальный урон который может быть нанесен снаряжению цели </div>
        </div> 
        
        <!-- шанс сломать эквип -->
        <div>
            <div class="iconParams"></div> <span class="valueParams">${weapon.equip_critical_damage}</span>
            <div class="paramsTip"> Шанс вывести случайный модуль из строя </div>
        </div> 
        
        <!-- иконка типа атаки -->
        <div style="background-image: url(/assets/components/inventory/img/${weapon.type_attack}.png); margin-left: 15px">
            <div class="paramsTip"> Тип атаки оружия: <span class="importantly">${weapon.type_attack}</span> </div>
        </div> 
        
        <!-- иконка артилерии -->
        <div style="background-image: url(${artilleryIcon}); margin-left: 20px">
            <div class="paramsTip"> Орудие <span class="importantly">${noArtillery}</span> может стрелять через препятвия </div>
        </div> 
        
        <!-- пустая клетка -->
        <div></div> 
    `
    } else {
        attackParams.innerHTML = ``;
    }
}

function FillDefendPanel(unit) {
    let defendParams = document.getElementById("infoParamsDefend");

    defendParams.innerHTML = `

        <!-- текущие хп/макс хп -->
        <div style="grid-column-start: 1; grid-column-end: 3;">
            <div class="iconParams"></div> <span class="valueParams"> ${unit.hp} / ${unit.max_hp} </span>
            <div class="paramsTip"> Текущие/макс количество жизней </div>
        </div>
        <!-- востановление хп -->
        <div>
            <div class="iconParams"></div> <span class="valueParams"> ${unit.recovery_HP} </span>   
            <div class="paramsTip"> Количество жизней востанавливаемых каждый ход </div>     
        </div>
        
        <!-- макс энергии -->
        <div style="grid-column-start: 1; grid-column-end: 3;">
            <div class="iconParams"></div> <span class="valueParams"> ${unit.max_power} </span>
            <div class="paramsTip"> Максимальное количество эннергии в акамуляторах </div>     
        </div>
        <!-- востановление энергии -->
        <div>
            <div class="iconParams"></div> <span class="valueParams"> ${unit.recovery_power} </span>
            <div class="paramsTip"> Количество эннергии востанавливаемых каждый ход </div>     
        </div>

        <!-- броня -->
        <div>
            <div class="iconParams"></div> <span class="valueParams"> ${unit.armor} </span>
            <div class="paramsTip"> Броня в асбсолютных очках <span style="color: #9f8450">(пример: урон 5, броня 2, будет нанесено 3 очка урона)</span> </div>     
        </div>
        <!-- уязвимости -->
        <div style="grid-column-start: 2; grid-column-end: 4;">
            <div class="iconParams"></div> 
            <span class="valueParams" style="font-size: 10px; font-weight: 900;"> 
                <span style="color: #ffff00">${unit.vul_to_explosion}%</span>\\
                <span style="color: #ff0008">${unit.vul_to_thermo}%</span>\\
                <span style="color: #b5bdb9">${unit.vul_to_kinetics}%</span> 
            </span>
            
            <div class="paramsTip">
            Уязвимости корпуса по типу урона: <br>
            <span style="color: #ffff00">explosion</span>, <span style="color: #ff0008">thermo</span>, <span style="color: #b5bdb9">kinetic</span> <br>
            <span style="color: #9f8450">(пример: урон термо 10, уязвимость к термо урону 20%, получаемый урон 12хп)</span>
            </div>     
        </div>
    `
}

function FillNavPanel(unit) {
    let navParams = document.getElementById("infoParamsNav");

    let globalSpeed = '';
    if (unit.body.mother_ship)
        globalSpeed = `<div class="iconParams"></div> <span class="valueParams">${unit.speed * 30}</span>`;

    let wallHack = '/assets/components/inventory/img/wall_hack_off.png';
    let noWallHack = 'не';
    if (unit.wall_hack) {
        wallHack = '/assets/components/inventory/img/wall_hack_on.png';
        noWallHack = '';
    }

    navParams.innerHTML = `

        <div>
            <div class="iconParams"></div> <span class="valueParams">${unit.speed}</span>
            <div class="paramsTip"> Скорость юнита в бою. Измеряется в количестве клеток которые он может преодолеть за 1 ход </div>     
        </div> 
        <div>
            ${globalSpeed}
            <div class="paramsTip"> Скорость МазерШипа вне боя. Измеряется в км/ч </div>     
        </div> 
        
        <div>
            <div class="iconParams"></div> <span class="valueParams">${unit.initiative}</span>
            <div class="paramsTip"> Инициатива движения. Опеределяет кто первый будет ходить в бою. </div>     
        </div> 
        
        <div>
            <div class="iconParams"></div> <span class="valueParams">${unit.range_view}</span>
            <div class="paramsTip"> Дальность обзора. Опеределяет как далеко может видеть юнит в бою </div>     
        </div> 
        
        <div style="background-image: url(${wallHack})">
            <div class="paramsTip"> Юнит <span class="importantly">${noWallHack}</span> может смотреть через препятвия </div>     
        </div> 
        <div></div> 
        
        <div></div> 
        <div></div> 
        <div></div> 
    `;
}