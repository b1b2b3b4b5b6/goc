package errtool

//Errpanic is
func Errpanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//Assert is
func Assert(val bool) {
	if !val {
		panic("assert fail")
	}
}

func Ignore(err error) {

}
