package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testUserec         *bbs.Userec
	testBoardSummary6  *bbs.BoardSummary
	testBoardSummary7  *bbs.BoardSummary
	testBoardSummary8  *bbs.BoardSummary
	testBoardSummary11 *bbs.BoardSummary

	testArticleSummary0 *bbs.ArticleSummary
	testArticleSummary1 *bbs.ArticleSummary
	testContent1        []byte
)

func initTestVars() {
	if testBoardSummary6 != nil {
		return
	}

	testUserec = &bbs.Userec{
		Version:  4194,
		UUserID:  bbs.UUserID("SYSOP"),
		Username: "SYSOP",
		Realname: []byte{ // CodingMan
			0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e,
		},
		Nickname: []byte{0xaf, 0xab}, // 神

		Uflag:        33557088,
		Userlevel:    536871943,
		Numlogindays: 2,
		Numposts:     0,
		Firstlogin:   1600681288,
		Lastlogin:    1600756094,
		Lasthost:     "59.124.167.226",
		/*
			Address: []byte{ //新竹縣子虛鄉烏有村543號
				0xb7, 0x73, 0xa6, 0xcb, 0xbf, 0xa4, 0xa4, 0x6c, 0xb5, 0xea,
				0xb6, 0x6d, 0xaf, 0x51, 0xa6, 0xb3, 0xa7, 0xf8, 0x35, 0x34,
				0x33, 0xb8, 0xb9,
			},
		*/
		Over18:   true,
		Pager:    ptttype.PAGER_ON,
		Career:   []byte{0xa5, 0xfe, 0xb4, 0xba, 0xb3, 0x6e, 0xc5, 0xe9}, //全景軟體
		LastSeen: 1600681288,
	}

	testBoardSummary6 = &bbs.BoardSummary{
		BBoardID: bbs.BBoardID("6_ALLPOST"),
		BrdAttr:  ptttype.BRD_POSTMASK,
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "ALLPOST",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xb8, 0xf3, 0xaa, 0x4f, 0xa6, 0xa1, 0x4c, 0x4f, 0x43, 0x41,
			0x4c, 0xb7, 0x73, 0xa4, 0xe5, 0xb3, 0xb9,
		},
		BoardType: []byte{0xa1, 0xb7},
		BM:        []bbs.UUserID{},
	}

	testBoardSummary7 = &bbs.BoardSummary{
		BBoardID: bbs.BBoardID("7_deleted"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "deleted",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xb8, 0xea, 0xb7, 0xbd, 0xa6, 0x5e, 0xa6, 0xac, 0xb5, 0xa9,
		},
		BoardType: []byte{0xa1, 0xb7},
		BM:        []bbs.UUserID{},
	}

	testBoardSummary8 = &bbs.BoardSummary{
		BBoardID: bbs.BBoardID("8_Note"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "Note",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xb0, 0xca, 0xba, 0x41, 0xac, 0xdd, 0xaa, 0x4f, 0xa4, 0xce,
			0xba, 0x71, 0xa6, 0xb1, 0xa7, 0xeb, 0xbd, 0x5a,
		},
		BoardType: []byte{0xa1, 0xb7},
		BM:        []bbs.UUserID{},
	}

	testBoardSummary11 = &bbs.BoardSummary{
		BBoardID: bbs.BBoardID("11_EditExp"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "EditExp",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xbd, 0x64, 0xa5, 0xbb, 0xba, 0xeb, 0xc6, 0x46, 0xa7, 0xeb,
			0xbd, 0x5a, 0xb0, 0xcf,
		},
		BoardType: []byte{0xa1, 0xb7},
		BM:        []bbs.UUserID{},
	}

	testArticleSummary0 = &bbs.ArticleSummary{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  "1Vo_M_CDSYSOP",
		IsDeleted:  false,
		Filename:   "M.1607202239.A.30D",
		CreateTime: 1607202239,
		MTime:      1607202238,
		Owner:      "SYSOP",
		Title: []byte{
			0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0x20, 0xa7,
			0xda, 0xac, 0x4f, 0xbd, 0xd6, 0xa1, 0x48, 0xa1,
			0xe3,
		},
		Class: []byte{0xb0, 0xdd, 0xc3, 0x44},
	}

	testArticleSummary1 = &bbs.ArticleSummary{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  "1Vo_f30DSYSOP",
		IsDeleted:  false,
		Filename:   "M.1607203395.A.00D",
		CreateTime: 1607203395,
		MTime:      1607203394,
		Owner:      "SYSOP",
		Title: []byte{
			0x5b, 0xa4, 0xdf, 0xb1, 0x6f, 0x5d, 0x20, 0xb5,
			0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1,
			0xe3,
		},

		Filemode: ptttype.FILE_MARKED,

		Class: []byte{0xa4, 0xdf, 0xb1, 0x6f},
	}

	testContent1 = []byte{
		0xa7, 0x40, 0xaa, 0xcc, 0x3a, 0x20, 0x53, 0x59, 0x53,
		0x4f, 0x50, 0x20, 0x28, 0x29, 0x20, 0xac, 0xdd, 0xaa,
		0x4f, 0x3a, 0x20, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49,
		0x0a, 0xbc, 0xd0, 0xc3, 0x44, 0x3a, 0x20, 0x5b, 0xb0,
		0xdd, 0xc3, 0x44, 0x5d, 0x20, 0xa7, 0xda, 0xac, 0x4f,
		0xbd, 0xd6, 0xa1, 0x48, 0xa1, 0xe3, 0x0a, 0xae, 0xc9,
		0xb6, 0xa1, 0x3a, 0x20, 0x53, 0x75, 0x6e, 0x20, 0x44,
		0x65, 0x63, 0x20, 0x20, 0x36, 0x20, 0x30, 0x35, 0x3a,
		0x30, 0x33, 0x3a, 0x35, 0x37, 0x20, 0x32, 0x30, 0x32,
		0x30, 0x0a, 0x0a, 0xa7, 0xda, 0xac, 0x4f, 0xbd, 0xd6,
		0xa1, 0x48, 0xa1, 0xe3, 0x0a, 0x0a, 0xa7, 0xda, 0xa6,
		0x62, 0xad, 0xfe, 0xb8, 0xcc, 0xa1, 0x48, 0xa1, 0xe3,
		0x0a, 0x0a, 0xa7, 0xda, 0xac, 0xb0, 0xa4, 0xb0, 0xbb,
		0xf2, 0xb7, 0x7c, 0xa6, 0x62, 0xb3, 0x6f, 0xb8, 0xcc,
		0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3, 0x0a, 0x0a, 0x2d,
		0x2d, 0x0a, 0xa1, 0xb0, 0x20, 0xb5, 0x6f, 0xab, 0x48,
		0xaf, 0xb8, 0x3a, 0x20, 0xa7, 0xe5, 0xbd, 0xf0, 0xbd,
		0xf0, 0x20, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x28,
		0x70, 0x74, 0x74, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
		0x2e, 0x74, 0x65, 0x73, 0x74, 0x29, 0x2c, 0x20, 0xa8,
		0xd3, 0xa6, 0xdb, 0x3a, 0x20, 0x31, 0x37, 0x32, 0x2e,
		0x31, 0x38, 0x2e, 0x30, 0x2e, 0x31, 0x0a,
	}
}
