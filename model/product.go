package main

type Product struct {
	ProductId       int64          `json:"productId"`
	Description     *string        `json:"description"`
	Name            string         `json:"name"`
	IsSellable      bool           `json:"IsSellable"`
	Sellers         []Seller       `json:"sellers"`
	SellerPriceInfo map[string]int `json:"sellerPriceInfo"`
}

type Seller struct {
	SellerId   int64  `json:"sellerId"`
	SellerName string `json:"sellerName"`
	Price      int    `json:"price"`
}

func (p *Product) ToPriceInfoBySellers() {
	sellerPriceInfo := make(map[string]int)
	for i := range p.Sellers {
		sellerPriceInfo[p.Sellers[i].SellerName] = p.Sellers[i].Price
	}
	p.SellerPriceInfo = sellerPriceInfo
}
