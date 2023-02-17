<br />

## About

<div>
  <br />
  <p>User Capacity is a tool for evaluating how much time a developer will be able to spend on tickets during a sprint. It gathers data from Everhour, and a calendar feeds to determine this figure and sends it to Jira.</p>
  <br />
</div>

## Installation

```bash
git clone https://github.com/rossw19/usercapacity.git
cd usercapacity
go get ./...
go build cmd/capacity.go
```

<br />

## Configuration

<div>
  <br />
  <p>
    The tool uses a yaml configuration file to determine the parameters for the application (similar to a dot-env). I have chosen this design decision so that key-value pairs can be added, along with user mapping between Everhour, Jira and the calendar feed. To start, rename the sample configuration file and fill in the values for API keys and API URLs. Some values are slightly ambiguous and explained below.
  </p>
  <br />
</div>

| Key                                 | Description                                                                                                             |
|-------------------------------------|-------------------------------------------------------------------------------------------------------------------------|
| `api_url_everhour`                  | The base URL for the [Everhour API](https://everhour.docs.apiary.io/)                                                   |
| `api_url_jira`                      | The base URL for the [Jira API](https://docs.atlassian.com/software/jira/docs/api/REST/1000.824.0/)                     |
| `api_url_ics`                       | The base URL for the calendar feed                                                                                      |
| `api_email_jira`                    | The email address for the [Jira API](https://docs.atlassian.com/software/jira/docs/api/REST/1000.824.0/) authentication |
| `api_auth_everhour`                 | The API key for the [Everhour API](https://everhour.docs.apiary.io/)                                                    |
| `api_auth_jira`                     | The API key for the [Jira API](https://docs.atlassian.com/software/jira/docs/api/REST/1000.824.0/)                      |
| `api_version_everhour`              | The version of the [Everhour API](https://everhour.docs.apiary.io/)                                                     |
| `logging`                           | Enables logging to `var/log/debug.log` in root directory of application                                                 |
| `application_context_calendar_days` | The number of days a sprint takes to complete, including non-working days                                               |
| `application_context_working_days`  | The number of days a sprint takes to complete, excluding non-working days                                               |
| `application_context_average_over`  | The number of sprints the average is calculated over                                                                    |

<br />

## User Mapping

<div>
  <br />
  <p>
    Additionally, there is a user mapping configuration between each service. It contains the user identities from each feed. The calendar feed is mapped by name, as the service I use does not provide an employee ID.
  </p>
  <br />
</div>

| Key          | Description                                                                                                             |
|--------------|-------------------------------------------------------------------------------------------------------------------------|
| `everhourId` | The user ID for the [Everhour API](https://everhour.docs.apiary.io/)                                                     |
| `jiraId`     | The user ID for the [Jira API](https://docs.atlassian.com/software/jira/docs/api/REST/1000.824.0/)                      |
| `name`       | The name of the user in the calendar feed                                                                               |

<br />

## Usage

<div>
  <br />
  <p>
    After the application has been configured and built, it can be run by executing the file created when building the application. Additionally, a cron can be made to fire the application when required. Running the program during a workday when time is tracked will yield less accurate results.
  </p>
  <br />
</div>

## Contact

<br />

Ross Wilson - [ross.wilson.190298@gmail.com](mailto:ross.wilson.190298@gmail.com)

<br />
