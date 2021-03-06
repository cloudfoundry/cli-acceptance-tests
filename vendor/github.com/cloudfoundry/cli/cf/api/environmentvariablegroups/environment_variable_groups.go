package environmentvariablegroups

import (
	"fmt"
	"strings"

	"github.com/cloudfoundry/cli/cf/configuration/coreconfig"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/net"
)

//go:generate counterfeiter . EnvironmentVariableGroupsRepository

type EnvironmentVariableGroupsRepository interface {
	ListRunning() (variables []models.EnvironmentVariable, apiErr error)
	ListStaging() (variables []models.EnvironmentVariable, apiErr error)
	SetStaging(string) error
	SetRunning(string) error
}

type CloudControllerEnvironmentVariableGroupsRepository struct {
	config  coreconfig.Reader
	gateway net.Gateway
}

func NewCloudControllerEnvironmentVariableGroupsRepository(config coreconfig.Reader, gateway net.Gateway) (repo CloudControllerEnvironmentVariableGroupsRepository) {
	repo.config = config
	repo.gateway = gateway
	return
}

func (repo CloudControllerEnvironmentVariableGroupsRepository) ListRunning() (variables []models.EnvironmentVariable, apiErr error) {
	var raw_response interface{}
	url := fmt.Sprintf("%s/v2/config/environment_variable_groups/running", repo.config.APIEndpoint())
	apiErr = repo.gateway.GetResource(url, &raw_response)
	if apiErr != nil {
		return
	}

	variables, err := repo.marshalToEnvironmentVariables(raw_response)
	if err != nil {
		return nil, err
	}

	return variables, nil
}

func (repo CloudControllerEnvironmentVariableGroupsRepository) ListStaging() (variables []models.EnvironmentVariable, apiErr error) {
	var raw_response interface{}
	url := fmt.Sprintf("%s/v2/config/environment_variable_groups/staging", repo.config.APIEndpoint())
	apiErr = repo.gateway.GetResource(url, &raw_response)
	if apiErr != nil {
		return
	}

	variables, err := repo.marshalToEnvironmentVariables(raw_response)
	if err != nil {
		return nil, err
	}

	return variables, nil
}

func (repo CloudControllerEnvironmentVariableGroupsRepository) SetStaging(staging_vars string) error {
	return repo.gateway.UpdateResource(repo.config.APIEndpoint(), "/v2/config/environment_variable_groups/staging", strings.NewReader(staging_vars))
}

func (repo CloudControllerEnvironmentVariableGroupsRepository) SetRunning(running_vars string) error {
	return repo.gateway.UpdateResource(repo.config.APIEndpoint(), "/v2/config/environment_variable_groups/running", strings.NewReader(running_vars))
}

func (repo CloudControllerEnvironmentVariableGroupsRepository) marshalToEnvironmentVariables(raw_response interface{}) ([]models.EnvironmentVariable, error) {
	var variables []models.EnvironmentVariable
	for key, value := range raw_response.(map[string]interface{}) {
		stringvalue, err := repo.convertValueToString(value)
		if err != nil {
			return nil, err
		}
		variable := models.EnvironmentVariable{Name: key, Value: stringvalue}
		variables = append(variables, variable)
	}
	return variables, nil
}

func (repo CloudControllerEnvironmentVariableGroupsRepository) convertValueToString(value interface{}) (string, error) {
	stringvalue, ok := value.(string)
	if !ok {
		floatvalue, ok := value.(float64)
		if !ok {
			return "", fmt.Errorf("Attempted to read environment variable value of unknown type: %#v", value)
		}
		stringvalue = fmt.Sprintf("%d", int(floatvalue))
	}
	return stringvalue, nil
}
