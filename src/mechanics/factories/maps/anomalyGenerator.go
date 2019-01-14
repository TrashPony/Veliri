package maps

import (
	"../../factories/boxes"
	"../../factories/gameTypes"
	"../../gameObjects/boxInMap"
	"../../gameObjects/map"
	"../../gameObjects/resource"
	"math/rand"
)

type Anomaly struct {
	q     int
	r     int
	MapID int `json:"map_id"`
	power int // сила это растояние на котором можно увидить аномалию на сканере со сканером нулевого радиуса, общая дальность сила аномалии + радиус сканера
	Type  int

	/* награда */
	box      *boxInMap.Box
	resource *resource.Map
	text     *AnomalyText
}

type AnomalyText struct {
	Pages []Page // все страницы диалога
}

type Page struct {
	Text string // текст страницы
	Asc  []Ask  // варианты отетов
}

type Ask struct {
	Text   string // текст ответа
	ToPage int    // страница на которую ведет ответ
}

func (a *Anomaly) GetQ() int {
	return a.q
}

func (a *Anomaly) GetR() int {
	return a.r
}

func (a *Anomaly) GetPower() int {
	return a.power
}

func (a *Anomaly) GetLoot() (*boxInMap.Box, *resource.Map, *AnomalyText) {
	return a.box, a.resource, a.text
}

func anomalyGenerator(mp *_map.Map, m *MapStore) {
	i := 0

	for i < 5 {

		typeAnomaly := rand.Intn(4)

		power := rand.Intn(6) // радиус

		q := rand.Intn(mp.QSize)
		r := rand.Intn(mp.RSize)
		coordinatePlace, _ := mp.GetCoordinate(q, r)

		if coordinatePlace.Move {
			if m.anomaly[mp.Id] == nil {
				m.anomaly[mp.Id] = make([]*Anomaly, 0)
			}

			anomaly := &Anomaly{q: q, r: r, Type: typeAnomaly, power: power, MapID: mp.Id}

			// коробка с ресурсами 2+лвл
			if typeAnomaly == 0 {
				anomaly.box = boxes.Boxes.GetAnomalyRandomBox(typeAnomaly, gameTypes.Boxes.GetRandomBox())
			}

			// коробка с чертежом
			if typeAnomaly == 1 {
				anomaly.box = boxes.Boxes.GetAnomalyRandomBox(typeAnomaly, gameTypes.Boxes.GetRandomBox())
			}

			// руда
			if typeAnomaly == 2 {
				anomaly.resource = gameTypes.Resource.GetRandomMapResource()
			}

			// текст
			if typeAnomaly == 3 {
				/// todo невероятный говнокод но ради фана

				pages := make([]Page, 0)

				ask1 := make([]Ask, 0)
				ask1 = append(ask1, Ask{Text: "Прочитать запись", ToPage: 2})
				ask1 = append(ask1, Ask{Text: "О-о-о нет я сваливаю", ToPage: 0})
				pages = append(pages, Page{Text: "Вы находите старый ржавый не на что не похожий информационный пакет, вы попытаись " +
					"подколючится к нему и считать информацию но сходу не удалось расшифровать запись, спустя не" +
					" продолжительное время для человека и целую вечность для машины вы смогли расшифровать информацию " +
					"и удивились тому насколько глубока мысль тех кто оставил этот пакет здесь когда то очень давно. <br><br> " +
					"Информация в пакете гласила:", Asc: ask1})

				ask2 := make([]Ask, 0)
				ask2 = append(ask2, Ask{Text: "Понятно", ToPage: 0}) // если 0 то close
				pages = append(pages, Page{Text: "\"Ты пидор\"", Asc: ask2})

				anomaly.text = &AnomalyText{Pages: pages}
			}

			m.anomaly[mp.Id] = append(m.anomaly[mp.Id], anomaly)
			i++
		}
	}
}

func (m *MapStore) GetAllMapAnomaly(mapID int) []*Anomaly {
	return Maps.anomaly[mapID]
}

func (m *MapStore) GetMapAnomaly(mapID, q, r int) *Anomaly {
	for _, anomaly := range Maps.anomaly[mapID] {
		if anomaly.q == q && anomaly.r == r {
			return anomaly
		}
	}

	return nil
}
