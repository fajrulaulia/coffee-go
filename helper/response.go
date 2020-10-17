package coffeego

import (
	core "coffeego/helper/core"
	"net/http"
)

// ReponseOnServerError Should be Exported
func ReponseOnServerError(w http.ResponseWriter) {
	w.WriteHeader(500)
	output := core.Response([]string{"code:number|500", "message:string|Internal Server Error, Please contact Administrator"})
	w.Write(output)
}
