package main

import (
	"fmt"
	"time"

	"github.com/mattevans/dinero"
)

func PairRate(a, b string, sumVal float64) float64 {
	client := dinero.NewClient("8990687eff3648c5b53612f49fc5782a", a, 20*time.Minute) //создаст клиента NewClient(ID клиента, базовая валюта, время обновления )
	rsp, err := client.Rates.Get(b)                                                   // запишет в rsp курc базовой валюты к валюте b
	if err != nil {
		return 0.0
	}
	return sumVal * *rsp //количество базовой валюты переведём по курсу в валюту b
}

func main() {
	a, b := "USD", "RUB" //EUR, USD, RUB
	var sumVal = 100.0
	fmt.Println(PairRate(a, b, sumVal))
}
