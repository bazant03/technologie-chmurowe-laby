## Marek Ruszecki - TCH Lab14 (część obowiązkowa)
**Uruchomienie stack-a LEMP z użyciem Docker Compose**

Aplikacja zawiera 4 mikrousługi:
- mysql
- php-fpm
- phpmyadmin-fpm
- nginx 

Do sieci backend podłączone są:
- mysql
- php-fpm
- phpmyadmin-fpm
- (nginx)

Do frontend podłączony jest tylko nginx.

Domyslnie phpmyadmin zawiera serwer webowy aby móc "hostować się niezależnie od reszty usług".<br>
W tym wypadku gdy korzystamy z nginx, mamy już serwer webowy wieć możemy skorzystać z phpmyadmin w wersji fpm,<br>
która zawiera kod usługi oraz interpreter ale nie ma serwera.<br> 
Pliki źródłowe z kontenera phpmyadmin są udostępniane do nginx poprzez wspólny wolumen, <br>
dzięki czemu może on hostować treść phpmyadmin, którego kontener pełni funkcję mikrousługi. <br>
Nginx łączy się z phpmyadmin (i drugą "aplikacją") z wykorzystaniem protokołu FastCGI.<br>

Obraz php został poszerzony o bibliotekę PDO do obsługi połączenia z bazą danych mysql.