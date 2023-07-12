package apirepository

import (
	"log"
)

// GetKrakenPriceBTC  is a function which is parsing the response
func GetKrakenPriceBTC(resp map[string]interface{}) string {
	if results, found := resp["result"].(map[string]interface{}); found {
		if xbtusdt, ok := results["XBTUSDT"].(map[string]interface{}); ok {
			if a, ok := xbtusdt["a"].([]interface{}); ok {
				if len(a) > 0 {
					if price, ok := a[0].(string); ok {
						return price
					}
				}
			}
		}
	} else {
		log.Print("Failed to extract price from KrakenPriceBTC")
		return ""
	}
	return ""
}

// GetKrakenPriceETH  is a function which is parsing the response
func GetKrakenPriceETH(resp map[string]interface{}) string {
	if results, found := resp["result"].(map[string]interface{}); found {
		if ethusdt, ok := results["ETHUSDT"].(map[string]interface{}); ok {
			if a, ok := ethusdt["a"].([]interface{}); ok {
				if len(a) > 0 {
					if price, ok := a[0].(string); ok {
						return price
					}
				}
			}
		}
	} else {
		log.Print("Failed to extract price from KrakenPriceETH")
		return ""
	}
	return ""
}

// GetCoinBasePrice  is a function which is parsing the response
func GetCoinBasePrice(resp map[string]interface{}) string {
	data := resp["data"].(map[string]interface{})
	return data["amount"].(string)
}
