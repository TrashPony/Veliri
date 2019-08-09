let windowsState = {};

function SetWindowsState(state) {

    if (!state) return;
    windowsState = state;

    let awaitReady = function (id, state) {
        if (document.getElementById(id)) {
            setWindow(document.getElementById(id), state)
        } else {
            setTimeout(() => awaitReady(id, state), 50); //wait 50 ms, then try again
        }
    };

    for (let resolution in state) {

        if (resolution === window.screen.availWidth + ':' + window.screen.availHeight) {

            for (let id in state[resolution]) {

                let window = document.getElementById(id);
                let currentState = state[resolution][id];

                if (window && currentState.open) {
                    setWindow(window, currentState)
                } else if (window && !currentState.open) {
                    window.remove();
                } else if (!window && currentState.open) {

                    // общие сервисы
                    if (id === 'inventoryBox') {
                        InitInventoryMenu(null, 'constructor');
                        awaitReady(id, currentState);
                    }

                    if (id === 'GlobalViewMap') {
                        GlobalViewMap();
                        awaitReady(id, currentState);
                    }

                    if (id === 'UsersStatus') {
                        UsersStatus();
                        awaitReady(id, currentState);
                    }

                    if (id === 'marketBox') {
                        InitMarketMenu(true);
                        awaitReady(id, currentState);
                    }

                    // могут быть открыты только на базе, впрочем отруливается бекендом, но надо избежать ошибки
                    if (id === 'processorRoot') {
                        InitProcessorRoot();
                        awaitReady(id, currentState);
                    }

                    if (id === 'wrapperInventoryAndStorage') {
                        InitInventoryMenu(null, 'storage');
                        awaitReady(id, currentState);
                    }

                    if (id === 'Workbench') {
                        InitWorkbench();
                        awaitReady(id, currentState);
                    }

                    if (id === 'DepartmentOfEmployment') {
                        let departmentID = id;
                        setTimeout(function () { // этот костыль тут из за диалога обучения который идет не через стандартный путь
                            if (!document.getElementById('DepartmentOfEmployment')) {
                                InitDepartmentOfEmployment();
                                awaitReady(departmentID, currentState);
                            }
                        }, 200)
                    }

                    // могут быть открыты только на глобалке, впрочем отруливается бекендом, но надо избежать ошибки
                    if (id === "Inventory") {
                        InitInventoryMenu(null, 'inventory');
                        awaitReady(id, currentState);
                    }
                }
            }
        }
    }
}

function setState(id, left, top, height, weight, open) {

    chat.send(JSON.stringify({
        event: "setWindowState",
        resolution: window.screen.availWidth + ':' + window.screen.availHeight,
        name: id,
        left: Number(left),
        top: Number(top),
        height: Number(height),
        width: Number(weight),
        open: open,
    }));

    let state = {
        left: Number(left),
        top: Number(top),
        height: Number(height),
        width: Number(weight),
        open: open,
    };

    if (!windowsState.hasOwnProperty(window.screen.availWidth + ':' + window.screen.availHeight)) windowsState[window.screen.availWidth + ':' + window.screen.availHeight] = {}
    windowsState[window.screen.availWidth + ':' + window.screen.availHeight][id] = state;

    if (!open) {
        if (document.getElementById(id)) document.getElementById(id).remove();
    }
}

function setWindow(window, state) {
    window.style.left = state.left + "px";
    window.style.top = state.top + "px";
    window.style.height = state.height + "px";
    window.style.width = state.width + "px";

    if ($(window).data("resize")) {
        $(window).data("resize")(null, null, $(window))
    }
}

function openWindow(id, window) {
    let state = checkWindowState(id);
    if (state) {
        state.open = true;
        setWindow(window, state);
        setState(id, state.left, state.top, state.height, state.width, true);
    } else {
        setState(id, $(window).position().left, $(window).position().top, $(window).height(), $(window).width(), true);
    }
}

function checkWindowState(id) {

    if (windowsState && windowsState[window.screen.availWidth + ':' + window.screen.availHeight] &&
        windowsState[window.screen.availWidth + ':' + window.screen.availHeight][id]) {
        return windowsState[window.screen.availWidth + ':' + window.screen.availHeight][id];
    }

    return null
}