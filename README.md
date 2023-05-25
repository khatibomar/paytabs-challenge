# paytabs-challenge

## Requirements
to be able to run and test this API you need `make` and `go` installed
### Installing Go
To install Go simply follow guide on their [official website](https://go.dev/dl/)
### Installing Make
#### Windows
first download [scoop](https://scoop.sh/) by following [official guide](https://github.com/ScoopInstaller/Install#readme)
after installing scoop, run the following command in powershell
```bash
scoop install make
```
#### Linux
Every linux distro come with it's own package manager, so if you use Arch BTW, you can install make using `yay`
```bash
yay -S make
```
for Debian based distros install using
```bash
sudo apt install make
```

## Run
after having `make` and `go` installed we ready to Go ( get it ?  )

### running tests
to run tests simply run in terminal
```bash
make test
```
and let's pray you got same output as mine
```
?       github.com/khatibomar/paytabs-challenge/cmd/api/v1      [no test files]
?       github.com/khatibomar/paytabs-challenge/internal/customerrors   [no test files]
ok      github.com/khatibomar/paytabs-challenge/internal/datastructure  (cached)
ok      github.com/khatibomar/paytabs-challenge/internal/parser (cached)
?       github.com/khatibomar/paytabs-challenge/internal/validator      [no test files]
ok      github.com/khatibomar/paytabs-challenge/internal/store  0.035s
```
if not send me an email :)

### running API
to run the API run the following
```bash
make go/tidy
```
then enter `y` then enter, wait for `go` to get packages, after that, if all succeed, run
```bash
make port=8855 run/api/v1
```
> change `8855` to any available port 
then you should get this output
```
go run ./cmd/api/v1/... -port=8855
2023-05-25T16:19:44.624+0300    INFO    v1/main.go:40   Ingesting accounts into the store...
2023-05-25T16:19:44.626+0300    INFO    v1/main.go:45   Ingesting accounts done...
2023-05-25T16:19:44.626+0300    INFO    v1/server.go:48 starting server {"addr": ":8855", "env": "development"}
```
## Endpoints
### Account
#### List all accounts
Since we seeded 500 accounts into our store, we have 500 account by default
```
curl --location 'http://localhost:8855/v1/accounts/'
```
![image](https://github.com/khatibomar/paytabs-challenge/assets/35725554/8b9a7889-4ff5-43d2-99f5-46cf5707b7a7)
> you may get accounts in different order, because I am using a map without sorting
#### Create an account
```bash
curl --location 'http://localhost:8855/v1/accounts/' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Omar",
    "balance": 1999888777
}'
```
my name is `Omar` and I have `1999888777` in my balance ( don't think I am rich it's in LBP not USD :3 ) 
![image](https://github.com/khatibomar/paytabs-challenge/assets/35725554/6dd793bf-53dc-4064-95dd-c866cbf8c46e)
we can see that account created succesfully, with guid of `0d865cb9-0793-4579-93c6-a0899c89f2d0`
#### Get account information
```
curl --location 'http://localhost:8855/v1/accounts/2a3e1307-0c09-47fe-b6e5-ce8630c152bf'
```
![image](https://github.com/khatibomar/paytabs-challenge/assets/35725554/946e8e14-da14-4a55-929a-7f7b89615276)
> change guid to get info for different account

#### Deposit
```bash
curl --location 'http://localhost:8855/v1/deposit' \
--header 'Content-Type: application/json' \
--data '{
    "id": "2a3e1307-0c09-47fe-b6e5-ce8630c152bf",
    "amount": 100
}'
```
![image](https://github.com/khatibomar/paytabs-challenge/assets/35725554/bbbb85a6-be57-4212-9361-5eb8acd05c79)
Dynava now have +100 on her old balance, let's take it away from her 3:)

### Withdraw
```bash
curl --location 'http://localhost:8855/v1/withdraw' \
--header 'Content-Type: application/json' \
--data '{
    "id": "2a3e1307-0c09-47fe-b6e5-ce8630c152bf",
    "amount": 100
}'
```
Dynava now have +100 on her old balance
![image](https://github.com/khatibomar/paytabs-challenge/assets/35725554/cbb88689-71ab-4a95-83ed-0ba31a38eecc)

That's our simple API, have fun and try to break it by passing negative amounts, or wrong GUIDs

## Improvements
1- Add more endpoints ( Full CRUD ) operations
2- Write more concurrent tests
3- Use repository pattern with real database
4- Make it more secure ( rate limiting, SSL, AUTH )
5- Make it distributed ( by adding service discovery, coordination, and replicating )
6- Create CLI menu App to be easier to test by collaborators :) or web frontend if I have time
