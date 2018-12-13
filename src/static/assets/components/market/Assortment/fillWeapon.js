let weaponTypes;

function fillWeapon(types) {
    weaponTypes = types;
    let filterBlock = document.getElementById("weaponCategoryItem");
    filterBlock.onclick = function (){
        openAmmoOneScroll("Оружие", this)
    }
}