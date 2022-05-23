<?php
use PhpParser\Error;
use PhpParser\NodeDumper;
use PhpParser\ParserFactory;

$start = microtime(true);
require_once __DIR__ . '/vendor/autoload.php';

$code = file_get_contents($argv[1]);

$parser = (new ParserFactory)->create(ParserFactory::PREFER_PHP7);
try {
    $ast = $parser->parse($code);
} catch (Error $error) {
    echo "Parse error: {$error->getMessage()}\n";
    return;
}

$dumper = new NodeDumper;
echo $dumper->dump($ast) . "\n";
echo 'Execute time: '.(microtime(true) - $start)."\n";
printf('Memory usage: %s Mb', memory_get_peak_usage(true)/1024/1024);
