start:
	@go build -o boodbood .
	@echo "copying files..."
	@cp boodbood.sh ~/.config/boodbood/
	@echo "moving binary to /usr/local/bin/"
	@mv boodbood /usr/local/bin/
	@echo "Done!"
