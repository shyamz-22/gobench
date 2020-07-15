# gobench
A small application to showcase how to use benchmark

## Setup

Install [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat?tab=doc)

```bash
> go get golang.org/x/perf/cmd/benchstat
```

## Benchmarking

### Step 1:

```bash
> cd gobench
> go test -run=None -bench=BenchmarkJson  -count=10 > old.txt
```

### Step 2:

Comment Line 23  and UnComment Line 24 in marshal_test.go


```bash
> cd gobench
> go test -run=None -bench=BenchmarkJson  -count=10 > new.txt
```

### Step 3:

Now compare the two benchmarks

```bash
> cd gobench
> benchstat old.txt new.txt
```

```
name    old time/op    new time/op    delta
Json-8     405ns ± 1%     414ns ± 0%   +2.25%  (p=0.000 n=9+9)

name    old alloc/op   new alloc/op   delta
Json-8      176B ± 0%      112B ± 0%  -36.36%  (p=0.000 n=10+10)

name    old allocs/op  new allocs/op  delta
Json-8      3.00 ± 0%      2.00 ± 0%  -33.33%  (p=0.000 n=10+10)
```

Watch video to understand more about [p-value](https://www.youtube.com/watch?v=YSwmpAmLV2s)