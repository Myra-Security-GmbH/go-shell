package event

//
// CloneSubscriber ...
//
func CloneSubscriber(s Subscriber) Subscriber {
	newHandler := make(map[string][]Func)

	switch h := s.(type) {
	case *GenericSubscriber:
		for key, ll := range h.handler {
			newHandler[key] = []Func{}

			for _, v := range ll {
				newHandler[key] = append(newHandler[key], v)
			}
		}

	}

	return &GenericSubscriber{
		handler: newHandler,
	}
}
