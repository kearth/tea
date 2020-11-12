<?php

/**
 * Tea Framework
 * Version 1.0.0
 *
 * Copyright 2018 - 2020, Kearth
 * PHP 7.2.3
 */

declare(strict_types=1);

error_reporting(E_ALL);

use Tea\Framework\Bootstrap as Bootstrap;

define("APP_ROOT", __DIR__ . "/app");

require(__DIR__ . '/vendor/autoload.php');

Bootstrap::run(__DIR__ . "/conf/");

