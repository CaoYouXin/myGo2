package files

const MAIN = `func main() {

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.InitRouter(gin.Default()),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen and Serve: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown: %v\n", err)
	}
	log.Println("Server Existing In 5 Seconds.")
}`
