package cerror

func WithCodeAndMessage(err error, code int, msg string) error {
	return WithStatusCode(WithUserMessage(err, msg), code)
}
