# Backend minimo para el frontend con Gin

Este backend expone lo minimo que el frontend necesita:

- `POST /login`
- `GET /health`

Usuario disponible por defecto:

- `admin`
- `admin`

## Ejecutar

```bash
cd backend
go run .
```

## Probar login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin"}'
```
