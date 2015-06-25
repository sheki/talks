func main() {

	flag.Parse()
	mgoSession, _ := mgo.Dial(*mgoAddrFlag)

	ship.SimpleStart(
		&inject.Object{Value: scribe.NewHTTPScribeClient()},
		&inject.Object{Value: mgoSession},
		&inject.Object{Value: parse.NewLogger()},
		&inject.Object{Value: parse.EnvFromFlag(*envFlag)},
	)

	defer ship.Stop()
	....
	// rest of main
}
