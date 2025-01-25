<h1 align="center">Commandtrein</h1>

<p align="center">Commandtrein is a command-line interface (CLI) tool designed to access and display timetables and route information for SNCB (Belgian Railways) directly from your terminal.</p>

<p align="center"><a href="https://github.com/Kaya-Sem/commandtrein/wiki">Explore the documentation</a></p>

<p align="center"> • <a href="https://github.com/MDeLuise/plant-it/#features-highlight">Features highlights</a> • <a href="https://github.com/MDeLuise/plant-it/#quickstart">Quickstart</a></p>

![commandtrein](https://github.com/user-attachments/assets/f4343bf1-d8e4-4151-a2d3-4e4289307ad3)

# commandtrein

[![Build](https://github.com/Kaya-Sem/commandtrein/actions/workflows/build.yml/badge.svg)](https://github.com/Kaya-Sem/commandtrein/actions/workflows/build.yml)



## Features
- Timetables: retrieve and display the current timetable for any SNCB station.
- Routes: Get detailed connections and travel times between two SNCB stations.
- Shortcuts: Configure and use shortcuts for frequently used routes.

![commandtrein](https://github.com/user-attachments/assets/f4343bf1-d8e4-4151-a2d3-4e4289307ad3)

# Documentation

[Installation](https://github.com/Kaya-Sem/commandtrein/wiki/Installation)

[Usage](https://github.com/Kaya-Sem/commandtrein/wiki/Usage)

[Shell completion](https://github.com/Kaya-Sem/commandtrein/wiki/Shell-Tab-Completion)

## Configuration
Commandtrein uses a YAML configuration file located at `~/.config/commandtrein/config.yaml`. This files stores your shortcuts for frequent routes.

Example configuration:
```yaml
shortcuts:
  work:
    station1: "Brussel-Midi"
    station2: "Antwerpen-Centraal"
``` 

## Changelog

[changelog.md](https://github.com/Kaya-Sem/commandtrein/blob/main/CHANGELOG.md)

## Roadmap
- **Stylistic Improvements:** I appreciate any stylistic advice to enhance the user experience and code quality.
- **Upcoming Features:**
  - Filtering station results directly within the CLI.
  - Flags for departure time and date


#### Acknowledgements

Commandtrein leverages the iRails API, an open-source API for accessing real-time data from SNCB.
