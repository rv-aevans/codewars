package main

type PriceExample struct {
	model_number          string
	portfolio_code        string
	index_type            string
	base_price_us         int
	base_price_ca         int
	entry_base_multiplier float64
	entry_min_multiplier  float64
	entry_max_multiplier  float64
	mid_base_multiplier   float64
	mid_min_multiplier    float64
	mid_max_multiplier    float64
	high_base_multiplier  float64
	high_min_multiplier   float64
	high_max_multiplier   float64
	entry_multiplier      float64
	mid_multiplier        float64
	high_multiplier       float64
}

var p = PriceExample{
	model_number:          "ARTIEV20",
	portfolio_code:        "AV",
	index_type:            "entry",
	base_price_us:         100,
	base_price_ca:         120,
	entry_base_multiplier: 1.2,
	entry_min_multiplier:  .8,
	entry_max_multiplier:  1.4,
	mid_base_multiplier:   1.4,
	mid_min_multiplier:    1,
	mid_max_multiplier:    1.5,
	high_base_multiplier:  1.1,
	high_min_multiplier:   1,
	high_max_multiplier:   2,
	entry_multiplier:      1,
	mid_multiplier:        1.2,
	high_multiplier:       3,
}

func main() {
	// basePrice := 1
	// // priceExample, err := svc.Db.GetPriceByZipAndModel()
	// // country, err := svc.Db.GetCountyByZip
	// country := "us"

	// switch country {
	// case "us":
	// 	basePrice = p.base_price_us
	// case "ca":
	// 	basePrice = p.base_price_ca
	// }
}
