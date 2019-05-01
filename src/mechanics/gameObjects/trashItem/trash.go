package trashItem

type TrashItem struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Size        float32 `json:"size"`
	Description string  `json:"description"`
}
