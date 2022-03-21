Generate a bcrypted and/or encoded password.

## Usage

`camo` processes stdin and turns a string into an encrypted password.

```
Usage of ./camo:
  -cost int
        The cost weight, range of 4-31 (default 10)
  -encode
        Base64 encode the resulting hash (default false)
  -secret
        Add the resulting value to a k8s secret and print that out (default false)
```

Note: If you are passing the `--secret` flag you will have to edit the resulting
yaml to update the `name`, `namespace` and `username` fields to match your needs.

### With a Docker image (no binary install)

```
PASS=foobar
docker run -e CLEAR_PASSWORD="$PASS" ghcr.io/warehouse-13/camo:latest
```

### With released binary

```
echo -n foobar | camo
```
