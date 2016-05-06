package id

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"CSD": ut.Currency{Currency: "CSD", DisplayName: "Dinar Serbia (2002–2006)", Symbol: "CSD"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Franc Luksemburg", Symbol: "LUF"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Shilling Tanzania", Symbol: "TZS"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Rupiah Indonesia", Symbol: "Rp"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar Kuwait", Symbol: "KWD"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Satuan Hitung (UF) Cile", Symbol: "CLF"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Rubel Rusia (1991–1998)", Symbol: "RUR"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra Sao Tome dan Principe", Symbol: "STD"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi Ghana (1979–2007)", Symbol: "GHC"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirham Maroko", Symbol: "MAD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina Papua Nugini", Symbol: "PGK"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Leu Rumania (1952–2006)", Symbol: "ROL"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Lev Bulgaria (1879–1952)", Symbol: "BGO"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Cruzeiro Brasil (1942–1967)", Symbol: "BRZ"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Hwan Korea Selatan (1953–1962)", Symbol: "KRH"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Dinar Baru Yugoslavia (1994–2002)", Symbol: "YUM"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Kupon Larit Georgia", Symbol: "GEK"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Tolar Slovenia", Symbol: "SIT"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Dinar Sudan (1992–2007)", Symbol: "SDD"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Syli Guinea", Symbol: "GNS"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Krona Islandia (1918–1981)", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge Kazakstan", Symbol: "KZT"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Escudo Portugal", Symbol: "PTE"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Karbovanet Ukraina", Symbol: "UAK"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Franc CFP", Symbol: "CFPF"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dolar Australia", Symbol: "AU$"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Mark Jerman", Symbol: "DEM"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha Malawi", Symbol: "MWK"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Escudo Mozambik", Symbol: "MZE"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dolar Barbados", Symbol: "BBD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dolar Bahama", Symbol: "BSD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca Makau", Symbol: "MOP"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Pound Saint Helena", Symbol: "SHP"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna Kroasia", Symbol: "HRK"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint Hungaria", Symbol: "HUF"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Manat Turkmenistan (1993–2009)", Symbol: "TMM"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Unit Keuangan Eropa", Symbol: "XBB"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Rubel Tajikistan", Symbol: "TJR"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr Etiopia", Symbol: "ETB"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dolar Hong Kong", Symbol: "HK$"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som Kirgistan", Symbol: "KGS"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Unit Investasi Meksiko", Symbol: "MXV"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha Zambia (1968–2012)", Symbol: "ZMK"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram Armenia", Symbol: "AMD"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Drachma Yunani", Symbol: "GRD"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Franc UIC Perancis", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Palladium", Symbol: "XPD"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Kwanza Angola (1977–1991)", Symbol: "AOK"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Riel Kamboja", Symbol: "KHR"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia Ukraina", Symbol: "UAH"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Shilling Uganda (1966–1987)", Symbol: "UGS"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Satuan Mata Uang Eropa", Symbol: "XEU"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar Bahrain", Symbol: "BHD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dolar Fiji", Symbol: "FJD"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yen Jepang", Symbol: "JP¥"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Won Korea Selatan (1945–1953)", Symbol: "KRO"}, "USN": ut.Currency{Currency: "USN", DisplayName: "Dolar AS (Hari berikutnya)", Symbol: "USN"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florin Aruba", Symbol: "AWG"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Mark Konvertibel Bosnia-Herzegovina", Symbol: "BAM"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Cruzeiro Baru Brasil (1967–1986)", Symbol: "BRB"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira Honduras", Symbol: "HNL"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Silver", Symbol: "XAG"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial Yaman", Symbol: "YER"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "Euro WIR", Symbol: "CHE"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Pound Kepulauan Falkland", Symbol: "FKP"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dolar Jamaika", Symbol: "JMD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Pound Lebanon", Symbol: "LBP"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyal Arab Saudi", Symbol: "SAR"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni Swaziland", Symbol: "SZL"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Bolivar Venezuela (1871–2008)", Symbol: "VEB"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Dong Vietnam (1978–1985)", Symbol: "VNN"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Austral Argentina", Symbol: "ARA"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real Brasil", Symbol: "R$"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Rubel Belarusia", Symbol: "BYR"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Denar Makedonia", Symbol: "MKD"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Rand Afrika Selatan (Keuangan)", Symbol: "ZAL"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Peseta Spanyol (akun)", Symbol: "ESA"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial Iran", Symbol: "IRR"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Dolar Selandia Baru", Symbol: "NZ$"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira Nigeria", Symbol: "NGN"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga Tonga", Symbol: "TOP"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Lira Turki (1922–2005)", Symbol: "TRL"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dolar Zimbabwe (1980–2008)", Symbol: "ZWD"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Mvdol Bolivia", Symbol: "BOV"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dolar Kanada", Symbol: "CA$"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colon Kosta Rika", Symbol: "CRC"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical Mozambik", Symbol: "MZN"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guarani Paraguay", Symbol: "PYG"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dolar Trinidad dan Tobago", Symbol: "TTD"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Franc Gold Perancis", Symbol: "XFO"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Dinar Konvertibel Yugoslavia (1990–1992)", Symbol: "YUN"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dolar Belize", Symbol: "BZD"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Sucre Ekuador", Symbol: "ECS"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Franc Komoro", Symbol: "KMF"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa Panama", Symbol: "PAB"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat Myanmar", Symbol: "MMK"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metical Mozambik (1980–2006)", Symbol: "MZM"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu Rumania", Symbol: "RON"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni Tajikistan", Symbol: "TJS"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Socialist Lev Bulgaria", Symbol: "BGM"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Cruzeiro Brasil (1990–1993)", Symbol: "BRE"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Satuan Nilai Tetap Ekuador", Symbol: "ECV"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dolar Kepulauan Cayman", Symbol: "KYD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinar Tunisia", Symbol: "TND"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Zloty Polandia (1950–1995)", Symbol: "PLZ"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Franc CFA BCEAO", Symbol: "CFA"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand Afrika Selatan", Symbol: "ZAR"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Peso Argentina (1983–1985)", Symbol: "ARP"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Pound Israel", Symbol: "ILP"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Talonas Lituania", Symbol: "LTT"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Cordoba Nikaragua", Symbol: "NIO"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Pound Siprus", Symbol: "CYP"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde Haiti", Symbol: "HTG"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Dolar Kepulauan Solomon", Symbol: "SBD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Pound Sudan Selatan", Symbol: "SSP"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Peso Filipina", Symbol: "PHP"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Dolar Rhodesia", Symbol: "RHD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Dinar Serbia", Symbol: "RSD"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Koruna Slovakia", Symbol: "SKK"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Inti Peru", Symbol: "PEI"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone Sierra Leone", Symbol: "SLL"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Peso Uruguay (Unit Diindeks)", Symbol: "UYI"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Pound Inggris", Symbol: "£"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Franc Monegasque", Symbol: "MCF"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Rufiyaa Maladewa (1947–1981)", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupee Nepal", Symbol: "NPR"}, "USS": ut.Currency{Currency: "USS", DisplayName: "Dolar AS (Hari yang sama)", Symbol: "USS"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Tala Samoa", Symbol: "WST"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Satuan Hitung Eropa (XBD)", Symbol: "XBD"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Dinar Yaman", Symbol: "YDD"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Peso Dominika", Symbol: "DOP"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Kroon Estonia", Symbol: "EEK"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Franc Prancis", Symbol: "FRF"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Shilling Somalia", Symbol: "SOS"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupee India", Symbol: "Rs"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Krona Islandia", Symbol: "ISK"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial Qatar", Symbol: "QAR"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat Azerbaijan", Symbol: "AZN"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviano Bolivia (1863–1963)", Symbol: "BOL"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan Tiongkok", Symbol: "CN¥"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Shekel Baru Israel", Symbol: "₪"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Franc Kongo", Symbol: "CDF"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unit Nilai Nyata Kolombia", Symbol: "COU"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaire Zaire (1971–1993)", Symbol: "ZRZ"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Guilder Antilla Belanda", Symbol: "ANG"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Pound Gibraltar", Symbol: "GIP"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupee Sri Lanka", Symbol: "LKR"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Guilder Suriname", Symbol: "SRG"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Peso Guinea-Bissau", Symbol: "GWP"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Shekel Israel", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Franc Malagasi", Symbol: "MGF"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dolar Singapura", Symbol: "SGD"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dolar Amerika Serikat", Symbol: "US$"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Peso Uruguay", Symbol: "UYU"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka Bangladesh", Symbol: "BDT"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum Bhutan", Symbol: "BTN"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti Lesotho", Symbol: "LSL"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Financial Franc Luksemburg", Symbol: "LUL"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Dolar Zimbabwe (2008)", Symbol: "ZWR"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Cruzeiro Brasil (1993–1994)", Symbol: "BRR"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Peso Silver Meksiko (1861–1992)", Symbol: "MXP"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial Oman", Symbol: "OMR"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dolar Suriname", Symbol: "SRD"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Franc Guinea", Symbol: "GNF"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats Latvia", Symbol: "LVL"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupee Mauritius", Symbol: "MUR"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Peso Uruguay (1975–1993)", Symbol: "UYP"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Cruzado Brasil (1986–1989)", Symbol: "BRC"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Denar Makedonia (1992–1993)", Symbol: "MKN"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Pound Suriah", Symbol: "SYP"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat Turkimenistan", Symbol: "TMT"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Mark Jerman Timur", Symbol: "DDM"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Sol Peru (1863–1965)", Symbol: "PES"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupee Seychelles", Symbol: "SCR"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Dolar Zimbabwe (2009)", Symbol: "ZWL"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Koruna Cheska", Symbol: "CZK"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Lira Malta", Symbol: "MTL"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Dong Vietnam", Symbol: "₫"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "Dana RINET", Symbol: "XRE"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Hard Koruna Cheska", Symbol: "CSK"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Pound Sudan", Symbol: "SDG"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Hard Dinar Yugoslavia (1966–1990)", Symbol: "YUD"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Peso Bolivia", Symbol: "BOP"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Peso Kuba", Symbol: "CUP"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik Mongolia", Symbol: "MNT"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Krona Swedia", Symbol: "SEK"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dolar Karibia Timur", Symbol: "EC$"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Dinar Reformasi Yugoslavia (1992–1993)", Symbol: "YUR"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dolar Bermuda", Symbol: "BMD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Peso Kolombia", Symbol: "COP"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Cupon Moldova", Symbol: "MDC"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Satuan Hitung Eropa (XBC)", Symbol: "XBC"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek Albania", Symbol: "ALL"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Franc Belgia (keuangan)", Symbol: "BEL"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dolar Guyana", Symbol: "GYD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Rubel Soviet", Symbol: "SUR"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Pound Sudan (1957–1998)", Symbol: "SDP"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolivar Venezuela", Symbol: "VEF"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula Botswana", Symbol: "BWP"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Dinar Kroasia", Symbol: "HRD"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupee Pakistan", Symbol: "PKR"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Franc Rwanda", Symbol: "RWF"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Dinar Baru Bosnia-Herzegovina (1994–1997)", Symbol: "BAN"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Markka Finlandia", Symbol: "FIM"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu Moldova", Symbol: "MDL"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Polandia Zloty", Symbol: "PLN"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Franc Mali", Symbol: "MLF"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Dolar Baru Taiwan", Symbol: "NT$"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Hak Khusus Menggambar", Symbol: "XDR"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha Zambia", Symbol: "ZMW"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Dinar Bosnia-Herzegovina (1992–1994)", Symbol: "BAD"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Peso Cile", Symbol: "CLP"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Pound Mesir", Symbol: "EGP"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariary Madagaskar", Symbol: "MGA"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Rubel Latvia", Symbol: "LVR"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Nuevo Sol Peru", Symbol: "PEN"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afgani Afganistan (1927–2002)", Symbol: "AFA"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Franc Belgia", Symbol: "BEF"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal Guatemala", Symbol: "GTQ"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Lira Italia", Symbol: "ITL"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirham Uni Emirat Arab", Symbol: "AED"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afgani Afganistan", Symbol: "AFN"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Peso Argentina (1881–1970)", Symbol: "ARM"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Peseta Spanyol (konvertibel)", Symbol: "ESB"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won Korea Utara", Symbol: "KPW"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dolar Namibia", Symbol: "NAD"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Shilling Uganda", Symbol: "UGX"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Unit Gabungan Eropa", Symbol: "XBA"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "Franc WIR", Symbol: "CHW"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa Eritrea", Symbol: "ERN"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Franc Burundi", Symbol: "BIF"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Shilling Kenya", Symbol: "KES"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won Korea Selatan", Symbol: "₩"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya Mauritania", Symbol: "MRO"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Guilder Belanda", Symbol: "NLG"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Peseta Andorra", Symbol: "ADP"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Kwanza Baru Angola (1990–2000)", Symbol: "AON"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Kwanza Angola yang Disesuaikan Lagi (1995–1999)", Symbol: "AOR"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Franc Belgia (konvertibel)", Symbol: "BEC"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Krone Denmark", Symbol: "DKK"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Peseta Spanyol", Symbol: "ESP"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas Lituania", Symbol: "LTL"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Colon El Savador", Symbol: "SVC"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Manat Azerbaijan (1993–2006)", Symbol: "AZM"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Franc Konvertibel Luksemburg", Symbol: "LUC"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Peso Meksiko", Symbol: "MX$"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Peso Konvertibel Kuba", Symbol: "CUC"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip Laos", Symbol: "LAK"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som Uzbekistan", Symbol: "UZS"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaire Baru Zaire (1993–1998)", Symbol: "ZRN"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Escudo Guinea Portugal", Symbol: "GWE"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platinum", Symbol: "XPT"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Kode Mata Uang Pengujian", Symbol: "XTS"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Mata Uang Tidak Dikenal", Symbol: "XXX"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Peso Argentina", Symbol: "ARS"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Rubel Baru Belarus (1994–1999)", Symbol: "BYB"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Escudo Tanjung Verde", Symbol: "CVE"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari Georgia", Symbol: "GEL"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Escudo Cile", Symbol: "CLE"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Franc Jibuti", Symbol: "DJF"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Lira Turki", Symbol: "TRY"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Escudo Timor", Symbol: "TPE"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Schilling Austria", Symbol: "ATS"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Hard Lev Bulgaria", Symbol: "BGL"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinar Algeria", Symbol: "DZD"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Ekuele Guinea Ekuatorial", Symbol: "GQE"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Pound Malta", Symbol: "MTP"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Peso Ley Argentina (1970–1983)", Symbol: "ARL"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dolar Brunei", Symbol: "BND"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi Gambia", Symbol: "GMD"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar Yordania", Symbol: "JOD"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu Vanuatu", Symbol: "VUV"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Emas", Symbol: "XAU"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza Angola", Symbol: "AOA"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dolar Liberia", Symbol: "LRD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rufiyaa Maladewa", Symbol: "MVR"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rubel Rusia", Symbol: "RUB"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Franc CFA BEAC", Symbol: "FCFA"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Franc Swiss", Symbol: "CHF"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi Ghana", Symbol: "GHS"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Pound Irlandia", Symbol: "IEP"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar Irak", Symbol: "IQD"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Cruzado Baru Brasil (1989–1990)", Symbol: "BRN"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Cordoba Nikaragua (1988–1991)", Symbol: "NIC"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Krone Norwegia", Symbol: "NOK"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano", Symbol: "BOB"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Ringgit Malaysia", Symbol: "MYR"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Franc Maroko", Symbol: "MAF"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht Thailand", Symbol: "฿"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev Bulgaria", Symbol: "BGN"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinar Libya", Symbol: "LYD"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Kyat Burma", Symbol: "BUK"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}}
