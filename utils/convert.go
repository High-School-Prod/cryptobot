package utils

func ParseCryptoName(cryptoId int) string {
	switch cryptoId {
	case 1:
		return "BTC"
	case 74:
		return "DOGE"
	default:
		return ""
	}
}

func ParseCryptoId(cryptoName string) int {
	switch cryptoName {
	case "BTC":
		return 1
	case "DOGE":
		return 74
	default:
		return 0
	}
}

func ParseFrequencyName(frequencyId int) string {
	switch frequencyId {
	case 1:
		return "10m"
	case 2:
		return "30m"
	case 3:
		return "1h"
	case 4:
		return "2h"
	case 5:
		return "6h"
	case 6:
		return "12h"
	case 7:
		return "24h"
	default:
		return ""
	}
}

func ParseFrequencyId(frequencyName string) int {
	switch frequencyName {
	case "10m":
		return 1
	case "30m":
		return 2
	case "1h":
		return 3
	case "2h":
		return 4
	case "6h":
		return 5
	case "12h":
		return 6
	case "24h":
		return 7
	default:
		return 0
	}
}
