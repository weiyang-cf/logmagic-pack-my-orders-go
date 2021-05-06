package orderHandler

type Packer interface {
	PackOrder(OrderRequest) (ShipmentRecord, error)
}

type OrderHandler struct {
	ContainerSpecs []ContainerSpec
}

// TODO: OrderHandler should implement Packer interface
// TODO: feel free to include other methods and functions

func (handler OrderHandler) PackOrder(request OrderRequest) (ShipmentRecord, error) {
	return ShipmentRecord{}, nil
}