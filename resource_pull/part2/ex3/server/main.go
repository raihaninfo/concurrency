package main

func main() {
	srv := newTCPServer("8081")
	srv.Start()
}
