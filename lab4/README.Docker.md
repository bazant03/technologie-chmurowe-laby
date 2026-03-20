<pre>
Obraz można zbudować poleceniem: docker build -t local/web100:v1.0.0 .
I uruchomić kontener poleceniem: docker run -d --rm -p 8000:80 --name=web100 local/web100:v1.0.0
Aby połączyć się ze stroną wystarczy wpisać w adresie przeglądarki http://localhost:8000
(Serwer w kontenerze nasłuchuje na porcie 80)

Zbudowany obraz posiada 3 warstwy:
1.Z obrazu bazowego
2.Aktualizacja systemu, pobranie apache ... (Polecenie RUN)
3.Skopiowanie pliku src/index.html do obrazu (Polecenie COPY)

Pliki zostały wygenerowane z użyciem docker init

W katalogu src/ znajduje się hostowana przykładowa "strona internetowa"
</pre>