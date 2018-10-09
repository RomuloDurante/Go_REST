package api

import (
	"net/http"

	"github.com/RomuloDurante/Go_REST/api/controller"
)

//Post ...
func Post(w http.ResponseWriter, r *http.Request, c *controller.Controller) error {

	data, err := c.ReadData()
	if err != nil {
		return err
	}

	body, err := c.GetBody(r)
	if err != nil {
		return err
	}

	msg, err := c.CreateItem(data, body)

	if err != nil {
		return err
	}

	w.Write(msg)
	return nil
}
