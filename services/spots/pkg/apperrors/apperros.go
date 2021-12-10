package apperrors

import "github.com/matiasvarela/errors"

var (
	NotFound                   = errors.Define("not_found")
	InvalidInput               = errors.Define(("invalid_input"))
	Internal                   = errors.Define("internal")
	AlreadyHostingAOnlineEvent = errors.Define("already_hosting_a_online_event")
)
