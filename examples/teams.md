# User API

Please refer to `client.md` before reviewing these examples.

## Usage

> List Teams

```go
result, _, err := client.UserService.ListTeams(map[string]string{})
if err != nil {
    log.Fatalf("List Teams err: %s", err)
}

log.Println("List Team: ", result)
```

> List Teams with a name

```go
result, _, err := a.AWXClient.AWX.TeamService.ListTeams(map[string]string{
    "name": "my-team-name",
})
if err != nil {
    log.Fatalf("List Teams err: %s", err)
}

log.Println("List Team: ", result)
```

> Associate user and team. User will be added as member

```go
err := a.AWXClient.AWX.TeamService.AddTeamUser(teamId, map[string]interface{}{
	"id": userId,
})
if err != nil {
    log.Fatalf("Fail to associate user %d and team %d, err: %s", userId, teamId, err)
}
```

> Disassociate user and team

```go
err := a.AWXClient.AWX.TeamService.RemoveTeamUser(teamId, map[string]interface{}{
	"id": userId,
})
if err != nil {
    log.Fatalf("Fail to disassociate user %d and team %d, err: %s", userId, teamId, err)
}
```
