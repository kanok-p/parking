package repository

type PageOption struct {
	Page     int      `form:"page" validate:"min=0"`
	PerPage  int      `form:"per_page" validate:"min=0"`
	Filters  []string `form:"filters"`
	Sorts    []string `form:"sorts"`
	Language string   `form:"-"`
}

type Park struct {
	ID     string `bson:"_id"`
	Index  int    `bson:"index"`
	Car    string `bson:"car"`
	Colour string `bson:"colour"`
	Status string `bson:"status"`
}
