function CreateEquipInBody(slotData) {
    if (slotData && slotData.equip && (slotData.equip.x_attach > 0 || slotData.equip.y_attach > 0)) {
        let MSIcon = document.getElementById("MSIcon");
        let equipSprite = document.createElement("div");

        equipSprite.className = "equipSprites";
        equipSprite.style.backgroundImage = "url(/assets/units/equip/" + slotData.equip.name + ".png)";
        equipSprite.style.top = (slotData.y_attach - slotData.equip.y_attach) / 2 + "px";
        equipSprite.style.left = (slotData.x_attach - slotData.equip.x_attach) / 2 + "px";

        MSIcon.appendChild(equipSprite);
    }
}