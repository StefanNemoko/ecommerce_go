package helpers

import (
	"strconv"
	"strings"
)

func RetrieveIdFromUri(uri string) (int64, error) {
	// the ID should always be at the end of the uri.
	uriParts := strings.Split(uri, "/")

	id := uriParts[len(uriParts)-1]

	// Remove any query parameters from uri
	strippedUri := strings.Split(id, "?")[0]

	//Parse to int64 and return with errors (if any)
	return strconv.ParseInt(strippedUri, 0, 64)
}
