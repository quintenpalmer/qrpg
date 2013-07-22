package engine

type QrpgError struct {
	message string
}

func NewQrpgError(message string) QrpgError {
	return QrpgError{message}
}

func (enginee QrpgError) Error() string {
	return enginee.message
}
