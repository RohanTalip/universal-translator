package dsb

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.yy"}}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "maj", 7: "jul", 8: "awg", 9: "sep", 10: "okt", 11: "now", 2: "feb", 3: "měr", 4: "apr", 6: "jun", 12: "dec", 1: "jan"}, Narrow: ut.CalendarMonthFormatNameValue{6: "j", 11: "n", 12: "d", 1: "j", 4: "a", 5: "m", 8: "a", 9: "s", 10: "o", 2: "f", 3: "m", 7: "j"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "januar", 2: "februar", 3: "měrc", 6: "junij", 7: "julij", 9: "september", 4: "apryl", 5: "maj", 8: "awgust", 10: "oktober", 11: "nowember", 12: "december"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "nje", 1: "pón", 2: "wał", 3: "srj", 4: "stw", 5: "pět", 6: "sob"}, Narrow: ut.CalendarDayFormatNameValue{6: "s", 0: "n", 1: "p", 2: "w", 3: "s", 4: "s", 5: "p"}, Short: ut.CalendarDayFormatNameValue{3: "sr", 4: "st", 5: "pě", 6: "so", 0: "nj", 1: "pó", 2: "wa"}, Wide: ut.CalendarDayFormatNameValue{6: "sobota", 0: "njeźela", 1: "pónjeźele", 2: "wałtora", 3: "srjoda", 4: "stwórtk", 5: "pětk"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "dopołdnja", "pm": "wótpołdnja"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "dop.", "pm": "wótp."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "dopołdnja", "pm": "wótpołdnja"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "pó Kristusowem naroźenju", Abbrev: "pó Chr.n.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "pśed Kristusowym naroźenim", Abbrev: "pś.Chr.n.", Narrow: ""}}}}
