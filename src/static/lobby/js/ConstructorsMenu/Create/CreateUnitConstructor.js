function CreateUnitConstructor() {
    var lobbyMenu = document.getElementById("lobby");

    var unitConstructor = document.createElement("div");
    unitConstructor.id = "unitConstructor";
    lobbyMenu.appendChild(unitConstructor);

    var constructorTable = document.createElement("table");
    constructorTable.id = "constructorTable";
    var constructorTr = document.createElement("tr");
    var tdUnitParams = document.createElement("td");
    tdUnitParams.className = "ConstructorTD";
    var tdUnitMenu = document.createElement("td");
    tdUnitMenu.className = "ConstructorTD";
    var tdTabDetailMenu = document.createElement("td");
    tdTabDetailMenu.className = "ConstructorTD";

    constructorTable.appendChild(constructorTr);
    constructorTr.appendChild(tdUnitParams);
    constructorTr.appendChild(tdUnitMenu);
    constructorTr.appendChild(tdTabDetailMenu);

    var unitMenu = CreateUnitMenu();
    var unitParams =  CreateUnitParams();
    var tabDetailMenu = CreateTabDetailMenu();

    tdUnitParams.appendChild(unitParams);
    tdUnitMenu.appendChild(unitMenu);
    tdTabDetailMenu.appendChild(tabDetailMenu);

    unitConstructor.appendChild(constructorTable);

    lobby.send(JSON.stringify({
        event: "GetDetailOfUnits"
    }));

    return unitConstructor;
}