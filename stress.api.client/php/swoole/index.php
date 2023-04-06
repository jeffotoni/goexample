<?php
// main.php
//require_once 'vendor/autoload.php';

use Swoole\Http\Server;
use Swoole\Http\Request;
use Swoole\Http\Response;
use Swoole\Coroutine\Http\Client;

$server = new Server('127.0.0.1', 8080);

$server->on('start', function (Server $server) {
    echo "Swoole HTTP server is started at http://127.0.0.1:8080\n";
});

$server->on('request', function (Request $request, Response $response) {
    if ($request->server['request_method'] == 'GET' && $request->server['request_uri'] == '/v1/client/get') {
        $token = 'x3939393939x39393';
        $url = 'http://127.0.0.1:3000/v1/customer/get';

        $data = fetchCustomerData($url, $token);
        $response->header('Content-Type', 'application/json');
        $response->end(json_encode($data));
    } else {
        $response->status(404);
        $response->end('Not Found');
    }
});

$server->start();

function fetchCustomerData(string $url, string $token): array
{
    $urlParts = parse_url($url);
    $client = new Client($urlParts['host'], $urlParts['port']);
    $client->setHeaders(['Authorization' => "Bearer $token"]);
    $client->get($urlParts['path']);

    $body = $client->getBody();
    $client->close();

    return json_decode($body, true);
}

