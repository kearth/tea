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

require(__DIR__ . '/framework/bootstrap.php');

Bootstrap::run(__DIR__);

