# SUSE Manager - api automation - *autoapi*

This program automates spacewalk/uyuni/SUSE-Manager xmlrpc api calls by reading out from an user-given 
configuration file in yaml format. The api call's output can be used as input for next api call and so on.

The program is written in go.

## Updates: 31. 07. 2020
forked xmlrpc and changed import to use forked repo.
added a new 'dependmethod' which is taking the output from method1 and loop through it as input for method2. 
This 'dependmethod' is needed for example:
  If method 1 outputs a list of active systems in a given group.
  Then method 2 uses the serverid and seeks for every single serverid the relevant errata as output 
  The finalmethod will then creates apply errata job for each serverid with their relevant errata.
  For this purpose a new 'opotion' in yaml config file has been introduced (see config.yml)
    
    ```options:
         meth2_depend_meth1: true```


## __Benefits__
* No need to ask scripter to create many python/perl/etc. script just to automate some api calls. You can do it by yourself. __Save time, be flexible and be independant__ :-)
* The yaml config file allows to puzzle your desired api calls with arbitrary input and output vars, and automate it as you want. 

* The program automates api calls with your desired input parameters and values, including api url, username, password etc. Higher security as the file should be root user rw only.
* the program can use output from previous api call to be used as input for the next api call.

__For example: Call-A ouptuts list of serverid, Call-B need serverid to find installable patches and returns a list of patches and finally the last Call-C will use the list of serverid and list of patches to schedule installation jobs.__

## Prerequisites:
* SUSE Manager 4.0.x (tested)
* go 1.14 (from packagehub for sles15sp1)

## 3rd party go-lib needed:
```cd $GOPATH```
* Download the xmlrpc lib from my repo which is a fork.
  ```go get github.com/bjin01/go-xmlrpc```

* gopkg.in/yaml.v2 (for yaml file reading)
  ```go get gopkg.in/yaml.v2```


## __Usage__:
1. Download this github repo to your local machine:

```git clone https://github.com/bjin01/autoapi.git```

2. Prepare your configuration file in yaml format.
As an example look at the config.yml file

3. Run the program with your configuration file.
```./autoapi -config ./config.yml```

__Notes:__
You need to follow this rules in order to create your configuration file.
* the program supports up to 3 api calls. (listmethod1, listmethod2, finalmethod)
* for each api call you need to provide input variables and output variables.
* the input variables (input_map) will be handled as a map (dictionary) with name and value.
* the order of the input variables is handled through prefixed 1_, 2_, 3_ etc.
* the output variables will be read as a list. The order of the output var depends on the order in the configuration file.
  * look the section out_variablenames
* if boolean, datetime and or array is needed as input variable. Then you have to add the type to the variable name.
  * e.g. listmethod1.array.id means from the output of listmethod1, a list of id is needed as input for api call listmethod2
  * e.g. datetime.2020-07-30T21:30:00 means the input is a datetime format and the schedule date time is in ISO8601 format.
  * e.g. bool.true means the input variable is of type boolean and the value is true or false


__below configuration file shows an example how the parameters for 3 api calls are being used.__

```
server:
  apiurl: http://suma.is.great/rpc/api
  username: admin
  password: hiphiphurra

listmethod1:
  methodname: systemgroup.listActiveSystemsInGroup
  input_map:
    1_systemGroupName: test2
       
  out_variablenames:  
    - id

listmethod2:
  methodname: system.getRelevantErrataByType
  input_map:
    1_serverid: listmethod1.id
    2_advisoryType: "Security Advisory"

  out_variablenames: 
    - id
    - advisory_synopsis
    - advisory_name

finalmethod:
  methodname: system.scheduleApplyErrata
  input_map:
    1_serverid: listmethod1.array.id
    2_errataId: listmethod2.array.id
    3_earliestOccurrence: datetime.2020-07-30T21:30:00
    4_allowModules: bool.true

  out_variablenames: 
    - actionId
```

## limitations:
* the program can only take up to 3 api calls.
* the program can only accept up to 10 input parameters for each api call.

## Next coming:
* optimize codes
* continue testing more api calls.
* continues bug-fixing

## 3rd party go-lib modifications:
* I needed to change the default boolean constant "true" and "false" to "1" and "0" because spacewalk xmlrpc expects that.
In the client.go (github.com/SHyx0rmZ/go-xmlrpc) below code snippet has been added from me.
```
  if strings.Contains(buffer.String(), "<boolean>true</boolean>") {

		newstrings := strings.Replace(buffer.String(), "<boolean>true</boolean>", "<boolean>1</boolean>", 1)
		buffer.Reset()
		buffer.WriteString(newstrings)

	} else if strings.Contains(buffer.String(), "<boolean>false</boolean>") {

		newstrings := strings.Replace(buffer.String(), "<boolean>false</boolean>", "<boolean>0</boolean>", 1)
		buffer.Reset()
		buffer.WriteString(newstrings)

	}
  ```
* due to the fact that spacewalk xmlrpc expects datetime.iso8601 format "20060102T15:04:05" I had to change the format as below in client.go (github.com/SHyx0rmZ/go-xmlrpc) file.
```
  case reflect.Struct:
			if v.Type().PkgPath() != "time" || v.Type().Name() != "Time" {
				return nil, &Error{"Invalid type " + v.Kind().String()}
			}

			t := arg.(time.Time)
			results = append(results, value{DateTimeTag: t.Format("20060102T15:04:05")})
```