package ar

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MZE": ut.Currency{Currency: "MZE", DisplayName: "اسكود موزمبيقي", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "سول جديد البيرو", Symbol: "PEN"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "أفغاني - 1927-2002", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "اسكودو الرأس الخضراء", Symbol: "CVE"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "جنيه إيرلندي", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "ليرة إيطالية", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "فرنك لوكسمبرج", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "فرنك سي إف بي", Symbol: "CFPF"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "شلن نمساوي", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "كرونة استونية", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "فرنك جزر القمر", Symbol: "ف.ج.ق.\u200f"}, "RON": ut.Currency{Currency: "RON", DisplayName: "ليو روماني", Symbol: "RON"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "فرنك وسط أفريقي", Symbol: "FCFA"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "البلاتين", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "فرنك بلجيكي", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "ناكفا أريتري", Symbol: "ERN"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "بيزيتا إسباني", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "اسكود برتغالي غينيا", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "هريفنيا أوكراني", Symbol: "UAH"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "دولار برمودي", Symbol: "BMD"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "كرونه أيسلندي", Symbol: "ISK"}, "STD": ut.Currency{Currency: "STD", DisplayName: "دوبرا ساو تومي وبرينسيبي", Symbol: "STD"}, "USS": ut.Currency{Currency: "USS", DisplayName: "دولار أمريكي (نفس اليوم)\u200f", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "دينار صربي", Symbol: "RSD"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "فرنك جيبوتي", Symbol: "DJF"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "ين ياباني", Symbol: "JP¥"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "شلن أوغندي", Symbol: "UGX"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "زائير زائيري", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "ليك ألباني", Symbol: "ALL"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "كونا كرواتي", Symbol: "HRK"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "دينار عراقي", Symbol: "د.ع.\u200f"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "تينغ كازاخستاني", Symbol: "KZT"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "جنيه مالطي", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "كوانزا أنجولي - 1977-1990", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "يورو", Symbol: "€"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "جنيه لبناني", Symbol: "ل.ل.\u200f"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "سوم أوزبكستاني", Symbol: "UZS"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "تولار سلوفيني", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "روبل طاجيكستاني", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "بوليفار فنزويلي - 1871-2008", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "بيزو كوبي", Symbol: "CUP"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "كرونة نرويجية", Symbol: "NOK"}, "YER": ut.Currency{Currency: "YER", DisplayName: "ريال يمني", Symbol: "ر.ي.\u200f"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "غيلدر أنتيلي هولندي", Symbol: "ANG"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "فرنك بروندي", Symbol: "BIF"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "جنيه سانت هيلين", Symbol: "SHP"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "دينار بحريني", Symbol: "د.ب.\u200f"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "جنيه إسترليني", Symbol: "£"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "روبية سريلانكية", Symbol: "LKR"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "اسكود تيموري", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "فضة", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "كروزادو برازيلي", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "دولار هونغ كونغ", Symbol: "HK$"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "روبل روسي - 1991-1998", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "ليرة تركي", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "دولار بليزي", Symbol: "BZD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "دولار فيجي", Symbol: "FJD"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "ليرة مالطية", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "ليو روماني قديم", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "دولار زمبابوي", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "فرنك بلجيكي مالي", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "بوليفيانو بوليفي", Symbol: "BOB"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "كواشا زامبي", Symbol: "ZMW"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "نوفو كروزايرو برازيلي - 1967-1986", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "ليتا ليتوانية", Symbol: "LTL"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "بيزو فضي مكسيكي - 1861-1992", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "دولار أمريكي (اليوم التالي)\u200f", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "ذهب", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "(XBD)وحدة الحساب الأوروبية", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "راند جنوب أفريقيا -مالي", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "دولار بربادوسي", Symbol: "BBD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "كرونة تشيكية", Symbol: "CZK"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "لارى جورجي", Symbol: "GEL"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "كوتزال جواتيمالا", Symbol: "GTQ"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "دولار ترينداد وتوباجو", Symbol: "TTD"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "ماركا فنلندي", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "دينار تونسي", Symbol: "د.ت.\u200f"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "سوم قيرغستاني", Symbol: "KGS"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "فرنك فرنسي ذهبي", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "راند جنوب أفريقيا", Symbol: "ZAR"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "أوستمارك ألماني شرقي", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "بير أثيوبي", Symbol: "ETB"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "جنيه جزر فوكلاند", Symbol: "FKP"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "سيدي غانا", Symbol: "GHS"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "(UIC)فرنك فرنسي", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "بيزو أرجنتيني", Symbol: "ARS"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "نايرا نيجيري", Symbol: "NGN"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "أرياري مدغشقر", Symbol: "MGA"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "جلدر هولندي", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "روبية باكستاني", Symbol: "ر.ب."}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "كوانزا أنجولي", Symbol: "AOA"}, "AON": ut.Currency{Currency: "AON", DisplayName: "كوانزا أنجولي جديدة - 1990-2000", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "كينا بابوا غينيا الجديدة", Symbol: "PGK"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "جلدر سورينامي", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "دولار جزر كيمن", Symbol: "KYD"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "دينار ليبي", Symbol: "د.ل.\u200f"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "فرنك فرنسي", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "دينار مقدوني", Symbol: "MKD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "دولار سنغافوري", Symbol: "SGD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "روبل سوفيتي", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "تالا ساموا", Symbol: "WST"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "باتاكا ماكاوي", Symbol: "MOP"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "جنيه جنوب السودان", Symbol: "ج.ج.س."}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "ليرة سورية", Symbol: "ل.س.\u200f"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "مانات تركمنستاني", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "جنيه قبرصي", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "روبية موريشيوسية", Symbol: "MUR"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "رينغيت ماليزي", Symbol: "MYR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "جواراني باراجواي", Symbol: "PYG"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "شلن صومالي", Symbol: "SOS"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "استرال أرجنتيني", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "كولن كوستا ريكي", Symbol: "CRC"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "وون كوريا الشمالية", Symbol: "KPW"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "روبية نيبالي", Symbol: "NPR"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "اسكود برتغالي", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "أوقية موريتانية", Symbol: "أ.م.\u200f"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "دينار يوغسلافي", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "فرنك بلجيكي قابل للتحويل", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "بيزو بوليفي", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "كيات بورمي", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "تالوناس ليتواني", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "فرنك لوكسمبرج قابل للتحويل", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "بيزو فلبيني", Symbol: "PHP"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "شلن أوغندي - 1966-1987", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "بيزو أوروجواي - 1975-1993", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "كوانزا أنجولي معدلة - 1995 - 1999", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "روبيه هندي", Symbol: "₹"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "دولار جامايكي", Symbol: "JMD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "روفيه جزر المالديف", Symbol: "MVR"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "بيزو مكسيكي", Symbol: "MX$"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "روبل بيلاروسي", Symbol: "BYR"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "فاتو فانواتو", Symbol: "VUV"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "دينار يمني", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "متكال موزمبيقي", Symbol: "MZN"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "الوحدة الأوروبية المركبة", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "تاكا بنجلاديشي", Symbol: "BDT"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "فرنك غينيا", Symbol: "GNF"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "حقوق السحب الخاصة", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "دينار يوغسلافي قابل للتحويل", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "بيزو أرجنتيني - 1983-1985", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "بيزو الدومنيكان", Symbol: "DOP"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "الوحدة الحسابية الأوروبية", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "دولار شرق الكاريبي", Symbol: "EC$"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "فرنك مدغشقر", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "ريال قطري", Symbol: "ر.ق.\u200f"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "دولار زمبابوي 2009", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "بيزو شيلي", Symbol: "CLP"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "دينار جزائري", Symbol: "د.ج.\u200f"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "جنيه مصري", Symbol: "ج.م.\u200f"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "دينار كويتي", Symbol: "د.ك.\u200f"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "دولار ليبيري", Symbol: "LRD"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "دينار سوداني", Symbol: "د.س.\u200f"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "بانغا تونغا", Symbol: "TOP"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "روبية سيشيلية", Symbol: "SCR"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "جنيه سوداني قديم", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "ليرة تركية", Symbol: "ل.ت."}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "دولار أسترالي", Symbol: "AU$"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "فلورن أروبي", Symbol: "AWG"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "ليمبيرا هنداروس", Symbol: "HNL"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "درهم مغربي", Symbol: "د.م.\u200f"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "زلوتي بولندي - 1950-1995", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "الوحدة المالية الأوروبية", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "درهم إماراتي", Symbol: "د.إ.\u200f"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "مارك البوسنة والهرسك قابل للتحويل", Symbol: "BAM"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "نولتوم بوتاني", Symbol: "BTN"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "درام أرميني", Symbol: "AMD"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "دينار كرواتي", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "وون كوريا الجنوبية", Symbol: "₩"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "مانات أذربيجان", Symbol: "AZN"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "ريال إيراني", Symbol: "ر.إ."}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "فرنك مالي", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "روبل بيلاروسي جديد - 1994-1999", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "فرنك كونغولي", Symbol: "CDF"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "ليو مولدوفي", Symbol: "MDL"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "دينار أردني", Symbol: "د.أ.\u200f"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "كواشا زامبي - 1968-2012", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "بيستا أندوري", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "بيزو كولومبي", Symbol: "COP"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "جنيه جبل طارق", Symbol: "GIP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "دلاسي جامبي", Symbol: "GMD"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "روبية إندونيسية", Symbol: "ر.إن."}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "جوردى هايتي", Symbol: "HTG"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "فورينت مجري", Symbol: "HUF"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "دولار روديسي", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "دولار أمريكي", Symbol: "US$"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "ليف بلغاري", Symbol: "BGN"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "فرنك لوكسمبرج المالي", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "كيات ميانمار", Symbol: "MMK"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "دولار كندي", Symbol: "CA$"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "مانات تركمانستان", Symbol: "TMT"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "زلوتي بولندي", Symbol: "PLN"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "بولا بتسواني", Symbol: "BWP"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "كرونة تشيكوسلوفاكيا", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "فرنك غرب أفريقي", Symbol: "CFA"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "بالبوا بنمي", Symbol: "PAB"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "مانات أذريبجاني", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "دولار باهامي", Symbol: "BSD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "فرنك سويسري", Symbol: "CHF"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "سيلي غينيا", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "دولار غيانا", Symbol: "GYD"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "ريال برازيلي", Symbol: "R$"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "ليون سيراليوني", Symbol: "SLL"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "كروزايرو برازيلي - 1990-1993", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "كود اختبار العملة", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "بوليفار فنزويلي", Symbol: "VEF"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "بالاديوم", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "مارك ألماني", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "شيكل إسرائيلي جديد", Symbol: "₪"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "توغروغ منغولي", Symbol: "MNT"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "دولار ناميبي", Symbol: "NAD"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "دولار سورينامي", Symbol: "SRD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "بيزو كوبي قابل للتحويل", Symbol: "CUC"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "جنيه إسرائيلي", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "كيب لاوسي", Symbol: "LAK"}, "THB": ut.Currency{Currency: "THB", DisplayName: "باخت تايلاندي", Symbol: "฿"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "شلن تنزاني", Symbol: "TZS"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "عملة غير معروفة", Symbol: "***"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "دينار البوسنة والهرسك", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "كرونة دانماركي", Symbol: "DKK"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "فرنك مغربي", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "روبل روسي", Symbol: "RUB"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "دولار جزر سليمان", Symbol: "SBD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "دونج فيتنامي", Symbol: "₫"}, "BND": ut.Currency{Currency: "BND", DisplayName: "دولار بروناي", Symbol: "BND"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "مفدول بوليفي", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "بيزو غينيا بيساو", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "رييال كمبودي", Symbol: "KHR"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "روبل لاتفيا", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "يوان صيني", Symbol: "CN¥"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "كولون سلفادوري", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "ليلانجيني سوازيلندي", Symbol: "SZL"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "أفغاني", Symbol: "AFN"}, "KES": ut.Currency{Currency: "KES", DisplayName: "شلن كينيي", Symbol: "KES"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "دولار نيوزيلندي", Symbol: "NZ$"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "جنيه سوداني", Symbol: "ج.س."}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "بيزو اوروغواي", Symbol: "UYU"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "سيدي غاني", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "دراخما يوناني", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "لاتس لاتفيا", Symbol: "LVL"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "دولار تايواني", Symbol: "NT$"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "وحدة النقد الأوروبية", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "اكويل جونينا غينيا الاستوائيّة", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "لوتي ليسوتو", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "كواشا مالاوي", Symbol: "MWK"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "كوردوبة نيكاراجوا", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "ريال سعودي", Symbol: "ر.س.\u200f"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "دينار صربي قديم", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ريال عماني", Symbol: "ر.ع.\u200f"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "سوموني طاجيكستاني", Symbol: "TJS"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "زائير زائيري جديد", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "كرونة سويدية", Symbol: "SEK"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "قرطبة نيكاراغوا", Symbol: "NIO"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "فرنك رواندي", Symbol: "RWF"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "كرونة سلوفاكية", Symbol: ""}}
