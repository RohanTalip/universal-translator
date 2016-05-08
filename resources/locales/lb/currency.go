package lb

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BUK": ut.Currency{Currency: "BUK", DisplayName: "Birmanesche Kyat", Symbol: "BUK"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Spuenesch Peseta", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fidschi-Dollar", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Chileneschen Unidad de Fomento", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Tschechesch Kroun", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Israelescht Pond", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad-an-Tobago-Dollar", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Lëtzebuerger Frang", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Mexikaneschen Unidad de Inversion (UDI)", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kambodschanesche Riel", Symbol: "KHR"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghanesch Afghani", Symbol: "AFN"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angolaneschen Neie Kwanza (1990–2000)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Chilenesche Peso", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iranesch Rial", Symbol: "IRR"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Lettesche Rubel", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Südsudanesescht Pond", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Jugoslaweschen Neien Dinar (1994–2002)", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentinesche Peso (1983–1985)", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libanesescht Pond", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Maltesesch Lira", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Sowjetesche Rubel", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tuneseschen Dinar", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italienesch Lira", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Ruanda-Frang", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguayesche Peso", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Südafrikanesche Rand", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Serbeschen Dinar (2002–2006)", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "El-Salvador-Colón", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "US Dollar (Nächsten Dag)", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "US Dollar (Selwechten Dag)", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Onbekannt Währung", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Mexikanesche Peso", Symbol: "MX$"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentineschen Austral", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosnien an Herzegowina Dinar (1992–1994)", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kubanesche Peso (konvertibel)", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Equatorialguinea-Ekwele", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuwait-Dinar", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberianeschen Dollar", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldawesche Leu", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Onze Sëlwer", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbeschen Dinar", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israeleschen Neie Schekel", Symbol: "₪"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudanesescht Pond", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Simbabwe-Dollar (2008)", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kap-Verde-Escudo", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Estnesch Kroun", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Irescht Pond", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Kenia-Schilling", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentinesche Peso", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litauesche Litas", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokkaneschen Dirham", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Ostkaribeschen Dollar", Symbol: "EC$"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Onz Palladium", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha (1968–2012)", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laoteschen Kip", Symbol: "LAK"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afghanesch Afghani (1927–2002)", Symbol: "AFA"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolanesche Kwanza", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "DDR-Mark", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algereschen Dinar", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambia-Dalasi", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroatesche Kuna", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Islännesch Kroun", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Turkmenistan-Manat (1993–2009)", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinea-Bissau Peso", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyana-Dollar", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordaneschen Dinar", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadschikistan-Somoni", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaire-Zaïre (1971–1993)", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswanesch Pula", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR-Euro", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Honduras-Lempira", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Kroateschen Dinar", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nicaraguanesche Córdoba (1988–1991)", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norwegesch Kroun", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Europäesch Währungseenheet (XEU)", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnien an Herzegowina Konvertéierbar Mark", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinea-Frang", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Rumänesche Leu", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Europäesch Währungseenheet (XBB)", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bulgaresch Lew (1962–1999)", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Brasilianesche Cruzeiro (1942–1967)", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tadschikistan-Rubel", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Uganda-Schilling (1966–1987)", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mosambikanesche Metical (1980–2006)", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaire-Neien Zaïre (1993–1998)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaresch Lew", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peruaneschen Inti", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatu-Vatu", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Antillen-Gulden", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Brasilianesche Cruzeiro (1993–1994)", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoaneschen Tala", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP-Frang", Symbol: "CFPF"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominikanesche Peso", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadorianesche Sucre", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finnesch Mark", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Ungaresche Forint", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychellen-Rupie", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Surinamesche Gulden", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongoleschen Tögrög", Symbol: "MNT"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Philippinnesche Peso", Symbol: "PHP"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongaeschen Paʻanga", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Usbekistan-Sum", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Litaueschen Talonas", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapur-Dollar", Symbol: "SGD"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukraineschen Hrywnja", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemen-Rial", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda-Dollar", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasilianesche Real", Symbol: "R$"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Wäissrussesche Rubel", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Dänesch Kroun", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigerianeschen Naira", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Sloweneschen Tolar", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angolanesche Kwanza Reajustado (1995–1999)", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Aruba-Florin", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri-Lanka-Rupie", Symbol: "LKR"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomonen-Dollar", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbados-Dollar", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Renminbi Yuan", Symbol: "CN¥"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Testwährung", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanesche Lek", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Griichesch Drachme", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indesch Rupie", Symbol: "₹"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruaneschen Neie Sol", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Neien Taiwan-Dollar", Symbol: "NT$"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivianesche Boliviano", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Onz Platin", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanadeschen Dollar", Symbol: "CA$"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritréieschen Nakfa", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptescht Pond", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Spuenesch Peseta (konvertibel)", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Lëtzebuerger Finanz-Frang", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rhodeseschen Dollar", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Syrescht Pond", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Simbabwe-Dollar (2009)", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Hollännesche Gulden", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Madagaskar-Ariary", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamaesche Balboa", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "St. Helena-Pond", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolanesche Kwanza (1977–1990)", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belize-Dollar", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komore-Frang", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Russesche Rubel (1991–1998)", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Onze Gold", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA-Frang (BCEAO)", Symbol: "CFA"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Jugoslaweschen Dinar (1966–1990)", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanaeschen Cedi", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katar-Riyal", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slowakesch Kroun", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "VAE-Dirham", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armeneschen Dram", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Kaiman-Dollar", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malayseschen Ringgit", Symbol: "MYR"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mosambikaneschen Escudo", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somalia-Schilling", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Éisträichesche Schilling", Symbol: "öS"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Brasilianesche Cruzeiro (1990–1993)", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Zypern-Pond", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Lëtzebuerger Frang (konvertibel)", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltar-Pond", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Nordkoreanesche Won", Symbol: "KPW"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldiven-Rupie", Symbol: "MVR"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Peruanesche Sol (1863–1965)", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra-leonesche Leone", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET Funds", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalesesch Rupie", Symbol: "NPR"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Neiséiland-Dollar", Symbol: "NZ$"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritius-Rupie", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Europäesch Rechnungseenheet (XBD)", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Georgesche Kupon Larit", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Mexikanesche Sëlwer-Peso (1861–1992)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haitianesch Gourde", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguayesche Peso (1975–1993)", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Kolumbianesche Peso", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Tschechoslowakesch Kroun", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesesch Rupiah", Symbol: "IDR"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauretaneschen Ouguiya", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunei-Dollar", Symbol: "BND"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Bolivianesche Peso", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Däitsch Mark", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irakeschen Dinar", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Malesche Frang", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolivianseche Mvdol", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japanesche Yen", Symbol: "¥"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Rumänesche Leu (1952–2006)", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA-Frang (BEAC)", Symbol: "FCFA"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Franséische Gold-Frang", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australeschen Dollar", Symbol: "AU$"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Brasilianesche Cruzado (1986–1989)", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaika-Dollar", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Marokkanesche Frang", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papua-Neiguinéiesche Kina", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudi-Rial", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorranesch Peseta", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Franséische Frang", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lettesche Lats", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibia-Dollar", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgesche Lari", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omanesche Rial", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Sudanesescht Pond (1957–1998)", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Tierkesch Lira (1922–2005)", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vietnameseschen Dong", Symbol: "₫"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Franséischen UIC-Frang", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Simbabwe-Dollar (1980–2008)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Brasilianesche Cruzado Novo (1989–1990)", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macau-Pataca", Symbol: "MOP"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mosambikanesche Metical", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezolanesche Bolívar (1871–2008)", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Europäesch Rechnungseenheet (XBC)", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Uganda-Schilling", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahama-Dollar", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falkland-Pond", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong-Kong-Dollar", Symbol: "HK$"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Sudaneseschen Dinar (1992–2007)", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundi-Frang", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongo-Frang", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kubanesche Peso", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Dschibuti-Frang", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Britescht Pond", Symbol: "£"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghanaeschen Cedi (1979–2007)", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinameschen Dollar", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Ukrainesche Karbovanetz", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Jemen-Dinar", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Aserbaidschan-Manat (1993–2006)", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belsche Finanz-Frang", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Schwäizer Frang", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmaresche Kyat", Symbol: "MMK"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timor-Escudo", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Südafrikanesche Rand (Finanz)", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Aserbaidschan-Manat", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Ethiopescht Birr", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kasacheschen Tenge", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Mazedoneschen Denar", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistanesch Rupie", Symbol: "PKR"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Polneschen Zloty (1950–1995)", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Jugoslaweschen Dinar (konvertibel)", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "US-Dollar", Symbol: "$"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Brasilianesche Cruzeiro Novo (1967–1986)", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Wäissrussesche Rubel (1994–1999)", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa-Rica-Colón", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgisesche Som", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libeschen Dinar", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swasilännesche Lilangeni", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahrain-Dinar", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Tierkesch Lira", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladesch-Taka", Symbol: "BDT"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belsche Frang (konvertibel)", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Polneschen Zloty", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tansania-Schilling", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR-Frang", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemaltekesche Quetzal", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaraguanesche Córdoba", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "São-toméeschen Dobra", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmenistan-Manat", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Sonnerzéiungsrecht", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutan-Ngultrum", Symbol: "BTN"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Verrechnungseenheete fir Ecuador", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Madagaskar-Frang", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Maltesescht Pond", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawi-Kwacha", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguayeschen Guaraní", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russesche Rubel", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Südkoreanesche Won", Symbol: "₩"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portugiseschen Escudo", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezolanesche Bolívar", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belsche Frang", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guinéiesche Syli", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Europäesch Rechnungseenheet", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Spuenesch Peseta (A–Konten)", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portugisesch-Guinea Escudo", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Schwedesch Kroun", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thailännesche Baht", Symbol: "฿"}}
