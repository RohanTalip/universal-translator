package fy

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"DEM": ut.Currency{Currency: "DEM", DisplayName: "Dútske mark", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambiaanske dalasi", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruaanske nieuwe sol", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Braziliaanske cruzeiro", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Nederlânske gûne", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Russyske roebel (1991–1998)", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somalyske shilling", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Sûd-Soedaneeske pûn", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Uruguayaanske peso en geïndexeerde eenheden", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongoalske tugrik", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Sileenske escudo", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Nij-Seelânske dollar", Symbol: "NZ$"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Santomese dobra", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Sûd-Afrikaanske rand", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australyske dollar", Symbol: "AU$"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Zwitserse frank", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israëlyske nieuwe shekel", Symbol: "₪"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaikaanske dollar", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Monegaskyske frank", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzaniaanske shilling", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Goud", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbadaanske dollar", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritaanske ouguiya", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armeense dram", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Soedaneeske pûn (1957–1998)", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "Amerikaanske dollar (folgjende dei)", Symbol: ""}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "Sucre", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Bulgaarse socialistyske lev", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Ecuadoraanske unidad de valor constante (UVC)", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Lúksemboargske frank", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwandese frank", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA-franc BCEAO", Symbol: "CFA"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bulgaarse harde lev", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Bruneise dollar", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Ethiopyske birr", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Keniaanske shilling", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberiaanske dollar", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portugeeske escudo", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Alde Roemeenske leu", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Soedaneeske dinar", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Salvadoraanske colón", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosnyske dinar", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Europeeske rekkenienheid (XBC)", Symbol: "XBC"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Simbabwaanske dollar", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Syrysk pûn", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guinese syli", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Iraakse dinar", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipynske peso", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Bulgaarse lev (1879–1952)", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltarees pûn", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Soedaneeske pûn", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Saïreeske Saïre", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Braziliaanske cruzado", Symbol: ""}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Nije Bosnyske dinar (1994–1997)", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeriaanske naira", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Argentynske peso (1881–1970)", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Lúksemboargske convertibele franc", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iraanske rial", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritiaanske roepie", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmeense manat", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Boliviaanske mvdol", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Arubaanske gulden", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Tsjechoslowaakse harde koruna", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Spaanske peseta", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Letse roebel", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Marokkaanske franc", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macause pataca", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Sloveenske tolar", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentynske peso (1983–1985)", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djiboutiaanske frank", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "Peruaanske sol", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR franc", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyaanske dollar", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmarese kyat", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Turkmeense manat (1993–2009)", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angolese kwanza reajustado (1995–1999)", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litouwse litas", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kubaanske convertibele peso", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kaapverdyske escudo", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algerynske dinar", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Sovjet-roebel", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguayaanske peso", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Europeeske rekkenienheid (XBD)", Symbol: "XBD"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Wit-Russyske roebel", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Braziliaanske real", Symbol: "R$"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Sineeske yuan renminbi", Symbol: "CN¥"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fiji-dollar", Symbol: "FJ$"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Surinaamske gulden", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Oekraïense karbovanetz", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Palladium", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda-dollar", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Braziliaanske cruzeiro (1990–1993)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Sileenske peso", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Deenske kroon", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghaanske afghani", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominikaanske peso", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgyske lari", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamese balboa", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Simbabwaanske dollar (2009)", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutaanske ngultrum", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawyske kwacha", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omaanske rial", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongaanske paʻanga", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Caymaneilânske dollar", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costaricaanske colón", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hongkongske dollar", Symbol: "HK$"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Maltese lire", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaraguaanske córdoba", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thaise baht", Symbol: "฿"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Oegandese shilling", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Sileenske unidades de fomento", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguayaanske guarani", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Kroatyske dinar", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Meksikaanske unidad de inversion (UDI)", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Roemeenske leu", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saoedi-Arabyske riyal", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomon-dollar", Symbol: "SI$"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Sint-Heleenske pûn", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "European Currency Unit", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Litouwse talonas", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Tsjechyske kroon", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swazyske lilangeni", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tadzjikistaanske roebel", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kubaanske peso", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portugees-Guinese escudo", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Macedonyske denar", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platina", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Verenigde Arabyske Emiraten-dirham", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malagassyske ariary", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mozambikaanske metical", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad en Tobago-dollar", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswaanske pula", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angolese nieuwe kwanza (1990–2000)", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belizaanske dollar", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordaanske dinar", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldavyske leu", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Alde Mozambikaanske metical", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistaanske roepie", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkse lira", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoaanske tala", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundese frank", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Franse gouden franc", Symbol: "XFO"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemalteekse quetzal", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Birmese kyat", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Cyprysk pûn", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komorese frank", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Servyske dinar", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Braziliaanske cruzeiro novo (1967–1986)", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Joegoslavyske noviy-dinar", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA-frank", Symbol: "FCFA"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laotiaanske kip", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libyske dinar", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Maltees pûn", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mozambikaanske escudo", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Simbabwaanske dollar (2008)", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnyske convertibele mark", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russyske roebel", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptysk pûn", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Spaanske peseta (account A)", Symbol: ""}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Alde Sûd-Koreaanske won (1945–1953)", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Malagassyske franc", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Europeeske gearfoege ienheid", Symbol: "XBA"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "ADB-rekkenienheid", Symbol: "XUA"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemenityske rial", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadoraanske sucre", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Equatoriaal-Guinese ekwele guineana", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Noard-Koreaanske won", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Sûd-Koreaanske won", Symbol: "₩"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolese kwanza (1977–1990)", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Kolombiaanske peso", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidad de Valor Real", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finse markka", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Grykse drachme", Symbol: ""}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Macedonyske denar (1992–1993)", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolese kwanza", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinee-Bissause peso", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Sûd-Koreaanske hwan (1953–1962)", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Koeweitse dinar", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Lúksemboargske finansjele franc", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Moldavyske cupon", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Meksikaanske sulveren peso (1861–1992)", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Franse UIC-franc", Symbol: "XFU"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrese nakfa", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Joegoslavyske convertibele dinar", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Turkse lire", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Nije Taiwanese dollar", Symbol: "NT$"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lankaanske roepie", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviaanske boliviano", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Brits pûn", Symbol: "£"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Poalske zloty", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singaporese dollar", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Fenezolaanske bolivar (1871–2008)", Symbol: ""}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Alde Fietnameeske dong (1978–1985)", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Sulver", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahreinse dinar", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Hongaarse forint", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinaamske dollar", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "Amerikaanske dollar (zelfde dei)", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Special Drawing Rights", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Argentynske peso ley (1970–1983)", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Yslânske kroon", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Poalske zloty (1950–1995)", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatuaanske vatu", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Valutacode voor testdoeleinden", Symbol: "XTS"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Eastenrykse schilling", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Hondurese lempira", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japanse yen", Symbol: "JP¥"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Maleisyske ringgit", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguayaanske peso (1975–1993)", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bengalese taka", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Meksikaanske peso", Symbol: "MX$"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "East-Dútske ostmark", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR euro", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Franske franc", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Israëlysk pûn", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Wit-Russyske nieuwe roebel (1994–1999)", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbeidzjaanske manat", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamaanske dollar", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timorese escudo", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Sûd-Afrikaanske rand (finansjeel)", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorrese peseta", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slowaakse koruna", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Fenezolaanske bolivar", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Sambiaanske kwacha (1968–2012)", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Azerbeidzjaanske manat (1993–2006)", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rhodesyske dollar", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Noarske kroon", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papuaanske kina", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunesyske dinar", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "East-Karibyske dollar", Symbol: "EC$"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Sambiaanske kwacha", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanese lek", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belgyske frank (finansjeel)", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Congolese frank", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokkaanske dirham", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivyske rufiyaa", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierraleoonse leone", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Amerikaanske dollar", Symbol: "US$"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP-franc", Symbol: "XPF"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belgyske frank", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibyske dollar", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Joegoslavyske harde dinar", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Saïreeske nije Saïre", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgizyske som", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haïtiaanske gourde", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazachstaanske tenge", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesothaanske loti", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychelse roepie", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belgyske frank (convertibel)", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesyske roepia", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalese roepie", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "Fietnameeske dong", Symbol: "₫"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Alde Servyske dinar", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Nederlânsk-Antilliaanske gûne", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falklâneilânske pûn", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afghani (1927–2002)", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libaneeske pûn", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Letse lats", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nicaraguaanske córdoba (1988–1991)", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Oekraïense hryvnia", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Unbekende muntienheid", Symbol: "XXX"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Jemenityske dinar", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Canadese dollar", Symbol: "C$"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Braziliaanske cruzado novo", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Ierske pûn", Symbol: "IEP"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaarse lev", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Estlânske kroon", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanese cedi", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinese franc", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indiase roepie", Symbol: "₹"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italiaanske lire", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Sweedske kroon", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Braziliaanske cruzeiro (1942–1967)", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peruaanske inti", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katarese rial", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Oegandese shilling (1966–1987)", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET-fondsen", Symbol: "XRE"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroatyske kuna", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Boliviaanske peso", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Spaanske peseta (convertibele account)", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Georgyske kupon larit", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghanese cedi (1979–2007)", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kambodjaanske riel", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Malinese franc", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadzjikistaanske somoni", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviaanske boliviano (1863–1963)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentynske peso", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Oezbekistaanske sum", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Europeeske monetaire ienheid", Symbol: "XBB"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Joegoslavyske herfoarme dinar (1992–1993)", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentynske austral", Symbol: ""}}
