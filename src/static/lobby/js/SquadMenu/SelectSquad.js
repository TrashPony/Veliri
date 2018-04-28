function SelectSquad(squad) {

    var idParse = squad.id.split(':'); // "id:squad"

    lobby.send(JSON.stringify({
        event: "SelectSquad",
        squad_id: Number(idParse[0])
    }));

    if (squad.matherShip !== null && squad.matherShip.id !== 0) {
        var sliderContent = document.getElementById("sliderContent");
        if (sliderContent.matherShips) {
            for (var i = 0; i < sliderContent.matherShips.length; i++) { // прокручиваем слайдер до нужного ))
                if (sliderContent.matherShips[i].id === squad.matherShip.id) {

                    var tmpMatherShips = sliderContent.matherShips[0];
                    sliderContent.matherShips[0] = sliderContent.matherShips[i];
                    sliderContent.matherShips[i] = tmpMatherShips;

                    RemoveUnitBox();
                    NextSlide(sliderContent);
                    ConfigurationMatherShip(sliderContent.matherShips[0]);
                }
            }
        }
    } else {
        // todo удаление параметров и картинок
    }
}