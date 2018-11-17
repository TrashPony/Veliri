package order

import "time"

type Order struct {
	Id        int
	IdUser    int
	Price     int    /* цена за еденицу */
	Count     int    /* количество итемов */
	Type      string /* buy/sell */
	MinBuyOut int    /* минимальное количество для покупки */
	TypeItem  string /* body, ammo, weapon, equip */
	IdItem    int    /* ид итема */
	Expires   time.Time
	PlaceName string /* место продажи */
	PlaceID   int    /* ид места продажи */
	Item      interface{}
}
