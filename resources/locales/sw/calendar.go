package sw

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "Apr", 6: "Jun", 9: "Sep", 11: "Nov", 12: "Des", 1: "Jan", 2: "Feb", 7: "Jul", 8: "Ago", 10: "Okt", 3: "Mac", 5: "Mei"}, Narrow: ut.CalendarMonthFormatNameValue{2: "F", 12: "D", 4: "A", 5: "M", 6: "J", 7: "J", 8: "A", 9: "S", 1: "J", 3: "M", 10: "O", 11: "N"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "Septemba", 1: "Januari", 4: "Aprili", 6: "Juni", 7: "Julai", 8: "Agosti", 12: "Desemba", 2: "Februari", 3: "Machi", 5: "Mei", 10: "Oktoba", 11: "Novemba"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "Jumatano", 4: "Alhamisi", 5: "Ijumaa", 6: "Jumamosi", 0: "Jumapili", 1: "Jumatatu", 2: "Jumanne"}, Narrow: ut.CalendarDayFormatNameValue{1: "M", 2: "T", 3: "W", 4: "T", 5: "F", 6: "S", 0: "S"}, Short: ut.CalendarDayFormatNameValue{5: "Ijumaa", 6: "Jumamosi", 0: "Jumapili", 1: "Jumatatu", 2: "Jumanne", 3: "Jumatano", 4: "Alhamisi"}, Wide: ut.CalendarDayFormatNameValue{2: "Jumanne", 3: "Jumatano", 4: "Alhamisi", 5: "Ijumaa", 6: "Jumamosi", 0: "Jumapili", 1: "Jumatatu"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "saa sita za usiku", "noon": "saa sita za mchana", "pm": "PM", "afternoon1": "mchana", "evening1": "jioni", "am": "AM", "morning1": "alfajiri", "morning2": "asubuhi", "night1": "usiku"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "usiku", "pm": "pm", "morning1": "alfajiri", "morning2": "asubuhi", "afternoon1": "mchana", "midnight": "saa sita za usiku", "am": "am", "noon": "saa sita za mchana", "evening1": "jioni"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "saa sita za mchana", "morning2": "asubuhi", "afternoon1": "mchana", "midnight": "saa sita za usiku", "pm": "PM", "morning1": "alfajiri", "evening1": "jioni", "night1": "usiku"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "Baada ya Kristo", Abbrev: "AD", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Kabla ya Kristo", Abbrev: "BC", Narrow: ""}}}}
