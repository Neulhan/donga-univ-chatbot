package main

const haksikEndpoint string = "http://www.donga.ac.kr/gzSub_007005005.aspx"
const libraryEndpoint string = "http://168.115.33.207/WebSeat/Domian5.asp"
const informationEndpoint string = "http://www.donga.ac.kr/gzSub_007005008.aspx"
const weatherEndpoint string = "http://apis.data.go.kr/1360000/VilageFcstInfoService/getVilageFcst"
const weatherUltraEndpoint string = "http://apis.data.go.kr/1360000/VilageFcstInfoService/getUltraSrtNcst"
const weatherKey string = "D6vdl5myJ0a0xl%2FDzf61U3kriBw%2FiCmMqJoIAoJ3Mg4kFkmXD1Wj6WwehGkJFW%2FIXJ4hd6l%2BAlVrB4DXKlsL2w%3D%3D"

var categoryMatching = map[string]string{
	"POP": "강수확률",
	"PTY": "강수형태",
	"R06": "6시간 강수량",
	"REH": "습도",
	"S06": "6시간 신적설",
	"SKY": "하늘상태",
	"T3H": "3시간 기온",
	"UUU": "풍속(동서)",
	"VEC": "풍향",
	"VVV": "풍속(남북)",
	"WSD": "풍속",
	"WAV": "파고",
	"TMX": "낮 최고기온",
	"TMN": "아침 최저기온",
	"T1H": "기온",
	"RN1": "1시간 강수량",
}
var timeMapping = map[string]string{
	"0":  "23",
	"1":  "23",
	"2":  "02",
	"3":  "02",
	"4":  "02",
	"5":  "05",
	"6":  "05",
	"7":  "05",
	"8":  "08",
	"9":  "08",
	"10": "08",
	"11": "11",
	"12": "11",
	"13": "11",
	"14": "14",
	"15": "14",
	"16": "14",
	"17": "17",
	"18": "17",
	"19": "19",
	"20": "20",
	"21": "20",
	"22": "20",
	"23": "20",
	"24": "20",
}

const sunnyURL string = "http://222.237.78.242:9091/assets/sunny.png"
const cloudyURL string = "http://222.237.78.242:9091/assets/cloudy.png"
const rainyURL string = "http://222.237.78.242:9091/assets/rainy.png"
const snowyURL string = "http://222.237.78.242:9091/assets/snowy.png"
