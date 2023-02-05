package currency

// Currencies
// ¤
//  $ ¢ € ₠ £ ₨ ₹ ₵ ₡
//₳ ฿ ₣ ₲ ₭ ₥ ₦ ₱ ₽ ₴ ₮ ₩ ¥
//  ₢ ₫ ₯ ₪ ₧ ₰

var Symbols = map[string]string{
	"USD": "$",
	"EUR": "€",
	"GBP": "£",
	"TRY": "₹",
	// ₵ ₡
	//"Cent": "¢",
	//"Eurocent": "₠",
	//₨
	//₳ ฿ ₣ ₲ ₭ ₥ ₦ ₱ ₽ ₴ ₮ ₩ ¥
}
