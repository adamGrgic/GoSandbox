package main

import(
    "fmt"
    "log"
    "bytes"
    "sync"
    "golang.org/x/crypto/ssh"
    "bufio"
)

type SSHClient struct {
    Host string
    Username string
    Password string
}

func sshToServer(client SSHClient, wg *sync.WaitGroup, output chan<- string) {
    defer wg.Done()

    config := &ssh.ClientConfig{
		User: client.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(client.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

    connection, err := ssh.Dial("tcp", client.Host, config)
	if err != nil {
		output <- fmt.Sprintf("Error connecting to %s: %v", client.Host, err)
		return
	}
	defer connection.Close()

	// Open a new session
	session, err := connection.NewSession()
	if err != nil {
		output <- fmt.Sprintf("Error creating session for %s: %v", client.Host, err)
		return
	}
	defer session.Close()

	// Run the command
	stdout, err := session.StdoutPipe()
	if err != nil {
		output <- fmt.Sprintf("Error getting stdout for %s: %v", client.Host, err)
		return
	}

	// Start the command execution
    cmd := fmt.Sprintf("echo hello from %s", client.Host)
	if err := session.Start(cmd); err != nil {
		output <- fmt.Sprintf("Error running command on %s: %v", client.Host, err)
		return
	}

	// Read the output
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		output <- fmt.Sprintf("[%s]: %s", client.Host, scanner.Text())
	}

	// Wait for the command to complete
	if err := session.Wait(); err != nil {
		output <- fmt.Sprintf("Error waiting for command on %s: %v", client.Host, err)
	}
}

func main() {
    fmt.Println("Running ping pong module!")

    // the purpose of this module is to help develop my skills in network
    // programming with Go

    // boilerplate (? - not sure exactly what I want to put here yet)


    // (X) LEVEL 0: create an ssh connection to a server and return a list of files
    // at the root directory

    // Level 1: Create a command (from level 0?) that can connect to and run a script on an Ubuntu server to startup

    fmt.Println("Creating SSH client configuration object")
    // Define the SSH client configuration
    // + create an ssh pointer
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

    // Feels like I need a better understanding of using the defer keyword
    // because I don't know understand at what point the connection actually
    // closes. Because here we close the client but then we use client again
    // so I don't think I understand what's going on here.

    // ok so yeah it does execute after the OUTTER function is done executing
	defer client.Close()

	// Open a session for running a command
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
    defer session.Close()

    // Use either A or B

    // A
    //  (X)todo: understand defer keyword better
    //	defer session.Close()
    //  ... essentialy runs last after the outside function has been run

    // B
    // Capture command output
    var stdout, stderr bytes.Buffer
    session.Stdout = &stdout
    session.Stderr = &stderr
    // end
    //


    // Info on executing multiple commands successively:
    // https://chatgpt.com/c/67818c07-ce40-8004-8e77-14fc2a1115ac
	// Execute a remote command
	err = session.Run("echo I am a message from 192.168.0.26") // Replace "ls -l" with your command
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

    fmt.Printf("Output of command 1:\n%s\n", stdout.String())


    // level 0.1: SSH into a list of servers and evaluate the responses for
    // successful and non-successful connections (display a table in the
    // terminal, perhaps?)
    //  +  this is the initialization of your module

        // 192.168.0.26
        // 192.168.0.28
        // 192.168.0.32
        // 192.168.0.35

    // todo: add this to config file
    clients := []SSHClient{
        {"192.168.0.26:22", "adam", "Hockey7232!"},
        {"192.168.0.28:22", "adam", "Hockey7232!"},
        {"192.168.0.32:22", "adam", "Hockey7232!"},
        {"192.168.0.35:22", "adam", "Hockey7232!"},
    }

    output := make(chan string)
    var wg sync.WaitGroup

    for _, client := range clients  {
        wg.Add(1)
        go sshToServer(client, &wg, output)
    }

    // Close the output channel when all sessions are done
	go func() {
		wg.Wait()
		close(output)
	}()

	// Print output from the channel
	for msg := range output {
		fmt.Println(msg)
	}






    // Level 2: Write responses from each connection to individual log files

    // Level 3: Simulate signal interruptions and how that would be handled

    // Level 4:


    // Level 5:

    fmt.Println("Program end")

}
