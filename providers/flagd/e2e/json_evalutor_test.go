package e2e

import (
	"flag"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"testing"

	"github.com/cucumber/godog"
	"github.com/open-feature/go-sdk-contrib/tests/flagd/pkg/integration"
	"github.com/open-feature/go-sdk/pkg/openfeature"
)

func TestJsonEvaluatorInRPC(t *testing.T) {
	if testing.Short() {
		// skip e2e if testing -short
		t.Skip()
	}

	flag.Parse()

	name := "flagd-json-evaluator.feature"

	testSuite := godog.TestSuite{
		Name: name,
		ScenarioInitializer: integration.InitializeFlagdJsonScenario(func() openfeature.FeatureProvider {
			return flagd.NewProvider(flagd.WithPort(8013))
		}),
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../../flagd-testbed/gherkin/flagd-json-evaluator.feature"},
			TestingT: t, // Testing instance that will run subtests.
			Strict:   true,
		},
	}

	if testSuite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run evaluation tests")
	}
}

func TestJsonEvaluatorInProcess(t *testing.T) {
	if testing.Short() {
		// skip e2e if testing -short
		t.Skip()
	}

	flag.Parse()

	name := "flagd-json-evaluator.feature"

	testSuite := godog.TestSuite{
		Name: name,
		ScenarioInitializer: integration.InitializeFlagdJsonScenario(func() openfeature.FeatureProvider {
			return flagd.NewProvider(flagd.WithInProcessResolver(), flagd.WithPort(9090))
		}),
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../../flagd-testbed/gherkin/flagd-json-evaluator.feature"},
			TestingT: t, // Testing instance that will run subtests.
			Strict:   true,
		},
	}

	if testSuite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run evaluation tests")
	}
}