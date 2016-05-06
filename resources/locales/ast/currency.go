package ast

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MZE": ut.Currency{Currency: "MZE", DisplayName: "Escudu mozambicanu", Symbol: "MZE"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "pesu uruguayu", Symbol: "UYU"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats letón", Symbol: "LVL"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rufiyaa maldiviana", Symbol: "MVR"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Córdoba nicaraguanu (1988–1991)", Symbol: "NIC"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Florín surinamés", Symbol: "SRG"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dólar barbadianu", Symbol: "BBD"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Marcu alemán", Symbol: "DEM"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Escudo de Guinea portuguesa", Symbol: "GWE"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won surcoreanu", Symbol: "₩"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brasilanu", Symbol: "R$"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dólar canadiense", Symbol: "CA$"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar iraquín", Symbol: "IQD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti de Lesothu", Symbol: "LSL"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Shilling somalín", Symbol: "SOS"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari xeorxanu", Symbol: "GEL"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dólar hongkonés", Symbol: "HK$"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Won surcoreanu (1945–1953)", Symbol: "KRO"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia de Sri Lanka", Symbol: "LKR"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Francu francés", Symbol: "FRF"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Francu monegascu", Symbol: "MCF"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat turcomanu", Symbol: "TMT"}, "USN": ut.Currency{Currency: "USN", DisplayName: "Dólar d’EE.XX. (día siguiente)", Symbol: "USN"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afganí afganistanu", Symbol: "AFN"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dólar guyanés", Symbol: "GYD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Denar macedoniu", Symbol: "MKD"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Dinar convertible yugoslavu (1990–1992)", Symbol: "YUN"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dólar liberianu", Symbol: "LRD"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Unidá d’inversión mexicana", Symbol: "MXV"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical mozambicanu", Symbol: "MZN"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial omanianu", Symbol: "OMR"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Dinar nuevu de Bosnia-Herzegovina (1994–1997)", Symbol: "BAN"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Francu belga (financieru)", Symbol: "BEL"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev búlgaru", Symbol: "BGN"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivianu de Bolivia", Symbol: "BOB"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dólar singapuranu", Symbol: "SGD"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Oru", Symbol: "XAU"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaire nuevu zairiegu (1993–1998)", Symbol: "ZRN"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghanianu", Symbol: "GHS"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone sierralleonés", Symbol: "SLL"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Llibra siria", Symbol: "SYP"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi ghanianu (1979–2007)", Symbol: "GHC"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupiah indonesia", Symbol: "IDR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayu", Symbol: "PYG"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rublu rusu", Symbol: "RUB"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Kwanza angolanu reaxustáu (1995–1999)", Symbol: "AOR"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat azerbaixanu", Symbol: "AZN"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "francu suizu", Symbol: "CHF"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Dólar del Bancu Popular Chinu", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Grivna ucraína", Symbol: "UAH"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolívar venezolanu", Symbol: "VEF"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Rublu taxiquistanín", Symbol: "TJR"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Austral arxentín", Symbol: "ARA"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu moldavu", Symbol: "MDL"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Corona eslovaca", Symbol: "SKK"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni suazilandés", Symbol: "SZL"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som kirguistanín", Symbol: "KGS"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Llibra maltesa", Symbol: "MTP"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu vanuatuanu", Symbol: "VUV"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "Sucre", Symbol: "XSU"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Unidá de cuenta chilena (UF)", Symbol: "CLF"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Tolar eslovenu", Symbol: "SIT"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Manat turcomanu (1993–2009)", Symbol: "TMM"}, "USS": ut.Currency{Currency: "USS", DisplayName: "Dólar d’EE.XX. (mesmu día)", Symbol: "USS"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar arxelín", Symbol: "DZD"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Peseta española (cuenta A)", Symbol: "ESA"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "Fondos RINET", Symbol: "XRE"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar baḥreiní", Symbol: "BHD"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa eritréu", Symbol: "ERN"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Ekwele de Guinea Ecuatorial", Symbol: "GQE"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "corona islandesa", Symbol: "ISK"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Pesu de plata mexicanu (1861–1992)", Symbol: "MXP"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Shilling tanzanianu", Symbol: "TZS"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Kwanza nuevu angolanu (1990–2000)", Symbol: "AON"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Cruzeiru brasilanu (1942–1967)", Symbol: "BRZ"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "llibra esterlina", Symbol: "£"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Dinar croata", Symbol: "HRD"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Dinar de Bosnia-Herzegovina (1992–1994)", Symbol: "BAD"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Leu rumanu (1952–2006)", Symbol: "ROL"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambianu", Symbol: "GMD"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Dracma griegu", Symbol: "GRD"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Dinar reformáu yugoslavu (1992–1993)", Symbol: "YUR"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dólar xamaicanu", Symbol: "JMD"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Zloty polacu (1950–1995)", Symbol: "PLZ"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyal saudita", Symbol: "SAR"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Escudu timorés", Symbol: "TPE"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "francu CFA BCEAO", Symbol: "CFA"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand sudafricanu", Symbol: "ZAR"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Cruzeiro nuevu brasilanu (1967–1986)", Symbol: "BRB"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Rublu bielorrusu", Symbol: "BYR"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidá de valor real colombiana", Symbol: "COU"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Florín neerlandés", Symbol: "NLG"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Llira maltesa", Symbol: "MTL"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metical mozambicanu (1980–2006)", Symbol: "MZM"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha malauianu", Symbol: "MWK"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial qatarín", Symbol: "QAR"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "francu CFP", Symbol: "CFPF"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "corona sueca", Symbol: "SEK"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra de Santu Tomé y Príncipe", Symbol: "STD"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florín arubanu", Symbol: "AWG"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "corona danesa", Symbol: "DKK"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia mauriciana", Symbol: "MUR"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Inti peruanu", Symbol: "PEI"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirḥam marroquín", Symbol: "MAD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina papuana", Symbol: "PGK"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Unidá de cuenta europea (XBD)", Symbol: "XBD"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Lev socialista búlgaru", Symbol: "BGM"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Cruzado nuevu brasilanu (1989–1990)", Symbol: "BRN"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Pesu cubanu convertible", Symbol: "CUC"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dólar caimanés", Symbol: "KYD"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Unidá monetaria europea", Symbol: "XBB"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirḥam de los Emiratos Árabes Xuníos", Symbol: "AED"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colón costarricanu", Symbol: "CRC"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna croata", Symbol: "HRK"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariary malgaxe", Symbol: "MGA"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Cruzado brasilanu (1986–1989)", Symbol: "BRC"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Dólar nuevu taiwanés", Symbol: "NT$"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Francu CFA centroafricanu", Symbol: "FCFA"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar kuwaitianu", Symbol: "KWD"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "llibra sudanesa (1957–1998)", Symbol: "SDP"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Francu financieru luxemburgués", Symbol: "LUL"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dólar neozelandés", Symbol: "NZ$"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia seixelesa", Symbol: "SCR"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "llibra sudanesa", Symbol: "SDG"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Peseta andorrana", Symbol: "ADP"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Pesu arxentín (1881–1970)", Symbol: "ARM"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Francu belga", Symbol: "BEF"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Peseta española", Symbol: "ESP"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Unidá Compuesta Europea", Symbol: "XBA"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Dinar fuerte yugoslavu (1966–1990)", Symbol: "YUD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa panamiegu", Symbol: "PAB"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolivianos mvdol", Symbol: "BOV"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum butanés", Symbol: "BTN"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Escudu chilenu", Symbol: "CLE"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde haitianu", Symbol: "HTG"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Llibra irlandesa", Symbol: "IEP"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Llibra israelina", Symbol: "ILP"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Llibra libanesa", Symbol: "LBP"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Francu malgaxe", Symbol: "MGF"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nixeriana", Symbol: "NGN"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Lev búlgaru (1879–1952)", Symbol: "BGO"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Sucre ecuatorianu", Symbol: "ECS"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Francu oru francés", Symbol: "XFO"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira hondurana", Symbol: "HNL"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Francu ruandés", Symbol: "RWF"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Pesu uruguayu (Unidaes indexaes)", Symbol: "UYI"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Dong vietnamín (1978–1985)", Symbol: "VNN"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint húngaru", Symbol: "HUF"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malasiu", Symbol: "MYR"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Dólar rodesianu", Symbol: "RHD"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Llira turca", Symbol: "TRY"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Rublu letón", Symbol: "LVR"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha zambianu (1968–2012)", Symbol: "ZMK"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Drechos especiales de xiru", Symbol: "XDR"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "francu congolés", Symbol: "CDF"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Pesu dominicanu", Symbol: "DOP"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr etíope", Symbol: "ETB"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni taxiquistanín", Symbol: "TJS"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Corona fuerte checoslovaca", Symbol: "CSK"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "corona noruega", Symbol: "NOK"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht tailandés", Symbol: "฿"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Paladiu", Symbol: "XPD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbiu", Symbol: "RSD"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Rand sudafricanu (financieru)", Symbol: "ZAL"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Francu xibutianu", Symbol: "DJF"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guineanu", Symbol: "GNS"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Francu comoranu", Symbol: "KMF"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge kazaquistanín", Symbol: "KZT"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Divisa desconocida", Symbol: "XXX"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Peseta española (cuenta convertible)", Symbol: "ESB"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Xequel israelín (1980–1985)", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Córdoba nicaraguanu", Symbol: "NIO"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dólar estaunidense", Symbol: "$"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Dinar yemenín", Symbol: "YDD"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Lek albanés (1946–1965)", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Pesu cubanu", Symbol: "CUP"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupia paquistanina", Symbol: "PKR"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som uzbequistanín", Symbol: "UZS"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Plata", Symbol: "XAG"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Dinar nuevu yugoslavu (1994–2002)", Symbol: "YUM"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dólar bruneyanu", Symbol: "BND"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Cruzeiro brasilanu (1990–1993)", Symbol: "BRE"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "pesu chilenu", Symbol: "CLP"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libiu", Symbol: "LYD"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Francu luxemburgués", Symbol: "LUF"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Karbovanets ucraína", Symbol: "UAK"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Pesu bolivianu", Symbol: "BOP"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Cruzeiro brasilanu (1993–1994)", Symbol: "BRR"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Xequel nuevu israelín", Symbol: "₪"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won norcoreanu", Symbol: "KPW"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dólar namibianu", Symbol: "NAD"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dólar australianu", Symbol: "A$"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Pesu de Guinea-Bisáu", Symbol: "GWP"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel camboyanu", Symbol: "KHR"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laosianu", Symbol: "LAK"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Unidá de cuenta europea (XBC)", Symbol: "XBC"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Unidá de divisa europea", Symbol: "XEU"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Dólar zimbabuanu (2008)", Symbol: "ZWR"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Guílder de les Antilles Neerlandeses", Symbol: "ANG"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Pesu Ley arxentín (1970–1983)", Symbol: "ARL"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Shilling kenianu", Symbol: "KES"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong vietnamín", Symbol: "₫"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Colón salvadorianu", Symbol: "SVC"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "unidá de cuenta ADB", Symbol: "XUA"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Kwanza angolanu (1977–1991)", Symbol: "AOK"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Pesu arxentín (1983–1985)", Symbol: "ARP"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu rumanu", Symbol: "RON"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "llibra de Santa Lena", Symbol: "SHP"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar xordanu", Symbol: "JOD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca de Macáu", Symbol: "MOP"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dólar surinamés", Symbol: "SRD"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dólar de Trinidá y Tobagu", Symbol: "TTD"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afganí afganistanu (1927–2002)", Symbol: "AFA"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka bangladexí", Symbol: "BDT"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Unidá ecuatoriana de valor constante", Symbol: "ECV"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "llibra malviniana", Symbol: "FKP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Shilling ugandés", Symbol: "UGX"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunecín", Symbol: "TND"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Dólar zimbabuanu (2009)", Symbol: "ZWL"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaire zairiegu (1971–1993)", Symbol: "ZRZ"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dólar bahamés", Symbol: "BSD"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial iranín", Symbol: "IRR"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Denar macedoniu (1992–1993)", Symbol: "MKN"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platín", Symbol: "XPT"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dólar salomonés", Symbol: "SBD"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Francu belga (convertible)", Symbol: "BEC"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Corona checa", Symbol: "CZK"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "llibra exipciana", Symbol: "EGP"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Pesu mexicanu", Symbol: "MX$"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dólar del Caribe Oriental", Symbol: "EC$"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "Francu WIR", Symbol: "CHW"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan chinu", Symbol: "CN¥"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupia india", Symbol: "₹"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "paʻanga tonganu", Symbol: "TOP"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial yemenín", Symbol: "YER"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza angolanu", Symbol: "AOA"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Francu burundianu", Symbol: "BIF"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Marcu d’Alemaña Oriental", Symbol: "DDM"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupia nepalesa", Symbol: "NPR"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marcu convertible de Bosnia-Herzegovina", Symbol: "BAM"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dólar fixanu", Symbol: "FJD"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat de Myanmar", Symbol: "MMK"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "pesu arxentín", Symbol: "ARS"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Kyat birmanu", Symbol: "BUK"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Rublu nuevu bielorrusu (1994–1999)", Symbol: "BYB"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Pesu uruguayu (1975–1993)", Symbol: "UYP"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yen xaponés", Symbol: "¥"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritanu", Symbol: "MRO"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Bolívar venezolanu (1871–2008)", Symbol: "VEB"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Códigu monetariu de prueba", Symbol: "XTS"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Rublu rusu (1991–1998)", Symbol: "RUR"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Francu UIC francés", Symbol: "XFU"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Manat azerbaixanu (1993–2006)", Symbol: "AZM"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dólar bermudianu", Symbol: "BMD"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Cupón moldavu", Symbol: "MDC"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "sol nuevu peruanu", Symbol: "PEN"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Llira italiana", Symbol: "ITL"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Francu convertible luxemburgués", Symbol: "LUC"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik mongol", Symbol: "MNT"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Llira turca (1922–2005)", Symbol: "TRL"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Chelín austriacu", Symbol: "ATS"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Lev fuerte búlgaru", Symbol: "BGL"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Llibra xipriota", Symbol: "CYP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal guatemalianu", Symbol: "GTQ"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoanu", Symbol: "WST"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Sol peruanu (1863–1965)", Symbol: "PES"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanés (1992–2007)", Symbol: "SDD"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Kupon larit xeorxanu", Symbol: "GEK"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas lituanu", Symbol: "LTL"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "francu marroquín", Symbol: "MAF"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Rupia maldiviana (1947–1981)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek albanés", Symbol: "ALL"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Francu malianu", Symbol: "MLF"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dólar zimbabuanu (1980–2008)", Symbol: "ZWD"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram armeniu", Symbol: "AMD"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Corona estonia", Symbol: "EEK"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "francu guineanu", Symbol: "GNF"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Talonas lituanu", Symbol: "LTT"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviano de Bolivia (1863–1963)", Symbol: "BOL"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudu cabuverdianu", Symbol: "CVE"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "llibra sursudanesa", Symbol: "SSP"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha zambianu", Symbol: "ZMW"}, "COP": ut.Currency{Currency: "COP", DisplayName: "pesu colombianu", Symbol: "COP"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "pesu filipín", Symbol: "PHP"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloty polacu", Symbol: "PLN"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Rublu soviéticu", Symbol: "SUR"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula botsuaniana", Symbol: "BWP"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Marcu finlandés", Symbol: "FIM"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "llibra de Xibraltar", Symbol: "GIP"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Hwan surcoreanu (1953–1962)", Symbol: "KRH"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Escudu portugués", Symbol: "PTE"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Shilling ugandés (1966–1987)", Symbol: "UGS"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dólar belizianu", Symbol: "BZD"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "Euru WIR", Symbol: "CHE"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Dinar serbiu (2002–2006)", Symbol: "CSD"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Corona islandesa (1918–1981)", Symbol: ""}}
