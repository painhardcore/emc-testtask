package main

type car struct {
	_msgpack struct{} `msgpack:",asArray"`

	Serialnumber       uint64             `json:"serialNumber"`
	OwnerName          string             `json:"ownerName"`
	ModelYear          uint64             `json:"modelYear"`
	Code               string             `json:"code"`
	VehicleCode        string             `json:"vehicleCode"`
	Engine             engine             `json:"engine"`
	FuelFigures        fuelFigures        `json:"fuelFigures"`
	PerformanceFigures performanceFigures `json:"performanceFigures"`
	Manufacturer       string             `json:"manufacturer"`
	Model              string             `json:"model"`
	ActivationCode     string             `json:"activationCode"`
}
type engine struct {
	Capacity         uint16 `json:"capacity"`
	NumCylinders     uint8  `json:"numCylinders"`
	MaxRpm           uint16 `json:"maxRpm"`
	ManufacturerCode string `json:"manufacturerCode"`
}

type fuelFigures struct {
	Speed            uint16  `json:"speed"`
	Mpg              float64 `json:"mpg"`
	UsageDescription string  `json:"usageDescription"`
}
type performanceFigures struct {
	OctaneRating     uint16       `json:"octaneRating"`
	Acceleration     acceleration `json:"acceleration"`
}
type acceleration struct {
	Mph     uint16  `json:"mph"`
	Seconds float64 `json:"seconds"`
}
