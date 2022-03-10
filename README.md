Generate a bcrypted password.

Usage:

```
PASS=foobar
docker run -e CLEAR_PASSWORD="$PASS" ghcr.io/warehouse-13/camo:latest
```

Output will be something like:

```
JDIkMTIkQlVJSDVtSDlkR3ZBakZJM3cyRGRWdU0xMkUubUtHTUpodEF5MVJnMUpXWXl5Y1RjbkhDLksK
```

This is base64 encoded, you can go ahead and throw this into a k8s secret yaml.

_I will make this smarter soon, like write it in python and have an
option to spit out a kube secret, but this will do for now._
