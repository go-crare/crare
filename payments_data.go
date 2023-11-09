package crare

var SupportedCurrencies = make(map[string]Currency)

func init() {
	if err := defaultJson.Unmarshal(dataCurrencies, &SupportedCurrencies); err != nil {
		panic(err)
	}
	dataCurrencies = []byte{}
}
