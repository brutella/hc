package accessory

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/service"
)

type Accessory struct {
    Id int                   `json:"aid"`
    Services []*service.Service `json:"services"`
    
    Info *service.AccessoryInfo `json:"-"`
    idCount int
}


func New(info model.Info) *Accessory {
    i := service.NewInfo(info)
    a := &Accessory{
        idCount: 1,
        Info: i,
    }
    
    a.AddService(i.Service)
    
    return a
}

func (a *Accessory) SetId(id int) {
    a.Id = id
}

func (a *Accessory) GetId()int {
    return a.Id
}

func (a *Accessory) GetServices()[]model.Service {
    result := make([]model.Service, 0)
    for _, s := range a.Services {
        result = append(result, s)
    }
    return result
}

func (a *Accessory) Name() string {
    return a.Info.Name.Name()
}

func (a *Accessory) SerialNumber() string {
    return a.Info.Serial.SerialNumber()
}

func (a *Accessory) Manufacturer() string {
    return a.Info.Manufacturer.Manufacturer()
}

func (a *Accessory) Model() string {
    return a.Info.Model.Model()
}

// Adds a service to the accessory and updates the ids of the service and the corresponding characteristics
func (a *Accessory) AddService(s *service.Service) {
    s.SetId(a.idCount)
    a.idCount += 1
    
    for _, c := range s.GetCharacteristics() {
        c.SetId(a.idCount)
        a.idCount += 1
    }
    
    a.Services = append(a.Services, s)
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
