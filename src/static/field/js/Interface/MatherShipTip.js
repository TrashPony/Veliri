function MatherShipTip(matherShip) {
    var matherShipTip = document.getElementById("matherShipTip").style;

    if (matherShip) {
        document.getElementById("matherShipTipType").innerHTML = "<spen class='Value'>" + matherShip.info.type + "</spen>";
        document.getElementById("matherShipArmor").innerHTML = "<spen class='Value'>" + matherShip.info.armor + "</spen>";
        document.getElementById("matherShipOwner").innerHTML = "<spen class='Value'>" + matherShip.info.hp + "</spen>";
        document.getElementById("matherShipArmor").innerHTML = "<spen class='Value'>" + matherShip.info.owner + "</spen>";
        document.getElementById("matherShipRangeView").innerHTML = "<spen class='Value'>" + matherShip.info.range_view + "</spen>";

        matherShipTip.display = "block"; // Показываем слой
    }
}