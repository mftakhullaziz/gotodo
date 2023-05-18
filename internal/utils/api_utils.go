package utils

// ListEndpoints function to show all registered apis
func ListEndpoints(endpoints []string) {
	log := LoggerParent()
	log.Infoln("Registered endpoints:")
	for _, e := range endpoints {
		log.Infoln(e)
	}
}
