package wae

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: ""}, AD: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: ""}}, Time: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "Brá", 7: "Hei", 11: "Win", 12: "Chr", 1: "Jen", 3: "Mär", 5: "Mei", 9: "Her", 10: "Wím", 2: "Hor", 4: "Abr", 8: "Öig"}, Narrow: ut.CalendarMonthFormatNameValue{9: "H", 10: "W", 1: "J", 2: "H", 3: "M", 4: "A", 6: "B", 7: "H", 11: "W", 12: "C", 5: "M", 8: "Ö"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{6: "Bráčet", 7: "Heiwet", 10: "Wímánet", 11: "Wintermánet", 1: "Jenner", 2: "Hornig", 4: "Abrille", 9: "Herbštmánet", 12: "Chrištmánet", 3: "Märze", 5: "Meije", 8: "Öigšte"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Män", 2: "Ziš", 3: "Mit", 4: "Fró", 5: "Fri", 6: "Sam", 0: "Sun"}, Narrow: ut.CalendarDayFormatNameValue{5: "F", 6: "S", 0: "S", 1: "M", 2: "Z", 3: "M", 4: "F"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{2: "Zištag", 3: "Mittwuč", 4: "Fróntag", 5: "Fritag", 6: "Samštag", 0: "Sunntag", 1: "Mäntag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "v. Chr.", Narrow: ""}}}}
