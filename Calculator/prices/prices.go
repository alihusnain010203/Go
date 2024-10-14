package prices

type taxInclude struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxInclude(taxRate float64) *taxInclude {
	return &taxInclude{
		TaxRate:           taxRate,
		InputPrices:       []float64{100, 200, 300},	
		TaxIncludedPrices: make(map[string]float64),
	}
}

func (tax *taxInclude) CalculateTaxIncludedPrices() {
	for _, price := range tax.InputPrices {
		tax.TaxIncludedPrices[" "] = price * (1 + tax.TaxRate)
	}
}
