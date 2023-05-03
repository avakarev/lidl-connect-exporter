# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Changed
- Bump go 1.19 => 1.20
- Bump go deps

## [1.1.1] - 2023-02-07
### Changed
- Bump go 1.18 => 1.19
- Bump go deps

## [1.1.0] - 2022-09-15
### Added
- Implement lidl_connect_consumption_expires_in_sec metric
### Changed
- Reduce scraping intervals:
    * balance: 12 hours -> 30 minutes
    * tariff: 24 hours -> 1 hour

## [1.0.0] - 2022-09-14
### Added
- Initial implementation

[Unreleased]: https://github.com/avakarev/lidl-connect-exporter/compare/v1.1.1...HEAD
[1.1.1]: https://github.com/avakarev/lidl-connect-exporter/compare/1.1.0...1.1.1
[1.1.0]: https://github.com/avakarev/lidl-connect-exporter/compare/1.0.0...1.1.0
[1.0.0]: https://github.com/avakarev/lidl-connect-exporter/releases/tag/v1.0.0
