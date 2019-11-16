function RadarWork(data) {
    // если обьект создается то метка не нужна
    if (data.action_mark === "createRadarMark" && data.action_object !== "createObj") {
        CreateMark(data.radar_mark, data.x, data.y)
    }

    if (data.action_mark === "removeRadarMark") {
        RemoveMark(data.radar_mark)
    }

    if (data.action_mark === "hideRadarMark") {
        HideMark(data.radar_mark)
    }

    if (data.action_mark === "unhideRadarMark") {
        UnhideMark(data.radar_mark, data.x, data.y)
    }

    if (data.action_object === "createObj") {
        CreateRadarObject(data.radar_mark, data.object)
    }

    if (data.action_object === "updateObj") {
        UpdateObject(data.radar_mark, data.object)
    }

    if (data.action_object === "removeObj") {
        RemoveRadarObject(data.radar_mark)
    }
}