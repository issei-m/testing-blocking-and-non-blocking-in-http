<?php

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;
use Symfony\Component\HttpClient\HttpClient;
use Symfony\Contracts\HttpClient\Exception\TransportExceptionInterface;

require __DIR__ . '/../vendor/autoload.php';

$app = AppFactory::create();

$app->get('/', function (Request $request, Response $response, $args) {
    $sleepTime = $request->getQueryParams()['sleep'] ?? '';

    try {
        $httpResponse = HttpClient::create()->request('GET', "http://rust:3000/?sleep={$sleepTime}");

        if ($httpResponse->getStatusCode() === 200) {
            $response->getBody()->write($httpResponse->getContent());

            return $response;
        }
    } catch (TransportExceptionInterface $e) {
    }

    $response = $response->withStatus(500);
    $response->getBody()->write('Failed to fetch data');

    return $response;
});

$app->run();
