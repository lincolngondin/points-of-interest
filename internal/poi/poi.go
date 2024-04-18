package poi

import "math"

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

func (poi *POI) Distance(ref *point) float64 {
	xS := float64((poi.X - ref.X) * (poi.X - ref.X))
	yS := float64((poi.Y - ref.Y) * (poi.Y - ref.Y))
	return math.Sqrt(xS + yS)
}
