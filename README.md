# THIS REPO IS DEFUNCT
These tests have been merged into the [integration](https://github.com/cloudfoundry/cli/tree/master/integration) of the CF CLI. This repo is left here for historical purposes.

CLI Acceptance Tests
====
These are high-level tests for the [Cloud Foundry
CLI](https://github.com/cloudfoundry/cli) that make assertions about the
behavior of the `cf` binary.

These tests require that a `cf` binary built from the latest source is
available in your `PATH`.

# New Suite:
These are the notes for the `integration` suite tests.

# How to run:
Running is simple:

```
ginkgo -p -r -randomizeAllSpecs -slowSpecThreshold=120 ./integration
```

Customizations (based on environment variables):

- `CF_API` - Sets the CF API URL these tests will be using. Will `--skip-ssl-validation` by default. Should default to `api.bosh-lite.com` if not set.
- `SKIP_SSL_VALIDATION` - Sets the CF API URL these tests will be using. Will `--skip-ssl-validation` by default. Should default to `api.bosh-lite.com` if not set.



# Old Suite:
These are the notes for the `gats` directory tests.

### Installing from source

1. Install [Go](https://golang.org/dl)
1. Ensure your `$GOPATH` [is set correctly](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable)
1. Get the cli source code: `go get -u github.com/cloudfoundry/cli` (ignore the "no buildable Go source files" warning)
1. Run `go get -u github.com/jteeuwen/go-bindata/...`
1. Run `bin/build` in `$GOPATH/src/cloudfoundry/cli`
1. Copy `$GOPATH/src/cloudfoundry/cli/out/cf` to a location in your `PATH`

### Running the suite

These tests are similar to the [CF Acceptance
Tests](https://github.com/cloudfoundry/cf-acceptance-tests), and use the same
configuration and test helpers.

To run the tests (example given is for [bosh-lite](https://github.com/cloudfoundry/bosh-lite)):

```
cat > gats_config.json <<EOF
{
  "api": "api.bosh-lite.com",
  "apps_domain": "bosh-lite.com",
  "skip_ssl_validation": true,
  "use_http": true,
  "admin_password": "admin",
  "admin_user": "admin",
  "existing_user": "admin",
  "existing_user_password": "admin"
}
EOF
export CONFIG=$PWD/gats_config.json
ginkgo -r
```
