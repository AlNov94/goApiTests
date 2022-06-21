package tests

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	"goApiTests/internal/api"
	"goApiTests/internal/dto"
	"goApiTests/internal/repository"
	"os"
	"testing"
)

type TestSuite struct {
	suite.Suite
}

func TestRunSuite(t *testing.T) {
	runner.Run(t, "Just Link", func(t provider.T) {
		t.SetIssue("https://pkg.go.dev/github.com/stretchr/testify")
	})

	runner.Run(t, "With Pattern", func(t provider.T) {
		_ = os.Setenv("ALLURE_ISSUE_PATTERN", "https://pkg.go.dev/github.com/stretchr/%s")
		t.SetIssue("testify")
	})
	t.Parallel()
	suite.RunSuite(t, new(TestSuite))
}

func (ts *TestSuite) TestDatabase(t provider.T) {
	assert := assert.New(t)
	u := repository.GetUserRepository().FindUserByID(1)
	assert.Equal(u.FirstName, "Joe")
}

func (ts *TestSuite) TestApi(t provider.T) {
	assert := assert.New(t)
	var u dto.UserResponse
	api.GetMockService().GetMockUser("1", &u)
	assert.Equal(u.FirstName, "Joe")
}
