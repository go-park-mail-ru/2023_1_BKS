type HttpServer struct {
	app app.Application
}

func CreateHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}