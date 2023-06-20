package utils

// ListEndpoints function to show all registered api
func ListEndpoints(endpoints []string) {
	log := LoggerParent().Log
	log.Infoln("URL Endpoints:")
	for _, e := range endpoints {
		log.Infoln(e)
	}
}
