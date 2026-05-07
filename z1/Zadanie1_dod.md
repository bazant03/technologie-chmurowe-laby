## Marek Ruszecki - TCH Zad1 dodatkowe

**Aplikacja wyśwetlająca dane pogodowe dla wybranego miasta korzystając z OpenWeather API**
- Obraz jest dostępny na platformy linux/amd64 i linux/arm64
- Kod aplikacji pobierany jest z repozytorium GitHub
- Cache jest przechowywany w registry


**Przykład użycia**

Budowanie obrazu:

```bash
docker buildx build \
    -f Dockerfile_nb \
    -t docker.io/[repo]/weather-multiarch:v1.0.0 \
    --ssh default \
    --platform linux/amd64,linux/arm64 \
    --output type=image,push=true,oci-mediatypes=true \
    --cache-to type=registry,ref=docker.io/[repo]/weather-multiarch:cache,mode=max \
    --cache-from type=registry,ref=docker.io/[repo]/weather-multiarch:cache \
    . 
```

Pobranie obrazu (z mojego repozytorium):
```bash
docker pull s101505/weather-multiarch:v1.0.0
```

Uruchomienie kontenera (w pliku .env należy podać klucz OPENWEATHER_API_KEY):

```bash
docker run -d -p 8080:80 --rm --env-file .env --name=weather s101505/weather-multiarch:v1.0.0
```

Sprawdzenie logów:

```bash
docker logs weather
```

**Zrzuty ekranu znajdują się w sprawozdaniu przesłanym na moodle.**