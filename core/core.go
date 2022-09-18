package core

import (
	"fmt"
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/model"
	"rosswilson/usercapacity/utility"
	"time"
)

type core struct {
	apiContext api.Context
	handler    model.ModelHandler
	strategies struct {
		everhourStrategy *api.EverhourStrategy
		jiraStrategy     *api.JiraStrategy
	}
}

func (c *core) Run() {
	dates := utility.CreateDates(-21, time.Now())
	c.apiContext = api.CreateApiContext()

	c.createHandlers()
	c.prepareEverhourStartegy()

	userResp, userTime := c.executeEverhourStrategy(*dates)
	userModel := model.CreateEverhourUserModel(userResp)

	c.handler.Handle(userModel)
	dataModel := model.CreateEverhourTimeModel(userTime, userModel)
	c.handler.Handle(dataModel)
}

func (c *core) createHandlers() {
	jiraHandler := model.CreateJiraHandler(nil)
	mathHandler := model.CreateMathHandler(jiraHandler)
	timeHandler := model.CreateEverhourTimeHandler(mathHandler)
	userHandler := model.CreateEverhourUserHandler(timeHandler)
	c.handler = userHandler
}

func (c *core) prepareEverhourStartegy() {
	c.strategies.everhourStrategy = api.CreateEverhourStrategy()
	c.apiContext.SetApiStrategy(c.strategies.everhourStrategy)
}

func (c core) executeEverhourStrategy(dates utility.Dates) ([]byte, []byte) {
	c.strategies.everhourStrategy.SetRequestUri("/team/users")
	c.apiContext.ExecuteApi()
	userResponseBody := c.strategies.everhourStrategy.GetResponseBody()

	c.strategies.everhourStrategy.SetRequestUri(fmt.Sprintf("/dashboards/users?date_gte=%s&date_lte=%s", dates.GetFrom(), dates.GetTo()))
	c.apiContext.ExecuteApi()
	timeResponseBody := c.strategies.everhourStrategy.GetResponseBody()

	return userResponseBody, timeResponseBody
}

func (c *core) prepareJiraStrategy() {
	c.strategies.jiraStrategy = api.CreateJiraStrategy()
	c.apiContext.SetApiStrategy(c.strategies.everhourStrategy)
}

func (c core) executeJiraStrategy(url string) {
	c.apiContext.ExecuteApi()
}

func CreateCore() *core {
	return &core{}
}
