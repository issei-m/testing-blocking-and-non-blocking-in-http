<?php

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;

require __DIR__ . '/../vendor/autoload.php';

$app = AppFactory::create();

$app->get('/', function (Request $request, Response $response, $args) {
    $sleepTime = $request->getQueryParams()['sleep'] ?? '10';
    sleep(intval($sleepTime));

    $response->getBody()->write("{$sleepTime} seconds have passed");

    return $response;
});

$app->run();
