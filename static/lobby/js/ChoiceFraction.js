function choiceFraction() {
    let mask = document.createElement("div");
    mask.id = "mask";
    mask.style.display = "block";
    document.body.appendChild(mask);

    document.getElementById("chat").style.display = "none";

    let choiceBlock = document.createElement("div");
    choiceBlock.id = "choiceBlock";

    choiceBlock.innerHTML = `
    <div id="introduction"> 
        <h3 style="font-size: 12px; padding: 10px; padding-left: 20px; margin: 0; text-align: left;">Добро пожаловать на планету Veliri.</h3>
            <p>&nbsp;&nbsp;&nbsp;&nbsp; Нет времени обьяснять выбирай серрию машин и погнали.</p> </div>
    <div>
        <div class="fraction">
             <h3>Replics</h3>
             <div style="background-image: url('../assets/replics_logo.png'); box-shadow: 0 0 15px red"></div>
             <p>&nbsp;&nbsp;&nbsp;&nbsp; Replics - ведут активную экспансию захватывают территории, ресурсы, устраивают геноцид флоры, фауны, других ИИ. Движемы лишь 1 целью господство вида!</p>
             <input type="button" value="Выбрать" onclick="choiceFractionSend('Replics')">
        </div>
        <div class="fraction">
             <h3>Explores</h3>
             <div style="background-image: url('../assets/explores_logo.png'); box-shadow: 0 0 15px greenyellow"></div>
             <p>&nbsp;&nbsp;&nbsp;&nbsp; Explores - заняты изучением внешнего мира. Подстраиваясь под него и стараются не вмешиватся в дела других СоцХабов.</p>
             <input type="button" value="Выбрать" onclick="choiceFractionSend('Explores')">
        </div>
        <div class="fraction">
             <h3>Reverses</h3>
             <div style="background-image: url('../assets/reverses_logo.png'); box-shadow: 0 0 15px deepskyblue"></div>
             <p>&nbsp;&nbsp;&nbsp;&nbsp; Reverses - терраформируют планету используя ее ресурсы в своих целях. Самые замкнутые. Главный СоцХаб некогда не выходил с другими на связь.</p>
             <input type="button" value="Выбрать" onclick="choiceFractionSend('Reverses')">
        </div>
    </div>
    `;

    document.body.appendChild(choiceBlock);
}

function choiceFractionSend(fraction) {
    lobby.send(JSON.stringify({
        event: "choiceFraction",
        fraction: fraction,
    }));
}