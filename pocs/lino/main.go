package main

// todo: figure out how to get up down keys working again

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	status int
	err    error
	foo    string

	cursor      int
	taskqueue   []queue
	menuoptions []MenuItem
}

type queue struct {
	job         string
	description string
}

var url = "http://localhost:43128/foo"

func checkServer() tea.Msg {

	// Create an HTTP client and make a GET request.
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)

	if err != nil {
		// There was an error making our request. Wrap the error we received
		// in a message and return it.
		return errMsg{err}
	}
	// We received a response from the server. Return the HTTP status code
	// as a message.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return errMsg{err}
	}

	log.Println(string(body))
	return statusMsg(res.StatusCode)
}

type statusMsg int

type errMsg struct{ err error }

type foo string

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e errMsg) Error() string { return e.err.Error() }

func (m model) Init() tea.Cmd {
	log.Println("making moves")

	return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println("Update ran")

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			log.Println("enter was pressed")
			return m, nil
			// The "up" and "k" keys move the cursor up
		case "up", "k":
			log.Println("up was pressed")
			log.Println("debug: cursor %d", m.cursor)
			if m.cursor > 0 {
				log.Println("debug: cursor is greater than 0")
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			log.Println("down was pressed")
			log.Println("debug: cursor %d", m.cursor)
			if m.cursor < len(m.menuoptions)-1 {
				log.Println("debug: cursor is less than length of options")
				m.cursor++
			}

		}

	case statusMsg:
		// The server returned a status message. Save it to our model. Also
		// tell the Bubble Tea runtime we want to exit because we have nothing
		// else to do. We'll still be able to render a final view with our
		// status message.
		m.status = int(msg)
		return m, tea.Println("status message was given")

	case errMsg:
		// There was an error. Note it in the model. And tell the runtime
		// we're done and want to quit.
		m.err = msg
		return m, tea.Quit

	case foo:

		log.Println("dog")
		m.foo = string(msg)

		return m, tea.Println("Foo was ran")
	}

	return m, nil
}

type MenuItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (m model) View() string {
	// If there's an error, print it out and don't do anything else.
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
	}

	s := "what's your perogative?\n\n"

	menuoptions := []MenuItem{
		{Title: "Default Startup", Description: "Default startup method for your build"},
		{Title: "Theme 1", Description: "Another common theme"},
	}

	for i, choice := range menuoptions {
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	// When the server responds with a status, add it to the current line.
	// if m.status > 0 {
	//    s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	// }

	// Send off whatever we came up with above for rendering.
	return "\n" + s + "\n\n"
}

func main() {

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	startupFile, err := os.Open("theme-options.json")
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer startupFile.Close()

	bytes, err := io.ReadAll(startupFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	//menuoptions := []MenuItem{
	//    {Title : "Default Startup",Description : "Default startup method for your build"},
	//    {Title : "Theme 1", Description : "Another common theme"},
	//}

	var jsonmenuoptions []MenuItem
	if err := json.Unmarshal(bytes, &jsonmenuoptions); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	log.SetOutput(logFile)
	if _, err := tea.NewProgram(model{menuoptions: jsonmenuoptions}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
