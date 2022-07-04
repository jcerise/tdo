package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"tdo/todo"
	"testing"
)

func NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add command",
		Run:   addRun,
	}
}

type AddCmdTestSuite struct {
	suite.Suite
	command  *cobra.Command
	datafile string
}

func (suite *AddCmdTestSuite) SetupSuite() {
	suite.datafile = "../test_data.json"
	// Viper takes an env key, upper-cases it, and appends the prefix TDO
	os.Setenv("TDO_DATAFILE", "../test_data.json")
}

func (suite *AddCmdTestSuite) TearDownSuite() {
	os.Unsetenv("TDO_DATAFILE")
	err := os.Remove(suite.datafile)
	if err != nil {
		log.Fatalf("could not remove file %v", suite.datafile)
	}
}

func (suite *AddCmdTestSuite) SetupTest() {
	suite.command = NewAddCmd()
}

func (suite *AddCmdTestSuite) TestAddTodo() {
	suite.command.SetArgs([]string{"test item"})
	suite.command.Execute()

	expectedItem := todo.Item{
		Text:     "test item",
		Position: 1,
		Priority: 2,
		Done:     false,
	}

	// Attempt to read from the file created, to check for the added value
	items, err := todo.ReadItems(suite.datafile)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), items[0], expectedItem)
	assert.Equal(suite.T(), 1, len(items))
}

func TestAddTestSuite(t *testing.T) {
	suite.Run(t, new(AddCmdTestSuite))
}
