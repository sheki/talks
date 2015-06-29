type Harness struct {
	*testship.FatHarness `inject:""`
	Log                  logger.Logger        `inject:""`
	Secrets              *config.ParseSecrets `inject:""`
	Mongo                *mongo.Service       `inject:""`
	CacheService         *cache.Service       `inject:""`
	Cassandra            cassandra.Service    `inject:""`
}

// NewHarness creates a test harness
func NewHarness(t *testing.T) *Harness {
	testship.Fat(
		t,
		h,
		&facebook.TestClient{},
		&scribe.ScribeHiveLogger{},
		&twitter.TestClient{},
		mockV8.Config,
		redis.NewTestService(&FakeRedis{}),
	)
	return h
}
