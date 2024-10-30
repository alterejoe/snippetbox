run:
	templ generate
	go run ./cmd/web/ --uniqueid=$(UUID)
	
