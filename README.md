# One Billion Row Challenge - My Playground

## Very short description

Read https://github.com/gunnarmorling/1brc for the full original description. This is just my own personal playground that I can share with friends.

In short:

- Write a program that reads a really large CSV file (a billion rows). Each row is in the format: `Location;temperature`. The location is a string, the temperature a floating point number with at most one decimal digit after the dot.

  Example of a file with 8 rows:

  ```plain
  Mandalay;23.2
  Cairns;18.8
  Fresno;1.6
  Abidjan;35.8
  Porto;30.3
  Nouadhibou;23.0
  Wau;16.5
  Frankfurt;-0.6
  ```

  Locations can be repeated in the file; e.g., you can encounter more than one datapoint for `Mandalay`. 
  
- You may assume that all lines are well-formed: all lines have two parts, separated by `;` and the second part is a number. There are no empty lines, except for the trailing `\n`. (Though in my experience, checking for parsing errors doesn't add any significant overhead.)

- Process these lines, so that per location, you determine the minimum temperature, the maximum, and the average.

- Output your results **in alphabetically sorted order**, e.g., given the above example `Abidjan` first and `Wau` last. Per result display the name of the location, the minimum temperature, the average, and the maximum. Temperatures must be displayed with a precision of 2 digits; e.g., `22.39`.

- **Make that program as fast as possible**. At the time of writing this, I managed to come up with an alrorithm that reduces the runtime when comparing to a simple naive algorithm. But I'm still playing around with it, I don't have "the right approach".

## Naive approach (for reference)

My "naive" approach works as follows.

- The input is read from file, line by line.
- Each line is split into two parts: a location and a temperature datapoint.
- The result is stored in an unsorted map (that's Go's default), keyed by location. The map value is a structure that tracks the minimum temperature, the maximum, the total and the number of datapoints for this location.
- While updating the map's value, the minimum and maximum may need adjusting. The total and number of occurrences are updated:
  ```
  // pseudo-code when a temperature is recorded for a location
  results[location].minimum = min(results[location].mininum, temperature)
  results[location].maximum = max(results[location].maximum, temperature)
  results[location].total += temperature
  results[location].n++
  ```
- Before reporting, the map's keys (i.e., locations) are collected in an array and alpha-sorted. 
- Finally the report is generated in the order of the sorted locations. The average temperature is computed as total divided by number of occurrences: `average := map[location].total / map[location].n`.

## Input data

The repository holds two datasets: `data/100.csv` and `data/2m.csv` of resp. 100 and 2.000.000 entries. The challenge is however about a billion.

To create larger files, run:

```sh
cat data/2m.csv data/2m.csv data/2m.csv data/2m.csv data/2m.csv           > data/10m.csv
cat data/10m.csv data/10m.csv data/10m.csv data/10m.csv data/10m.csv      > data/50m.csv
cat data/50m.csv data/50m.csv                                             > data/100m.csv
cat data/100m.csv data/100m.csv data/100m.csv data/100m.csv data/100m.csv > data/500m.csv
cat data/500m.csv data/500m.csv                                           > data/1b.csv 
```

The dataset to work with is `data/1b.csv` and should now hold 1.000.000.000 lines:

```sh
wc -l data/1b.csv
 1000000000 data/1b.csv
```

## Is the output correct?

I have provided reference files that you can `diff` against: `data/100.csv.out` (the output for the processing of `data/100.csv`), `data/2m.csv.out` (output for processing `data/2m.csv`), `data/10m.csv.out` etc.. 

The files were produced by `naive.go` which uses the `printf`-like formatting `"%-30s %.2f %.2f %.2f\n"` for each location's name, minimum temperature, average, and maximum. If your solution generates output in the same format, then diffing `data/$SIZE-csv.out` against your ouput should produce no differences.

Unless of course `naive.go` doesn't properly work. In that case let me know.
