package error

func ErrorPanic(err error){
	if err != nil {
		panic(err)
	}
}