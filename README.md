# Harpocrates

A secure key/value vault (localstorage) that makes use of distributed secrets to secure data in the vault. Harpocrates gives you a more secure way of saving your data in a manner where access to the data is only possible via distributed shared secrets. It is based off Shamir Secrets Sharing cryptography algorithm and Scdb for localStorage. Inspired by `Hashicorp Vault`, Harpocrates provides easy and faster implementation, accessibilty and flexibity of usage.

## Purpose

To ensure data storage have multiple, yet related secrets for authentication, limit risks resulting from a compromised secret.
Think of it like a Bank Account with multiple signatories, every participant has to sign before money can be taken out.

Readmore about [Distributed Secrets (Shamir Secrets Sharing)](https://khelechy.medium.com/understanding-distributed-secret-sharing-shamirs-secret-sharing-e76af7f4f6a5)

## Availability

Harpocrates is ready for use as both a Library and a Command Line Tool

## Dependencies

- golang +v1.19
- go-scdb
- hashicorp shamir

## Features

- Mount(parts)
- Unseal(secrets)
- Seal(secrets)
- Set(key, value)
- Get(key)

## Quick Start

- Ensure you have golang +v1.19 installed. You can check the [official instructions](https://go.dev/doc/install) for how
  to do that.

### As a Library

#### Install the package
```shell
go get github.com/khelechy/harpocrates
```


#### Import the package in your code

```go
import (
    ...
	"github.com/khelechy/harpocrates"
	...
)
```



#### Mount Harpocrates Vault on your environment `harpocrates.Mount(5)`
```go
harpocrates.Mount(5)
```

>**Mount(parts)** function takes an integer "parts" as a parameter, which indicates how many parts you want the generated secrets to be broken into,
>"parts" should not be less than 5 and not greater than 12.


The localStorage `harpocrates_db` folder and a `keys.json` file is generated in the root directory, which contains all the broken down secrets (distributed secrets), which would be used to `Unseal` and `Seal` the vault.


```json
{
  "keys": [
    "5105c79604545210e5fd30c1269ac1189e748aea93bae50db2d36bde8886adc143",
    "db206494f721b82e307417f54d18fc9396f12c067fd9d117a1f2e30fc5fb99f429",
    "94c7bcdf37cd0c2d48b19bf334b6e4c352a2729c55721fc9ef18aebfd94efe3ff8",
    "f7df31e448edb6613fd27528820845522120def4021ea4c13e5d906e8f74765bde",
    "25c662dc69637703eb04151fea9aa9c7f822c9444192d56414151bae2192ba6baa"
  ]
}
```



#### Unseal Harpocrates Vault on your environment `harpocrates.Unseal(secrets)`
Unsealing the vault makes it possible to now store and retrieve data securely, unlocks `Get` and `Set` operations.

```go
var secrets = []string{"f7df31e448edb6613fd27528820845522120def4021ea4c13e5d906e8f74765bde", "db206494f721b82e307417f54d18fc9396f12c067fd9d117a1f2e30fc5fb99f429", "25c662dc69637703eb04151fea9aa9c7f822c9444192d56414151bae2192ba6baa"}
harpocrates.Unseal(secrets)
```
>Unseal(secrets) function takes an array of `any three secrets`, combines and validates their authenticity, goes ahead to unseal the vault.




#### Seal Harpocrates Vault on your environment `harpocrates.Seal(secrets)`
Sealing the vault makes it impossible to store and retrieve data securely, locks `Get` and `Set` operations.

```go
var secrets = []string{"f7df31e448edb6613fd27528820845522120def4021ea4c13e5d906e8f74765bde", "db206494f721b82e307417f54d18fc9396f12c067fd9d117a1f2e30fc5fb99f429", "25c662dc69637703eb04151fea9aa9c7f822c9444192d56414151bae2192ba6baa"}
harpocrates.Seal(secrets)
```
>Seal(secrets) function takes an array of `any three secrets`, combines and validates their authenticity, goes ahead to seal the vault.

>**Important** `Unseal` must be called after `Mount`, so that the vault would available for `Set`and `Get` operations.




#### Storing data on Harpocrates Vault `harpocrates.Set(key, value)`
```go
harpocrates.Set("name", "kelechi")
```
>Get(key, value) function takes two string input indicating the key and value(data)




#### Retrieving data from Harpocrates Vault `harpocrates.Get(key)`
```go
harpocrates.Get("name")
```
>Get(key) function takes a string input indicating the key and returns the corresponding value if found.