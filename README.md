# Password hashing and verification

This package provides a simple way to hash and verify passwords using bcrypt.

## Installation

```bash
go get github.com/nlgolib/pwd
```

## Usage

```go
hash, err := pwd.Hash("password")
```

```go
verified := pwd.Verify("password", hash)
```
