package orderHandler

type OrderRequest struct {
	Id       string
	Products []OrderedProduct
}

type OrderedProduct struct {
	Id              string
	Name            string
	OrderedQuantity int
	Dimensions      Dimensions
	UnitPrice       float32
}
