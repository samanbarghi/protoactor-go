package actor

func makeMiddlewareChain(middleware []func(ReceiveFunc) ReceiveFunc, actorReceiver ReceiveFunc) ReceiveFunc {
	if len(middleware) == 0 {
		return nil
	}

	h := middleware[len(middleware)-1](actorReceiver)
	for i := len(middleware) - 2; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}
