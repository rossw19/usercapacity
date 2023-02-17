<br />
<div align="center">
  <h1>
    usercapacity
  </h1>
  <br />
  <br />
  <br />
  <h3 align="center">Evaluate developer capacity for sprints</h3>
  <br />
</div>

## About

<div>
  <br />
  <p>User Capacity is a tool for evaluating how much time a developer will be able to spend on tickets during a sprint. It gathers data from Everhour, and a calendar feeds to determine this figure and sends it to Jira.</p>
  <br />
</div>

## Settings

<div>
  <br />
  <p>
    The tool uses a yaml configuration file to determine the parameters for the application (similar to a dot-env). I have chosen this design decision so that key-value pairs can be added, along with user mapping between Everhour, Jira and the calendar feed. To start, rename the sample configuration file and fill in the values for API keys and API URLs. Some values are slightly ambiguous and explained below.
  </p>
  <br />
</div>

```bash
git clone https://github.com/rossw19/usercapacity.git
cd usercapacity
go get ./...
go build cmd/capacity.go
```

<br />

## User Mapping

<div>
  <br />
  <p>
    Additionally, there is a user mapping configuration between each service. It contains the user identities from each feed. The calendar feed is mapped by name, as the service I use does not provide an employee ID.
  </p>
  <br />
</div>

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
