package ro

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"CZK": ut.Currency{Currency: "CZK", DisplayName: "coroană cehă", Symbol: "CZK"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupie indoneziană", Symbol: "IDR"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldovenesc", Symbol: "MDL"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano bolivian", Symbol: "BOB"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxemburghez", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa maldiviană", Symbol: "MVR"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruvian", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "liră turcească", Symbol: "TRY"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dolar din Trinidad-Tobago", Symbol: "TTD"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "liră din Gibraltar", Symbol: "GIP"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croată", Symbol: "HRK"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colon costarican", Symbol: "CRC"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivian", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "pesetă spaniolă", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram armenesc", Symbol: "AMD"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanez", Symbol: "OMR"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nord-coreean", Symbol: "KPW"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laoțian", Symbol: "LAK"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "cordoba nicaraguană", Symbol: "NIO"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar din Bahrain", Symbol: "BHD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dolar din Belize", Symbol: "BZD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawiană", Symbol: "MWK"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba nicaraguană (1988–1991)", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu românesc", Symbol: "RON"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolivar venezuelean", Symbol: "VEF"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azer (1993–2006)", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodgian", Symbol: "KHR"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marocan", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigeriană", Symbol: "NGN"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar Bosnia-Herțegovina (1992–1994)", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "liră din Insulele Falkland", Symbol: "FKP"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilian", Symbol: "CLP"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "șiling ugandez (1966–1987)", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unitate de cont europeană (XBD)", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra Sao Tome și Principe", Symbol: "STD"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angoleză", Symbol: "AOA"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum din Bhutan", Symbol: "BTN"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar iordanian", Symbol: "JOD"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupie din Sri Lanka", Symbol: "LKR"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayan", Symbol: "UYU"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "gulden din Antilele Olandeze", Symbol: "ANG"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre Ecuador", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint maghiar", Symbol: "HUF"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "metical Mozambic vechi", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zlot polonez (1950–1995)", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentinian (1983–1985)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rublă belarusă", Symbol: "BYR"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "dinar Serbia și Muntenegru (2002–2006)", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa eritreeană", Symbol: "ERN"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo Mozambic", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc rwandez", Symbol: "RWF"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula Botswana", Symbol: "BWP"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brazilian", Symbol: "BRL"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc djiboutian", Symbol: "DJF"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dolar din Insulele Solomon", Symbol: "SBD"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dolar din Caraibele de Est", Symbol: "XCD"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka din Bangladesh", Symbol: "BDT"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc elvețian", Symbol: "CHF"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "liră cipriotă", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dolar din Bermuda", Symbol: "BMD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongol", Symbol: "MNT"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turkmen", Symbol: "TMT"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA BCEAO", Symbol: "CFA"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franc belgian", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar iugoslav convertibil", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "marcă est-germană", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "șechel israelian nou", Symbol: "ILS"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazahă", Symbol: "KZT"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc financiar luxemburghez", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmen (1993–2009)", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc francez de aur", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dolar Zimbabwe (1980–2008)", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso Guineea-Bissau", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "liră sud-sudaneză", Symbol: "SSP"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "liră siriană", Symbol: "SYP"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "carboavă ucraineană", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "monedă necunoscută", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litu lituanian", Symbol: "LTL"}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupie indiană", Symbol: "INR"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dolar namibian", Symbol: "NAD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina din Papua-Noua Guinee", Symbol: "PGK"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar sloven", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro brazilian (1993–1994)", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi Ghana (1979–2007)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentinian", Symbol: "ARS"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothian", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "liră malteză", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azer", Symbol: "AZN"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranian", Symbol: "IRR"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupie pakistaneză", Symbol: "PKR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "liră Insula Sf. Elena", Symbol: "SHP"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoană", Symbol: "WST"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "liră sterlină", Symbol: "GBP"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "liră irlandeză", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "sol nou peruvian", Symbol: "PEN"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dolar rhodesian", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "coroană suedeză", Symbol: "SEK"}, "USS": ut.Currency{Currency: "USS", DisplayName: "dolar american (aceeași zi)", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta spaniolă (cont A)", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hryvna ucraineană", Symbol: "UAH"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinar iugoslav greu", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc comorian", Symbol: "KMF"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgaș", Symbol: "MGA"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone din Sierra Leone", Symbol: "SLL"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "dolar nou din Taiwan", Symbol: "TWD"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dolar american", Symbol: "USD"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unitate compusă europeană", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "drepturi speciale de tragere", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatemalez", Symbol: "GTQ"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca din Macao", Symbol: "MOP"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "coroană slovacă", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum Uzbekistan", Symbol: "UZS"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgani afgan", Symbol: "AFN"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubanez convertibil", Symbol: "CUC"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drahmă grecească", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dolar guyanez", Symbol: "GYD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letonian", Symbol: "LVL"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filipinez", Symbol: "PHP"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolez", Symbol: "CDF"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni din Swaziland", Symbol: "SZL"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC francez", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brazilian (1990–1993)", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "marcă finlandeză", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunisian", Symbol: "TND"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "liră turcească (1922–2005)", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "paladiu", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sud-african", Symbol: "ZAR"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "EUR"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira honduriană", Symbol: "HNL"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "șiling somalez", Symbol: "SOS"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algerian", Symbol: "DZD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupie din Seychelles", Symbol: "SCR"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar irakian", Symbol: "IQD"}, "USN": ut.Currency{Currency: "USN", DisplayName: "dolar american (ziua următoare)", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unitate de monedă europeană", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "coroană norvegiană", Symbol: "NOK"}, "KES": ut.Currency{Currency: "KES", DisplayName: "șiling kenyan", Symbol: "KES"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertibil luxemburghez", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocan", Symbol: "MAD"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dolar australian", Symbol: "AUD"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zlot polonez", Symbol: "PLN"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "liră sudaneză", Symbol: "SDG"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinar iugoslav nou", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "șiling austriac", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta spaniolă (cont convertibil)", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croat", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kârgâz", Symbol: "KGS"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "leka albaneză", Symbol: "ALL"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dolar din Hong Kong", Symbol: "HKD"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sud-coreean", Symbol: "KRW"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rublă Letonia", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panameză", Symbol: "PAB"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga tongană", Symbol: "TOP"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc belgian (financiar)", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dolar jamaican", Symbol: "JMD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rublă sovietică", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "coroană daneză", Symbol: "DKK"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol bolivian", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dolar fijian", Symbol: "FJD"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colon El Salvador", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham din Emiratele Arabe Unite", Symbol: "AED"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA BEAC", Symbol: "FCFA"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unitate de cont europeană (XBC)", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "coroană islandeză", Symbol: "ISK"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso mexican de argint (1861–1992)", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "leu românesc (1952–2006)", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dolar Singapore", Symbol: "SGD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "dinar macedonean", Symbol: "MKD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc burundez", Symbol: "BIF"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "leva bulgărească", Symbol: "BGN"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birman", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malaiezian", Symbol: "MYR"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupie nepaleză", Symbol: "NPR"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso Uruguay (1975–1993)", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc belgian (convertibil)", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruvian (1863–1965)", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rublă rusească", Symbol: "RUB"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dolar Zimbabwe (2009)", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupie mauritiană", Symbol: "MUR"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat din Myanmar", Symbol: "MMK"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guarani paraguayan", Symbol: "PYG"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dolar surinamez", Symbol: "SRD"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dolar Zimbabwe (2008)", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "liră italiană", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platină", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "aur", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari georgian", Symbol: "GEL"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar kuweitian", Symbol: "KWD"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadjic", Symbol: "TJS"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "șiling ugandez", Symbol: "UGX"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso columbian", Symbol: "COP"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "liră sudaneză (1957–1998)", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yemenit", Symbol: "YER"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dolar Barbados", Symbol: "BBD"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo din Capul Verde", Symbol: "CVE"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar Venezuela (1871–2008)", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "cod monetar de test", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan chinezesc", Symbol: "CNY"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rublă Tadjikistan", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde din Haiti", Symbol: "HTG"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "coroană estoniană", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "dolar din Brunei", Symbol: "BND"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dolar liberian", Symbol: "LRD"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saudit", Symbol: "SAR"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thailandez", Symbol: "THB"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dolar din Insulele Cayman", Symbol: "KYD"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexican", Symbol: "MXN"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "gulden olandez", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar sârbesc", Symbol: "RSD"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanez", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghanez", Symbol: "GHS"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc Madagascar", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "șiling tanzanian", Symbol: "TZS"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libian", Symbol: "LYD"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "liră egipteană", Symbol: "EGP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi din Gambia", Symbol: "GMD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dolar din Bahamas", Symbol: "BSD"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sud-african (financiar)", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: "CFPF"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar Yemen", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unitate monetară europeană", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marcă germană", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominican", Symbol: "DOP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc francez", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "zair nou", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "pesetă andorrană", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dolar canadian", Symbol: "CAD"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubanez", Symbol: "CUP"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr etiopian", Symbol: "ETB"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "liră israeliană", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc Mali", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatarian", Symbol: "QAR"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "gulden Surinam", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florin aruban", Symbol: "AWG"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong vietnamez", Symbol: "VND"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dolar neozeelandez", Symbol: "NZD"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambian (1968–2012)", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "liră libaneză", Symbol: "LBP"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen japonez", Symbol: "JPY"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritană", Symbol: "MRO"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambican", Symbol: "MZN"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambian", Symbol: "ZMW"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc guineean", Symbol: "GNF"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu din Vanuatu", Symbol: "VUV"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argint", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marcă convertibilă din Bosnia și Herțegovina", Symbol: "BAM"}}
