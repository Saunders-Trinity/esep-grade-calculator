# ESEP Grade Calculator

A small Go library for calculating a course's final letter grade from assignment,
exam, and essay grades. The calculator averages each category and applies these
weights:

| Category | Weight |
| --- | ---: |
| Assignments | 50% |
| Exams | 35% |
| Essays | 15% |

## Requirements

- Go 1.24 or newer

## Run the tests

Run these commands from the `grade-calculator` directory:

```sh
go test ./...
```

To report statement coverage:

```sh
go test -cover ./...
```

To create a coverage profile and inspect uncovered lines in a browser:

```sh
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Use the library

Create a calculator, add grades using one of the exported grade types, then ask
for the final letter grade:

```go
calculator := esepunittests.NewGradeCalculator()
calculator.AddGrade("Homework 1", 92, esepunittests.Assignment)
calculator.AddGrade("Midterm", 86, esepunittests.Exam)
calculator.AddGrade("Research essay", 95, esepunittests.Essay)

finalGrade := calculator.GetFinalGrade()
```

The package name is `esepunittests`; the module path is
`esep/grade-calculator`.
