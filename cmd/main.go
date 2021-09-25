package main

func main() {
	srv := initHttpServer()
	srv.Logger.Fatal(srv.Start("localhost:8080"))
}
