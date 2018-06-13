function RemoveEquipCell(equipBox) {
    if (equipBox !== null) {
        equipBox.equip = null;
        equipBox.id = null;
        equipBox.style.backgroundImage = null;
        equipBox.onclick = null;
        equipBox.onmouseover = null;
        equipBox.onmouseout = null;
    }
}