package binding

var (
	JSON     = jsonBinding{}
	Form     = formBinding{}
	MutiForm = formMultipartBinding{}
	Query    = queryBinding{}
	Header   = headerBinding{}
)
