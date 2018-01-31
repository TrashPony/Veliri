function SelectChassis(chassis) {
    var picUnit = document.getElementById("picUnit");
    var chassisUnitBox = document.getElementById("chassisElement");
    chassisUnitBox.chassis = chassis.chassis;
    chassisUnitBox.style.backgroundImage = chassis.style.backgroundImage;

    var picChassis = document.getElementById("picChassis");

    if (!picChassis){
        picChassis = chassisUnitBox.cloneNode(false);
        picChassis.id = "picChassis";
        picChassis.className = "picDetail chassis";
        picChassis.style.backgroundImage = chassis.style.backgroundImage;
    }

    picUnit.appendChild(picChassis);
}

function SelectWeapon(weapon) {
    var picUnit = document.getElementById("picUnit");
    var weaponUnitBox = document.getElementById("weaponElement");
    weaponUnitBox.weapon = weapon.weapon;
    weaponUnitBox.style.backgroundImage = weapon.style.backgroundImage;

    var picWeapon = document.getElementById("picWeapon");

    if (!picWeapon) {
        picWeapon = weaponUnitBox.cloneNode(false);
        picWeapon.id = "picWeapon";
        picWeapon.className = "picDetail weapon";
        picWeapon.style.backgroundImage = weapon.style.backgroundImage;
    }

    picUnit.appendChild(picWeapon);
}

function SelectTower() {
    
}

function SelectBody() {
    
}

function SelectRadar() {
    
}