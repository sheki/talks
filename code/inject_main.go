func main() {

	flag.Parse()
	mgoSession, _ := mgo.Dial(*mgoAddrFlag)

	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: scribe.NewHTTPScribeClient()},
		&inject.Object{Value: mgoSession},
		&inject.Object{Value: parse.NewLogger()},
		&inject.Object{Value: parse.EnvFromFlag(*envFlag)},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// rest of main
}
