package orderHandler

type ShipmentRecord struct {
	OrderId     string
	TotalVolume Volume
	Containers  []Container
}

type Volume struct {
	Unit  string
	Value float32
}

type Container struct {
	ContainerType      string
	ContainingProducts []PackedProduct
}

type PackedProduct struct {
	Id       string
	Quantity int
}
