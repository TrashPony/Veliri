let idModales = [
    'chat', 'inventoryBox', 'marketBox', 'wrapperInventoryAndStorage', 'processorRoot',
    'Workbench', 'UsersStatus', 'GlobalViewMap', 'DepartmentOfEmployment', 'Inventory',
];

$(document).ready(function () {
    document.addEventListener('mousedown', function () {
        let elements = document.querySelectorAll(':hover');

        for (let i = 0; i < elements.length; i++) {
            if (idModales.indexOf(elements[i].id) !== -1) {
                modalUpper(elements[i].id);
                break
            }
        }
    });
});

function modalUpper(id) {
    document.getElementById(id).style.zIndex = '10';
    idModales.forEach(function (item) {
        if (item !== id && document.getElementById(item)) {
            document.getElementById(item).style.zIndex = '0';
        }
    });
}