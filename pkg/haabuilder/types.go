package haabuilder

type Pack struct {
	Code           string `json:"code"`
	CycleCode      string `json:"cycle_code"`
	HaaBuilderCode string
	Size           int `json:"size,omitempty"`
}

var Packs = []Pack{
	{
		HaaBuilderCode: "1",
		Code:           "core",
		CycleCode:      "core",
	},
	// dwl
	{
		HaaBuilderCode: "2",
		Code:           "dwl",
		CycleCode:      "dwl",
	},
	{
		HaaBuilderCode: "3",
		Code:           "tmm",
		CycleCode:      "dwl",
	},
	{
		HaaBuilderCode: "4",
		CycleCode:      "dwl",
		Code:           "tece",
	},
	{
		HaaBuilderCode: "5",
		CycleCode:      "dwl",
		Code:           "bota",
	},
	{
		HaaBuilderCode: "6",
		CycleCode:      "dwl",
		Code:           "uau",
	},
	{
		HaaBuilderCode: "7",
		CycleCode:      "dwl",
		Code:           "wda",
	},
	{
		HaaBuilderCode: "8",
		CycleCode:      "dwl",
		Code:           "litas",
	},
	// ptc
	{
		HaaBuilderCode: "9",
		CycleCode:      "ptc",
		Code:           "ptc",
	},
	{
		HaaBuilderCode: "10",
		CycleCode:      "ptc",
		Code:           "eotp",
	},
	{
		HaaBuilderCode: "11",
		CycleCode:      "ptc",
		Code:           "tuo",
	},
	{
		HaaBuilderCode: "12",
		CycleCode:      "ptc",
		Code:           "apot",
	},
	{
		HaaBuilderCode: "13",
		CycleCode:      "ptc",
		Code:           "tpm",
	},
	{
		HaaBuilderCode: "14",
		CycleCode:      "ptc",
		Code:           "bsr",
	},
	{
		HaaBuilderCode: "15",
		CycleCode:      "ptc",
		Code:           "dca",
	},
	// tfa
	{
		HaaBuilderCode: "16",
		CycleCode:      "tfa",
		Code:           "tfa",
	},
	{
		HaaBuilderCode: "17",
		CycleCode:      "tfa",
		Code:           "tof",
	},
	{
		HaaBuilderCode: "18",
		CycleCode:      "tfa",
		Code:           "tbb",
	},
	{
		HaaBuilderCode: "19",
		CycleCode:      "tfa",
		Code:           "hote",
	},
	{
		HaaBuilderCode: "20",
		CycleCode:      "tfa",
		Code:           "tcoa",
	},
	{
		HaaBuilderCode: "21",
		CycleCode:      "tfa",
		Code:           "tdoy",
	},
	{
		HaaBuilderCode: "22",
		CycleCode:      "tfa",
		Code:           "sha",
	},
	// tcu
	{
		HaaBuilderCode: "23",
		CycleCode:      "tcu",
		Code:           "tcu",
	},
	{
		HaaBuilderCode: "24",
		CycleCode:      "tcu",
		Code:           "tsn",
	},
	{
		HaaBuilderCode: "25",
		CycleCode:      "tcu",
		Code:           "wos",
	},
	{
		HaaBuilderCode: "26",
		CycleCode:      "tcu",
		Code:           "fgg",
	},
	{
		HaaBuilderCode: "27",
		CycleCode:      "tcu",
		Code:           "uad",
	},
	{
		HaaBuilderCode: "28",
		CycleCode:      "tcu",
		Code:           "icc",
	},
	{
		HaaBuilderCode: "29",
		CycleCode:      "tcu",
		Code:           "bbt",
	},
	// tda
	{
		HaaBuilderCode: "30",
		CycleCode:      "tde",
		Code:           "tde",
	},
	{
		HaaBuilderCode: "31",
		CycleCode:      "tde",
		Code:           "sfk",
	},
	{
		HaaBuilderCode: "32",
		CycleCode:      "tde",
		Code:           "tsh",
	},
	{
		HaaBuilderCode: "36",
		CycleCode:      "tde",
		Code:           "dsm",
	},
	{
		HaaBuilderCode: "37",
		CycleCode:      "tde",
		Code:           "pnr",
	},
	{
		HaaBuilderCode: "38",
		CycleCode:      "tde",
		Code:           "wgd",
	},
	{
		HaaBuilderCode: "39",
		CycleCode:      "tde",
		Code:           "woc",
	},
	// tic
	{
		HaaBuilderCode: "46",
		CycleCode:      "tic",
		Code:           "tic",
	},
	{
		HaaBuilderCode: "54",
		CycleCode:      "tic",
		Code:           "itd",
	},
	{
		HaaBuilderCode: "55",
		CycleCode:      "tic",
		Code:           "def",
	},
	{
		HaaBuilderCode: "57",
		CycleCode:      "tic",
		Code:           "hhg",
	},
	{
		HaaBuilderCode: "58",
		CycleCode:      "tic",
		Code:           "lif",
	},
	{
		HaaBuilderCode: "59",
		CycleCode:      "tic",
		Code:           "lod",
	},
	{
		HaaBuilderCode: "60",
		CycleCode:      "tic",
		Code:           "itm",
	},
	// eoe
	{
		HaaBuilderCode: "62",
		CycleCode:      "eoe",
		Code:           "eoec",
	},
	{
		HaaBuilderCode: "63",
		CycleCode:      "eoe",
		Code:           "eoep",
	},
	// tsk
	{
		HaaBuilderCode: "66",
		CycleCode:      "tsk",
		Code:           "tskc",
	},
	{
		HaaBuilderCode: "65",
		CycleCode:      "tsk",
		Code:           "tskp",
	},
	// fhv
	{
		HaaBuilderCode: "68",
		CycleCode:      "fhv",
		Code:           "fhvc",
	},
	{
		HaaBuilderCode: "69",
		CycleCode:      "fhv",
		Code:           "fhvp",
	},
	// return
	{
		HaaBuilderCode: "33",
		CycleCode:      "return",
		Code:           "rtnotz",
	},
	{
		HaaBuilderCode: "34",
		CycleCode:      "return",
		Code:           "rtdwl",
	},
	{
		HaaBuilderCode: "35",
		CycleCode:      "return",
		Code:           "rtptc",
	},
	{
		HaaBuilderCode: "40",
		CycleCode:      "return",
		Code:           "rttfa",
	},
	{
		HaaBuilderCode: "61",
		CycleCode:      "return",
		Code:           "rttcu",
	},
	// side
	{
		HaaBuilderCode: "47",
		CycleCode:      "side_stories",
		Code:           "guardians",
	},
	{
		HaaBuilderCode: "48",
		CycleCode:      "side_stories",
		Code:           "hotel",
	},
	{
		HaaBuilderCode: "49",
		CycleCode:      "side_stories",
		Code:           "blob",
	},
	{
		HaaBuilderCode: "50",
		CycleCode:      "side_stories",
		Code:           "lol",
	},
	{
		HaaBuilderCode: "51",
		CycleCode:      "side_stories",
		Code:           "cotr",
	},
	{
		HaaBuilderCode: "52",
		CycleCode:      "side_stories",
		Code:           "coh",
	},
	// 53 -> La Cabale de Myarlathotep -> not in arkhamdb
	{
		HaaBuilderCode: "56",
		CycleCode:      "side_stories",
		Code:           "wog",
	},
	{
		HaaBuilderCode: "64",
		CycleCode:      "side_stories",
		Code:           "mtt",
	},
	{
		HaaBuilderCode: "67",
		CycleCode:      "side_stories",
		Code:           "fof",
	},
	// investigator
	{
		HaaBuilderCode: "45",
		CycleCode:      "investigator",
		Code:           "nat",
	},
	{
		HaaBuilderCode: "44",
		CycleCode:      "investigator",
		Code:           "har",
	},
	{
		HaaBuilderCode: "43",
		CycleCode:      "investigator",
		Code:           "win",
	},
	{
		HaaBuilderCode: "42",
		CycleCode:      "investigator",
		Code:           "jac",
	},
	{
		HaaBuilderCode: "41",
		CycleCode:      "investigator",
		Code:           "ste",
	},
}
