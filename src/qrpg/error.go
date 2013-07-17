package qrpg

type QrpgError struct {
	message string
}

func NewQrpgError(message string) QrpgError {
	return QrpgError{message}
}

func (qrpge QrpgError) Error() string {
	return qrpge.message
}
