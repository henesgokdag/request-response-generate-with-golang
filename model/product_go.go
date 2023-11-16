package main

import (
	"request-response-generate-with-protobuf/generated_go"
)

type ProductGoResponse go_proto.ProductResponse

func (p *Product) TogoResponse() ProductGoResponse {
	id := p.ProductId
	name := p.Name
	sellable := p.IsSellable
	return ProductGoResponse{
		ProductId:       &id,
		Description:     p.Description,
		Name:            &name,
		IsSellable:      &sellable,
		Sellers:         SellerSliceToGoResponse(p.Sellers),
		SellerPriceInfo: SellerPriceInfoToResponse(p.SellerPriceInfo),
	}
}

func SellerPriceInfoToResponse(info map[string]int) map[string]int32 {
	sellerPriceInfo := make(map[string]int32)
	for k, v := range info {
		sellerPriceInfo[k] = int32(v)
	}
	return sellerPriceInfo
}

func SellerSliceToGoResponse(sellers []Seller) []*go_proto.Seller {
	sellersTogoResponse := make([]*go_proto.Seller, 0, len(sellers))

	for i := range sellers {
		goResponse := sellers[i].ToGoResponse()
		sellersTogoResponse = append(sellersTogoResponse, &goResponse)
	}
	return sellersTogoResponse
}

func (s *Seller) ToGoResponse() go_proto.Seller {
	id := s.SellerId
	name := s.SellerName
	price := int32(s.Price)
	return go_proto.Seller{
		SellerId:   &id,
		SellerName: &name,
		Price:      &price,
	}
}
