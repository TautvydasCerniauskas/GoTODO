# This is a simple TODO app build in golang

The purpose of this program is to explore golang capabilities for a simple application.
Sqlite is used to persist TODOs

## Supported commands

```
./GoTODO list <--all>
./GoTODO init
./GoTODO delete <id>
./GoTODO update <id> <content>
./GoTODO done <id>
```

### Add

```
./GoTODO add <name>
```

Add command is a special case which supports different types of severity.
Prefix of: 
  * ! = HIGH 
  * # - LOW
  * No prefix = MEDIUM
