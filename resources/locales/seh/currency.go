package seh

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan Renminbi chinês", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya da Mauritânia", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Franco ruandês", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra de São Tomé e Príncipe", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirém dos Emirados Árabes Unidos", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Xelim queniano", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dólar australiano", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone de Serra Leoa", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Iene japonês", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Cuacha do Maláui", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metical antigo de Moçambique", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Xelim somali", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Cuacha zambiano (1968–2012)", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical de Moçambique", Symbol: "MTn"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Libra de Santa Helena", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dólar norte-americano", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand sul-africano", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi de Gana (1979–2007)", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dólar liberiano", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirém marroquino", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Xelim da Tanzânia", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dólar do Zimbábue", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula botsuanesa", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Franco do Djibuti", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi de Gâmbia", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rúpia indiana", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia de Maurício", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Cuanza angolano", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dólar canadense", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa da Eritréia", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia das Seychelles", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Libra britânica", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dólar da Namíbia", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Libra sudanesa antiga", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Franco CFA BCEAO", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Franco do Burundi", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira nigeriana", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Dinar sudanês", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Franco congolês", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Syli da Guiné", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinar líbio", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni da Suazilândia", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Franco CFA BEAC", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar bareinita", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Escudo cabo-verdiano", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti do Lesoto", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Xelim ugandense (1966–1987)", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Cuacha zambiano", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr etíope", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Rial saudita", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Franco de Madagascar", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinar argelino", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Libra egípcia", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Franco de Comores", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinar tunisiano", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Franco suíço", Symbol: ""}}
