# SUSE Manager - api automation - *autoapi*

This program automates spacewalk/uyuni/SUSE-Manager xmlrpc api calls by reading from an user-given 
configuration file in yaml format. The api call output can be used as input for the next api call and so on.

The program is written in go.

## Updates: 14. 08. 2020
* bugfixe: created a new checkprint function to print out api call error but not exit the program. 
    Sometimes the api call output can be empty but to continue proceed to the end of list.
* added a new sample yml for content-lifecycle-build and promote 
* renamed the method name to 'method1' 'method2' instead old name which was listmethodx. 


## Updates: 31. 07. 2020
* forked xmlrpc and changed import to use forked repo.
* added a new 'dependmethod' which is taking the output from method1 and loop through it as input for method2. 
* This 'dependmethod' is needed for example:
    If method 1 outputs a list of active systems in a given group.
    Then method 2 uses the output from method1 (serverid) and queries for every serverid the relevant errata as output in list. 
    Then finalmethod will create apply-errata job for each serverid with their relevant errata.
    For this purpose a new 'opotion' in yaml config file has been introduced (see config.yml) to trigger the 'dependmethod'
    
```
    options:
         meth2_depend_meth1: true
```

## Get api method and parameters from API doc
* In order to run the program you have to compose the config.yml file with method name, input and output parameters you wish.
* Go to SUSE Manager Web UI - left side menue tree -> Help -> API -> Overview
* Select a method namespace e.g. systemgroup
* Select in the namespace systemgroup a method e.g. ```listAdministrators```
* Now you can find documented input parameters: 
    ```
    sessionKey
    systemGroupName
    ```

__The sessionKey is not needed as this will be automatically added by the internal logic.__
You need to copy paste systemGroupName into the config.yml as input e.g.
```
method1:
  methodname: systemgroup.listActiveSystemsInGroup
  input_map:
    1_systemGroupName: test2
       
  out_variablenames:  
    - id
```
__Attention:__ 
* the methodname is case-sensitive.
* The input parameter names must not be the same as in the doc but the value has to be in the correct type. If input parameters are not given correctly the api call will fail.
* For correct ordering of the input parameters you have to prefix it with 1_, 2_, 3_ etc.
* The order of output parameters is done based on order of the lines.

## __Benefits__
* No need to ask scripter to create many python/perl/etc. script just to automate some api calls. You can do it by yourself. __Save time, be flexible and be independant__ :-)
* The yaml config file allows to puzzle your desired api calls with arbitrary input and output vars, and automate it as you want. 

* The program automates api calls with your desired input parameters and values, including api url, username, password etc. Higher security as the file should be root user rw only.
* the program can use output from previous api call to be used as input for the next api call.

__Example: Call-A ouptuts list of serverid, Call-B need serverid to find installable patches and returns a list of patches and finally the last Call-C will use the list of serverid and list of patches to schedule installation jobs.__ (see sample-configs/patch-active-systems-in-group.yml)

## Prerequisites:
* SUSE Manager 4.0.x (tested)
* go 1.14 (from packagehub for sles15sp1)

## For developers: 3rd party go-lib needed:
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
You can copy the binary ```autoapi``` to ```/usr/local/sbin``` or run it from the directory where you git cloned the repo to.
```
# cp autoapi /usr/local/sbin/
# autoapi -config config.yml

```
or

```./autoapi -config sample-configs/patch-active-systems-in-group.yml```

__Notes:__
You need to follow this rules in order to create your configuration file.
* the program supports up to 3 api calls. (method1, method2, finalmethod)
* for each api call you need to provide input variables and output variables.
* the input variables (input_map) will be handled as a map (dictionary) with name and value.
* the order of the input variables is handled through prefixed 1_, 2_, 3_ etc.
* the output variables will be read as a list. The order of the output var depends on the order in the configuration file.
  * look the section out_variablenames
* if boolean, datetime and or array is needed as input variable. Then you have to add the type to the variable name.
  * e.g. method1.array.id means from the output of method1, a list of id is needed as input for api call method2
  * e.g. datetime.2020-07-30T21:30:00 means the input is a datetime format and the schedule date time is in ISO8601 format.
  * e.g. bool.true means the input variable is of type boolean and the value is true or false


__below configuration file shows an example how the parameters for 3 api calls are being used.__

```
server:
  apiurl: http://suma.is.great/rpc/api
  username: admin
  password: hiphiphurra

method1:
  methodname: systemgroup.listActiveSystemsInGroup
  input_map:
    1_systemGroupName: test2
       
  out_variablenames:  
    - id


method2:
  methodname: system.getRelevantErrata
  input_map:
    1_serverid: method1.id

  out_variablenames: 
    - id
    - advisory_synopsis
    - advisory_name
    - advisory_type

finalmethod:
  methodname: system.scheduleApplyErrata
  options:
    meth2_depend_meth1: true
  input_map:
    1_serverid: method1.id
    2_errataId: method2.array.id
    3_earliestOccurrence: datetime.2020-08-23T09:45:00

  out_variablenames: 
    - actionId
```

## limitations:
* the program can only take up to 3 api calls.
* the program can only accept up to 5 input parameters for each api call.

## Next coming:
* optimize code
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