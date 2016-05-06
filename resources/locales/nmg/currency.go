package nmg

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"SHP": ut.Currency{Currency: "SHP", DisplayName: "Mɔn má Saint Lina", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Fraŋ Guiné", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dɔ́llɔ Liberia", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Mɔn Marɔk", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Mɔn bó Chinois", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Mɔn Malawi", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Mɔn Afrik yí sí", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Mɔn Angola", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Mɔn Gambia", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Mɔn B ´Arabe", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Fraŋ Suisse", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Mɔn Kapvɛrt", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Fraŋ bó Kɔmɔr", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Mɔn Leɔne", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Mɔn Ligangeni", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Fraŋ bó Kongolɛ̌", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Mɔn Algeria", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Mɔn India", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Mɔn Sudan (1957–1998)", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Mɔn Sao tomé na prinship", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Mɔn Tunisia", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Mɔn Zambia", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Fraŋ Burundi", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dɔ́llɔ Kanada", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Fraŋ Jibuti", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Mɔn Madagaskar", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Mɔn Ngɛ̄lɛ̄n", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naïra Nigeria", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mɔn Moriss", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Fraŋ CFA BCEAO", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Mɔn Botswana", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Mɔn Saudi Arabia", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Mɔn Seychɛlle", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Mɔn Somalía", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dɔ́llɔ Ɔstralia", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Mɔn Libya", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Fraŋ Rwanda", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Mɔn Sudan", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Mɔn Ethiopia", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Mɔn Gana", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dɔ́llɔ Namibia", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dɔ́llɔ Amɛŕka", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Mɔn Lesoto", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Mɔn Tanzania", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dɔ́llɔ Zimbabwǝ (1980–2008)", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Fraŋ CFA BEAC", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Mɔn Ägyptɛn", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Mɔn Erytré", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Mɔn Japɔn", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Mɔn Kɛnya", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Mɔn Zambia (1968–2012)", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Mɔn Bahrein", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mɔn Moritania", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mɔn Mozambik", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Mɔn Uganda", Symbol: ""}}
