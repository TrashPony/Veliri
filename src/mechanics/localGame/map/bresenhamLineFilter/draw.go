package bresenhamLineFilter

import (
	"../coordinate"
	"../../"
)

func Draw(xStart, yStart int, endCoordinate *coordinate.Coordinate, game *localGame.Game, filter string) (*coordinate.Coordinate) {
	var x, y, dx, dy, incx, incy, pdx, pdy, es, el, err int

	dx = endCoordinate.X - xStart //проекция на ось x
	dy = endCoordinate.Y - yStart //проекция на ось y

	incx = sign(dx)

	incx = sign(dx)
	// Определяем, в какую сторону нужно будет сдвигаться. Если dx < 0, т.е. отрезок идёт
	// справа налево по иксу, то incx будет равен -1.
	// Это будет использоваться в цикле постороения.
	incy = sign(dy)
	// Аналогично. Если рисуем отрезок снизу вверх -
	// это будет отрицательный сдвиг для y (иначе - положительный).

	if dx < 0 { //далее мы будем сравнивать: "if (dx < dy)"
		dx = -dx //поэтому необходимо сделать dx = |dx|; dy = |dy|
	} //эти две строчки можно записать и так: dx = Math.abs(dx); dy = Math.abs(dy);
	if dy < 0 {
		dy = -dy
	}

	if dx > dy { //определяем наклон отрезка:
		/*
		 * Если dx > dy, то значит отрезок "вытянут" вдоль оси икс, т.е. он скорее длинный, чем высокий.
		 * Значит в цикле нужно будет идти по икс (строчка el = dx;), значит "протягивать" прямую по иксу
		 * надо в соответствии с тем, слева направо и справа налево она идёт (pdx = incx;), при этом
		 * по y сдвиг такой отсутствует.
		 */
		pdx = incx
		pdy = 0
		es = dy
		el = dx
	} else { //случай, когда прямая скорее "высокая", чем длинная, т.е. вытянута по оси y
		pdx = 0
		pdy = incy
		es = dx
		el = dy //тогда в цикле будем двигаться по y
	}

	x = xStart
	y = yStart
	err = el / 2

	gameCoordinate, find := game.GetMap().GetCoordinate(x, y)

	if find && !gameCoordinate.View && filter == "View"{
		return gameCoordinate
	}

	if find && !gameCoordinate.Attack && filter == "Target"{
		return gameCoordinate
	}

	//все последующие точки возможно надо сдвигать, поэтому первую ставим вне цикла
	for i := 0; i < el; i++ { //идём по всем точкам, начиная со второй и до последней

		err -= es
		if err < 0 {
			err += el
			x += incx //сдвинуть прямую (сместить вверх или вниз, если цикл проходит по иксам)
			y += incy //или сместить влево-вправо, если цикл проходит по y
		} else {
			x += pdx //продолжить тянуть прямую дальше, т.е. сдвинуть влево или вправо, если
			y += pdy //цикл идёт по иксу; сдвинуть вверх или вниз, если по y
		}

		gameCoordinate, find := game.GetMap().GetCoordinate(x, y)

		if find && !gameCoordinate.View && filter == "View" {
			return gameCoordinate
		} else {
			if x == endCoordinate.X && y == endCoordinate.Y && filter == "View" {
				return gameCoordinate
			}
		}

		if find && !gameCoordinate.Attack && filter == "Target" {
			return gameCoordinate
		} else {
			if x == endCoordinate.X && y == endCoordinate.Y && filter == "Target" {
				return gameCoordinate
			}
		}
	}

	return nil
}

// Этот код "рисует" все 9 видов отрезков. Наклонные (из начала в конец и из конца в начало каждый), вертикальный и горизонтальный - тоже из начала в конец и из конца в начало, и точку.
func sign(x int) int { //возвращает 0, если аргумент (x) равен нулю; -1, если x < 0 и 1, если x > 0.
	if x == 0 {
		return 0
	} else {
		if x < 0 {
			return -1
		} else {
			return 1
		}
	}
}
