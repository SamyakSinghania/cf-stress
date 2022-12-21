This `autograder` evaluates all the `go assignments` on the basis of running `go test` commands on the bash for all `_test.go` files.

Firstly all the submissions are uploaded in a folder named `Recruitment Task`. Using `os/exec package`, go test commands are run on the bash for each submission. 

A `map` is created with its `key` as `username` and `value` as the `number of tasks` successfully completed.

If the output received on `stdout` starts with `PASS`, the values for that username is `incremented` on the map. All the data of this map is then copied to a user defined struct array so that it can be sorted in a decreasing order.

Then this key value pair is written on a file named `results.txt`

If any error happens while running the `go test` command, the error is printed on `stdout`

This `autograder` can be effectively used to evaluate the `PCLUB` summer project submissions. 

Some demo submissions are provided in the `Recruitment Task` folder along with the results in the `results.txt` file.

Transfer the submissions to the `Recruitment Task` folder and run the `main.go` file.

The results will be uploaded in the `results.txt` file.
