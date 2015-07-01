type Handler struct {
	Scribe *scribe.Client `inject:""`
	Log    logger.Logger  `inject:""`
}

// ServeHTTP a sample implementation
func (h *Handler) ServeHTTP(w ResponseWriter, r *Request) {
	params := extractParams(r)
	h.Scribe.Log(params)
	h.Log("everything ok")
	w.Write(res)
}
