package enums

type emailStatus int

const (
	emailBounced emailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)
