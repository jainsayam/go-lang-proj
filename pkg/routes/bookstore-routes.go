package routes

import	"github.com/gorilla/mux"
import "github.com/jainsayam/go-bookstore/pkg/controllers"


var RegisterBookStoreRoutes=func (router *mux.Router)  {
router.HandleFunc("/book/",controllers.CreateBook).Methhods("POST")	
}