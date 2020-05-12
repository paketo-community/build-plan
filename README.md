# Build Plan Cloud Native Buildpack
This buildpack is meant to be used to write a Build Plan generated from the contents of an `plan.toml` in the base of the application directory. The `plan.toml` is meant to reflect the contents of [Build Plan (TOML)](https://github.com/buildpacks/spec/blob/master/buildpack.md#build-plan-toml) that is currently supported by `packit`.

## Usage

To package this buildpack for consumption:

```
$ ./scripts/package.sh
```

This builds the buildpack's Go source using `GOOS=linux` by default. You can
supply another value as the first argument to `package.sh`.

## `plan.toml`

```toml
[[provides]]
name = "<dependency name>"

[[requires]]
name = "<dependency name>"
version = "<dependency version>"

[requires.metadata]
# buildpack-specific data
```

If you are looking for concrete definitions on what these fields do inside of `packit` you can check the documentation [here](https://godoc.org/github.com/paketo-buildpacks/packit#BuildPlan). For the definition from the Cloud Native Buildpack specification itself you can check out the documentation [here](https://godoc.org/github.com/paketo-buildpacks/packit#BuildPlan).
To package this buildpack for consumption:
