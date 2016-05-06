package bem

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "Oga", 9: "Sep", 2: "Feb", 3: "Mac", 6: "Jun", 7: "Jul", 10: "Okt", 11: "Nov", 12: "Dis", 1: "Jan", 4: "Epr", 5: "Mei"}, Narrow: ut.CalendarMonthFormatNameValue{2: "F", 3: "M", 5: "M", 6: "J", 9: "S", 11: "N", 12: "D", 1: "J", 4: "E", 7: "J", 8: "O", 10: "O"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{5: "Mei", 11: "Novemba", 1: "Januari", 3: "Machi", 4: "Epreo", 6: "Juni", 7: "Julai", 8: "Ogasti", 9: "Septemba", 10: "Oktoba", 2: "Februari", 12: "Disemba"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue(nil), Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{6: "Pachibelushi", 0: "Pa Mulungu", 1: "Palichimo", 2: "Palichibuli", 3: "Palichitatu", 4: "Palichine", 5: "Palichisano"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "uluchelo", "pm": "akasuba"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "uluchelo", "pm": "akasuba"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Before Yesu", Abbrev: "BC", Narrow: ""}}}}
