package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error
}

// Error えらーを文字列に変換する
func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// Unwrap エラーをラップしているエラーを返す
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
