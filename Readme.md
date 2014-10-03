# GOHAP

This is an implementation of the **H**omeKit **A**ccessory **P**rotocol.


## IO with Virtual Devices

On linux, IO is done via files (e.g. [GPIO](https://developer.ridgerun.com/wiki/index.php/How_to_use_GPIO_signals) input is written to a file). Files can read and write to those files.

The HomeKit brige is a gateway to the HomeKit universe. It could act like a daemon which watches for accessory changes and delivers that to the HomeKit clients. An acccesory can offer services which are specified by the *HAP*.

- Info service: info (name, manufacturer, ...) about accessory, every bridge must have it too
- Thermostat service: read and write to temperature
- Switch service: dis-/enable a switch
- ...

Those services are made of characteristics like

- name
- current temperature
- target temperature
- on
- ...

Characteristics can be read and write by clients, and the bridge is the mediator to the accessory. The interaction between accessory and bridge should be based on files, very similar to virtual devices on linux. The accessory and bridge watches the file system. This allows multiple accessories to use the same bridge as a gateway.

There must be a contract between accessory and bridge on the file system representation.

**Thermostat**

| file | value |
| ---- | ----- |
| type | thermostat |
| serial-number| string |
| name | string |
| manufacturer| string |
| model| string |
| unit | celsius |
| temperature| float |
| target-temperature| float |
| mode| off | heating | cooling |
| target-mode| off | heating | cooling |
| humidity| float |
| target-humidity | float |

**Switch**

| file | value |
| ---- | ----- |
| type | switch |
| serial-number| string |
| name | string |
| manufacturer| string |
| model| string |
| on | 0 | 1 |

The files must be located in a specific folder. On startup, the bridge scans the folder and create a model representation. On file changes, the bridge re-initialized the model and notifies the client for new values by updating the Bonjour TXT record *s#*.

The accessory 