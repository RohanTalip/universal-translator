package sbp

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Mup", 3: "Msh", 4: "Mun", 5: "Mag", 6: "Muj", 9: "Mye", 11: "Mus", 12: "Muh", 2: "Mwi", 7: "Msp", 8: "Mpg", 10: "Mok"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{12: "Muhaano", 1: "Mupalangulwa", 3: "Mushende", 5: "Mushende Magali", 6: "Mujimbi", 7: "Mushipepo", 8: "Mupuguto", 10: "Mokhu", 2: "Mwitope", 4: "Munyi", 9: "Munyense", 11: "Musongandembwe"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Jtt", 2: "Jnn", 3: "Jtn", 4: "Alh", 5: "Iju", 6: "Jmo", 0: "Mul"}, Narrow: ut.CalendarDayFormatNameValue{2: "J", 3: "J", 4: "A", 5: "I", 6: "J", 0: "M", 1: "J"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{3: "Jumatano", 4: "Alahamisi", 5: "Ijumaa", 6: "Jumamosi", 0: "Mulungu", 1: "Jumatatu", 2: "Jumanne"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "Lwamilawu", "pm": "Pashamihe"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Lwamilawu", "pm": "Pashamihe"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Ashanali uKilisito", Abbrev: "AK", Narrow: ""}}}}
