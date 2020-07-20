# SUSE Manager - api automation - *autoapi*

This program automates spacewalk/uyuni/SUSE-Manager xmlrpc api calls by reading out from an user-given 
configuration file in yaml format. The api call's output can be used as input for next api call and so on.

The program is written in go.

## __Benefits__
* No need to ask scripter to create many python/perl/etc. script just to automate some api calls. You can do it by yourself. __Save time, be flexible and be independant__ :-)
* The yaml config file allows to puzzle your desired api calls with arbitrary input and output vars, and automate it as you want. 

* The program automates api calls with your desired input parameters and values, including api url, username, password etc. Higher security as the file should be root user rw only.
* the program can use output from previous api call to be used as input for the next api call.

__For example: Call-A ouptuts list of serverid, Call-B need serverid to find installable patches and returns a list of patches and finally the last Call-C will use the list of serverid and list of patches to schedule installation jobs.__

## Prerequisites:
* SUSE Manager 4.0.x (tested)
* ....

## __Usage__:
1. Download this github repo to your local machine:

```git clone https://github.com/bjin01/autoapi.git```

2. Copy the binary autoapi to your preferred binary directory which is in your path.

3. Prepare your configuration file in yaml format.
As an example look at the config.yml file

4. Run the program with your configuration file.
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
