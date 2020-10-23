# coffee-go
coffee-go written by golang and integrated by mongo-db, this is a standarization
for development API service using golang and mongodb.

## Setup for Development
you'r system should be installed depedencies like `go`, `make`, `docker-engine` and `docker-compose`

```bash
# Clone project and change directory to project folder
$ git clone https://github.com/fajrulaulia/coffee-go.git && cd coffee-go/

# build and run mongodb service
$ make service-up 

# Run Application on Development Mode 
$ go run cmd/main.go

```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.


## Author
- Fajrul Aulia


## License
[apache-2.0](https://choosealicense.com/licenses/https://choosealicense.com/licenses/apache-2.0/)
