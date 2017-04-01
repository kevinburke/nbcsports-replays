test:
	go test ./...

serve:
	go run main/main.go

watch:
	go install github.com/jmhodges/justrun
	justrun -c 'make serve' main.go template.go
