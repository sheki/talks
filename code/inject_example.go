// Handler serves a HTTP request
type Handler struct {
	Env    env.ServerEnv  `inject:""`
	Scribe *scribe.Client `inject:""`
	Log    logger.Logger  `inject:""`
	Mongo  *mgo.Session   `inject:"secondary"`
}

// ServeHTTP a sample implementation
func (h *Handler) ServeHTTP(w ResponseWriter, r *Request) {
	params := extractParams(r)
	res := mongoFetch(h.Mongo, params)
	h.Scribe.Log(params)
	w.Write(res)
}


