
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
  abigen --abi contracts/Counter/Counter.abi --pkg main --type Counter --out CounterContract.go
  ```
