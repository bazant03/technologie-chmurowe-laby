## Marek Ruszecki - TCH Lab5

**Warstwa build**
- Wykorzystuje obraz scratch i rootfs systemu alpine
- Hostowana aplikacja została napisana w Go
- Aplikacja jest kompilowana statycznie

**Warstwa finalna**
- Wykorzystuje oficjalny obraz nginx (debian-based)
- Nginx pracuje jako reverse proxy

**Przykład użycia**

Budowanie obrazu:

```bash
docker build --build-arg VERSION=v1.0.0 -t local/multistage:v1.0.0 .
```

Uruchomienie kontenera:

```bash
docker run -d -p 8080:80 --rm --name=multi local/multistage:v1.0.0 
```

Test połączenia:
```bash
curl -i http://localhost:8080
```
