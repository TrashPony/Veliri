package order

import (
	"time"
)

type Order struct {
	Id        int
	IdUser    int    `json:"-"`
	Price     int    /* цена за еденицу */
	Count     int    /* количество итемов */
	Type      string /* buy/sell */
	MinBuyOut int    /* минимальное количество для покупки */
	TypeItem  string /* body, ammo, weapon, equip */
	IdItem    int    /* ид итема */
	Expires   time.Time
	PlaceName string /* место продажи */
	PlaceID   int    /* ид места продажи */
	PathJump  int    `json:"path_jump"`
	Item      interface{}

	ItemSize float32 /* сколько весит 1 итем нужен что бы класть его в склад */
	ItemHP   int     /* количество хп итема, нужен что бы класть его в склад */
}
