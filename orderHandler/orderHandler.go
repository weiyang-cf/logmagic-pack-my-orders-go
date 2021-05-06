package orderHandler

type Packer interface {
	PackOrder(OrderRequest) (ShipmentRecord, error)
}

type OrderHandler struct {
	ContainerSpecs []ContainerSpec
}

func (handler OrderHandler) PackOrder(request OrderRequest) (ShipmentRecord, error) {
	// TODO: OrderHandler should implement Packer interface. Feel free to include other files, methods and functions...
	return ShipmentRecord{}, nil
}