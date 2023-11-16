package main

import (
	response "request-response-generate-with-protobuf/generated_gogo/response/gogo_proto"
)

func (p *Product) ToGogoResponse() response.ProductResponse {
	return response.ProductResponse{
		ProductId:       p.ProductId,
		Description:     p.Description,
		Name:            p.Name,
		IsSellable:      p.IsSellable,
		Sellers:         SellerSliceToGogoResponse(p.Sellers),
		SellerPriceInfo: SellerPriceInfoToResponse(p.SellerPriceInfo),
	}
}

func SellerSliceToGogoResponse(sellers []Seller) []response.Seller {
	sellersToGogoResponse := make([]response.Seller, 0, len(sellers))

	for i := range sellers {
		sellersToGogoResponse = append(sellersToGogoResponse, sellers[i].ToGogoResponse())
	}

	return sellersToGogoResponse
}

func (s *Seller) ToGogoResponse() response.Seller {
	return response.Seller{
		SellerId:   s.SellerId,
		SellerName: s.SellerName,
		Price:      int32(s.Price),
	}
}
