package main

import "strconv"

func MapAustralianStateToPostCode(entity Entity, billOrShipping string) interface{} {
	state, _ := entity.GetValue(billOrShipping).(string)
	stateMap := map[string]string{
		"ACT": "0200",
		"NSW": "1000",
		"NT":  "0800",
		"QLD": "4000",
		"SA":  "5000",
		"TAS": "7000",
		"VIC": "3000",
		"WA":  "6000",
	}

	return stateMap[state]
}

func MapAustralianPostCodeToStateZoneID(fieldName string) func(Entity) interface{} {
	// (191, 13, 'Australian Capital Territory', 'ACT', 1),
	// (192, 13, 'New South Wales', 'NSW', 1),
	// (193, 13, 'Northern Territory', 'NT', 1),
	// (194, 13, 'Queensland', 'QLD', 1),
	// (195, 13, 'South Australia', 'SA', 1),
	// (196, 13, 'Tasmania', 'TAS', 1),
	// (197, 13, 'Victoria', 'VIC', 1),
	// (198, 13, 'Western Australia', 'WA', 1),

	return func(entity Entity) interface{} {

		postCodeStr, _ := entity.GetValue(fieldName).(string)
		postCode, err := strconv.Atoi(postCodeStr)
		if err != nil {
			return 0 // or handle error as you deem fit
		}

		switch {
		case (postCode >= 1000 && postCode <= 1999) ||
			(postCode >= 2000 && postCode <= 2599) ||
			(postCode >= 2619 && postCode <= 2898) ||
			(postCode >= 2921 && postCode <= 2999):
			return 192 // NSW
		case (postCode >= 3000 && postCode <= 3999) ||
			(postCode >= 8000 && postCode <= 8999):
			return 197 // VIC
		case (postCode >= 4000 && postCode <= 4999) ||
			(postCode >= 9000 && postCode <= 9999):
			return 194 // QLD
		case (postCode >= 5000 && postCode <= 5799) ||
			(postCode >= 5800 && postCode <= 5999):
			return 195 // SA
		case (postCode >= 6000 && postCode <= 6797) ||
			(postCode >= 6800 && postCode <= 6999):
			return 198 // WA
		case (postCode >= 7000 && postCode <= 7799) ||
			(postCode >= 7800 && postCode <= 7999):
			return 196 // TAS
		case (postCode >= 800 && postCode <= 899) ||
			(postCode >= 900 && postCode <= 999):
			return 193 // NT
		case (postCode >= 2000 && postCode <= 2999) ||
			(postCode >= 2600 && postCode <= 2618) ||
			(postCode >= 2900 && postCode <= 2920):
			return 191 // ACT
		default:
			return 0
		}
	}
}

func MapCountryToCode(fieldName string) func(Entity) interface{} {

	return func(entity Entity) interface{} {
		country, _ := entity.GetValue(fieldName).(string)
		countryMap := map[string]int{
			"AF": 1,
			"AL": 2,
			"DZ": 3,
			"AS": 4,
			"AD": 5,
			"AO": 6,
			"AI": 7,
			"AQ": 8,
			"AG": 9,
			"AR": 10,
			"AM": 11,
			"AW": 12,
			"AU": 13,
			"AT": 14,
			"AZ": 15,
			"BS": 16,
			"BH": 17,
			"BD": 18,
			"BB": 19,
			"BY": 20,
			"BE": 21,
			"BZ": 22,
			"BJ": 23,
			"BM": 24,
			"BT": 25,
			"BO": 26,
			"BA": 27,
			"BW": 28,
			"BV": 29,
			"BR": 30,
			"IO": 31,
			"BN": 32,
			"BG": 33,
			"BF": 34,
			"BI": 35,
			"KH": 36,
			"CM": 37,
			"CA": 38,
			"CV": 39,
			"KY": 40,
			"CF": 41,
			"TD": 42,
			"CL": 43,
			"CN": 44,
			"CX": 45,
			"CC": 46,
			"CO": 47,
			"KM": 48,
			"CG": 49,
			"CK": 50,
			"CR": 51,
			"CI": 52,
			"HR": 53,
			"CU": 54,
			"CY": 55,
			"CZ": 56,
			"DK": 57,
			"DJ": 58,
			"DM": 59,
			"DO": 60,
			"TL": 61,
			"EC": 62,
			"EG": 63,
			"SV": 64,
			"GQ": 65,
			"ER": 66,
			"EE": 67,
			"ET": 68,
			"FK": 69,
			"FO": 70,
			"FJ": 71,
			"FI": 72,
			"FR": 74,
			"GF": 75,
			"PF": 76,
			"TF": 77,
			"GA": 78,
			"GM": 79,
			"GE": 80,
			"DE": 81,
			"GH": 82,
			"GI": 83,
			"GR": 84,
			"GL": 85,
			"GD": 86,
			"GP": 87,
			"GU": 88,
			"GT": 89,
			"GN": 90,
			"GW": 91,
			"GY": 92,
			"HT": 93,
			"HM": 94,
			"HN": 95,
			"HK": 96,
			"HU": 97,
			"IS": 98,
			"IN": 99,
			"ID": 100,
			"IR": 101,
			"IQ": 102,
			"IE": 103,
			"IL": 104,
			"IT": 105,
			"JM": 106,
			"JP": 107,
			"JO": 108,
			"KZ": 109,
			"KE": 110,
			"KI": 111,
			"KP": 112,
			"KR": 113,
			"KW": 114,
			"KG": 115,
			"LA": 116,
			"LV": 117,
			"LB": 118,
			"LS": 119,
			"LR": 120,
			"LY": 121,
			"LI": 122,
			"LT": 123,
			"LU": 124,
			"MO": 125,
			"MK": 126,
			"MG": 127,
			"MW": 128,
			"MY": 129,
			"MV": 130,
			"ML": 131,
			"MT": 132,
			"MH": 133,
			"MQ": 134,
			"MR": 135,
			"MU": 136,
			"YT": 137,
			"MX": 138,
			"FM": 139,
			"MD": 140,
			"MC": 141,
			"MN": 142,
			"MS": 143,
			"MA": 144,
			"MZ": 145,
			"MM": 146,
			"NA": 147,
			"NR": 148,
			"NP": 149,
			"NL": 150,
			"AN": 151,
			"NC": 152,
			"NZ": 153,
			"NI": 154,
			"NE": 155,
			"NG": 156,
			"NU": 157,
			"NF": 158,
			"MP": 159,
			"NO": 160,
			"OM": 161,
			"PK": 162,
			"PW": 163,
			"PA": 164,
			"PG": 165,
			"PY": 166,
			"PE": 167,
			"PH": 168,
			"PN": 169,
			"PL": 170,
			"PT": 171,
			"PR": 172,
			"QA": 173,
			"RE": 174,
			"RO": 175,
			"RU": 176,
			"RW": 177,
			"KN": 178,
			"LC": 179,
			"VC": 180,
			"WS": 181,
			"SM": 182,
			"ST": 183,
			"SA": 184,
			"SN": 185,
			"SC": 186,
			"SL": 187,
			"SG": 188,
			"SK": 189,
			"SI": 190,
			"SB": 191,
			"SO": 192,
			"ZA": 193,
			"GS": 194,
			"ES": 195,
			"LK": 196,
			"SH": 197,
			"PM": 198,
			"SD": 199,
			"SR": 200,
			"SJ": 201,
			"SZ": 202,
			"SE": 203,
			"CH": 204,
			"SY": 205,
			"TW": 206,
			"TJ": 207,
			"TZ": 208,
			"TH": 209,
			"TG": 210,
			"TK": 211,
			"TO": 212,
			"TT": 213,
			"TN": 214,
			"TR": 215,
			"TM": 216,
			"TC": 217,
			"TV": 218,
			"UG": 219,
			"UA": 220,
			"AE": 221,
			"GB": 222,
			"US": 223,
			"UM": 224,
			"UY": 225,
			"UZ": 226,
			"VU": 227,
			"VA": 228,
			"VE": 229,
			"VN": 230,
			"VG": 231,
			"VI": 232,
			"WF": 233,
			"EH": 234,
			"YE": 235,
			"YU": 236,
			"ZM": 237,
			"ZW": 238,
		}

		return countryMap[country]
	}
}
