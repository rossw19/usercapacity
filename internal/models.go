package internal

import (
	"os"
	"rosswilson/usercapacity/model"
	"rosswilson/usercapacity/utility"
)

func createModels(userResp []byte, timeResp []byte, clock utility.Clocker) []model.Modeler {
	userModel := model.CreateEverhourUserModel(nil, userResp)
	timeModel := model.CreateEverhourTimeModel(userModel, timeResp)
	mathModel := model.CreateMathModel(timeModel, clock)
	filterModel := model.CreateFilterModel(mathModel)
	jiraModel := model.CreateJiraModel(filterModel)

	return []model.Modeler{userModel, timeModel, mathModel, filterModel, jiraModel}
}

func bubbleModel(models []model.Modeler) model.Modeler {
	handler := model.CreateHandler(models)
	lastModel, err := handler.Handle().GetLastModel()

	if err != nil {
		utility.GetLogger().Write(err)
		os.Exit(1)
	}

	return lastModel
}
