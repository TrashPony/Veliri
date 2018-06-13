function getRadius(xCenter, yCenter, Radius) { // метод брезенхейма

    var coordinates = [];

    circle(xCenter, yCenter, Radius, false, coordinates); // метод отрисовывает только растовый полукруг что бы получить полную фигуруз надо у и х поменять местами и прогнать еще раз
    circle(yCenter, xCenter, Radius, true, coordinates);
    fillingCircle(xCenter, yCenter, Radius, coordinates);

    return removeDuplicates(coordinates);
}

function circle(xCenter, yCenter, radius, invert, coordinates) {

    var x, y, delta;
    x = 0;
    y = radius;
    delta = 3 - 2 * radius;

    while (x < y) { // инопланетные технологии взятые из С++ для формирования растовых окружностей алгоритмом Брезенхэма
        // https://ru.wikibooks.org/wiki/%D0%A0%D0%B5%D0%B0%D0%BB%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D0%B8_%D0%B0%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC%D0%BE%D0%B2/%D0%90%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC_%D0%91%D1%80%D0%B5%D0%B7%D0%B5%D0%BD%D1%85%D1%8D%D0%BC%D0%B0

        putCoordinates(x, y, xCenter, yCenter, invert, coordinates);
        putCoordinates(x, y, xCenter, yCenter, invert, coordinates);
        if (delta < 0) {
            delta += 4 * x + 6
        } else {
            delta += 4 * (x - y) + 10;
            y--
        }
        x++
    }

    if (x === y) {
        putCoordinates(x, y, xCenter, yCenter, invert, coordinates)
    }
}

function putCoordinates(x, y, xCenter, yCenter, invert, coordinates) {
    if (!invert) { // метод отрисовывает только растовый полукруг что бы получить полную фигуруз надо у и х поменять местами и прогнать еще раз
        coordinates.push({X: Number(xCenter) + x, Y: Number(yCenter) + y});
        coordinates.push({X: Number(xCenter) + x, Y: Number(yCenter) - y});
        coordinates.push({X: Number(xCenter) - x, Y: Number(yCenter) + y});
        coordinates.push({X: Number(xCenter) - x, Y: Number(yCenter) - y});
    } else {
        coordinates.push({X: Number(yCenter) + y, Y: Number(xCenter) + x});
        coordinates.push({X: Number(yCenter) - y, Y: Number(xCenter) + x});
        coordinates.push({X: Number(yCenter) + y, Y: Number(xCenter) - x});
        coordinates.push({X: Number(yCenter) - y, Y: Number(xCenter) - x});
    }
}

function fillingCircle(xCenter, yCenter, radius, coordinates) {
    var zx = xCenter - radius;
    var zy = yCenter - radius;

    for (var y = zy; y <= (radius * 2 + radius) + yCenter; y++) {

        var xMin = xMaxMin(y, coordinates, true);
        var xMax = xMaxMin(y, coordinates, false);

        for (var x = zx; x <= (radius * 2 + (radius - 1)) + xCenter; x++) {
            if (xMin < x && xMax > x) {
                coordinates.push({X: Number(x), Y: Number(y)});
            }
        }
    }
}

function xMaxMin(y, coordinates, min) {
    var xMax, xMin;

    for (var i = 0; i < coordinates.length; i++) {
        if (i === 0) {
            xMax = coordinates[i].X;
            xMin = coordinates[i].X;
        } else {
            if (coordinates[i].Y === y) {
                if (coordinates[i].X > xMax) {
                    xMax = coordinates[i].X
                }
                if (coordinates[i].X < xMin) {
                    xMin = coordinates[i].X
                }
            }
        }
    }

    if (min) {
        return xMin

    } else {
        return xMax
    }
}

function removeDuplicates(coordinates) {
    var circleCoordinates = {};

    for (var i = 0; i < coordinates.length; i++) {

        if (circleCoordinates.hasOwnProperty(coordinates[i].X)) {
            circleCoordinates[coordinates[i].X][coordinates[i].Y] = coordinates[i];
        } else {
            circleCoordinates[coordinates[i].X] = {};
            circleCoordinates[coordinates[i].X][coordinates[i].Y] = coordinates[i];
        }
    }

    return circleCoordinates
}