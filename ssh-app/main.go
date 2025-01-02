package main

import(
    "fmt"
    "log"
    "golang.org/x/crypto/ssh"
)

func main() {
    fmt.Println("running ssh app")

    // Define the SSH client configuration
	config := &ssh.ClientConfig{
		User: "adam", // Replace with the SSH username
		Auth: []ssh.AuthMethod{
			ssh.Password("Hockey7232!"), // Replace with the SSH password or use a private key
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: Not secure for production
	}

	// Replace with the target SSH server and port (e.g., "example.com:22")
    server := "192.168.0.26:22"

	// Establish the connection
	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	// Open a session for running a command
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	// Execute a remote command
	output, err := session.CombinedOutput("ls -l") // Replace "ls -l" with your command
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	// Print the command output
	fmt.Println(string(output))

}
