package api

import (
	"net/http"

	"github.com/RomuloDurante/Go_REST/api/controller"
)

// Get ...
func Get(w http.ResponseWriter, r *http.Request, c *controller.Controller) (err error) {

	data, err := c.GetItem()
	if err != nil {
		return err
	}
	// set content type as JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.Write(data)

	return nil
}
