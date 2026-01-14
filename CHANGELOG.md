# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- **BREAKING**: `Timeout` field type changed from `time.Duration` to `int` (seconds)
  - Affects `ChatCompletionRequest.Timeout`
  - Affects `EmbeddingRequest.Timeout`
  - Affects `Config.Timeout`
  - Default timeout changed from `600 * time.Second` to `600` (seconds)

### Added
- Unit tests for timeout JSON serialization

### Fixed
- Fixed third-party platform receiving incorrect timeout values (nanoseconds instead of seconds)

## [0.1.1] - Previous Release

### Added
- Initial implementation of LLM client

## [0.1.0] - Initial Release

### Added
- Basic chat completion support
- Embedding support
- Model listing support
