package consts

const (
	USER_SESSION_KEY = "user"
	CAS_TICKET       = "ticket"
	CAS_SERVICE      = "service"
	start  int8 = 1 << iota
	resume
	stop
)
