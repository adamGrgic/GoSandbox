package main
// objective: create a method that deploys go routines as agents that accept
// requests and schedules different tasks
import(
    "fmt"
    "net/http"
)


func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// isValidToken := auth.ValidateAuthorizationToken(r.Header.Get("Authorization"))
		// if !isValidToken {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }
		next.ServeHTTP(w, r) // Call the next handler if authorized
	})
}

func PingSystem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG FROM AGENT"))
}

func CreateAgent(w http.ResponseWriter, r *http.Request) {
    fmt.Println("creating agent...")






    w.Write([]byte("Agent Created"))
}

func addCORSHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust the origin as needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight (OPTIONS) request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in the chain
		handler.ServeHTTP(w, r)
	})
}


func setupSystemRouter() http.Handler {
	mux := http.NewServeMux()

	// Add your handlers to the mux
	mux.Handle("/", AuthenticationMiddleware(http.HandlerFunc(PingSystem)))

	// Wrap the mux with the CORS middleware
	return addCORSHeaders(mux)
}

func main() {
    fmt.Println("starting up agent application...")

	// create a location or server where readings can be retrieved from
	port := "8080"
	customHandler := setupSystemRouter()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: customHandler,
	}

	go func() {
		fmt.Println("Starting Server on port " + port)
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	select {}


}
