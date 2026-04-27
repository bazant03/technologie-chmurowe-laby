## Marek Ruszecki - TCH Zad1

**Aplikacja wyśwetlająca dane pogodowe dla wybranego miasta korzystając z OpenWeather API**
- Napisana w języku Rust
- Jest "samodzielnym" plikiem wykonywalnym
- Uruchamiania na obrazie scratch

**Przykład użycia**

Budowanie obrazu:

```bash
docker build -t local/weather:v1.0.0 .
```

Uruchomienie kontenera (w pliku .env należy podać klucz OPENWEATHER_API_KEY):

```bash
docker run -d -p 8080:80 --rm --env-file .env --name=weather local/weather:v1.0.0
```

Sprawdzenie logów:

```bash
docker logs weather
```

**Zrzuty ekranu znajdują się w sprawozdaniu przesłanym na moodle.**