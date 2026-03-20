FROM ubuntu:latest 

# Wpisanie informacji o autorze w zalecany sposób
LABEL org.opencontainers.image.authors="Ruszecki <s101505@pollub.edu.pl>"

# Aktualizacja Ubuntu, instalacja pakietu Apache
# następnie usunięcie katalogu z danymi konfiguracyjnymi repozytoriów apt
# oraz uzupełnienie konfiguracji serwera http (bez tego nie działa)
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends apache2 && \
    rm -rf /var/lib/apt/lists/* && \
    echo "ServerName localhost" >> /etc/apache2/apache2.conf

COPY src/index.html /var/www/html/index.html
    
EXPOSE 80/tcp

#Nie uruchamiamy procesu w tle, ponieważ wtedy kontener od razu by się wyłączył
ENTRYPOINT ["apache2ctl", "-D", "FOREGROUND"]
