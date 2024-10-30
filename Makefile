run:
	templ generate
	go run ./cmd/web/ -addr=":3000"  --unique-id=$(UUID) 
	
