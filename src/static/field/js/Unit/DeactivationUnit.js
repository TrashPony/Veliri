function DeactivationUnit(unit) {
    //todo сделать так что бы юзеру было понятно что этот юнит уже ходил
    unit.body.tint = 0x757575;
}

function ActivationUnit(unit) {
    unit.body.tint = 0xFFFFFF;
}