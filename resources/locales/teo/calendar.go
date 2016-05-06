package teo

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "Kwa", 4: "Dun", 7: "Jol", 8: "Ped", 9: "Sok", 1: "Rar", 2: "Muk", 5: "Mar", 6: "Mod", 10: "Tib", 11: "Lab", 12: "Poo"}, Narrow: ut.CalendarMonthFormatNameValue{4: "D", 5: "M", 7: "J", 8: "P", 10: "T", 11: "L", 3: "K", 2: "M", 6: "M", 9: "S", 12: "P", 1: "R"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "Orara", 8: "Opedel", 10: "Otibar", 12: "Opoo", 6: "Omodok’king’ol", 7: "Ojola", 9: "Osokosokoma", 11: "Olabor", 2: "Omuk", 3: "Okwamg’", 4: "Odung’el", 5: "Omaruk"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Aar", 3: "Uni", 4: "Ung", 5: "Kan", 6: "Sab", 0: "Jum", 1: "Bar"}, Narrow: ut.CalendarDayFormatNameValue{4: "U", 5: "K", 6: "S", 0: "J", 1: "B", 2: "A", 3: "U"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "Nakaejuma", 1: "Nakaebarasa", 2: "Nakaare", 3: "Nakauni", 4: "Nakaung’on", 5: "Nakakany", 6: "Nakasabiti"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "Ebongi", "am": "Taparachu"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Taparachu", "pm": "Ebongi"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Kabla ya Christo", Abbrev: "KK", Narrow: ""}}}}
