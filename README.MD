# Go Clean Architecture


### How To Run This Project
```bash
#move to directory
cd $GOPATH/src/github.com/<username>

# Clone into YOUR $GOPATH/src
git clone https://github.com/samayamnag/go-clean-arch.git

#move to project
cd go-clean-arch

#copy config and ensure all values are correct
cp config_example.json config/config.json
cp config_test_example.json config/test.json - For test cases

# Install Dependencies
dep ensure

# Run Project
go run main.go

```

Or

```bash
# GET WITH GO GET
go get github.com/samayamnag/go-clean-arch

# Go to directory

cd $GOPATH/src/github.com/go-clean-arch

# Install Dependencies
dep ensure


# Run Project
go run main.go
```

### Avialable routes
```bash
1. localhost:<port>users - GET -> All Users
2. localhost:<port>/users/ - POST -> Create User
3. localhost:<port>/users/{id} - GET -> Get User
4. localhost:<port>/users/{id} - PUT -> Update User
5. localhost:<port>/users/{id} - DELETE -> Delete User
```