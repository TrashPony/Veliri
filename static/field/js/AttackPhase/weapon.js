function fireWeapon(unit, unitInTargetCoordinate, resultAction, target) {
    if (unit && unitInTargetCoordinate) { // юнит стреляет и юзер видит все
        return Fire(unit, unitInTargetCoordinate)
    } else if (unit && target) {
        return Fire(unit, GetXYCenterHex(resultAction.target.q, resultAction.target.r), "coordinate")
    }

    if (unit && !target) { // видит кто стреляет, не видит куда стреляет
        // end_watch_attack последняя координата видимости полета снаряда
        return Fire(unit, GetXYCenterHex(resultAction.end_watch_attack.q, resultAction.end_watch_attack.r), "inFog")
    }

    if (!unit && target) { // не видит кто стреляет, видит куда
        // start_watch_attack первая координата откуда видно стрельбу
        let weapon = resultAction.weapon_slot.weapon;

        if (unitInTargetCoordinate) {
            return OutFogFire(GetXYCenterHex(resultAction.start_watch_attack.q, resultAction.start_watch_attack.r), unitInTargetCoordinate, weapon)
        } else {
            return OutFogFire(GetXYCenterHex(resultAction.start_watch_attack.q, resultAction.start_watch_attack.r), GetXYCenterHex(resultAction.target.q, resultAction.target.r), weapon, "coordinate")
        }
    }
}