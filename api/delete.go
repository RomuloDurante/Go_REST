package api

import (
	"net/http"
	"github.com/RomuloDurante/Go_REST/api/controller"
)

//Delete ...
func Delete(w http.ResponseWriter, r *http.Request, c *controller.Controller) error {

	// data, err := c.GetAllObj(nil)

	// if err != nil {
	// 	return err
	// }

	// data, err = c.GetAllObj(data)
	// if err != nil {
	// 	return err
	// }

	// c.CreateData(data)

	w.Write([]byte("Data was deleted."))

	return nil
}
