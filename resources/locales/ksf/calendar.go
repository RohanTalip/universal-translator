package ksf

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "ŋ4", 6: "ŋ6", 9: "ŋ9", 12: "ŋ12", 3: "ŋ3", 2: "ŋ2", 5: "ŋ5", 7: "ŋ7", 8: "ŋ8", 10: "ŋ10", 11: "ŋ11", 1: "ŋ1"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "ŋwíí akǝ ntɛk", 12: "ŋwíí akǝ ntɛk di bɛ́ɛ", 4: "ŋwíí akǝ nin", 5: "ŋwíí akǝ táan", 6: "ŋwíí akǝ táafɔk", 8: "ŋwíí akǝ táaraa", 9: "ŋwíí akǝ táanin", 1: "ŋwíí a ntɔ́ntɔ", 2: "ŋwíí akǝ bɛ́ɛ", 3: "ŋwíí akǝ ráá", 7: "ŋwíí akǝ táabɛɛ", 11: "ŋwíí akǝ ntɛk di bɔ́k"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "sam", 0: "sɔ́n", 1: "lǝn", 2: "maa", 3: "mɛk", 4: "jǝǝ", 5: "júm"}, Narrow: ut.CalendarDayFormatNameValue{3: "m", 4: "j", 5: "j", 6: "s", 0: "s", 1: "l", 2: "m"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "lǝndí", 2: "maadí", 3: "mɛkrɛdí", 4: "jǝǝdí", 5: "júmbá", 6: "samdí", 0: "sɔ́ndǝ"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "sárúwá", "pm": "cɛɛ́nko"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "sárúwá", "pm": "cɛɛ́nko"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "di Yɛ́sus aká yálɛ", Abbrev: "d.Y.", Narrow: ""}}}}
