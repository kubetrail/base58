# base58
encoder decoder in base58

## installation
download code to a folder, cd to it and run `go install`

## usage
To encode pass input as arguments
```bash
base58 this is a test
jo91waLQA1NNeBmZKUF
```

or pass input via STDIN
```bash
echo -n this is a test | base58
jo91waLQA1NNeBmZKUF
```

> Pl. note that the STDIN is ignored if arguments are provided

To decode use flag `-d`
```bash
base58 -d jo91waLQA1NNeBmZKUF
this is a test
```

or pass input via STDIN
```bash
echo -n jo91waLQA1NNeBmZKUF | base58 -d
this is a test
```

```bash
echo hello | base58 | base58 -d
hello
```