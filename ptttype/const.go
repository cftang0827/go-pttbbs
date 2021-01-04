package ptttype

import "os"

const (
	//////////
	//pttstruct.h
	//
	//SHMSIZE is computed in cache/shm.go NewSHM
	//////////
	IDLEN   = 12 /* Length of bid/uid */
	IPV4LEN = 15 /* a.b.c.d form */

	PASS_INPUT_LEN = 8 /* Length of valid input password length.
	   For DES, set to 8. */
	PASSLEN = 14 /* Length of encrypted passwd field */
	REGLEN  = 38 /* Length of registration data */

	REALNAMESZ = 20 /* Size of real-name field */
	NICKNAMESZ = 24 /* SIze of nick-name field */
	EMAILSZ    = 50 /* Size of email field */
	ADDRESSSZ  = 50 /* Size of address field */
	CAREERSZ   = 40 /* Size of career field */
	PHONESZ    = 20 /* Size of phone field */

	USERNAMESZ = 24 /* Size of Username in MailQueue */
	RCPTSZ     = 50 /* Size of RCPT in MailQueue */

	PASSWD_VERSION = 4194

	BTLEN = 48 /* Length of board title */

	TTLEN = 64 /* Length of title */
	FNLEN = 28 /* Length of filename */

	// Prefix of str_reply and str_forward usually needs longer.
	// In artitlcle list, the normal length is (80-34)=46.
	DISP_TTLEN = 46

	STRLEN = 80 /* Length of most string data */

	FAVMAX   = 1024 /* Max boards of Myfavorite */
	FAVGMAX  = 32   /* Max groups of Myfavorite */
	FAVGSLEN = 8    /* Max Length of Description String */

	ONEKEY_SIZE = int('~')

	ANSILINELEN = 511 /* Maximum Screen width in chars */

	/* USHM_SIZE 比 MAX_ACTIVE 大是為了防止檢查人數上限時, 又同時衝進來
	 * 會造成找 shm 空位的無窮迴圈.
	 * 又, 因 USHM 中用 hash, 空間稍大時效率較好. */
	USHM_SIZE = MAX_ACTIVE * 41 / 40

	MAX_BMs = 4
)

const (

	//mbbsd/register.c line: 415
	CLEAN_USER_EXPIRE_RANGE_MIN = 365 * 12 * 60 // 180 days.

	//mbbsd/user.c line: 42
	DIR_TMP   = "tmp"
	DIR_HOME  = "home"
	DIR_BOARD = "boards"
)

const (
	//mbbsd/board.c line: 1477
	USE_REAL_DESC_FOR_HIDDEN_BOARD_IN_MYFAV = false
)

var (
	//mbbsd/register.c line: 381
	FN_FRESH_POSTFIX = ".fresh"
	FN_FRESH         = BBSHOME + string(os.PathSeparator) + FN_FRESH_POSTFIX
)

var (
	//ptt/article_list.go
	N_SCREEN_BUFFER = 100 //to determine whether we should provide bottom.
)

var (
	ALOHA_MSG = []byte{ //<<上站通知>> -- 我從山中來～
		0x3c, 0x3c, 0xa4, 0x57, 0xaf, 0xb8, 0xb3, 0x71,
		0xaa, 0xbe, 0x3e, 0x3e, 0x20, 0x2d, 0x2d, 0x20,
		0xa7, 0xda, 0xb1, 0x71, 0xa4, 0x73, 0xa4, 0xa4,
		0xa8, 0xd3, 0xa1, 0xe3,
	}
)
