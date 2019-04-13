package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
)

type message struct {
	userID  int
	gameID  int
	message interface{}
}

type Message struct {
	Event            string                  `json:"event"`
	IdGame           int                     `json:"id_game"`
	UnitID           int                     `json:"unit_id"`
	EquipID          int                     `json:"equip_id"`
	IdTarget         string                  `json:"id_target"`
	TypeUnit         string                  `json:"type_unit"`
	Q                int                     `json:"q"`
	R                int                     `json:"r"`
	ToQ              int                     `json:"to_q"`
	ToR              int                     `json:"to_r"`
	TargetQ          int                     `json:"target_q"`
	TargetR          int                     `json:"target_r"`
	EquipType        int                     `json:"equip_type"`
	NumberSlot       int                     `json:"number_slot"`
	Seconds          int                     `json:"seconds"`
	AmmoSlots        map[int]*inventory.Slot `json:"ammo_slots"`
	Slot             int                     `json:"slot"`
	DiplomacyUsers   []*localGame.Pact       `json:"diplomacy_user"`
	Accept           bool                    `json:"accept"`
	ToUser           string                  `json:"to_user"`
	Credits          int                     `json:"credits"`
	Slots            map[int]*inventory.Slot `json:"slots"`
	DiplomacyRequest *diplomacyRequest       `json:"diplomacy_request"`
}

type ErrorMessage struct {
	Event string `json:"event"`
	Error string `json:"error"`
}