.PHONY : format install build

#Website/Restfull API Build and Run
run:
	@echo "Running server..."
	go run main.go

init:
	@echo "Initializing dependencies..."
	go mod init
	go mod tidy
	
install:
	@echo "Downloading dependencies..."
	go mod download

build:
	@echo "building binary..."
	go build main.go

start:
	@echo "Starting server..."
	./main

clean:
	@echo "Cleaning..."
	rm -rf main.exe
# live reload using nodemon: npm -g i nodemon
auto:
	@echo "Running server with nodemon..."
	nodemon --exec go run main.go

sol:
	@echo "Make Blockchain Script..."
	solcjs --optimize --abi ./app/contracts/Smartsign.sol -o build
	solcjs --optimize --bin ./app/contracts/Smartsign.sol -o build
	abigen --abi=./build/app_contracts_Smartsign_sol_Smartsign.abi --bin=./build/app_contracts_Smartsign_sol_Smartsign.bin --pkg=api --out=./app/contracts/Smartsign.go