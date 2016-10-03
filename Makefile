test:
	go test ./...

serve:
	go run main.go template.go

watch:
	go install github.com/jmhodges/justrun
	justrun -c 'make serve' main.go template.go
