package types

type RawPlace struct {
	ID   int    `name:"#"`
	Name string `name:"Name"`
}

type RawTerritory struct {
	ID   int    `name:"#"`
	Code string `name:"Name"`

	Region int `name:"PlaceName{Region}"`
	Zone   int `name:"PlaceName{Zone}"`
	Place  int `name:"PlaceName"`
	Map    int `name:"Map"`
}

type RawMap struct {
	ID        int `name:"#"`
	X         int `name:"Offset{X}"`
	Y         int `name:"Offset{Y}"`
	Region    int `name:"PlaceName{Region}"`
	Place     int `name:"PlaceName"`
	SubPlace  int `name:"PlaceName{Sub}"`
	Territory int `name:"TerritoryType"`
}
