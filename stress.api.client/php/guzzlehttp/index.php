<?php
/** 
 * Foi testado usando  php -S localhost:8080 main.php > /dev/null 2>&1 &
 * E foi usado também nginx para nosso web server
 * php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
 * php -r "if (hash_file('sha384', 'composer-setup.php') === '55ce33d7678c5a611085589f1f3ddf8b3c52d662cd01d4ba75c0ee0459970c2200a51f492d557530c71c15d8dba01eae') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;"
 * php composer-setup.php
 * php -r "unlink('composer-setup.php');"
 * sudo mv composer.phar /usr/local/bin/composer
 */
require_once 'vendor/autoload.php';

use GuzzleHttp\Client;
use GuzzleHttp\Promise;
use GuzzleHttp\Exception\RequestException;

// Verifique se a requisição é do tipo GET e tem o caminho correto
if ($_SERVER['REQUEST_METHOD'] === 'GET' && $_SERVER['REQUEST_URI'] === '/v1/client/get') {

$client = new Client();
$url1 = 'http://127.0.0.1:3000/v1/customer/get';
$token = 'x3939393939x39393';

try {
    $response = $client->get($url1, ['headers' => ['Authorization' => "Bearer $token"]]);
    echo $response->getBody();
    //$data1 = json_decode($response->getBody(), true);
    //echo($data1);

} catch (RequestException $e) {
    echo "RequestException: " . $e->getMessage() . "\n";
    if ($e->hasResponse()) {
        echo "Response: " . $e->getResponse()->getBody() . "\n";
    }
} catch (Exception $e) {
    echo "Exception: " . $e->getMessage() . "\n";
}

} else {
    header("HTTP/1.1 404 Not Found");
}
