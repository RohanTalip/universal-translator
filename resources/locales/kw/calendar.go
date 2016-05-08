package kw

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Gen", 2: "Hwe", 3: "Meu", 6: "Met", 8: "Est", 9: "Gwn", 11: "Du", 4: "Ebr", 5: "Me", 7: "Gor", 10: "Hed", 12: "Kev"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "mis Genver", 2: "mis Hwevrer", 4: "mis Ebrel", 7: "mis Gortheren", 8: "mis Est", 10: "mis Hedra", 12: "mis Kevardhu", 3: "mis Meurth", 5: "mis Me", 6: "mis Metheven", 9: "mis Gwynngala", 11: "mis Du"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "Gwe", 6: "Sad", 0: "Sul", 1: "Lun", 2: "Mth", 3: "Mhr", 4: "Yow"}, Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{2: "dy Meurth", 3: "dy Merher", 4: "dy Yow", 5: "dy Gwener", 6: "dy Sadorn", 0: "dy Sul", 1: "dy Lun"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "a.m.", "pm": "p.m."}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "a.m.", "pm": "p.m."}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "RC", Narrow: ""}}}}
