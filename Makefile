start:
	@go build -o boodbood .
	@cp boodbood.sh ~/.config/boodbood/
	@mv boodbood /usr/local/bin/
	@echo "Done!"
