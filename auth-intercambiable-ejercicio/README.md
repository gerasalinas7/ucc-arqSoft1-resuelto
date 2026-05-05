# Ejercicio para resolver en clase

Objetivo: que el `main` y los handlers no cambien cuando cambiamos el tipo de auth.

## Consigna

Completar las implementaciones que faltan para que el sistema soporte dos estrategias:

- `jwt`
- `apikey`

## Archivo a completar

- `auth/auth.go`

## .env

El proyecto ya trae este archivo:

```env
AUTH_MODE=jwt
AUTH_SECRET=secret-demo
```

## Pistas para la clase

- `Provider` es la interfaz que usa el `main`.
- `apikey` ya esta resuelto.
- Falta completar la parte de `jwt`.
- La idea es mostrar que el punto de variacion queda detras de una interfaz.

## Prueba esperada

1. `go run .`
2. Login en `/login`
3. Usar el token en `/profile`
4. Cambiar `AUTH_MODE=apikey` en `.env` y repetir
5. El `main.go` no se toca
