package accessory

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/service"
)

type Accessory struct {
    Id int              `json:"aid"`
    Services []model.Service `json:"services"`
    
    info service.AccessoryInfo
    idCount int
}

func NewAccessory() *Accessory {
    return &Accessory{
        idCount: 1,
    }
}

func (a *Accessory) SetId(id int) {
    a.Id = id
}

func (a *Accessory) GetId()int {
    return a.Id
}

func (a *Accessory) GetServices()[]model.Service {
    return a.Services
}

func (a *Accessory) GetName() string {
    return a.info.Name.Name()
}

func (a *Accessory) GetSerialNumber() string {
    return a.info.Serial.SerialNumber()
}

func (a *Accessory) GetManufacturer() string {
    return a.info.Manufacturer.Manufacturer()
}

func (a *Accessory) GetModel() string {
    return a.info.Model.Model()
}

// Adds a service to the accessory and updates the ids of the service and the corresponding characteristics
func (a *Accessory) AddService(s model.Service) {
    s.SetId(a.idCount)
    a.idCount += 1
    
    for _, c := range s.GetCharacteristics() {
        c.SetId(a.idCount)
        a.idCount += 1
    }
    
    a.Services = append(a.Services, s)
    
    if info, ok := s.(service.AccessoryInfo); ok == true {
        a.info = info
    }
}

func (a *Accessory) Equal(other interface{}) bool {
    if accessory, ok := other.(*Accessory); ok == true {
        if len(a.Services) != len(accessory.Services) {
            return false
        }
        
        for i, s := range a.Services {
            if s.Equal(accessory.Services[i]) == false {
                return false
            }
        }
        
        return a.Id == accessory.Id
    }
    
    return false
}
