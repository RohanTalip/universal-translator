package prg

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, y 'mettas' d. MMMM", Long: "y 'mettas' d. MMMM", Medium: "dd.MM 'st'. y", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, y 'mettas' d. MMMM", Long: "y 'mettas' d. MMMM", Medium: "dd.MM 'st'. y", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "sil", 11: "lap", 3: "pūl", 6: "sīm", 8: "dag", 5: "zal", 7: "līp", 10: "spa", 12: "sal", 1: "rag", 2: "was", 4: "sak"}, Narrow: ut.CalendarMonthFormatNameValue{3: "P", 4: "S", 8: "D", 9: "S", 10: "S", 11: "L", 12: "S", 1: "R", 5: "Z", 6: "S", 7: "L", 2: "W"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "spallins", 11: "lapkrūtis", 12: "sallaws", 3: "pūlis", 7: "līpa", 8: "daggis", 9: "sillins", 6: "sīmenis", 1: "rags", 2: "wassarins", 4: "sakkis", 5: "zallaws"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "wis", 3: "pus", 4: "ket", 5: "pēn", 6: "sab", 0: "nad", 1: "pan"}, Narrow: ut.CalendarDayFormatNameValue{3: "P", 4: "K", 5: "P", 6: "S", 0: "N", 1: "P", 2: "W"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "panadīli", 2: "wisasīdis", 3: "pussisawaiti", 4: "ketwirtiks", 5: "pēntniks", 6: "sabattika", 0: "nadīli"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "ankstāinan", "pm": "pa pussideinan"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "BC", Narrow: ""}}}}
