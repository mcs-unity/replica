# Read me

## Table of Contents

- [Summary](#summary)
- [Prerequisites](#prerequisites)
  - [Required](#required)
- [Usage](#usage)
  - [Initiation](#initiation)
  - [Synchronization](#synchronization)
- [Required](#required)
- [Built With](#built-with)
- [Authors](#authors)
- [License](#license)

## Summary

### Prerequisites

Before diving into the code, ensure you have all the necessary tools required to successfully debug and compile the package. To know exactly what you need, go to [Required](#required) section.

#### Required

The following tools are required for anyone who wishes to compile or run replica.
Please ensure you meet the requirements below:

- [Golang compiler](https://golang.org/) - Min Version 1.24.0

### Usage

The package intends to resolve cases where data needs to be shared between services.
As such the package must only be used where data sharing between services are needed.

Example a device connected to service A and a sync is required and a such it's triggered
Service B,C and D receives the data. A service may then perform X actions based on the data
that it received.

#### Initiation

To initiate a ReplicaSet you must first create a replica.json file in you root directory.

``` json
// url: url stands for in this case a http server which may be used to connect to you may also use unix sockets.
// auh: you will stor your authentication key here that may be submitted to the url.

[
    {
        "url": "http://localhost:3000/online",
        "auth": "185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969"
    }
]
```

Once the replica.json is done you then you can move on to the next step which is to initiate the ReplicaSet
as shown bellow

``` go
// if something goes wrong the replica will throw a panic
rep := start()
```

#### Synchronization

Once the replicaSet has been created it's time to perform the first sync, which will
make sure that given services are online and compatible with the ReplicaSet. This
step should be performed once durning boot to ensure that all given services are
online and functioning correctly. However you may used it as per your needs.

``` go
func isOnline(r replica.IReplica) error {
    resp, err := http.Get("http://localhost:3000/online")
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            return err
        }
        return fmt.Errorf("unable to process request, %d:%s", resp.StatusCode, body)
    }

    // do not call concurrently as defer resp.Body.Close() most likely will be called
    // which might trigger an error
    if err := r.Online(resp.Body); err != nil {
        return err
    }

    return nil
}

if err := rep.Sync(isOnline, false); err != nil {
    log.Println(err) // handle error as per your preference
}
```

isOnline is a callback function and in this case all we do is to check if the replica is online
by making a http get request to online(can be any url and method). We then make a simple check
for 200 and then parse the body (writer) to the replicas Online function for verification.
The response must be a valid response payload see bellow.

```go
{
  Online: bool
  time: time.Time
}
```

Calling Online is not required when making a sync call where data is synchronized to another service.
It's only used to verify that the service is still online

### Built With

- [Golang](https://golang.org/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [Editor config](https://editorconfig.org/)
- [GIT](https://git-scm.com/)

### Authors

- **MCS Unity** - _Initial work_ - [MCS Unity](https://github.com/mcs-unity)

See also the list of [contributors](https://github.com/mcs-unity/replica/graphs/contributors)
who participated in this project.

### License

This project is licensed under the GNU GENERAL PUBLIC V3 - see the [LICENSE](LICENSE) file for details
