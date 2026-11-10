# ctrl-alt-grim

Log filter for people who live in a terminal.

```sh
tail -f app.log | go run . -level warn -grep timeout
```

Flags: `-level debug|info|warn|error`, `-grep text`.
