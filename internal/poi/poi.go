package poi

type POI struct {
	Name string `json:"nome_do_poi"`
	X    int64  `json:"x_coordenada"`
	Y    int64  `json:"y_coordenada"`
}

func NewPOI(name string, x, y int64) *POI {
	return &POI{
		Name: name,
		X:    x,
		Y:    y,
	}
}

func NewDefaultPOI() *POI {
	return &POI{
		Name: "",
		X:    -1,
		Y:    -1,
	}
}

func (poi *POI) IsValid() bool {
	return poi.Name != "" && poi.X >= 0 && poi.Y >= 0
}
