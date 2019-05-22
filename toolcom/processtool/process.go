package processtool

type retryFunc func() bool

//RetryUntilTrue is
func RetryUntilTrue(f retryFunc, times int) bool {
	for i := 0; i < times; i++ {
		if f() == true {
			return true
		}
	}
	return false
}

//RetryUntilFalse is
func RetryUntilFalse(f retryFunc, times int) bool {
	for i := 0; i < times; i++ {
		if f() == false {
			return false
		}
	}
	return true
}
