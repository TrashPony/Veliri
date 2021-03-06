package unit

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/ammo"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/effect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/getlantern/deepcopy"
	"math/rand"
	"sync"
)

type Unit struct {
	ID      int    `json:"id"`
	SquadID int    `json:"squad_id"`
	Owner   string `json:"owner"`
	OwnerID int    `json:"owner_id"`

	Body *detail.Body `json:"body"`

	GunRotate int  `json:"gun_rotate"`
	FreezeGun bool `json:"-"`

	Rotate int  `json:"rotate"`
	OnMap  bool `json:"on_map"`
	Leave  bool `json:"leave"`
	GameID int  `json:"game_id"`

	Defend bool `json:"defend"`

	HP           int  `json:"hp"`
	Power        int  `json:"power"`
	ActionPoints int  `json:"action_points"`
	Move         bool `json:"move"`
	FindHostile  bool `json:"find_hostile"`

	//-- боевые характиристики живучести/нацигации
	Speed           int  `json:"speed"`
	MinSpeed        int  `json:"min_speed"`
	Initiative      int  `json:"initiative"`
	MaxHP           int  `json:"max_hp"`
	Armor           int  `json:"armor"`
	EvasionCritical int  `json:"evasion_critical"`
	VulToKinetics   int  `json:"vul_to_kinetics"`
	VulToThermo     int  `json:"vul_to_thermo"`
	VulToExplosion  int  `json:"vul_to_explosion"`
	RangeView       int  `json:"range_view"`
	RangeRadar      int  `json:"range_radar"`
	Accuracy        int  `json:"accuracy"`
	MaxPower        int  `json:"max_power"`
	RecoveryPower   int  `json:"recovery_power"`
	RecoveryHP      int  `json:"recovery_HP"`
	WallHack        bool `json:"wall_hack"`

	Effects []*effect.Effect `json:"effects"`
	MS      bool             `json:"ms"`
	Units   map[int]*Slot    `json:"units"` // в роли ключей карты выступают

	Reload *ReloadAction `json:"reload"`

	/* покраска юнитов */
	BodyColor1   string `json:"body_color_1"`
	BodyColor2   string `json:"body_color_2"`
	WeaponColor1 string `json:"weapon_color_1"`
	WeaponColor2 string `json:"weapon_color_2"`

	/* путь к файлу готовой покраске, пока не реализовано */
	BodyTexture   string `json:"body_texture"`
	WeaponTexture string `json:"weapon_texture"`

	/* очередь глобальных коодинат по которому юнит еще будет идти */
	PointsPath []*coordinate.Coordinate

	/* путь по которому идет юнит */
	ActualPath     *[]*PathUnit `json:"actual_path"`
	ActualPathCell *PathUnit    `json:"actual_path_cell"`
	LastPathCell   *PathUnit    `json:"actual_path"`

	CurrentSpeed float64 `json:"current_speed"`
	HighGravity  bool    `json:"high_gravity"`
	Afterburner  bool    `json:"afterburner"`

	X   int     `json:"x"`    /* текущая координата на пиксельной сетке */
	Y   int     `json:"y"`    /* текущая координата на пиксельной сетке */
	ToX float64 `json:"to_x"` /* куда юнит двигается */
	ToY float64 `json:"to_y"` /* куда юнит двигается */
	// у мс мап ид ставится сразу как создается отряд, тоесть ситуация что отряд на глобалке без мап ид крайне мала, а юниты получают мап ид как выйдут из трюма
	MapID int `json:"map_id"`

	Evacuation      bool `json:"evacuation"`
	ForceEvacuation bool `json:"force_evacuation"`
	InSky           bool `json:"in_sky"` /* отряд по той или иной причине летит Оо */

	MoveChecker bool   `json:"move_checker"`
	MoveUUID    string `json:"-"`

	Inventory *inventory.Inventory `json:"inventory"` // в роли ключей карты выступают номера слотов где содержиться итем

	FollowUnitID int  `json:"follow_unit_id"`
	Return       bool `json:"returning"`

	FormationPos *coordinate.Coordinate `json:"formation_pos"`
	Formation    bool                   `json:"formation"`

	target   *Target
	targetMX sync.Mutex
}

type Target struct {
	Type   string `json:"type"` // box, unit, map
	ID     int    `json:"id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Follow bool   `json:"follow"`
}

type Bullet struct {
	UUID      string         `json:"uuid"`
	Weapon    *detail.Weapon `json:"-"`
	Ammo      *ammo.Ammo     `json:"ammo"`
	Rotate    int            `json:"rotate"`
	Artillery bool           `json:"artillery"`
	X         int            `json:"x"`
	Y         int            `json:"y"`
	Z         float64        `json:"z"` // определяет "высоту" пули (сильнее отдалять тени)
	Speed     int            `json:"speed"`
	Target    *Target        `json:"target"`
	OwnerID   int            `json:"owner_id"` // какой игрок стрелял
	MaxRange  int            `json:"max_range"`
	FirePos   int            `json:"-"`
	// состояние юнита на момент выстрела т.к. его состояние может изменится за время полету пути, а урон считать надо
	Unit *Unit `json:"-"`
}

type ShortUnitInfo struct {
	ID int `json:"id"`
	/* позиция */
	GunRotate int `json:"gun_rotate"`
	Rotate    int `json:"rotate"`
	X         int `json:"x"`
	Y         int `json:"y"`

	/*видимый фит*/
	Body *detail.Body `json:"body"`

	/* покраска юнитов */
	BodyColor1   string `json:"body_color_1"`
	BodyColor2   string `json:"body_color_2"`
	WeaponColor1 string `json:"weapon_color_1"`
	WeaponColor2 string `json:"weapon_color_2"`

	/* путь к файлу готовой покраске, пока не реализовано */
	BodyTexture   string `json:"body_texture"`
	WeaponTexture string `json:"weapon_texture"`

	/*ид владелдьца*/
	OwnerID         int       `json:"owner_id"`
	Owner           string    `json:"owner"`
	MapID           int       `json:"map_id"`
	Evacuation      bool      `json:"evacuation"`
	ForceEvacuation bool      `json:"force_evacuation"`
	InSky           bool      `json:"in_sky"` /* отряд по той или иной причине летит Оо */
	MoveChecker     bool      `json:"move_checker"`
	ActualPathCell  *PathUnit `json:"actual_path_cell"`

	HP int `json:"hp"` // TODO хп видно только тогда когда юнит в радаре
}

type PathUnit struct {
	X           int `json:"x"`
	Y           int `json:"y"`
	Rotate      int `json:"rotate"`
	RotateGun   int `json:"rotate_gun"`
	Millisecond int `json:"millisecond"`
	Speed       float64
	Traversed   bool `json:"traversed"`
	Animate     bool `json:"animate"` // включить или нет анимацию движения гусениц
}

type ReloadAction struct {
	AmmoID        int
	InventorySlot int
}

type Slot struct {
	Unit       *Unit `json:"unit"`
	NumberSlot int   `json:"number_slot"`
}

// возарщает урон простой и урон по снаряжению
func (unit *Unit) GetDamage() (int, int) {
	damage := rand.Intn(unit.GetWeaponSlot().Ammo.MaxDamage-unit.GetWeaponSlot().Ammo.MinDamage) + unit.GetWeaponSlot().Ammo.MinDamage
	// todo урон по снаряжению
	return damage, 0
}

func (unit *Unit) GetDistWeaponToTarget() int {
	unit.targetMX.Lock()
	defer unit.targetMX.Unlock()

	if unit.target == nil {
		return 9999
	}

	xWeapon, yWeapon := unit.GetWeaponPos()
	return int(game_math.GetBetweenDist(unit.target.X, unit.target.Y, xWeapon, yWeapon))
}

func (unit *Unit) GetWeaponRange() int {
	weaponSlot := unit.GetWeaponSlot()
	return weaponSlot.Weapon.Range
}

func (unit *Unit) GetWeaponMinRange() int {
	weaponSlot := unit.GetWeaponSlot()
	return weaponSlot.Weapon.MinAttackRange
}

func (unit *Unit) GetTarget() *Target {
	unit.targetMX.Lock()
	defer unit.targetMX.Unlock()
	return unit.target
}

func (unit *Unit) SetTarget(target *Target) {
	unit.targetMX.Lock()
	defer unit.targetMX.Unlock()
	unit.target = target
}

func (unit *Unit) SetFollowTarget(follow bool) {
	unit.targetMX.Lock()
	defer unit.targetMX.Unlock()
	unit.target.Follow = follow
}

func (unit *Unit) GetShortInfo() *ShortUnitInfo {
	if unit == nil || unit.Body == nil {
		return nil
	}

	var hostile ShortUnitInfo

	hostile.X = unit.X
	hostile.Y = unit.Y

	hostile.GunRotate = unit.GunRotate
	hostile.Rotate = unit.Rotate
	hostile.MapID = unit.MapID
	hostile.Evacuation = unit.Evacuation
	hostile.ForceEvacuation = unit.ForceEvacuation
	hostile.InSky = unit.InSky
	hostile.MoveChecker = unit.MoveChecker
	hostile.ActualPathCell = unit.ActualPathCell

	hostile.Body, _ = gameTypes.Bodies.GetByID(unit.Body.ID)
	hostile.OwnerID = unit.OwnerID
	hostile.Owner = unit.Owner
	hostile.ID = unit.ID

	hostile.BodyColor1 = unit.BodyColor1
	hostile.BodyColor2 = unit.BodyColor2
	hostile.BodyTexture = unit.BodyTexture

	hostile.WeaponColor1 = unit.WeaponColor1
	hostile.WeaponColor2 = unit.WeaponColor2
	hostile.WeaponTexture = unit.WeaponTexture
	hostile.HP = unit.HP

	if unit.GetWeaponSlot() != nil && unit.GetWeaponSlot().Weapon != nil {
		for _, weaponSlot := range hostile.Body.Weapons {
			if weaponSlot != nil {
				weaponSlot.Weapon, _ = gameTypes.Weapons.GetByID(unit.GetWeaponSlot().Weapon.ID)
			}
		}
	}

	copyEquips := func(realEquips *map[int]*detail.BodyEquipSlot, copyEquips *map[int]*detail.BodyEquipSlot) {
		for key, equipSlot := range *realEquips {

			var fakeSlot detail.BodyEquipSlot
			err := deepcopy.Copy(&fakeSlot, equipSlot)
			if err != nil {
				println(err.Error())
			}

			fakeSlot.HP = 0
			fakeSlot.Used = false
			fakeSlot.StepsForReload = 0
			fakeSlot.Target = nil

			(*copyEquips)[key] = &fakeSlot
		}
	}

	copyEquips(&unit.Body.EquippingI, &hostile.Body.EquippingI)
	copyEquips(&unit.Body.EquippingII, &hostile.Body.EquippingII)
	copyEquips(&unit.Body.EquippingIII, &hostile.Body.EquippingIII)
	copyEquips(&unit.Body.EquippingIV, &hostile.Body.EquippingIV)
	copyEquips(&unit.Body.EquippingV, &hostile.Body.EquippingV)

	return &hostile
}

func (unit *Unit) GetID() int {
	return unit.ID
}

func (unit *Unit) GetBody() *detail.Body {
	return unit.Body
}

func (unit *Unit) DelBody() {
	if unit.Body != nil {
		unit.Body = nil
	}
}

func (unit *Unit) DelEquip() {

}

func (unit *Unit) DelAmmo() {

}

func (unit *Unit) SetBody(body *detail.Body) {
	unit.Body = body
}

func (unit *Unit) SetEquip() {

}

func (unit *Unit) SetAmmo() {

}

func (unit *Unit) GetWatchZone() int {
	return unit.RangeView
}

func (unit *Unit) GetOwnerUser() string {
	return unit.Owner
}

func (unit *Unit) GetOnMap() bool {
	return unit.OnMap
}

func (unit *Unit) SetOnMap(bool bool) {
	unit.OnMap = bool
}

func (unit *Unit) GetWallHack() bool {
	return unit.WallHack
}

func (unit *Unit) GetAmmoCount() int { // по диз доку оружие в юните может быть только одно
	for _, weaponSlot := range unit.Body.Weapons {
		if weaponSlot.Weapon != nil {
			return weaponSlot.AmmoQuantity
		}
	}

	return 0
}

func (unit *Unit) GetWeaponSlot() *detail.BodyWeaponSlot { // по диз доку оружие в юните может быть только одно

	if unit.Body == nil {
		return nil
	}

	for _, weaponSlot := range unit.Body.Weapons {
		return weaponSlot
	}

	return nil
}

func (unit *Unit) GetWeaponPos() (int, int) {
	return unit.relativeToAbsolutePointWeapon(unit.X, unit.Y, unit.GetWeaponSlot().XAttach, unit.GetWeaponSlot().YAttach, unit.Rotate, false)
}

func (unit *Unit) GetWeaponFirePos() []*coordinate.Coordinate {

	realPos := make([]*coordinate.Coordinate, 0)
	weaponScale, _ := unit.getOffsetSpritesSize()

	for _, pos := range unit.GetWeaponSlot().Weapon.FirePositions {

		xWeapon, yWeapon := unit.relativeToAbsolutePointWeapon(
			unit.X, unit.Y,
			unit.GetWeaponSlot().XAttach, unit.GetWeaponSlot().YAttach,
			unit.Rotate, false) // я хз почему тут false но оно работает ¯\_(ツ)_/¯

		// берем координаты оружия без разворота, отнимаем сдвиг что бы получить 0:0 спрайта в игровой сети и прибавляем сдвиг для выстрела
		x := xWeapon - unit.GetWeaponSlot().Weapon.XAttach/int(1/weaponScale) + pos.X/int(1/weaponScale)
		y := yWeapon - unit.GetWeaponSlot().Weapon.YAttach/int(1/weaponScale) + pos.Y/int(1/weaponScale)

		newX, newY := game_math.RotatePoint(float64(x), float64(y), float64(xWeapon), float64(yWeapon), unit.GunRotate)

		realPos = append(realPos, &coordinate.Coordinate{X: int(newX), Y: int(newY)})
	}

	return realPos
}

func (unit *Unit) getOffsetSpritesSize() (float64, float64) {
	weaponScale := 0.0
	spriteOffset := 0.0
	if unit.Body.MotherShip {
		weaponScale = 0.25
		spriteOffset = 50
	} else {
		weaponScale = 0.20
		spriteOffset = 20
	}
	return weaponScale, spriteOffset
}

func (unit *Unit) relativeToAbsolutePointWeapon(x0, y0, x, y, rotate int, noRotate bool) (int, int) {
	// TODO координаты атача для оружия считавются правильно но слишком много условностей которые подбирались на фронте...

	weaponScale, spriteOffset := unit.getOffsetSpritesSize()

	xWeapon := x0 + int(((float64(x))/(1/weaponScale))-spriteOffset)
	yWeapon := y0 + int(((float64(y))/(1/weaponScale))-spriteOffset)

	if noRotate {
		return xWeapon, yWeapon
	} else {
		// поворачиваем точку крепления по отношению к телу
		newX, newY := game_math.RotatePoint(float64(xWeapon), float64(yWeapon), float64(x0), float64(y0), rotate)
		return int(newX), int(newY)
	}
}

func (unit *Unit) SetWeaponSlot(newWeaponSlot *detail.BodyWeaponSlot) {
	for i := range unit.Body.Weapons {
		unit.Body.Weapons[i] = newWeaponSlot
	}
}

func (unit *Unit) GetReactorEfficiency() int {
	// метод отдает эффективность реактора мобильной платформы
	// высчитвается как количество слотов в реакторе под торий / количество занятых слотов
	if unit.Body.MotherShip {
		fullCount := 0
		for _, slot := range unit.Body.ThoriumSlots {
			if slot.Count > 0 {
				fullCount++
			}
		}

		efficiency := (fullCount * 100) / len(unit.Body.ThoriumSlots)
		if unit.Afterburner {
			efficiency *= 2
		}
		return efficiency
	} else {
		// у бота нет слотов поэтому там всегда 100
		return 100
	}
}

func (unit *Unit) WorkOutPower(count float32) {
	if unit.Body.MotherShip {
		for _, slot := range unit.Body.ThoriumSlots {
			slot.WorkedOut += count
			if slot.WorkedOut >= 100 {
				slot.Count--
				slot.WorkedOut = 0
			}
		}
	} else {
		unit.Power -= int(count)
	}
}

func (unit *Unit) WorkOutMovePower() {
	// формула выроботки топлива, если работает только 1 ячейка из 3х то ее эффективность в 66% больше
	thorium := 1 / float32(100/unit.GetReactorEfficiency())

	if !unit.HighGravity && !unit.Afterburner { // если не форсах и не высокая гравитация, то не тратим топливо
		return
	}

	if unit.HighGravity && unit.Afterburner { // если активирован форсаж и высокая гравитация то топливо тратиться х15
		thorium = thorium * 15
	}

	if !unit.HighGravity && unit.Afterburner { // если активирован форсаж и низкая гравитация то топливо тратиться х5
		thorium = thorium * 5
	}

	unit.WorkOutPower(thorium)
}

func (unit *Unit) CheckViewCoordinate(x, y int) (bool, bool) {
	if !unit.OnMap && !unit.Body.MotherShip {
		return false, false
	}

	if unit.RangeView >= int(game_math.GetBetweenDist(unit.X, unit.Y, x, y)) {
		return true, true
	}

	if unit.RangeRadar >= int(game_math.GetBetweenDist(unit.X, unit.Y, x, y)) {
		return false, true
	}

	return false, false
}

func (unit *Unit) CalculateParams() {

	if unit.Body == nil {
		unit.Speed = 0
		unit.Initiative = 0
		unit.MaxHP = 0
		unit.Armor = 0
		unit.EvasionCritical = 0
		unit.VulToKinetics = 0
		unit.VulToThermo = 0
		unit.VulToExplosion = 0
		unit.RangeView = 0
		unit.Accuracy = 0
		unit.MaxPower = 0
		unit.RecoveryPower = 0
		unit.RecoveryHP = 0
		unit.WallHack = false

		return
	}

	// начальные параметры тела
	unit.Speed = unit.Body.Speed
	unit.Initiative = unit.Body.Initiative
	unit.MaxHP = unit.Body.MaxHP
	unit.Armor = unit.Body.Armor
	unit.EvasionCritical = unit.Body.EvasionCritical
	unit.VulToKinetics = unit.Body.VulToKinetics
	unit.VulToThermo = unit.Body.VulToThermo
	unit.VulToExplosion = unit.Body.VulToExplosion
	unit.RangeRadar = unit.Body.RangeRadar
	unit.RangeView = unit.Body.RangeView
	unit.Accuracy = unit.Body.Accuracy
	unit.MaxPower = unit.Body.MaxPower
	unit.RecoveryPower = unit.Body.RecoveryPower
	unit.RecoveryHP = unit.Body.RecoveryHP
	unit.WallHack = unit.Body.WallHack

	// смотрим пасивное обородование
	var checkEffect = func(equipEffect *effect.Effect, parameter *int) {
		if equipEffect.Type == "enhances" {
			if equipEffect.Percentages {
				*parameter += *parameter / 100 * equipEffect.Quantity
			} else {
				*parameter += equipEffect.Quantity
			}
		}

		if equipEffect.Type == "reduced" {
			if equipEffect.Percentages {
				*parameter += *parameter / 100 * equipEffect.Quantity
			} else {
				*parameter += equipEffect.Quantity
			}
		}
	}

	var checkParams = func(equipEffect *effect.Effect) {
		if equipEffect.Parameter == "speed" {
			checkEffect(equipEffect, &unit.Speed)
		}

		if equipEffect.Parameter == "initiative" {
			checkEffect(equipEffect, &unit.Initiative)
		}

		if equipEffect.Parameter == "max_hp" {
			checkEffect(equipEffect, &unit.MaxHP)
		}

		if equipEffect.Parameter == "armor" {
			checkEffect(equipEffect, &unit.Armor)
		}

		if equipEffect.Parameter == "evasion_critical" {
			checkEffect(equipEffect, &unit.EvasionCritical)
		}

		if equipEffect.Parameter == "vulnerability_to_kinetics" {
			checkEffect(equipEffect, &unit.VulToKinetics)
		}

		if equipEffect.Parameter == "vulnerability_to_thermo" {
			checkEffect(equipEffect, &unit.VulToThermo)
		}

		if equipEffect.Parameter == "vulnerability_to_explosion" {
			checkEffect(equipEffect, &unit.VulToExplosion)
		}

		if equipEffect.Parameter == "range_view" {
			checkEffect(equipEffect, &unit.RangeView)
		}

		if equipEffect.Parameter == "accuracy" {
			checkEffect(equipEffect, &unit.Accuracy)
		}

		if equipEffect.Parameter == "max_power" {
			checkEffect(equipEffect, &unit.MaxPower)
		}

		if equipEffect.Parameter == "recovery_power" {
			checkEffect(equipEffect, &unit.RecoveryPower)
		}

		if equipEffect.Parameter == "recovery_hp" {
			checkEffect(equipEffect, &unit.RecoveryHP)
		}
	}

	var checkPassiveEquip = func(equip map[int]*detail.BodyEquipSlot, gameUnit *Unit) {
		for _, slot := range equip {
			if slot.Equip != nil && !slot.Equip.Active {
				for _, equipEffect := range slot.Equip.Effects {
					if equipEffect != nil {
						checkParams(equipEffect)
					}
				}
			}
		}
	}

	checkPassiveEquip(unit.Body.EquippingI, unit)
	checkPassiveEquip(unit.Body.EquippingII, unit)
	checkPassiveEquip(unit.Body.EquippingIII, unit)
	checkPassiveEquip(unit.Body.EquippingIV, unit)
	checkPassiveEquip(unit.Body.EquippingV, unit)

	// смотрим наложеные в игре эфекты
	for _, unitEffect := range unit.Effects {
		checkParams(unitEffect)
	}

	// высчитывает повер рековери
	// TODO unit.RecoveryPower = unit.Body.RecoveryPower - (unit.Body.GetUsePower() / 4)
	// востанавление энергии зависит от используемой энергии, чем больше обородования тем выше штраф
	// todo штраф так же должен зависеть от скила пользвотеля
}
