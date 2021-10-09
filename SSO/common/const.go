package common

import "time"

const (
	SignTypePhonePassword = "phone"
	SignTypePhoneSms      = "sms"
	SignTypeEmailPassword = "email"
	SignTypeLark          = "lark"
)

const (
	CASErrInvalidRequest    = "INVALID_REQUEST"
	CASErrInvalidTicketSpec = "INVALID_TICKET_SPEC"
	CASErrInvalidTicket     = "INVALID_TICKET"
	CASErrInvalidService    = "INVALID_SERVICE"
	CASErrInternalError     = "INTERNAL_ERROR"
	CASErrUnauthorized      = "UNAUTHENTICATED"
)

const (
	DebugMode = "debug"
)

const (
	SESSION_NAME = "UNIQUE_SESSION"
	// CAS_COOKIE_NAME    = "CASTGC"
	// CAS_TGT_EXPIRES    = time.Hour
	// CAS_TICKET_EXPIRES = time.Minute * 5
	// DEFAULT_TIMEOUT    = 10000000

	SMS_CODE_EXPIRES = time.Minute * 3
)

const (
	SMS_TEMPO_CODE = "verificationCode"
)

const (
	LARK_OAUTH = "https://open.larksuite.com/open-apis/authen/v1/index?redirect_uri=%s&app_id=%s&state=%s"
)
