package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"os"
	"tdo/todo"
	"testing"
)

func NewDoneCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Done command",
		Run:   doneRun,
	}
}

type DoneCmdTestSuite struct {
	suite.Suite
	addCommand  *cobra.Command
	doneCommand *cobra.Command
	datafile    string
}

func (suite *DoneCmdTestSuite) SetupSuite() {
	suite.datafile = "../test_data.json"
	// Viper takes an env key, upper-cases it, and appends the prefix TDO
	os.Setenv("TDO_DATAFILE", "../test_data.json")
}

func (suite *DoneCmdTestSuite) TearDownSuite() {
	os.Unsetenv("TDO_DATAFILE")
	err := os.Remove(suite.datafile)
	if err != nil {
		log.Fatalf("could not remove file %v", suite.datafile)
	}
}

func (suite *DoneCmdTestSuite) SetupTest() {
	suite.addCommand = NewAddCmd()
	suite.doneCommand = NewDoneCmd()
}

func (suite *DoneCmdTestSuite) TearDownTest() {
	suite.addCommand = nil
	suite.doneCommand = nil
}

func (suite *DoneCmdTestSuite) TestMarkDone() {
	// First add a Todo item
	suite.addCommand.SetArgs([]string{"test item"})
	suite.addCommand.Execute()

	suite.doneCommand.SetArgs([]string{"1"})
	suite.doneCommand.Execute()

	items, err := todo.ReadItems(suite.datafile)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), true, items[0].Done)
}

func (suite *DoneCmdTestSuite) TestMarkDoneInvalid() {
	// First add a Todo item
	suite.addCommand.SetArgs([]string{"test item"})
	suite.addCommand.Execute()

	suite.doneCommand.SetArgs([]string{"2"})
	b := bytes.NewBufferString("")
	suite.doneCommand.SetOut(b)
	suite.doneCommand.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		suite.T().Fatal(err)
	}
	assert.Equal(suite.T(), "2 doesn't match any items\n", string(out))

	items, err := todo.ReadItems(suite.datafile)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), false, items[0].Done)
}

func (suite *DoneCmdTestSuite) TestMarkDoneAlreadyDone() {
	// First add a Todo item
	suite.addCommand.SetArgs([]string{"test item"})
	suite.addCommand.Execute()

	// Mark the Todo item as done
	suite.doneCommand.SetArgs([]string{"1"})
	suite.doneCommand.Execute()

	// Attempt to mark the Todo item done again
	suite.doneCommand.SetArgs([]string{"1"})
	b := bytes.NewBufferString("")
	suite.doneCommand.SetOut(b)
	suite.doneCommand.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		suite.T().Fatal(err)
	}
	assert.Equal(suite.T(), "item 1 is already marked as done\n", string(out))
}

func TestDoneTestSuite(t *testing.T) {
	suite.Run(t, new(DoneCmdTestSuite))
}
