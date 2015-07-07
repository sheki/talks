func main() {

	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: scribe.NewHTTPScribeClient()},
		&inject.Object{Value: parse.NewLogger()},
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
