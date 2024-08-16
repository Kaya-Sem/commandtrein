# commandtrein

[![Build](https://github.com/Kaya-Sem/commandtrein/actions/workflows/build.yml/badge.svg)](https://github.com/Kaya-Sem/commandtrein/actions/workflows/build.yml)

Commandtrein is a command-line interface (CLI) tool designed to access and display timetables and route information for SNCB (Belgian Railways) directly from your terminal.

See also: [commandlijn](https://github.com/Command-Transport/commandlijn)

## Features
- Timetables: retrieve and display the current timetable for any SNCB station.
- Routes: Get detailed connections and travel times between two SNCB stations.

![image](https://github.com/user-attachments/assets/33c21ed4-86ed-4a45-b6f4-7f047c9fd09a)


## Installation

#### 1. Install via `go install`

To quickly install commandtrein using go, run the following command:

```bash
go install github.com/Command-Transport/commandtrein@latest
```

This will download and install commandtrein and make it avaible in your `$GOPATH/bin` directory. Ensure that this directory is included in your system's `PATH` environment variable.

#### 2. Install from source
If you prefer to build commandtrein from source, follow these steps:

```bash
git clone https://github.com/Command-Transport/commandtrein.git
cd commandtrein
go build -o commandtrein
sudo mv commandtrein /usr/local/bin/
```

## Usage

Commandtrein supports the following actions:

#### Display timetable / liveboard for a station

```bash
commandtrein <station> # shows timetable for station
```

*Example:*
```bash
commandtrein Gent-Sint-Pieters
```
#### Find connections between two stations
```bash
commandtrein <station1> to <name2>
```
*Example:*
```bash
commandtrein Gent-Sint-Pieters to Zingem
```

this command will display the avaible connections, including departure and arrival times, duration, and platform information.



#### Acknowledgements

Commandtrein leverages the iRails API, an open-source API for accessing real-time data from SNCB. We extend our thanks to the iRails team for their excellent service.
