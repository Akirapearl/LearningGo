package main

// Requires to run as go run ./cmd/web (from root of the project)
import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// dependencies injection - 3.3
type application struct {
	logger *slog.Logger // more advanced way to log using log/slog
}

func main() {
	// Define a command-line flag with default value ":4000" w/some help text
	// Value will be stored in the addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Parse command-line flag
	// Reads value and assings it to addr variable
	// otherwise - always contain default value :4000
	// if errror - application will be terminated
	flag.Parse()

	// Use slog.New() -- Initialize a new structured logger
	// writes to the standard out stream & uses the default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// after (os.Stdout, &slog.HandlerOptions{
	//		Level: slog.LevelDebug, -- Minimum set by default avoids Debug
	//		AddSource: true, -- Include filename & line number of the calling source code
	//
	//}

	// to output logs to a file by using the tty - go run ./cmd/web >>/tmp/web.log

	// Initialize a new instance of the struct - containing dependencies (just the logger for now)
	app := &application{
		logger: logger,
	}

	// Less complete version for logging - log.Printf("starting server on %s", *addr)
	// use of slog.Any allows to avoid any problems at compile time related to
	// forgotten values i.e *addr, which would use !BADKEY

	logger.Info("starting server...", "addr", *addr)

	// Call the new app.routes() method to get the servemux containing the routes
	// and pass it to htpp.ListenandServe().
	err := http.ListenAndServe(*addr, app.routes())
	//log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1)
}
