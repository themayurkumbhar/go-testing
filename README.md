# Go -Tests Fundaments :white_check_mark:	

## 1. Getting started with go-testing :coffee:	

* Go installation provide the inbuilt `testing` and `profiling` tools
* types of test in go
  * basic test
    * unit, integration, end-to-end test
  * benchmark test
    * performace profiling
  * example test
    * mostly for documentation
* Writing you first test in go
  * Syntax
  ```go
    package main_test
    import (
        "testing"
    )
    func MethodToTest(testingPointer *testing.T) {
        // initialize
        // do something
        // assert the result
    }
  ```
  * Here the `*testing.T` in package `testing` provides all the required methods to testing in go
    * in Go there is no such thing as assert, you have to manully analyze the result with `if` and raise any error
* Once you write the basic test, lets run
    ```bash
    go test -v
    ```
  * This will execute the `tests` in current `go-module` and show the output as `PASS/FAIL`
    ```bash
    === RUN   TestMultiplication
    --- PASS: TestMultiplication (0.00s)
    === RUN   TestAddition
    --- PASS: TestAddition (0.00s)
    PASS
    ok      github.com/themayurkumbhar/go-testing   0.547s
    ```
  * You can check the source code of [app_test.go](app_test.go)
* How `go` understands the `test` files
  * Each `test` file shoule have `<sourceFilename>_test.go` 
  * Go `package` should have `<packageName>_test`
  * All the test methods should prefix with `Test`

* Testing related other `packages` :gear:	
  * `testing` package in go
  * `testing/quick` simplify the `blackbox` testing
  * `testing/iotest` contains `read/write` specific tests
  * `net/http/httptest` for `network` related testing
  * Commnunity testing projects
    * Testify : [stretchr/testify](https://github.com/stretchr/testify)
    * Ginkgo : [onsi/ginkgo](https://github.com/onsi/ginkgo)
    * GoConvey : [goconvey](http://goconvey.co/) Visualize you test reports
    * Httpexpect : [gavv/httpexpect](https://github.com/gavv/httpexpect) end-to-end web testing
    * Gomock : [gomock](https://github.com/golang/mock) mock objects
    * go-sqlmock : [DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)

---