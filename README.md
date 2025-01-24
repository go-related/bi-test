# Notes during developement

the code that call the contract correctly is found under blockchain folder not under internal.
<br> 
I am not deleting and cleaning things up but the main problem in the internal were 
not waiting for the tx and tx options related things.



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

