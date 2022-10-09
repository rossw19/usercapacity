package internal

import (
	"os"
	"rosswilson/usercapacity/model"
	"rosswilson/usercapacity/utility"
)

// Models that hold a reference to previous model
func createModels(userResp []byte, timeResp []byte, scheduleResp []byte, clock utility.Clocker) []model.Modeler {
	userModel := model.CreateEverhourUserModel(nil, userResp)
	timeModel := model.CreateEverhourTimeModel(userModel, timeResp)
	vacationModel := model.CreateVacationModel(timeModel, scheduleResp, clock)
	mathModel := model.CreateMathModel(vacationModel, clock)
	filterModel := model.CreateFilterModel(mathModel)
	jiraModel := model.CreateJiraModel(filterModel)

	return []model.Modeler{userModel, timeModel, vacationModel, mathModel, filterModel, jiraModel}
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
