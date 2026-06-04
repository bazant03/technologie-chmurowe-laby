<!DOCTYPE html>
<html lang="pl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Stronka</title>
</head>
<body>
    <p>Testowa strona</p>

    <?php
    // Dane konfiguracyjne 
    $host = 'mysql'; 
    $user = 'root';
    $password = 'to_nie_produkcja';
    $charset = 'utf8mb4';

    $dsn = "mysql:host=$host;charset=$charset";
    
    $options = [
        PDO::ATTR_ERRMODE            => PDO::ERRMODE_EXCEPTION, // Włącza rzucanie wyjątków przy błędach
        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,       // Zwraca wyniki jako tablice asocjacyjne
    ];

    try {
        // Próba nawiązania połączenia
        $pdo = new PDO($dsn, $user, $password, $options);
        echo '<p style="color: green;"><strong>Sukces:</strong> Połączenie z bazą MySQL zostało ustanowione!</p>';

        // Testowe dane - lista wszystkich baz z kontenera mysql
        $stmt = $pdo->query('SHOW DATABASES');
        
        echo '<h3>Dostępne bazy danych:</h3>';
        echo '<ul>';
        while ($row = $stmt->fetch()) {
            echo '<li>' . htmlspecialchars($row['Database']) . '</li>';
        }
        echo '</ul>';

    } catch (\PDOException $e) {
        // Jeśli nie uda się połączyć z bazą
        echo '<p style="color: red;"><strong>Błąd połączenia:</strong> ' . htmlspecialchars($e->getMessage()) . '</p>';
    }
    ?>
</body>
</html>