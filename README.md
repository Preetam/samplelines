# samplelines

A small Go program to sample 1 in N lines from stdin

## Usage

```
Usage of samplelines:
  -n uint
    	1 in n lines to sample. 1 means every line is printed to stdout. (default 1)
```

## Example

```
$ some_program | samplelines -n 10 # samples 1 in 10 lines
```
