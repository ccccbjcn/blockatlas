package market

type Currency string

const (
	AED Currency = "AED"
	AFN Currency = "AFN"
	ALL Currency = "ALL"
	AMD Currency = "AMD"
	ANG Currency = "ANG"
	AOA Currency = "AOA"
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	AWG Currency = "AWG"
	AZN Currency = "AZN"
	BAM Currency = "BAM"
	BBD Currency = "BBD"
	BDT Currency = "BDT"
	BGN Currency = "BGN"
	BHD Currency = "BHD"
	BIF Currency = "BIF"
	BMD Currency = "BMD"
	BND Currency = "BND"
	BOB Currency = "BOB"
	BRL Currency = "BRL"
	BSD Currency = "BSD"
	BTC Currency = "BTC"
	BTN Currency = "BTN"
	BWP Currency = "BWP"
	BYN Currency = "BYN"
	BYR Currency = "BYR"
	BZD Currency = "BZD"
	CAD Currency = "CAD"
	CDF Currency = "CDF"
	CHF Currency = "CHF"
	CLF Currency = "CLF"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	COP Currency = "COP"
	CRC Currency = "CRC"
	CUC Currency = "CUC"
	CUP Currency = "CUP"
	CVE Currency = "CVE"
	CZK Currency = "CZK"
	DJF Currency = "DJF"
	DKK Currency = "DKK"
	DOP Currency = "DOP"
	DZD Currency = "DZD"
	EGP Currency = "EGP"
	ERN Currency = "ERN"
	ETB Currency = "ETB"
	EUR Currency = "EUR"
	FJD Currency = "FJD"
	FKP Currency = "FKP"
	GBP Currency = "GBP"
	GEL Currency = "GEL"
	GGP Currency = "GGP"
	GHS Currency = "GHS"
	GIP Currency = "GIP"
	GMD Currency = "GMD"
	GNF Currency = "GNF"
	GTQ Currency = "GTQ"
	GYD Currency = "GYD"
	HKD Currency = "HKD"
	HNL Currency = "HNL"
	HRK Currency = "HRK"
	HTG Currency = "HTG"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	IMP Currency = "IMP"
	INR Currency = "INR"
	IQD Currency = "IQD"
	IRR Currency = "IRR"
	ISK Currency = "ISK"
	JEP Currency = "JEP"
	JMD Currency = "JMD"
	JOD Currency = "JOD"
	JPY Currency = "JPY"
	KES Currency = "KES"
	KGS Currency = "KGS"
	KHR Currency = "KHR"
	KMF Currency = "KMF"
	KPW Currency = "KPW"
	KRW Currency = "KRW"
	KWD Currency = "KWD"
	KYD Currency = "KYD"
	KZT Currency = "KZT"
	LAK Currency = "LAK"
	LBP Currency = "LBP"
	LKR Currency = "LKR"
	LRD Currency = "LRD"
	LSL Currency = "LSL"
	LTL Currency = "LTL"
	LVL Currency = "LVL"
	LYD Currency = "LYD"
	MAD Currency = "MAD"
	MDL Currency = "MDL"
	MGA Currency = "MGA"
	MKD Currency = "MKD"
	MMK Currency = "MMK"
	MNT Currency = "MNT"
	MOP Currency = "MOP"
	MRO Currency = "MRO"
	MUR Currency = "MUR"
	MVR Currency = "MVR"
	MWK Currency = "MWK"
	MXN Currency = "MXN"
	MYR Currency = "MYR"
	MZN Currency = "MZN"
	NAD Currency = "NAD"
	NGN Currency = "NGN"
	NIO Currency = "NIO"
	NOK Currency = "NOK"
	NPR Currency = "NPR"
	NZD Currency = "NZD"
	OMR Currency = "OMR"
	PAB Currency = "PAB"
	PEN Currency = "PEN"
	PGK Currency = "PGK"
	PHP Currency = "PHP"
	PKR Currency = "PKR"
	PLN Currency = "PLN"
	PYG Currency = "PYG"
	QAR Currency = "QAR"
	RON Currency = "RON"
	RSD Currency = "RSD"
	RUB Currency = "RUB"
	RWF Currency = "RWF"
	SAR Currency = "SAR"
	SBD Currency = "SBD"
	SCR Currency = "SCR"
	SDG Currency = "SDG"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	SHP Currency = "SHP"
	SLL Currency = "SLL"
	SOS Currency = "SOS"
	SRD Currency = "SRD"
	STD Currency = "STD"
	SVC Currency = "SVC"
	SYP Currency = "SYP"
	SZL Currency = "SZL"
	THB Currency = "THB"
	TJS Currency = "TJS"
	TMT Currency = "TMT"
	TND Currency = "TND"
	TOP Currency = "TOP"
	TRY Currency = "TRY"
	TTD Currency = "TTD"
	TWD Currency = "TWD"
	TZS Currency = "TZS"
	UAH Currency = "UAH"
	UGX Currency = "UGX"
	USD Currency = "USD"
	UYU Currency = "UYU"
	UZS Currency = "UZS"
	VEF Currency = "VEF"
	VND Currency = "VND"
	VUV Currency = "VUV"
	WST Currency = "WST"
	XAF Currency = "XAF"
	XAG Currency = "XAG"
	XAU Currency = "XAU"
	XCD Currency = "XCD"
	XDR Currency = "XDR"
	XOF Currency = "XOF"
	XPF Currency = "XPF"
	YER Currency = "YER"
	ZAR Currency = "ZAR"
	ZMK Currency = "ZMK"
	ZMW Currency = "ZMW"
	ZWL Currency = "ZWL"
)

var Currencies = map[Currency]bool{
	AED: true,
	AFN: true,
	ALL: true,
	AMD: true,
	ANG: true,
	AOA: true,
	ARS: true,
	AUD: true,
	AWG: true,
	AZN: true,
	BAM: true,
	BBD: true,
	BDT: true,
	BGN: true,
	BHD: true,
	BIF: true,
	BMD: true,
	BND: true,
	BOB: true,
	BRL: true,
	BSD: true,
	BTC: true,
	BTN: true,
	BWP: true,
	BYN: true,
	BYR: true,
	BZD: true,
	CAD: true,
	CDF: true,
	CHF: true,
	CLF: true,
	CLP: true,
	CNY: true,
	COP: true,
	CRC: true,
	CUC: true,
	CUP: true,
	CVE: true,
	CZK: true,
	DJF: true,
	DKK: true,
	DOP: true,
	DZD: true,
	EGP: true,
	ERN: true,
	ETB: true,
	EUR: true,
	FJD: true,
	FKP: true,
	GBP: true,
	GEL: true,
	GGP: true,
	GHS: true,
	GIP: true,
	GMD: true,
	GNF: true,
	GTQ: true,
	GYD: true,
	HKD: true,
	HNL: true,
	HRK: true,
	HTG: true,
	HUF: true,
	IDR: true,
	ILS: true,
	IMP: true,
	INR: true,
	IQD: true,
	IRR: true,
	ISK: true,
	JEP: true,
	JMD: true,
	JOD: true,
	JPY: true,
	KES: true,
	KGS: true,
	KHR: true,
	KMF: true,
	KPW: true,
	KRW: true,
	KWD: true,
	KYD: true,
	KZT: true,
	LAK: true,
	LBP: true,
	LKR: true,
	LRD: true,
	LSL: true,
	LTL: true,
	LVL: true,
	LYD: true,
	MAD: true,
	MDL: true,
	MGA: true,
	MKD: true,
	MMK: true,
	MNT: true,
	MOP: true,
	MRO: true,
	MUR: true,
	MVR: true,
	MWK: true,
	MXN: true,
	MYR: true,
	MZN: true,
	NAD: true,
	NGN: true,
	NIO: true,
	NOK: true,
	NPR: true,
	NZD: true,
	OMR: true,
	PAB: true,
	PEN: true,
	PGK: true,
	PHP: true,
	PKR: true,
	PLN: true,
	PYG: true,
	QAR: true,
	RON: true,
	RSD: true,
	RUB: true,
	RWF: true,
	SAR: true,
	SBD: true,
	SCR: true,
	SDG: true,
	SEK: true,
	SGD: true,
	SHP: true,
	SLL: true,
	SOS: true,
	SRD: true,
	STD: true,
	SVC: true,
	SYP: true,
	SZL: true,
	THB: true,
	TJS: true,
	TMT: true,
	TND: true,
	TOP: true,
	TRY: true,
	TTD: true,
	TWD: true,
	TZS: true,
	UAH: true,
	UGX: true,
	USD: true,
	UYU: true,
	UZS: true,
	VEF: true,
	VND: true,
	VUV: true,
	WST: true,
	XAF: true,
	XAG: true,
	XAU: true,
	XCD: true,
	XDR: true,
	XOF: true,
	XPF: true,
	YER: true,
	ZAR: true,
	ZMK: true,
	ZMW: true,
	ZWL: true,
}
