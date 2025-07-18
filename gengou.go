// Package gengou implements functions for converting dates to the
// Japanese era calendar scheme.
package gengou

import (
	"fmt"
	"slices"
	"time"
)

const Version = "0.1.0"

const Offset = 9 * 60 * 60

// Represents a Japanese era.
type Era struct {
	Name, Kana string
	Y, M, D    int
	Date       *time.Time
}

// All Japanese eras, from 大化 to 令和.
var Eras = []Era{
	{"大化", "たいか", 645, 1, 1, nil}, // date?
	{"白雉", "はくち", 650, 2, 15, nil},
	{"朱鳥", "しゅちょう（すちょう）", 686, 7, 20, nil},
	{"大宝", "たいほう（だいほう）", 701, 3, 21, nil},
	{"慶雲", "けいうん（きょううん）", 704, 5, 10, nil},
	{"和銅", "わどう", 708, 1, 11, nil},
	{"養老", "ようろう", 717, 11, 17, nil},
	{"神亀", "じんき", 724, 2, 4, nil},
	{"天平", "てんぴょう（てんびょう）", 729, 8, 5, nil},
	{"天平感宝", "てんぴょうかんぽう", 749, 4, 14, nil},
	{"天平勝宝", "てんぴょうしょうほう", 749, 7, 2, nil},
	{"天平宝字", "てんぴょうほうじ", 757, 8, 18, nil},
	{"天平神護", "てんぴょうしんご", 765, 1, 7, nil},
	{"神護景雲", "しんごけいうん", 767, 8, 16, nil},
	{"宝亀", "ほうき", 770, 10, 1, nil},
	{"天応", "てんおう", 781, 1, 1, nil},
	{"延暦", "えんりゃく", 782, 8, 19, nil},
	{"弘仁", "こうにん", 810, 9, 19, nil},
	{"天長", "てんちょう", 824, 1, 5, nil},
	{"承和", "じょうわ（しょうわ）", 834, 1, 3, nil},
	{"嘉祥", "かしょう（かじょう）", 848, 6, 13, nil},
	{"仁寿", "にんじゅ", 851, 4, 28, nil},
	{"斎衡", "さいこう", 854, 11, 30, nil},
	{"天安", "てんあん（てんなん）", 857, 2, 21, nil},
	{"貞観", "じょうがん", 859, 4, 15, nil},
	{"元慶", "がんぎょう（げんけい）", 877, 4, 16, nil},
	{"仁和", "にんな（じんな）", 885, 2, 21, nil},
	{"寛平", "かんぴょう（かんぺい）", 889, 4, 27, nil},
	{"昌泰", "しょうたい", 898, 4, 26, nil},
	{"延喜", "えんぎ", 901, 7, 15, nil},
	{"延長", "えんちょう", 923, 4, 11, nil},
	{"承平", "じょうへい（しょうへい）", 931, 4, 26, nil},
	{"天慶", "てんぎょう（てんきょう）", 938, 5, 22, nil},
	{"天暦", "てんりゃく（てんれき）", 947, 4, 22, nil},
	{"天徳", "てんとく", 957, 10, 27, nil},
	{"応和", "おうわ", 961, 2, 16, nil},
	{"康保", "こうほう", 964, 7, 10, nil},
	{"安和", "あんな（あんわ）", 968, 8, 13, nil},
	{"天禄", "てんろく", 970, 3, 25, nil},
	{"天延", "てんえん", 973, 12, 20, nil},
	{"貞元", "じょうげん（ていげん）", 976, 7, 13, nil},
	{"天元", "てんげん", 978, 11, 29, nil},
	{"永観", "えいかん", 983, 4, 15, nil},
	{"寛和", "かんな（かんわ）", 985, 4, 27, nil},
	{"永延", "えいえん（ようえん）", 987, 4, 5, nil},
	{"永祚", "えいそ", 989, 8, 8, nil},
	{"正暦", "しょうりゃく（じょうりゃく）", 990, 11, 7, nil},
	{"長徳", "ちょうとく", 995, 2, 22, nil},
	{"長保", "ちょうほう", 999, 1, 13, nil},
	{"寛弘", "かんこう", 1004, 7, 20, nil},
	{"長和", "ちょうわ", 1012, 12, 25, nil},
	{"寛仁", "かんにん", 1017, 4, 23, nil},
	{"治安", "じあん（ちあん）", 1021, 2, 2, nil},
	{"万寿", "まんじゅ", 1024, 7, 13, nil},
	{"長元", "ちょうげん", 1028, 7, 25, nil},
	{"長暦", "ちょうりゃく（ちょうれき）", 1037, 4, 21, nil},
	{"長久", "ちょうきゅう", 1040, 11, 10, nil},
	{"寛徳", "かんとく", 1044, 11, 24, nil},
	{"永承", "えいしょう（えいじょう）", 1046, 4, 14, nil},
	{"天喜", "てんぎ（てんき）", 1053, 1, 11, nil},
	{"康平", "こうへい", 1058, 8, 29, nil},
	{"治暦", "じりゃく（ちりゃく）", 1065, 8, 2, nil},
	{"延久", "えんきゅう", 1069, 4, 13, nil},
	{"承保", "じょうほう（しょうほう）", 1074, 8, 23, nil},
	{"承暦", "じょうりゃく（しょうりゃく）", 1077, 11, 17, nil},
	{"永保", "えいほう", 1081, 2, 10, nil},
	{"応徳", "おうとく", 1084, 2, 7, nil},
	{"寛治", "かんじ", 1087, 4, 7, nil},
	{"嘉保", "かほう", 1094, 12, 15, nil},
	{"永長", "えいちょう（ようちょう）", 1096, 12, 17, nil},
	{"承徳", "じょうとく（しょうとく）", 1097, 11, 21, nil},
	{"康和", "こうわ", 1099, 8, 28, nil},
	{"長治", "ちょうじ", 1104, 2, 10, nil},
	{"嘉承", "かじょう（かしょう）", 1106, 4, 9, nil},
	{"天仁", "てんにん", 1108, 8, 3, nil},
	{"天永", "てんえい", 1110, 7, 13, nil},
	{"永久", "えいきゅう", 1113, 7, 13, nil},
	{"元永", "げんえい", 1118, 4, 3, nil},
	{"保安", "ほうあん", 1120, 4, 10, nil},
	{"天治", "てんじ", 1124, 4, 3, nil},
	{"大治", "だいじ（たいじ）", 1126, 1, 22, nil},
	{"天承", "てんしょう（てんじょう）", 1131, 1, 29, nil},
	{"長承", "ちょうしょう（ちょうじょう）", 1132, 8, 11, nil},
	{"保延", "ほうえん", 1135, 4, 27, nil},
	{"永治", "えいじ", 1141, 7, 10, nil},
	{"康治", "こうじ", 1142, 4, 28, nil},
	{"天養", "てんよう", 1144, 2, 23, nil},
	{"久安", "きゅうあん", 1145, 7, 22, nil},
	{"仁平", "にんぺい（にんびょう）", 1151, 1, 26, nil},
	{"久寿", "きゅうじゅ", 1154, 10, 28, nil},
	{"保元", "ほうげん", 1156, 4, 27, nil},
	{"平治", "へいじ（びょうじ）", 1159, 4, 20, nil},
	{"永暦", "えいりゃく（ようりゃく）", 1160, 1, 10, nil},
	{"応保", "おうほう", 1161, 9, 4, nil},
	{"長寛", "ちょうかん", 1163, 3, 29, nil},
	{"永万", "えいまん（ようまん）", 1165, 6, 5, nil},
	{"仁安", "にんあん（にんなん）", 1166, 8, 27, nil},
	{"嘉応", "かおう", 1169, 4, 8, nil},
	{"承安", "じょうあん（しょうあん）", 1171, 4, 21, nil},
	{"安元", "あんげん", 1175, 7, 28, nil},
	{"治承", "じしょう（じじょう）", 1177, 8, 4, nil},
	{"養和", "ようわ", 1181, 7, 14, nil},
	{"寿永", "じゅえい", 1182, 5, 27, nil},
	{"元暦", "げんりゃく", 1184, 4, 16, nil},
	{"文治", "ぶんじ（もんじ）", 1185, 8, 14, nil},
	{"正治", "しょうじ", 1199, 4, 27, nil},
	{"建仁", "けんにん", 1201, 2, 13, nil},
	{"元久", "げんきゅう", 1204, 2, 20, nil},
	{"建永", "けんえい", 1206, 4, 27, nil},
	{"承元", "じょうげん（しょうげん）", 1207, 10, 25, nil},
	{"建暦", "けんりゃく", 1211, 3, 9, nil},
	{"建保", "けんぽう（けんほう）", 1213, 12, 6, nil},
	{"承久", "じょうきゅう（しょうきゅう）", 1219, 4, 12, nil},
	{"貞応", "じょうおう（ていおう）", 1222, 4, 13, nil},
	{"元仁", "げんにん", 1224, 11, 20, nil},
	{"嘉禄", "かろく", 1225, 4, 20, nil},
	{"安貞", "あんてい", 1227, 12, 10, nil},
	{"寛喜", "かんき", 1229, 3, 5, nil},
	{"貞永", "じょうえい（ていえい）", 1232, 4, 2, nil},
	{"天福", "てんぷく（てんふく）", 1233, 4, 15, nil},
	{"文暦", "ぶんりゃく（もんりゃく）", 1234, 11, 5, nil},
	{"嘉禎", "かてい", 1235, 9, 19, nil},
	{"暦仁", "りゃくにん（れきにん）", 1238, 11, 23, nil},
	{"延応", "えんおう（えんのう）", 1239, 2, 7, nil},
	{"仁治", "にんじ（にんち）", 1240, 7, 16, nil},
	{"寛元", "かんげん", 1243, 2, 26, nil},
	{"宝治", "ほうじ", 1247, 2, 28, nil},
	{"建長", "けんちょう", 1249, 3, 18, nil},
	{"康元", "こうげん", 1256, 10, 5, nil},
	{"正嘉", "しょうか", 1257, 3, 14, nil},
	{"正元", "しょうげん", 1259, 3, 26, nil},
	{"文応", "ぶんおう", 1260, 4, 13, nil},
	{"弘長", "こうちょう", 1261, 2, 20, nil},
	{"文永", "ぶんえい", 1264, 2, 28, nil},
	{"建治", "けんじ", 1275, 4, 25, nil},
	{"弘安", "こうあん", 1278, 2, 29, nil},
	{"正応", "しょうおう", 1288, 4, 28, nil},
	{"永仁", "えいにん", 1293, 8, 5, nil},
	{"正安", "しょうあん", 1299, 4, 25, nil},
	{"乾元", "けんげん", 1302, 11, 21, nil},
	{"嘉元", "かげん", 1303, 8, 5, nil},
	{"徳治", "とくじ", 1306, 12, 14, nil},
	{"延慶", "えんきょう（えんぎょう）", 1308, 10, 9, nil},
	{"応長", "おうちょう", 1311, 4, 28, nil},
	{"正和", "しょうわ", 1312, 3, 20, nil},
	{"文保", "ぶんぽう（ぶんほう）", 1317, 2, 3, nil},
	{"元応", "げんおう（げんのう）", 1319, 4, 28, nil},
	{"元亨", "げんこう", 1321, 2, 23, nil},
	{"正中", "しょうちゅう", 1324, 12, 9, nil},
	{"嘉暦", "かりゃく", 1326, 4, 26, nil},
	{"元徳", "げんとく", 1329, 8, 29, nil},
	{"元弘", "げんこう", 1331, 8, 9, nil},
	{"正慶", "しょうきょう（しょうけい）", 1332, 4, 28, nil},
	{"建武", "けんむ（けんぶ）", 1334, 1, 29, nil},
	{"延元", "えんげん", 1336, 2, 29, nil},
	{"暦応", "りゃくおう（れきおう）", 1338, 8, 28, nil},
	{"興国", "こうこく", 1340, 4, 28, nil},
	{"康永", "こうえい", 1342, 4, 27, nil},
	{"貞和", "じょうわ（ていわ）", 1345, 10, 21, nil},
	{"正平", "しょうへい", 1346, 12, 8, nil},
	{"観応", "かんおう（かんのう）", 1350, 2, 27, nil},
	{"文和", "ぶんな（ぶんわ）", 1352, 9, 27, nil},
	{"延文", "えんぶん", 1356, 3, 28, nil},
	{"康安", "こうあん", 1361, 3, 29, nil},
	{"貞治", "じょうじ（ていじ）", 1362, 9, 23, nil},
	{"応安", "おうあん", 1368, 2, 18, nil},
	{"建徳", "けんとく", 1370, 7, 24, nil},
	{"文中", "ぶんちゅう", 1372, 4, 1, nil},
	{"天授", "てんじゅ", 1375, 5, 27, nil},
	{"永和", "えいわ", 1375, 2, 27, nil},
	{"康暦", "こうりゃく", 1379, 3, 22, nil},
	{"弘和", "こうわ", 1381, 2, 10, nil},
	{"永徳", "えいとく", 1381, 2, 24, nil},
	{"元中", "げんちゅう", 1384, 4, 28, nil},
	{"至徳", "しとく", 1384, 2, 27, nil},
	{"嘉慶", "かきょう（かけい）", 1387, 8, 23, nil},
	{"康応", "こうおう", 1389, 2, 9, nil},
	{"明徳", "めいとく", 1390, 3, 26, nil},
	{"応永", "おうえい", 1394, 7, 5, nil},
	{"正長", "しょうちょう", 1428, 4, 27, nil},
	{"永享", "えいきょう", 1429, 9, 5, nil},
	{"嘉吉", "かきつ（かきち）", 1441, 2, 17, nil},
	{"文安", "ぶんあん", 1444, 2, 5, nil},
	{"宝徳", "ほうとく", 1449, 7, 28, nil},
	{"享徳", "きょうとく", 1452, 7, 25, nil},
	{"康正", "こうしょう", 1455, 7, 25, nil},
	{"長禄", "ちょうろく", 1457, 9, 28, nil},
	{"寛正", "かんしょう", 1460, 12, 21, nil},
	{"文正", "ぶんしょう（もんしょう）", 1466, 2, 28, nil},
	{"応仁", "おうにん", 1467, 3, 5, nil},
	{"文明", "ぶんめい", 1469, 4, 28, nil},
	{"長享", "ちょうきょう", 1487, 7, 20, nil},
	{"延徳", "えんとく", 1489, 8, 21, nil},
	{"明応", "めいおう", 1492, 7, 19, nil},
	{"文亀", "ぶんき", 1501, 2, 29, nil},
	{"永正", "えいしょう", 1504, 2, 30, nil},
	{"大永", "だいえい", 1521, 8, 23, nil},
	{"享禄", "きょうろく", 1528, 8, 20, nil},
	{"天文", "てんぶん", 1532, 7, 29, nil},
	{"弘治", "こうじ", 1555, 10, 23, nil},
	{"永禄", "えいろく", 1558, 2, 28, nil},
	{"元亀", "げんき", 1570, 4, 23, nil},
	{"文禄", "ぶんろく", 1592, 12, 8, nil},
	{"慶長", "けいちょう（きょうちょう）", 1596, 10, 27, nil},
	{"寛永", "かんえい", 1624, 2, 30, nil},
	{"正保", "しょうほう", 1644, 12, 16, nil},
	{"慶安", "けいあん", 1648, 2, 15, nil},
	{"承応", "じょうおう（しょうおう）", 1652, 9, 18, nil},
	{"明暦", "めいれき（みょうりゃく）", 1655, 4, 13, nil},
	{"万治", "まんじ", 1658, 7, 23, nil},
	{"寛文", "かんぶん", 1661, 4, 25, nil},
	{"延宝", "えんぽう", 1673, 9, 21, nil},
	{"天和", "てんな", 1681, 9, 29, nil},
	{"貞享", "じょうきょう", 1684, 2, 21, nil},
	{"元禄", "げんろく", 1688, 9, 30, nil},
	{"宝永", "ほうえい", 1704, 3, 13, nil},
	{"正徳", "しょうとく", 1711, 4, 25, nil},
	{"享保", "きょうほう（きょうほ）", 1716, 6, 22, nil},
	{"元文", "げんぶん", 1736, 4, 28, nil},
	{"寛保", "かんぽう（かんほう）", 1741, 2, 27, nil},
	{"延享", "えんきょう", 1744, 2, 21, nil},
	{"寛延", "かんえん", 1748, 7, 12, nil},
	{"宝暦", "ほうれき（ほうりゃく）", 1751, 10, 27, nil},
	{"明和", "めいわ", 1764, 6, 2, nil},
	{"安永", "あんえい", 1772, 11, 16, nil},
	{"天明", "てんめい", 1781, 4, 2, nil},
	{"寛政", "かんせい", 1789, 1, 25, nil},
	{"享和", "きょうわ", 1801, 2, 5, nil},
	{"文化", "ぶんか", 1804, 2, 11, nil},
	{"文政", "ぶんせい", 1818, 4, 22, nil},
	{"天保", "てんぽう（てんほう）", 1830, 12, 10, nil},
	{"弘化", "こうか", 1844, 12, 2, nil},
	{"嘉永", "かえい", 1848, 2, 28, nil},
	{"安政", "あんせい", 1854, 11, 27, nil},
	{"万延", "まんえん", 1860, 3, 18, nil},
	{"文久", "ぶんきゅう", 1861, 2, 19, nil},
	{"元治", "げんじ", 1864, 2, 20, nil},
	{"慶応", "けいおう", 1865, 4, 7, nil},
	{"明治", "めいじ", 1868, 9, 8, nil},
	{"大正", "たいしょう", 1912, 7, 30, nil},
	{"昭和", "しょうわ", 1926, 12, 25, nil},
	{"平成", "へいせい", 1989, 1, 8, nil},
	{"令和", "れいわ", 2019, 5, 1, nil},
}

func init() {
	loc := time.FixedZone("JST", Offset)
	for i, era := range Eras {
		date := time.Date(era.Y, time.Month(era.M), era.D, 0, 0, 0, 0, loc)
		Eras[i].Date = &date
	}
}

// Find returns the Era matching time t.
// Reports an error if the date is before the first era (大化).
func Find(t time.Time) (*Era, error) {
	for _, era := range slices.Backward(Eras) {
		if t.Equal(*era.Date) || t.After(*era.Date) {
			return &era, nil
		}
	}
	return nil, fmt.Errorf("era not found")
}

// EraYear returns the era name and year corresponding to t
func EraYear(t time.Time) string {
	era, err := Find(t)
	if err != nil {
		return fmt.Sprintf("%d年", t.Year())
	}
	if t.Year() == era.Y {
		return fmt.Sprintf("%s元年", era.Name)
	}
	loc := time.FixedZone("JST", Offset)
	start := time.Date(era.Y, 1, 1, 0, 0, 0, 0, loc)
	year := t.Year() - start.Year() + 1
	return fmt.Sprintf("%s%d年", era.Name, year)
}

// EraDate returns the full era name and date corresponding to t
func EraDate(t time.Time) string {
	eraYear := EraYear(t)
	return fmt.Sprintf("%s%d月%d日", eraYear, t.Month(), t.Day())
}

// EraDateWithZeros returns the era year and zero-prefixed month and day for t
func EraDateWithZeros(t time.Time) string {
	eraYear := EraYear(t)
	return fmt.Sprintf("%s%02d月%02d日", eraYear, t.Month(), t.Day())
}
