func main() {
	err := err2()
	fmt.Println(err)
}

func err2() error {
	err := err1()
	if err != nil {
		return stackerr.Wrap(err)
	}
	return nil
}

func err1() error {
	return stackerr.Wrap(errors.New("failure"))
}
