# slack-scripts
Collection of useful slack scripts written in go

- [slack-scripts](#slack-scripts)
  - [Deactivated Users](#deactivated-users)
## Deactivated Users

Get a list of deactivated users from slack and order it by most recent. Outputs as an organized json file in current working directory.

1. Create slack app api token `xoxp-foo`
2. Add read/list users permission to token
3. [Install go](https://golang.org/doc/install)
4. Pop your slack api token into `deactivated_users.go`
5. `go run deactivated_users.go`
