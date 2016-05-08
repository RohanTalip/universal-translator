package sw_CD

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "min", 6: "mst", 7: "msb", 8: "mun", 9: "mts", 1: "mkw", 2: "mpi", 3: "mtu", 11: "mkm", 5: "mtn", 10: "mku", 12: "mkb"}, Narrow: ut.CalendarMonthFormatNameValue{1: "k", 4: "i", 6: "s", 7: "s", 9: "t", 12: "m", 2: "p", 3: "t", 5: "t", 8: "m", 10: "k", 11: "m"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{4: "mwezi ya ine", 5: "mwezi ya tanu", 7: "mwezi ya saba", 8: "mwezi ya munane", 11: "mwezi ya kumi na moya", 12: "mwezi ya kumi ya mbili", 3: "mwezi ya tatu", 2: "mwezi ya pili", 6: "mwezi ya sita", 9: "mwezi ya tisa", 10: "mwezi ya kumi", 1: "mwezi ya kwanja"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "sit", 0: "yen", 1: "kwa", 2: "pil", 3: "tat", 4: "ine", 5: "tan"}, Narrow: ut.CalendarDayFormatNameValue{4: "i", 5: "t", 6: "s", 0: "y", 1: "k", 2: "p", 3: "t"}, Short: ut.CalendarDayFormatNameValue{1: "Jumatatu", 2: "Jumanne", 3: "Jumatano", 4: "Alhamisi", 5: "Ijumaa", 6: "Jumamosi", 0: "Jumapili"}, Wide: ut.CalendarDayFormatNameValue{6: "siku ya sita", 0: "siku ya yenga", 1: "siku ya kwanza", 2: "siku ya pili", 3: "siku ya tatu", 4: "siku ya ine", 5: "siku ya tanu"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "alfajiri", "night1": "usiku", "evening1": "jioni", "noon": "saa sita za mchana", "afternoon1": "mchana", "am": "AM", "morning2": "asubuhi", "midnight": "saa sita za usiku"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "saa sita za usiku", "afternoon1": "mchana", "morning2": "asubuhi", "night1": "usiku", "pm": "pm", "am": "am", "noon": "saa sita za mchana", "evening1": "jioni", "morning1": "alfajiri"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "alfajiri", "evening1": "jioni", "night1": "usiku", "midnight": "saa sita za usiku", "noon": "saa sita za mchana", "am": "ya asubuyi", "pm": "ya muchana", "morning2": "asubuhi", "afternoon1": "mchana"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
