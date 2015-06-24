type ScribeClient {
	Thrift	*ThriftPool `inject:""`
}

// ScribeClient start will be called after ThriftPool.Start
func (s *ScribeClient) Start() error {
	fmt.Println("starting scribe client")
	return nil
}

type ThriftPool struct {
	// 
}

// ThriftPool start will be called first.
func (t *ThriftPool) Start() error {
	fmt.Println("starting thrift pool")
	return t.tcpDial()
}

func (t *ThriftPool) Stop() error {
	fmt.Println("stopping thrift pool")
	return t.tcpCloseAll()
}
