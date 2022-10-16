# User Capacity

---

User Capacity is a tool for evaluating how much time a developer will be able to spend on tickets during a sprint. It
gathers data from [Everhour API](https://everhour.docs.apiary.io/) and a calendar feed to determine this figure, and
sends
the data
to the [Jira API](https://docs.atlassian.com/software/jira/docs/api/REST/1000.824.0/).

* [Installation](#installation)
* [Configuration](#configuration)
* [Usage](#usage)

## Installation

---

```
git clone https://github.com/rossw19/usercapacity.git
cd usercapacity
go get ./...
go build cmd/capacity.go
```

## Configuration

---

### General Application Settings

The tool uses a yaml configration file to determine the parameters for the tool. This is so key-value pairs can be
added, along with user mapping between [Everhour](https://everhour.com/),
[Jira](https://www.atlassian.com/software/jira) and the calendar feed. To get started, rename the `config.yml.sample`
to `config.yml` and fill in the values such as API keys and API urls. Some values are slightly ambiguous and explained
below.

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

### User Mapping

Additionally, there is a user mapping configuration between each service. It contains the user identities
from each feed. The calendar feed is mapped by name as the service I personally use does not provide an employee ID.

| Key          | Description                                                                                                             |
|--------------|-------------------------------------------------------------------------------------------------------------------------|
| `everhourId` | The user ID for the [Everhour API](https://everhour.docs.apiary.io/)                                                     |
| `jiraId`     | The user ID for the [Jira API](https://docs.atlassian.com/software/jira/docs/api/REST/1000.824.0/)                      |
| `name`       | The name of the user in the calendar feed                                                                               |

## Usage

---

After the application has been configured and built, it can be ran by executing the file created when building the
application or simply by setting up a cron to fire the application when required. Running the program during a workday
when time is being tracked will yield less accurate results. 