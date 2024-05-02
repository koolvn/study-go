# study-go

Repo with toy projects to study *Go*


## Generate ED25519 PEM keys
```bash
openssl genpkey -algorithm Ed25519 -out certs/jwt/ed25519_private.pem
openssl pkey -in certs/jwt/ed25519_private.pem -pubout -out certs/jwt/ed25519_public.pem
chmod 600 certs/jwt/ed25519_private.pem
chmod 644 certs/jwt/ed25519_public.pem
```