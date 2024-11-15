package fees

func ServiceFeesByCountry(reason, countryISOCode string) int {
	countryFees := map[string]map[string]int{
		"LAND_REGISTRY": {"US": 25, "KE": 2000},
	}

	return countryFees[reason][countryISOCode]
}
