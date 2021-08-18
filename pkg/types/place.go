package types

type Place struct {
	ID   int    `name:"#"`
	Name string `name:"Name"`
}

type Territory struct {
	ID   int    `name:"#"`
	Name string `name:"Name"`

	Region int `name:"PlaceName{Region}"`
	Zone   int `name:"PlaceName{Zone}"`
	Place  int `name:"PlaceName"`
	Map    int `name:"Map"`
}

type Map struct {
	ID        int `name:"#"`
	X         int `name:"Offset{X}"`
	Y         int `name:"Offset{Y}"`
	Region    int `name:"PlaceName{Region}"`
	Place     int `name:"PlaceName"`
	SubPlace  int `name:"PlaceName{Sub}"`
	Territory int `name:"TerritoryType"`
}
