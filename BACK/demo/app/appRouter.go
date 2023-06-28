package app

import "net/http"

func enableCore(reponseWriter *http.ResponseWriter){
	(*reponseWriter).Header().Set("Access-control-allow-Methods","*")
	(*reponseWriter).Header().Set("Access-control-allow-Origin","*")
	(*reponseWriter).Header().Set("Access-control-allow-Headers","*")
}

func (app Application) Setup() {
	app.datastore.Setup()
	app.Router.Use(app.route)
	app.Router.HandleFunc("/user/save",app.SaveUser).Methods(http.MethodPost, http.MethodOptions)
	app.Router.HandleFunc("/user/validate",app.ValidateUser).Methods(http.MethodPost, http.MethodOptions)
	app.Router.HandleFunc("/user/retrieveUsers",app.RetrieveUsers).Methods(http.MethodGet,http.MethodOptions)
	app.Router.HandleFunc("/user/delete/{id}",app.DeleteUser).Methods(http.MethodDelete,http.MethodOptions)
	app.Router.HandleFunc("/user/update",app.UpdateUser).Methods(http.MethodPut,http.MethodOptions)

}

func (app Application) route(next http.Handler)  http.Handler{
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		enableCore(&responseWriter)
		if (*request).Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(responseWriter, request)


	})
}