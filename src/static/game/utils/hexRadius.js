function getRadius(xCenter, yCenter, zCenter, Radius) { // метод брезенхейма
    let coordinates = [];
    for (let x = xCenter - Radius; x <= Number(xCenter) + Number(Radius); x++) {
        for (let y = yCenter - Radius; y <= Number(yCenter) + Number(Radius); y++) {
            for (let z = zCenter - Radius; z <= Number(zCenter) + Number(Radius); z++) {
                if (x + y + z === 0) {
                    coordinates.push({Q: Number(x) + (z-(z&1))/2, R: z});
                }
            }
        }
    }
    return coordinates
}

