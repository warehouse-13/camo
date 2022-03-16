Generate a bcrypted and encoded password.

You can go ahead and throw the output into a k8s secret yaml.

## Usage

### With a Docker image (no binary install)

```
PASS=foobar
docker run -e CLEAR_PASSWORD="$PASS" ghcr.io/warehouse-13/camo:latest
```

### With released binary

```
echo -n foobar | camo
```
