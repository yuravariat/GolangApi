package entities

//Hotel represents hotel entity
type Hotel struct {
	ID          int        `json:"id"`
	Lang        string     `json:"lang"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Created     *Timestamp `json:"created"`
}

//Hotels represents hotels list
type Hotels []Hotel
