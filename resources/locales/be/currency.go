package be

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"AWG": ut.Currency{Currency: "AWG", DisplayName: "арубанскі фларын", Symbol: "AWG"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "кувейцкі дынар", Symbol: "KWD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "шведская крона", Symbol: "SEK"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "сейшэльская рупія", Symbol: "SCR"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "суданскі фунт", Symbol: "SDG"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "сірыйскі фунт", Symbol: "SYP"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "аргенцінскае песа", Symbol: "ARS"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "мексіканскае песа", Symbol: "MX$"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "кубінскае песа", Symbol: "CUP"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "егіпецкі фунт", Symbol: "EGP"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "маўрыкійская рупія", Symbol: "MUR"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "сінгапурскі долар", Symbol: "SGD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "туніскі дынар", Symbol: "TND"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "балгарскі леў", Symbol: "BGN"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "кангалезскі франк", Symbol: "CDF"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "малайзійскі рынгіт", Symbol: "MYR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "аманскі рыал", Symbol: "OMR"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "долар Саламонавых Астравоў", Symbol: "SBD"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "сурынамскі долар", Symbol: "SRD"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "свазілендскі лілангені", Symbol: "SZL"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "цэнтральнаафрыканскі франк КФА", Symbol: "FCFA"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "ангольская кванза", Symbol: "AOA"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "новы ізраільскі шэкель", Symbol: "₪"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "паўднёваафрыканскі ранд", Symbol: "ZAR"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "іракскі дынар", Symbol: "IQD"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "эрытрэйская накфа", Symbol: "ERN"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "брытанскі фунт стэрлінгаў", Symbol: "£"}, "WST": ut.Currency{Currency: "WST", DisplayName: "самаанская тала", Symbol: "WST"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "беларускі рубель", Symbol: "р."}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "вату", Symbol: "VUV"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "кіна", Symbol: "PGK"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "пакістанская рупія", Symbol: "PKR"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "афганскі афгані", Symbol: "AFN"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "нарвежская крона", Symbol: "NOK"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "кубінскае канверсоўнае песа", Symbol: "CUC"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "лівійскі дынар", Symbol: "LYD"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "мараканскі дырхам", Symbol: "MAD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "малавійская квача", Symbol: "MWK"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "непальская рупія", Symbol: "NPR"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "расійскі рубель", Symbol: "₽"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "бермудскі долар", Symbol: "BMD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "калумбійскае песа", Symbol: "COP"}, "THB": ut.Currency{Currency: "THB", DisplayName: "тайскі бат", Symbol: "THB"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "новы тайваньскі долар", Symbol: "NT$"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "перуанскі новы соль", Symbol: "PEN"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "сербскі дынар", Symbol: "RSD"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "замбійская квача", Symbol: "ZMW"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "албанскі лек", Symbol: "ALL"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "камбаджыйскі рыэль", Symbol: "KHR"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "белізскі долар", Symbol: "BZD"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "філіпінскае песа", Symbol: "PHP"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "японская іена", Symbol: "¥"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "мальдыўская руфія", Symbol: "MVR"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "нігерыйская найра", Symbol: "NGN"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "бангладэшская така", Symbol: "BDT"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "багамскі долар", Symbol: "BSD"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "польскі злоты", Symbol: "PLN"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "лаоскі кіп", Symbol: "LAK"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "новазеландскі долар", Symbol: "NZD"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "гібралтарскі фунт", Symbol: "GIP"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "каморскі франк", Symbol: "KMF"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "бутанскі нгултрум", Symbol: "BTN"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "эскуда Каба-Вердэ", Symbol: "CVE"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "ямайскі долар", Symbol: "JMD"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "мазамбікскі метыкал", Symbol: "MZN"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "намібійскі долар", Symbol: "NAD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "в’етнамскі донг", Symbol: "₫"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "французскі ціхаакіянскі франк", Symbol: "CFPF"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "чылійскае песа", Symbol: "CLP"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ганскі седзі", Symbol: "GHS"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "паўднёвасуданскі фунт", Symbol: "SSP"}, "BND": ut.Currency{Currency: "BND", DisplayName: "брунейскі долар", Symbol: "BND"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "малдаўскі лей", Symbol: "MDL"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "харвацкая куна", Symbol: "HRK"}, "KES": ut.Currency{Currency: "KES", DisplayName: "кенійскі шылінг", Symbol: "KES"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "ліберыйскі долар", Symbol: "LRD"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "заходнеафрыканскі франк КФА", Symbol: "CFA"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "канверсоўная марка", Symbol: "BAM"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "бразільскі рэал", Symbol: "BRL"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "венгерскі форынт", Symbol: "HUF"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "м’янманскі к’ят", Symbol: "MMK"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "руандыйскі франк", Symbol: "RWF"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "саудаўскі рыял", Symbol: "SAR"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "барбадоскі долар", Symbol: "BBD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "швейцарскі франк", Symbol: "CHF"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "маўрытанская ўгія", Symbol: "MRO"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "панамскае бальбоа", Symbol: "PAB"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "леонэ", Symbol: "SLL"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "таджыкскі самані", Symbol: "TJS"}, "USD": ut.Currency{Currency: "USD", DisplayName: "долар ЗША", Symbol: "$"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "аўстралійскі долар", Symbol: "A$"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "паўночнакарэйская вона", Symbol: "KPW"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "еўра", Symbol: "€"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "катарскі рыал", Symbol: "QAR"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "чэшская крона", Symbol: "CZK"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "дацкая крона", Symbol: "DKK"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "іарданскі дынар", Symbol: "JOD"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "венесуальскі балівар", Symbol: "VEF"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "бурундзійскі франк", Symbol: "BIF"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "гаянскі долар", Symbol: "GYD"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "фунт Фалклендскіх астравоў", Symbol: "FKP"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "грузінскі лары", Symbol: "GEL"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "македонскі дэнар", Symbol: "MKD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "мангольскі тугрык", Symbol: "MNT"}, "YER": ut.Currency{Currency: "YER", DisplayName: "еменскі рыал", Symbol: "YER"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "армянскі драм", Symbol: "AMD"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "коста-рыканскі калон", Symbol: "CRC"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "долар Кайманавых астравоў", Symbol: "KYD"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "турэцкая ліра", Symbol: "TRY"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "уругвайскае песа", Symbol: "UYU"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "гамбійскі даласі", Symbol: "GMD"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "паўднёвакарэйская вона", Symbol: "₩"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "нікарагуанская кордаба", Symbol: "NIO"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "фунт Святой Алены", Symbol: "SHP"}, "STD": ut.Currency{Currency: "STD", DisplayName: "добра Сан-Тамэ і Прынсіпі", Symbol: "STD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "батсванская пула", Symbol: "BWP"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "кітайскі юань", Symbol: "CN¥"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "гвінейскі франк", Symbol: "GNF"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "інданезійская рупія", Symbol: "IDR"}, "INR": ut.Currency{Currency: "INR", DisplayName: "індыйская рупія", Symbol: "₹"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "іранскі рыал", Symbol: "IRR"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "кіргізскі сом", Symbol: "KGS"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "украінская грыўна", Symbol: "UAH"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "нідэрландскі антыльскі гульдэн", Symbol: "ANG"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "дамініканскае песа", Symbol: "DOP"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "шры-ланкійская рупія", Symbol: "LKR"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "самалійскі шылінг", Symbol: "SOS"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "долар Трынідада і Табага", Symbol: "TTD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "фіджыйскі долар", Symbol: "FJD"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "ганконгскі долар", Symbol: "HK$"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "гаіцянскі гурд", Symbol: "HTG"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "ліванскі фунт", Symbol: "LBP"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "малагасійскі арыяры", Symbol: "MGA"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "туркменскі манат", Symbol: "TMT"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "невядомая валюта", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "дырхем ААЭ", Symbol: "AED"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "джыбуційскі франк", Symbol: "DJF"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "казахстанскі тэнгэ", Symbol: "KZT"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "парагвайскі гуарані", Symbol: "PYG"}, "RON": ut.Currency{Currency: "RON", DisplayName: "румынскі лей", Symbol: "RON"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "танганская паанга", Symbol: "TOP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "угандыйскі шылінг", Symbol: "UGX"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "узбекскі сум", Symbol: "UZS"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "канадскі долар", Symbol: "CAD"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "эфіопскі быр", Symbol: "ETB"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "усходнекарыбскі долар", Symbol: "EC$"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "гандураская лемпіра", Symbol: "HNL"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "ісландская крона", Symbol: "ISK"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "гватэмальскі кетсаль", Symbol: "GTQ"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "азербайджанскі манат", Symbol: "AZN"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "бахрэйнскі дынар", Symbol: "BHD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "патака Макаа", Symbol: "MOP"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "танзанійскі шылінг", Symbol: "TZS"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "балівіяна", Symbol: "BOB"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "алжырскі дынар", Symbol: "DZD"}}
