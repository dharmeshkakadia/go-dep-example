go-dep-example
==============

0. `brew install dep`
1. `dep init`. This will generate following tree
    ```
    .
    ├── Gopkg.lock
    ├── Gopkg.toml
    ├── README.md
    └── vendor
    ```

2. Now import the new package in the code. We will import "github.com/franela/goreq". You will see your IDE complaining about not being able to find the package, since its not yet installed. We will do that next. You can confirm this by running `dep status`
    ```
    Lock inputs-digest mismatch due to the following packages missing from the lock:

    PROJECT                   MISSING PACKAGES
    github.com/franela/goreq  [github.com/franela/goreq]

    This happens when a new import is added. Run `dep ensure` to install the missing packages.
    ```
3. Run `dep ensure`. If you run `dep status` now it will confirm the package is installed.
    ```
    PROJECT                   CONSTRAINT  VERSION        REVISION  LATEST   PKGS USED
    github.com/franela/goreq  *           branch master  b5b0f5e   b5b0f5e  1
    ```

    Also, the vendor directory now will have the following files from the package:
    ```
    vendor
    └── github.com
        └── franela
            └── goreq
                ├── LICENSE
                ├── Makefile
                ├── README.md
                ├── goreq.go
                ├── goreq_test.go
                └── tags.go

    ```
