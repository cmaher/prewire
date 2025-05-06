# Prewire Requirements

## Overview
Prewire is a tool for generating wire provider sets from a more flexible and maintainable representation. It takes a `PreWireSet` variable and generates a static `WireSet` variable that can be used with Google's Wire dependency injection tool.

## Functional Requirements

### Input
1. The tool must look at an input Go package (the current directory).
2. It must identify a target wire set to build (a variable named `PreWireSet`).
3. The `PreWireSet` variable must be defined using the `prewire.NewSet()` function.
4. The `PreWireSet` can include direct providers via `prewire.Provider(key, provider)`.
5. The `PreWireSet` can include unionized sets via the `.Union()` method.

### Output
1. The tool must create a new generated file named `wire.go` in the same package.
2. The generated file must be a valid Go file.
3. The generated file must contain a static representation of the wire set (named `WireSet`).
4. The generated file must include all required imports.
5. The generated file must include a `go:generate` comment to regenerate it.

### Provider Functions
1. The tool must extract all provider functions from the `PreWireSet` variable.
2. Provider functions from unionized sets must be included in the generated `WireSet`.
3. Provider functions must be deduplicated to avoid duplicate entries in the generated `WireSet`.
4. Provider functions must maintain their package prefixes when appropriate.

### Imports
1. The generated file must include all necessary imports for the provider functions.
2. The generated file must not include unused imports.
3. The generated file must not include duplicate imports.
4. The tool must correctly resolve import paths for packages in various locations.
5. The tool must not make any package-specific special cases for imports.

### Error Handling
1. The tool must provide clear error messages when it fails to generate the wire file.
2. The tool must handle errors when parsing Go files, extracting provider functions, or writing the output file.
3. The tool must validate that the `PreWireSet` variable exists and is properly defined.

## Non-Functional Requirements

### Maintainability
1. The tool must work with any Go module, not just a specific one.
2. The tool must be able to find packages in various locations, including subdirectories of pkg.
3. The tool must not rely on hardcoded paths or package names.

### Usability
1. The tool must be easy to use with a simple command.
2. The tool must provide clear output indicating success or failure.
3. The generated file must be human-readable and well-formatted.

### Compatibility
1. The tool must be compatible with the latest version of Google's Wire tool.
2. The tool must work with standard Go project layouts.
3. The tool must work with non-standard project layouts or test environments.