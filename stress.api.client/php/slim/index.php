<?php

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;

use GuzzleHttp\Client;
use GuzzleHttp\Promise;
use GuzzleHttp\Exception\RequestException;

require __DIR__ . '/vendor/autoload.php';

$app = AppFactory::create();

$app->addErrorMiddleware(true, true, true);

$app->get('/v1/client/get', function (Request $request, Response $response, array $args) {

$client = new Client();
$url1 = 'http://127.0.0.1:3000/v1/customer/get';
$token = 'x3939393939x39393';

try {
    $response1 = $client->get($url1, ['headers' => ['Authorization' => "Bearer $token"]]);
    $data1 = json_decode($response1->getBody(), true);

    $response->getBody()->write(json_encode($data1));
    return $response
        ->withHeader('Content-Type', 'application/json')
        ->withStatus(200);

} catch (RequestException $e) {
    echo "RequestException: " . $e->getMessage() . "\n";
    if ($e->hasResponse()) {
        echo "Response: " . $e->getResponse()->getBody() . "\n";
    }
} catch (Exception $e) {
    $error = ['error' => $e->getMessage()];
    $response->getBody()->write(json_encode($error));
    return $response
        ->withHeader('Content-Type', 'application/json')
        ->withStatus(500);
}

});

$app->run();
