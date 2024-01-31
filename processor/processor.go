package processor

import (
	"DocumentManager/controller"
	"DocumentManager/errors"
	"DocumentManager/model"
	"DocumentManager/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ProcessFile(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		panic(errors.ErrInvalidCommand)
	}

	if controller.Store == nil {
		controller.InitStore()
		controller.InitUser()
		controller.InitDocumentDetails()
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputText := scanner.Text()

		command, err := addCommand(inputText)
		if err != nil {
			panic(errors.ErrInvalidCommand)

		}

		err = command.Connection.Execute(&command)
		if err != nil {
			panic(err)
		}
	}
}

func addCommand(inputText string) (model.Command, error) {
	inputCommand := strings.Split(inputText, " ")
	var command model.Command

	if len(inputCommand) == 0 {
		return command, errors.ErrInvalidCommand
	}

	command.CommandName = inputCommand[0]

	if command.CommandName == utils.CommandExit {
		os.Exit(1)
	}

	err := addConnection(&command)
	if err != nil {
		return command, err
	}

	err = addArguments(&command, inputText)
	if err != nil {
		return command, err
	}

	return command, nil
}

func addConnection(command *model.Command) error {

	switch command.CommandName {
	case utils.CommandCreateUser:
		command.Connection = controller.Store.CreateUser()
	case utils.CommandCreateDocument:
		command.Connection = controller.Store.CreateDocument()
	case utils.CommandUpdateDocument:
		command.Connection = controller.Store.UpdateDocument()
	case utils.CommandDeleteDocument:
		command.Connection = controller.Store.DeleteDocument()
	case utils.CommandGetUserDocuments:
		command.Connection = controller.Store.GetUserDocuments()
	case utils.CommandRevertVersion:
		command.Connection = controller.Store.RevertToVersion()
	default:
		return errors.ErrInvalidCommand

	}

	return nil
}

func addArguments(command *model.Command, inputText string) error {
	inputCommand := strings.Split(inputText, " ")

	switch command.CommandName {
	case utils.CommandCreateUser:
		command.Arguments = model.User{
			UserName: inputCommand[2],
			UserId:   inputCommand[1],
			Password: inputCommand[3],
		}
	case utils.CommandCreateDocument:
		command.Arguments = model.CreateDocument{
			Id:      inputCommand[1],
			Title:   inputCommand[2],
			UserId:  inputCommand[3],
			Content: inputCommand[4],
		}
	case utils.CommandUpdateDocument:
		command.Arguments = model.UpdateDocument{
			Id:       inputCommand[1],
			Title:    inputCommand[2],
			UserId:   inputCommand[3],
			Password: inputCommand[4],
			Content:  inputCommand[5],
		}
	case utils.CommandDeleteDocument:
		command.Arguments = model.DeleteDocument{
			Id:       inputCommand[1],
			UserId:   inputCommand[2],
			Password: inputCommand[3],
		}
	case utils.CommandGetUserDocuments:
		command.Arguments = model.GetUserDocuments{
			UserId:   inputCommand[1],
			Password: inputCommand[2],
		}
	case utils.CommandRevertVersion:
		versionId, err := strconv.Atoi(inputCommand[2])
		if err != nil {
			return errors.ErrInvalidCommand
		}

		command.Arguments = model.RevertToVersion{
			DocumentId: inputCommand[1],
			VersionId:  versionId,
			UserId:     inputCommand[3],
			Password:   inputCommand[4],
		}
	default:
		return errors.ErrInvalidCommand

	}

	return nil
}
