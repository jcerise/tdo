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

func (suite *AddCmdTestSuite) TestAddMultipleTodos() {
	suite.command.SetArgs([]string{"test item", "test item 2", "test item 3"})
	suite.command.Execute()

	expectedItem1 := todo.Item{
		Text:     "test item",
		Position: 1,
		Priority: 2,
		Done:     false,
	}

	expectedItem2 := todo.Item{
		Text:     "test item 2",
		Position: 2,
		Priority: 2,
		Done:     false,
	}

	expectedItem3 := todo.Item{
		Text:     "test item 3",
		Position: 3,
		Priority: 2,
		Done:     false,
	}

	expectedItems := []todo.Item{expectedItem1, expectedItem2, expectedItem3}

	// Attempt to read from the file created, to check for the added value
	items, err := todo.ReadItems(suite.datafile)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 3, len(items))

	for index, item := range items {
		assert.Equal(suite.T(), item, expectedItems[index])
	}
}

func (suite *AddCmdTestSuite) TestAddTodoWithPriority() {
	// TODO: For some reason if we don't redefine the flag here for the command, it is not available during execute
	//		 Figure out why.
	suite.command.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2, 3")

	suite.command.SetArgs([]string{"test item", "-p1"})
	suite.command.Execute()

	suite.command.SetArgs([]string{"test item 2", "-p3"})
	suite.command.Execute()

	expectedItemP1 := todo.Item{
		Text:     "test item",
		Position: 1,
		Priority: 1,
		Done:     false,
	}

	expectedItemP3 := todo.Item{
		Text:     "test item 2",
		Position: 2,
		Priority: 3,
		Done:     false,
	}

	// Attempt to read from the file created, to check for the added value
	items, err := todo.ReadItems(suite.datafile)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(items))

	assert.Equal(suite.T(), items[0], expectedItemP1)
	assert.Equal(suite.T(), items[1], expectedItemP3)

}

func TestAddTestSuite(t *testing.T) {
	suite.Run(t, new(AddCmdTestSuite))
}
