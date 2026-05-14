# Auth intercambiable resuelto

Ejemplo simple de autenticacion en Go donde el `main` no cambia y solo variamos `AUTH_MODE`.

Modos disponibles:

- `jwt`
- `apikey`

## .env

El proyecto ya trae este archivo:

```env
AUTH_MODE=jwt
AUTH_SECRET=secret-demo
```

## Probarlo

1. Entrar a la carpeta:

```bash
cd Go-examples/auth-intercambiable-resuelto
```

2. Levantarlo:

```bash
go run .
```

3. Hacer login:

```bash
curl -s -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

4. Probar el token en `/profile`:

```bash
curl -s http://localhost:8080/profile \
  -H "Authorization: Bearer TU_TOKEN"
```

5. Frenar el servidor, cambiar `AUTH_MODE=apikey` en `.env` y levantarlo otra vez:

```bash
go run .
```

6. Repetir el login y usar la API key devuelta como Bearer token en `/profile`.

Usuarios:

- `admin / admin123`
- `ana / clave123`

Verificacion automatica:

```bash
go test ./...
```
