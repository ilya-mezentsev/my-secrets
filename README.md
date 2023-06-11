# my-secrets
CLI tool for storing secrets

# Build:
```bash
$ cd source
$ GOMODCACHE=$(pwd)/pkg go build ./cmd/main.go
```

# Usage
1. Storing value by key: ```./main set key value```
2. Receiving value by key ```./main get key```

both require password (can be different for each key)

# Troubleshooting
If something went wrong set `VERBOSE_SECRETS` environment variable to `Y`
so application logs will appear in stdout
