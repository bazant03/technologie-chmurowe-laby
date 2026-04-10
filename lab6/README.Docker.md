## Marek Ruszecki - TCH Lab6 - modyfikacja poprzedniego zadania

**Warstwa build**
- Wykorzystuje obraz scratch i rootfs systemu alpine
- Hostowana aplikacja została napisana w Go
- Aplikacja jest kompilowana statycznie
- **Kod aplikacji pobierany z GitHub przez SSH**

**Warstwa finalna**
- Wykorzystuje oficjalny obraz nginx (debian-based)
- Nginx pracuje jako reverse proxy
- **Obraz ma etykietę automatycznie łączącą go z repo**

**Przykład użycia**

Budowanie obrazu:

```bash
docker build --ssh default --build-arg VERSION=v1.0.0 -t local/multistage2:lab6 .
```

Uruchomienie kontenera:

```bash
docker run -d -p 8080:80 --rm --name=multi2 local/multistage2:lab6
```

**Lub**

Pobranie obrazu z ghcr.io:

```bash
docker pull ghcr.io/bazant03/multistage2:lab6
```

Uruchomienie kontenera:

```bash
docker run -d -p 8080:80 --rm --name=multi2 ghcr.io/bazant03/multistage2:lab6
```

Test połączenia:
```bash
curl -i http://localhost:8080
```
