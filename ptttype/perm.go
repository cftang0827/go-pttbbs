package ptttype

type PERM uint32

const (
	PERM_INVALID       PERM = 000000000000
	PERM_BASIC         PERM = 000000000001 /* 基本權力       */
	PERM_CHAT          PERM = 000000000002 /* 進入聊天室     */
	PERM_PAGE          PERM = 000000000004 /* 找人聊天       */
	PERM_POST          PERM = 000000000010 /* 發表文章       */
	PERM_LOGINOK       PERM = 000000000020 /* 註冊程序認證   */
	PERM_MAILLIMIT     PERM = 000000000040 /* 信件無上限     */
	PERM_CLOAK         PERM = 000000000100 /* 目前隱形中     */
	PERM_SEECLOAK      PERM = 000000000200 /* 看見忍者       */
	PERM_XEMPT         PERM = 000000000400 /* 永久保留帳號   */
	PERM_SYSOPHIDE     PERM = 000000001000 /* 站長隱身術     */
	PERM_BM            PERM = 000000002000 /* 板主           */
	PERM_ACCOUNTS      PERM = 000000004000 /* 帳號總管       */
	PERM_CHATROOM      PERM = 000000010000 /* 聊天室總管     */
	PERM_BOARD         PERM = 000000020000 /* 看板總管       */
	PERM_SYSOP         PERM = 000000040000 /* 站長           */
	PERM_BBSADM        PERM = 000000100000 /* BBSADM         */
	PERM_NOTOP         PERM = 000000200000 /* 不列入排行榜   */
	PERM_VIOLATELAW    PERM = 000000400000 /* 違法通緝中     */
	PERM_ANGEL         PERM = 000001000000 /* 有資格擔任小天使 */
	PERM_NOREGCODE     PERM = 000002000000 /* 不允許認證碼註冊 */
	PERM_VIEWSYSOP     PERM = 000004000000 /* 視覺站長       */
	PERM_LOGUSER       PERM = 000010000000 /* 觀察使用者行蹤 */
	PERM_NOCITIZEN     PERM = 000020000000 /* 搋奪公權       */
	PERM_SYSSUPERSUBOP PERM = 000040000000 /* 群組長         */
	PERM_ACCTREG       PERM = 000100000000 /* 帳號審核組     */
	PERM_PRG           PERM = 000200000000 /* 程式組         */
	PERM_ACTION        PERM = 000400000000 /* 活動組         */
	PERM_PAINT         PERM = 001000000000 /* 美工組         */
	PERM_POLICE_MAN    PERM = 002000000000 /* 警察總管       */
	PERM_SYSSUBOP      PERM = 004000000000 /* 小組長         */
	PERM_OLDSYSOP      PERM = 010000000000 /* 退休站長       */
	PERM_POLICE        PERM = 020000000000 /* 警察 */
	// 32 個已經全部用光了。 後面沒有了。
)

const (
	NUMPERMS = 32
)

const (
	PERM_DEFAULT    PERM = PERM_BASIC | PERM_CHAT | PERM_PAGE
	PERM_MANAGER    PERM = PERM_ACCTREG | PERM_ACTION | PERM_PAINT
	PERM_ADMIN      PERM = PERM_ACCOUNTS | PERM_BOARD | PERM_SYSOP | PERM_SYSSUBOP | PERM_SYSSUPERSUBOP | PERM_MANAGER
	PERM_LOGINCLOAK PERM = PERM_SYSOP | PERM_ACCOUNTS
	PERM_SEEULEVELS PERM = PERM_SYSOP
	PERM_SEEBLEVELS PERM = PERM_SYSOP | PERM_BM
	PERM_NOTIMEOUT  PERM = PERM_SYSOP
	PERM_READMAIL   PERM = PERM_BASIC
	PERM_FORWARD    PERM = PERM_LOGINOK /* to do the forwarding */
	PERM_INTERNET   PERM = PERM_LOGINOK /* 身份認證過關的才能寄信到 Internet */
)

func (p PERM) HasUserPerm(perm PERM) bool {
	return p&perm != 0
}

func (p PERM) HasBasicUserPerm(perm PERM) bool {
	return p.HasUserPerm(PERM_BASIC) && p.HasUserPerm(perm)
}

func (p PERM) Hide() bool {
	return p == PERM_SYSOPHIDE
}
