package services

import (
	"context"
	"errors"
	productPB "go-grpc/pb/product"
)

type ProductService struct {
	productPB.UnimplementedProductServiceServer
}

type productDB struct {
	Id    uint32
	Name  string
	Price float64
	Stock uint32
}

var productDBs []*productDB

func (p *ProductService) GetProducts(context.Context, *productPB.Empty) (*productPB.Products, error) {
	// products := &productPB.Products{
	// 	Data: []*productPB.Product{
	// 		{
	// 			Id:    1,
	// 			Name:  "Metallica T-Shirts",
	// 			Price: 1200000.00,
	// 			Stock: 13,
	// 		},
	// 		{
	// 			Id:    2,
	// 			Name:  "Vans OldSkool",
	// 			Price: 980000.00,
	// 			Stock: 4,
	// 		},
	// 	},
	// }

	products := &productPB.Products{}
	for _, v := range productDBs {
		product := &productPB.Product{
			Id:    v.Id,
			Name:  v.Name,
			Price: v.Price,
			Stock: v.Stock,
		}
		products.Data = append(products.Data, product)
	}

	return products, nil
}

func (p *ProductService) GetProduct(ctx context.Context, id *productPB.Id) (*productPB.Product, error) {
	if id.Id == 0 {
		return nil, errors.New("id is nil")
	}
	for _, v := range productDBs {
		if v.Id == id.Id {
			return &productPB.Product{
				Id:    v.Id,
				Name:  v.Name,
				Price: v.Price,
				Stock: v.Stock,
			}, nil

		}
	}
	return nil, errors.New("data not found by id")
}

func (p *ProductService) CreateProduct(ctx context.Context, r *productPB.Product) (*productPB.Id, error) {
	errMessage := ""
	if len(r.Name) == 0 {
		errMessage = "Name value is nil"
	}
	if r.Price == 0 {
		if len(errMessage) == 0 {
			errMessage = "Price value is nil"
		} else {
			errMessage += ", Price value is nil"
		}
	}
	if r.Stock == 0 {
		if len(errMessage) == 0 {
			errMessage = "Stock value is nil"
		} else {
			errMessage += ", Stock value is nil"
		}
	}

	if len(errMessage) > 0 {
		return nil, errors.New(errMessage)
	}

	id := 0
	if len(productDBs) == 0 {
		id = 1
	} else {
		id = int(productDBs[len(productDBs)-1].Id) + 1
	}

	newProduct := &productDB{
		Id:    uint32(id),
		Name:  r.Name,
		Price: r.Price,
		Stock: r.Stock,
	}

	productDBs = append(productDBs, newProduct)

	returnID := &productPB.Id{
		Id: uint32(id),
	}
	return returnID, nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, r *productPB.Product) (*productPB.Status, error) {
	return nil, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, id *productPB.Id) (*productPB.Status, error) {
	return nil, nil
}
