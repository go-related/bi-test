
- Run migrations
  ```
  go run cmd/migrations/main.go
  ```
- Create a db entity
    ```
  go run -mod=mod entgo.io/ent/cmd/ent new User
  ```
- Creat migrations
    ```
  go generate ./ent
  ```
- Generate golag binding for smart contract
  ```shell
  abigen --abi contracts/CarRenting/CarRenting.abi --pkg main --type CarRenting --out CarRentingContract.go
  ```
