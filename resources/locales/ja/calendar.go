package ja

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "y年M月d日EEEE", Long: "y年M月d日", Medium: "y/MM/dd", Short: "y/MM/dd"}, AD: ut.CalendarDateFormat{Full: "y年M月d日EEEE", Long: "y年M月d日", Medium: "y/MM/dd", Short: "y/MM/dd"}}, Time: ut.CalendarDateFormat{Full: "H時mm分ss秒 zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "1月", 2: "2月", 4: "4月", 10: "10月", 11: "11月", 12: "12月", 3: "3月", 5: "5月", 6: "6月", 7: "7月", 8: "8月", 9: "9月"}, Narrow: ut.CalendarMonthFormatNameValue{7: "7", 8: "8", 9: "9", 1: "1", 3: "3", 4: "4", 10: "10", 11: "11", 12: "12", 2: "2", 5: "5", 6: "6"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "2月", 3: "3月", 4: "4月", 8: "8月", 11: "11月", 12: "12月", 1: "1月", 5: "5月", 6: "6月", 7: "7月", 9: "9月", 10: "10月"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "水", 4: "木", 5: "金", 6: "土", 0: "日", 1: "月", 2: "火"}, Narrow: ut.CalendarDayFormatNameValue{6: "土", 0: "日", 1: "月", 2: "火", 3: "水", 4: "木", 5: "金"}, Short: ut.CalendarDayFormatNameValue{4: "木", 5: "金", 6: "土", 0: "日", 1: "月", 2: "火", 3: "水"}, Wide: ut.CalendarDayFormatNameValue{0: "日曜日", 1: "月曜日", 2: "火曜日", 3: "水曜日", 4: "木曜日", 5: "金曜日", 6: "土曜日"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "真夜中", "am": "午前", "noon": "正午", "morning1": "朝", "afternoon1": "昼", "night1": "夜", "pm": "午後", "evening1": "夕方", "night2": "夜中"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "午後", "morning1": "朝", "afternoon1": "昼", "evening1": "夕方", "night1": "夜", "midnight": "真夜中", "noon": "正午", "night2": "夜中", "am": "午前"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"night1": "夜", "midnight": "真夜中", "am": "午前", "noon": "正午", "afternoon1": "昼", "pm": "午後", "morning1": "朝", "evening1": "夕方", "night2": "夜中"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "西暦", Abbrev: "西暦", Narrow: "AD"}, BC: ut.CalendarEraFormatNames{Full: "紀元前", Abbrev: "紀元前", Narrow: "BC"}}}}
