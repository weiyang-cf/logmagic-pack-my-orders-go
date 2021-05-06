package orderHandler

import (
	"fmt"
	"reflect"
	"testing"
)

// Helpers

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Error("Failure: An unexpected error occurred!")
	}
}

func assertEquals(t *testing.T, expected ShipmentRecord, actual ShipmentRecord) {
	if !reflect.DeepEqual(expected, actual) {
		t.Error(
			fmt.Sprintf(
				"FAILURE: Expected shipment record '%v', received '%v'",
				expected,
				actual),
		)
	}
}

// Fixture

var handler = OrderHandler{
	ContainerSpecs: []ContainerSpec{
		{
			ContainerType: "Cardboard A",
			Dimensions: Dimensions{
				Unit: "centimeter",
				Length: 30,
				Width: 30,
				Height: 30,
			},
		},
		{
			ContainerType: "Cardboard B",
			Dimensions: Dimensions{
				Unit: "centimeter",
				Length: 5,
				Width: 5,
				Height: 40,
			},
		},
	},
}

// Test cases

func TestPackSmallOrder(t *testing.T) {
	t.Log("SCENARIO: Given a small order, pack the ordered products into a single container with sufficient capacity.")

	request := OrderRequest{
		Id: "ORDER-001",
		Products: []OrderedProduct{
			{
				Id: "PRODUCT-001",
				Name: "GOOD FORTUNE COOKIES",
				OrderedQuantity: 9,
				UnitPrice: 13.4,
				Dimensions: Dimensions{
					Unit: "centimeter",
					Length: 10,
					Width: 10,
					Height: 30,
				},
			},
		},
	}

	expectedOutput := ShipmentRecord{
		OrderId: "ORDER-001",
		TotalVolume: Volume{
			Unit: "cubic centimeter",
			Value: 27000,
		},
		Containers: []Container{
			{
				ContainerType: "Cardboard A",
				ContainingProducts: []PackedProduct{
					{
						Id: "PRODUCT-001",
						Quantity: 9,
					},
				},
			},
		},
	}

	actualOutput, err := handler.PackOrder(request)

	assertNil(t, err)
	assertEquals(t, expectedOutput, actualOutput)
}

func TestPackBigOrder(t *testing.T) {
	t.Log("SCENARIO: Given a big order, pack it into a multiple containers.")

	request := OrderRequest{
		Id: "ORDER-002",
		Products: []OrderedProduct{
			{
				Id: "PRODUCT-001",
				Name: "GOOD FORTUNE COOKIES",
				OrderedQuantity: 17,
				UnitPrice: 13.4,
				Dimensions: Dimensions{
					Unit: "centimeter",
					Length: 10,
					Width: 10,
					Height: 30,
				},
			},
		},
	}

	expectedOutput := ShipmentRecord{
		OrderId: "ORDER-002",
		TotalVolume: Volume{
			Unit: "cubic centimeter",
			Value: 54000,
		},
		Containers: []Container{
			{
				ContainerType: "Cardboard A",
				ContainingProducts: []PackedProduct{
					{
						Id: "PRODUCT-001",
						Quantity: 9,
					},
				},
			},
			{
				ContainerType: "Cardboard A",
				ContainingProducts: []PackedProduct{
					{
						Id: "PRODUCT-001",
						Quantity: 8,
					},
				},
			},
		},
	}

	actualOutput, err := handler.PackOrder(request)

	assertNil(t, err)
	assertEquals(t, expectedOutput, actualOutput)
}

func TestUnpackableOrder(t *testing.T) {
	t.Log("SCENARIO: Given an order that cannot fit into any containers, return an error with a meaningful message.")

	request := OrderRequest{
		Id: "ORDER-003",
		Products: []OrderedProduct{
			{
				Id: "PRODUCT-002",
				Name: "BAD FORTUNE COOKIES",
				OrderedQuantity: 1,
				UnitPrice: 27.3,
				Dimensions: Dimensions{
					Unit: "centimeter",
					Length: 20,
					Width: 20,
					Height: 50,
				},
			},
		},
	}

	_, err := handler.PackOrder(request)

	if err == nil {
		t.Error("FAILURE: Expected an error to be returned, received nil!")
	} else {
		t.Log("Error:", err)
	}
}

func TestTallOrder(t *testing.T) {
	t.Log("SCENARIO: Given an order with a single product, select the appropriate container to pack the product.")

	request := OrderRequest{
		Id: "ORDER-004",
		Products: []OrderedProduct{
			{
				Id: "PRODUCT-003",
				Name: "TALL FORTUNE COOKIES",
				OrderedQuantity: 1,
				UnitPrice: 19.8,
				Dimensions: Dimensions{
					Unit: "centimeter",
					Length: 5,
					Width: 5,
					Height: 39.5,
				},
			},
		},
	}

	expectedOutput := ShipmentRecord{
		OrderId: "ORDER-004",
		TotalVolume: Volume{
			Unit: "cubic centimeter",
			Value: 1000,
		},
		Containers: []Container{
			{
				ContainerType: "Cardboard B",
				ContainingProducts: []PackedProduct{
					{
						Id: "PRODUCT-003",
						Quantity: 1,
					},
				},
			},
		},
	}

	actualOutput, err := handler.PackOrder(request)

	assertNil(t, err)
	assertEquals(t, expectedOutput, actualOutput)
}