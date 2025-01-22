package main

// @@@@@@@@@@ WIP @@@@@@@@@@@@@@
// go through basic code samples for the Go http/net package

// TODOS:
// add non configurable server examples
// add middleware

import (
	"fmt"
	"net/http"
)

// var mongoClient *mongo.Client
// var once sync.Once

// func getMongoClient() (*mongo.Client, error) {
// 	var err error
// 	once.Do(func() {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()
// 		clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
// 		mongoClient, err = mongo.Connect(ctx, clientOpts)
// 		if err != nil {
// 			log.Println("Error connecting to MongoDB:", err)
// 		}
// 	})
// 	return mongoClient, err
// }

func PingSystem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG"))
}

func FooSystem(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("FOO DOG YO"))
}

func ValidateAuthorizationToken(token string) bool {
	fmt.Println("Validating Auth Token")

	// client, err := getMongoClient()
	// if err != nil {
	// 	log.Println("Failed to get MongoDB client:", err)
	// 	return false
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// Access the database and collection using environment variables
	// databaseName := os.Getenv("MONGODB_DB")
	// collectionName := os.Getenv("MONGODB_MANIFEST")
	// if databaseName == "" || collectionName == "" {
	// 	log.Println("Environment variables MONGODB_DB or MONGODB_MANIFEST are not set")
	// 	return false
	// }

	// collection := client.Database(databaseName).Collection(collectionName)
	// filter := bson.M{"id": token}

	// var result bson.M
	// err = collection.FindOne(ctx, filter).Decode(&result)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		fmt.Println("No document found with the given client-id")
	// 		return false
	// 	}
	// 	log.Println("Error finding document:", err)
	// 	return false
	// }

	// fmt.Println("Document found:", result)
	return true
}

// CORS middleware function
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

func setupSystemRouter() http.Handler {
	mux := http.NewServeMux()

	// Add your handlers to the mux
	mux.Handle("/ping", AuthenticationMiddleware(http.HandlerFunc(PingSystem)))
    mux.Handle("/foo", AuthenticationMiddleware(http.HandlerFunc(FooSystem)))
	// Wrap the mux with the CORS middleware
	return addCORSHeaders(mux)
}

func main() {
	fmt.Println("Creating Server ...")

	// create a location or server where readings can be retrieved from
	port := "43128"
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
