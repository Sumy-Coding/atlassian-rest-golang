# Confluence REST service on Golang

https://developer.atlassian.com/server/confluence/confluence-rest-api-examples/

- Groovy version - https://github.com/AndriiMaliuta/confluence-rest-service-groovy
- Rust version - https://github.com/AndriiMaliuta/rust-atlas-rest

## Examples

### Confluence
```bash
./atlas --type confluence --action getPage --id "854950177"
./atlas --type confluence --action getPage --space "TEST" --title "Page A"
./atlas --type confluence --action getSpace --space "TEST"

```

### Jira
```bash
./atlas --type confluence --action getIssue --key "AAA-3"
```