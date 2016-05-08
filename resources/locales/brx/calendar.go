package brx

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue(nil), Narrow: ut.CalendarMonthFormatNameValue{10: "अ", 11: "न", 1: "ज", 2: "फे", 7: "जु", 8: "आ", 9: "से", 12: "दि", 3: "मा", 4: "ए", 5: "मे", 6: "जु"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{11: "नबेज्ब़र", 1: "जानुवारी", 2: "फेब्रुवारी", 3: "मार्स", 4: "एफ्रिल", 6: "जुन", 7: "जुलाइ", 10: "अखथबर", 12: "दिसेज्ब़र", 5: "मे", 8: "आगस्थ", 9: "सेबथेज्ब़र"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "सुनि", 0: "रबि", 1: "सम", 2: "मंगल", 3: "बुद", 4: "बिसथि", 5: "सुखुर"}, Narrow: ut.CalendarDayFormatNameValue{4: "बि", 5: "सु", 6: "सु", 0: "र", 1: "स", 2: "मं", 3: "बु"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "रबिबार", 1: "समबार", 2: "मंगलबार", 3: "बुदबार", 4: "बिसथिबार", 5: "सुखुरबार", 6: "सुनिबार"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "फुं", "pm": "बेलासे"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "फुं", "pm": "बेलासे"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "ईसा.पूर्व", Narrow: ""}}}}
