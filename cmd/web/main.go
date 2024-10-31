package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"sync"
)

type application struct {
	logger *slog.Logger
}

var mu sync.Mutex

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := path + "/index.html"
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}

	return f, nil
}

func main() {
	addr := flag.String(("addr"), ":3000", "HTTP network address")
	_ = flag.String("unique-id", "", "Unique ID for development")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		// AddSource: true,
	}))

	app := application{
		logger: logger,
	}

	logger.Info("Starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error((err.Error()))
	os.Exit(1)
}
